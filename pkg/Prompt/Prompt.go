package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func PromptProjecPath(str string) string {
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
