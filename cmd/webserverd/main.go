package main

import (
	"context"
	"fmt"
	"time"

	appConfig "github.com/audryus/crispy-octo-system/configs"
	"github.com/audryus/crispy-octo-system/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"

	"log"
)

type Ticket struct {
	Name    string
	Seating string
	Price   string
}

func getSeating(idx int) string {
	if idx%3 == 0 {
		return "Seated"
	}
	return "Stading"
}
func getPrice(idx int) string {
	return fmt.Sprintf("%.2f", (float32(idx+1) * 527.57))
}

func lero() []Ticket {
	var tickers []Ticket

	for i := 0; i < 5; i++ {
		tickers = append(tickers, Ticket{
			Name:    fmt.Sprintf("My Name %d", i+1),
			Seating: getSeating(i),
			Price:   getPrice(i),
		})
	}

	return tickers
}

func withLoggedRedirect(h fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Printf("tefff")
		return h(c)
	}
}

var replies []string

func main() {
	cfg := appConfig.New()
	httpd := cfg.HTTP

	//database.InitCockroachConnection(cfg.Cockroach)
	database.InitRedis(cfg.Redis)
	//database.InitSupabase(cfg.Supabase)

	engine := django.New("./web/app", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	app.Static("/static", "./web/static")

	app.Use(func(c *fiber.Ctx) error {
		name := "Nome de teste"

		data := make(map[string]string)
		data["Name"] = name

		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Tickets": lero(),
		})
	})

	app.Get("/signin", func(c *fiber.Ctx) error {
		return c.Render("signin/index", fiber.Map{})
	})

	app.Post("/signin", func(c *fiber.Ctx) error {
		email := c.FormValue("email")
		fmt.Printf("email: %s\n", email)
		_, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		response := fiber.Map{
			"reply": "RPL_SIGNIN",
		}

		/* err := database.Supa.Auth.SendMagicLink(ctx, email)
		if err != nil {
			fmt.Printf("Err %v", err)
		} */

		response["email"] = email

		return c.Render("signin/index", response)
	})

	app.Get("/signup", func(c *fiber.Ctx) error {
		invite := c.Query("convite")
		fmt.Printf("%s\n", invite)
		return c.Render("signup/index", fiber.Map{})
	})

	app.Get("/home", func(c *fiber.Ctx) error {
		return c.Render("home/index", fiber.Map{})
	})

	app.Get("/home/convite", func(c *fiber.Ctx) error {
		return c.Render("home/convite/index", fiber.Map{})
	})

	app.Post("/home/convite", func(c *fiber.Ctx) error {
		nome := c.FormValue("nome")
		email := c.FormValue("email")
		fmt.Printf("Nome: %s (%s)\n", nome, email)

		replies = make([]string, 0)
		replies = append(replies, "RLY_ERR")
		response := fiber.Map{
			"replies": replies,
		}
		header := c.Request().Header.Peek("HX-Request")
		fmt.Printf("Header: %s \n", header)
		c.Append("HX-Trigger", "{\"response\": \"RLY_ERR\", \"event2\": \"\"}")
		return c.Render("home/convite/form", response)
	})

	log.Fatal(app.Listen(fmt.Sprintf("0.0.0.0:%s", httpd.Port)))
}
