package pkg

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateMainFolder(Projectname, path string) string {
	Newpath := filepath.Join(path, Projectname)
	err := os.MkdirAll(Newpath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(fmt.Printf("%s created successfully!", Projectname))
	return Newpath
}

func CreatebackendFolder(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "Backend")
	err := os.MkdirAll(Newpath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Backend directory created successfully!")
	CreateHandlerFolder(Newpath)
	CreateMiddlewareFolder(Newpath)
	CreateModelsFolder(Newpath)
	CreateRoutesFolder(Newpath)
	CreateUtilsFolder(Newpath)
}

func CreateHandlerFolder(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "handler")
	err := os.MkdirAll(Newpath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("handler directory created successfully!")
}

func CreateMiddlewareFolder(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "middleware")
	err := os.MkdirAll(Newpath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("middleware directory created successfully!")
}

func CreateModelsFolder(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "models")
	err := os.MkdirAll(Newpath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("models directory created successfully!")
}

func CreateRoutesFolder(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "routes")
	err := os.MkdirAll(Newpath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("routes directory created successfully!")
}
func CreateUtilsFolder(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "utils")
	err := os.MkdirAll(Newpath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("utils directory created successfully!")
}
