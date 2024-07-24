package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	Projectname := PromptProjectName("Enter the name of the project :")
	fmt.Printf("The name of your project %s\n", Projectname)
	path := PromptPath("Enter the path :")
	fmt.Printf("It's your path %s\n", path)

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
