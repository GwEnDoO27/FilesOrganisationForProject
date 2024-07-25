package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var absolutepath = "/Users/gwendal/Desktop/"

	Projectname := PromptProjectName("Enter the name of the project :")
	fmt.Printf("The name of your project %s\n", Projectname)
	ProjectPath := CreateMainFolder(Projectname, absolutepath)

}



func PromptPath(str string) string {
	var s string

	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, str+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

func PromptProjectName(str string) string {
	var s string

	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, str+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)

}

func CreateMainFolder(Projectname, path string) string {
	Newpath := filepath.Join(path, Projectname)
	err := os.MkdirAll(Newpath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println("Directory created successfully!")
	return Newpath
}
