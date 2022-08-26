package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	header = `<!DOCTYPE html>
	​<html>​
	​ ​  <head>​
	​ ​    <meta http-equiv="content-type" content="text/html; charset=utf-8">​
	​ ​    <title>Markdown Preview Tool</title>​
	​ ​  </head>​
	​ ​  <body>​
	​`
	footer = `
	​ ​  </body>​
	​</html>​
	`
)

func main() {
	// Parse command line flags
	filename := flag.String("file", "", "Markdown file to preview")
	flag.Parse()

	// If user did not provide input file, show usage
	if *filename == "" {
		flag.Usage()
		return
	}

	if err := run(*filename); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

func run(filename string) error {
	// Read all the data from the input file and check for erros
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	htmlData := parseContent(input)

	outName := fmt.Sprintf("%s.html", filepath.Base(filename))
	fmt.Println("Output file:", outName)

	return saveHTML(outName, htmlData)
}

func parseContent(input []byte) []byte {
	// Parse the markdown file through blackfriday and bluemonday​
	// to generate a valid and safe HTML file
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)

	// Create a buffer of bytes to write to file
	var buffer bytes.Buffer

	// Write html to bytes buffer
	buffer.WriteString(header)
	buffer.Write(body)
	buffer.WriteString(footer)

	return buffer.Bytes()
}

func saveHTML(filename string, data []byte) error {
	// Write the bytes to the file
	return ioutil.WriteFile(filename, data, 0644)
}
