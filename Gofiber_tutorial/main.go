package main

import (
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

// database instance
var db *sqlx.DB

const jwtSecret = "infinitas"

// Database
const (
	host     = "localhost"
	port     = "3306" // default port
	user     = "root"
	password = "new_password"
	dbname   = "gofiber_tutorial"
)

type User struct {
	Id       int    `db:"id" json:"id" query:"id"`
	Username string `db:"username" json:"username" query:"name"`
	Password string `db:"password" json:"password" query:"password"`
}

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SigninRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	Users []User `json:"user"`
}

func GetUserList(c *fiber.Ctx) error {
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	result := UserResponse{}
	for rows.Next() {
		user := User{}
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			return err
		}
		// append user to userResponse
		result.Users = append(result.Users, user)
	}
	// return user in json format
	return c.Status(200).JSON(result)
}

func Signup(c *fiber.Ctx) error {
	request := SignupRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}
	if request.Username == "" || request.Password == "" {
		return fiber.ErrUnprocessableEntity
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	// check duplicate username
	// if request.Username == name {
	// 	return fiber.NewError(fiber.StatusUnprocessableEntity, "Duplicate username")
	// }

	// add new user
	query := "insert user (username , password) values (?,?)"
	result, err := db.Exec(query, request.Username, string(password))
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}
	id, err := result.LastInsertId() // id ล่าสุด
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}
	// response
	user := User{
		Id:       int(id),
		Username: request.Username,
		Password: string(password),
	}
	return c.JSON(user)
}
func Signin(c *fiber.Ctx) error {
	request := SigninRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}
	if request.Username == "" || request.Password == "" {
		return fiber.ErrUnprocessableEntity
	}
	user := User{}
	query := "select * from user where username=?"
	err = db.Get(&user, query, request.Username)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Incorrect username")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Incorrect password")
	}

	cliams := jwt.StandardClaims{
		Issuer:    strconv.Itoa(user.Id),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodES256, cliams)
	token, err := jwtToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(fiber.Map{
		"jwtToken": token,
	})

}

// connect database
func Connect() error {
	var err error
	db, err = sqlx.Open("mysql", user+":"+password+"@tcp"+"("+host+":"+port+")"+"/"+dbname)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil

}

func main() {
	// connect database
	if err := Connect(); err != nil {
		panic(err)
	}

	app := fiber.New(fiber.Config{
		Prefork: true,
	})
	v1 := app.Group("/v1")

	user := v1.Group("/user")
	user.Use("/hello", jwtware.New(jwtware.Config{
		SigningMethod: "HS256",
		SigningKey:    []byte(jwtSecret),
		SuccessHandler: func(c *fiber.Ctx) error {
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return fiber.ErrUnauthorized
		},
	}))
	user.Get("/list", GetUserList)
	user.Post("/signup", Signup)
	user.Post("/signin", Signin)

	app.Listen(":8000")
}
