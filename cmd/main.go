package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
	"github.com/weezqyd/hello-go/handlers"
	"log"
	"os"
)

func main() {
	var port, dburl string
	flag.StringVar(&port, "port", ":3000", "port to listen")
	flag.StringVar(&dburl, "dsn", "postgres://homestead:secret@192.168.10.10:5432/hello", "database dsn")
	flag.Parse()
	
	conn, err := pgx.Connect(context.Background(), dburl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	
	controller := &handlers.Controller{DB: conn}
	app := fiber.New()
	app.Get("/", controller.Welcome)
	app.Get("/widgets", controller.Widgets)
	
	log.Fatal(app.Listen(port))
}
