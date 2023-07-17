package router

/*  example

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func Fiber() {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	// middle ware
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("name", "test") // set varible
		fmt.Println("before")
		err := c.Next() // c.Next() => ทำงานลำดับต่อไป

		fmt.Println("after")
		return err
	})

	app.Use(requestid.New())

	// cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	// logger
	app.Use(logger.New(logger.Config{
		TimeZone: "Asia/Bangkok",
	}))

	// get
	app.Get("/hello", func(c *fiber.Ctx) error {
		name := c.Locals("name") // name คือ ตัวแปรจาก middleware
		return c.SendString(fmt.Sprintf("Get : Hello %v", name))
	})
	// post
	app.Post("/hello", func(c *fiber.Ctx) error {
		return c.SendString("POST : Hello World ")
	})

	// get params
	app.Get("/hello/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		return c.SendString("name : " + name)
	})

	// parameter optional
	app.Get("/hello/:name/:surname?", func(c *fiber.Ctx) error {
		name := c.Params("name")
		surname := c.Params("surname")
		return c.SendString("name : " + name + " , surname : " + surname)
	})

	// params int
	app.Get("/hello/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return fiber.ErrBadRequest
		}
		return c.SendString(fmt.Sprintf("id : %v", id))
	})

	// query
	app.Get("/query?name='test'", func(c *fiber.Ctx) error {
		name := c.Query("name")
		return c.SendString("name : " + name)
	})

	app.Get("/query/json?id=1&name=test", func(c *fiber.Ctx) error {
		person := Person{}
		c.QueryParser(&person)
		return c.JSON(person)
	})

	// Static file
	app.Static("/", "./static", fiber.Static{
		Index: "index.html",
	})

	// NewError
	app.Get("/error", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusNotFound, "not found")
	})

	// Group
	v1 := app.Group("/v1")
	v1.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello v1")
	})

	// Mount เหมือน app.Group แต่กำหนด path หลักได้เลย "/user"
	userApp := fiber.New()
	userApp.Get("/login", func(c *fiber.Ctx) error {
		return c.SendString("login")
	})

	app.Mount("/user", userApp)

	// server
	app.Server().MaxConnsPerIP = 1 // MaxConnsPerIP ยิง api ได้แค่ครั้งเดียว
	app.Get("/server", func(c *fiber.Ctx) error {
		time.Sleep(time.Second * 30)
		return c.SendString("server")
	})

	// env
	app.Get("/env", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"BaseURL":     c.BaseURL(),
			"Hostname":    c.Hostname(),
			"IP":          c.IP(), // IP Address
			"IPS":         c.IPs(),
			"OriginalURL": c.OriginalURL(),
			"Path":        c.Path(),
			"Protocol":    c.Protocol(),
			"Subdomains":  c.Subdomains(),
		})
	})

	// Body
	app.Post("/body", func(c *fiber.Ctx) error {
		fmt.Printf("Is JSON : %v\n", c.Is("json"))
		// fmt.Println(string(c.Body()))
		data := make(map[string]interface{})

		err := c.BodyParser(&data)
		if err != nil {
			return err
		}
		return c.SendString(fmt.Sprintf("data : %v\n", data))

	})

	app.Listen(":8000")
}
*/
