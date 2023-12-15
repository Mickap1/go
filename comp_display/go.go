package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

type RadientBackground struct {
	Background struct {
		Colors []struct {
			Color   string `json:"color"`
			Percent int    `json:"percent"`
		} `json:"colors"`
	} `json:"background"`
}

func main() {
	//Read the config.json file
	configData, err := os.ReadFile("config/config.json")
	if err != nil {
		fmt.Println("Error reading config.json:", err)
		return
	}

	//Parse the JSON data
	var config RadientBackground
	err = json.Unmarshal(configData, &config)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// var path_to_image string = "/images/tipek.png"
	// imgSrc := path.Join("/", path_to_image)

	//Display the data
	app := fiber.New()
	app.Static("/images", "./images")
	app.Get("/", func(c *fiber.Ctx) error {
		html := `<html>
				<head>
					<style>
						body {
							font-family: Arial, sans-serif;
							text-align: center;
							margin: 0;
							padding: 0;
							height: 100vh; /* Set to 100% viewport height */
							background: linear-gradient(to bottom, %s);
						}
						h1 {
							color: white; /* Set the text color to white or another contrasting color */
						}
						img {
							max-width: 100%;
							height: auto;
						}
					</style>
				</head>
				<body>
					<h1>Welcome to My Webpage</h1>
					<img src="/images/tipek.png" alt="My Image">
				</body>
			</html>`
		return c.Type("html").SendString(html)
	})

	nerr := app.Listen(":3000")
	if nerr != nil {
		panic(err)
	}
}
