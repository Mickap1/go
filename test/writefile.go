package main

import (
	"bufio"
	"fmt"
	"os"
)

func writeStructToFile(data Game, fileName string) error {
	// Open the file for writing
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	// Create a buffered writer
	writer := bufio.NewWriter(file)

	// fmt.Fprintf(writer, "Field2: %d\n", data.Field2)

	fmt.Fprintf(writer, "<!DOCTYPE html>\n")
	fmt.Fprintf(writer, "<html>\n")
	fmt.Fprintf(writer, "<head>\n")
	fmt.Fprintf(writer, "<title>%d / %d</title>\n", data.Point, data.Round)
	fmt.Fprintf(writer, "</head>\n")
	fmt.Fprintf(writer, "<body>\n")
	fmt.Fprintf(writer, "<h1>%s</h1>\n", data.Title)
	fmt.Fprintf(writer, "<img src=\"%s\" alt=\"Sample Image\" width=\"300\" height=\"200\">\n", data.ImagePath)
	fmt.Fprintf(writer, "<button onclick=\"buttonClicked('%s')\">%s</button>\n", data.Choice1, data.Choice1)
	fmt.Fprintf(writer, "<button onclick=\"buttonClicked('%s')\">%s</button>\n", data.Choice2, data.Choice2)
	fmt.Fprintf(writer, "<button onclick=\"buttonClicked('%s')\">%s</button>\n", data.Choice3, data.Choice3)
	fmt.Fprintf(writer, "<button onclick=\"buttonClicked('%s')\">%s</button>\n", data.Choice4, data.Choice4)
	fmt.Fprintf(writer, "<script>\n")
	fmt.Fprintf(writer, "function buttonClicked(buttonNumber) {\n")
	fmt.Fprintf(writer, "fetch('/handleClick', {\n")
	fmt.Fprintf(writer, "method: 'POST',\n")
	fmt.Fprintf(writer, "headers: {\n")
	fmt.Fprintf(writer, "'Content-Type': 'application/json',\n")
	fmt.Fprintf(writer, "},\n")
	fmt.Fprintf(writer, "body: JSON.stringify({ buttonNumber: buttonNumber }),\n")
	fmt.Fprintf(writer, "})\n")
	fmt.Fprintf(writer, ".then(response => response.json())\n")
	fmt.Fprintf(writer, ".then(data => {\n")
	fmt.Fprintf(writer, "alert(data.message);\n")
	fmt.Fprintf(writer, "})\n")
	fmt.Fprintf(writer, ".catch(error => {\n")
	fmt.Fprintf(writer, "console.error('Error:', error);\n")
	fmt.Fprintf(writer, "});\n")
	fmt.Fprintf(writer, "</script>\n")
	fmt.Fprintf(writer, "</body>\n")
	fmt.Fprintf(writer, "</html>\n")

	// Flush the buffer to ensure data is written to the file
	writer.Flush()

	fmt.Printf("Data written to %s\n", fileName)
	return nil
}
