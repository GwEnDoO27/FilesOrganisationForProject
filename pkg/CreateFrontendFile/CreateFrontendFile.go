package createfrontendfile

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateHTMLFile(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "index.html")
	file, err := os.Create(Newpath)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	for _, v := range html {
		fmt.Fprintln(file, v)
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

var html = []string{
	"<!DOCTYPE html>", "<html lang=\"en\">\n", "<head>", "    <meta charset=\"UTF-8\">", "    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">", "    <link rel=\"stylesheet\" href=\"./static/style.css\"/>", "    <title>Welcome to MyProject</title>", "</head>\n",
	"<body>", "    <h1>Welcome to MyProject</h1>", "    <p>This is a simple web page served by a Go web server.</p>", "</body>\n", "</html>\n",
}

func CreateCssFile(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "style.css")
	file, err := os.Create(Newpath)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	for _, v := range css {
		fmt.Fprintln(file, v)
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

var css = []string{
	"/* styles.css */\n", "body {", "    font-family: Arial, sans-serif;", "    margin: 0;", "    padding: 0;", "    background-color: #f4f4f4;", "}\n", "h1 {", "    color: #333;",
	"    text-align: center;", "    margin-top: 20px;", "}\n", "p {", "    color: #666;", "    text-align: center;", "    font-size: 1.2em;", "    margin: 20px;", "}\n",
	".container {\n", "    max-width: 800px;", "    margin: 0 auto;", "    padding: 20px;", "    background: #fff;", "    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);", "    border-radius: 8px;",
	"}\n",
}

func CreateJSFile(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "app.js")
	file, err := os.Create(Newpath)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	for _, v := range js {
		fmt.Fprintln(file, v)
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

var js = []string{
	"Code JavaScript principal de l'application",
}
