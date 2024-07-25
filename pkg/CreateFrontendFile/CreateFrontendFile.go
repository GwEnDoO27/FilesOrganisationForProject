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
	"<!DOCTYPE html>", "<html lang=\"en\">\n", "<head>", "    <meta charset=\"UTF-8\">", "    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">","    <link rel=\"stylesheet\" href=\"./static/style.css\"/>", "    <title>Welcome to MyProject</title>", "</head>\n",
	"<body>", "    <h1>Welcome to MyProject</h1>", "    <p>This is a simple web page served by a Go web server.</p>", "</body>\n", "</html>\n",
}
