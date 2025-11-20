package main

import (
    "log"
    "github.com/joho/godotenv"
    "go-proj/route"
    "github.com/gofiber/fiber/v2"
    "go-proj/configs"
)

func main() {
    godotenv.Load()

    configs.InitDB()

    app := fiber.New()

    route.InitRoutes(app)

    log.Fatal(app.Listen(":3000"))
}