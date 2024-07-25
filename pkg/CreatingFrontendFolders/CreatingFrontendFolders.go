package creatingfrontendfolders

import (
	createfrontendfile "FilesOrganisationForProject/pkg/CreateFrontendFile"
	"fmt"
	"os"
	"path/filepath"
)

func CreateFrontendFolder(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "Frontend")
	err := os.MkdirAll(Newpath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	createfrontendfile.CreateHTMLFile(Newpath)
	CreateStaticFolder(Newpath)
}

func CreateStaticFolder(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "static")
	err := os.MkdirAll(Newpath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	CreateCssFolder(Newpath)
	CreateJsFolder(Newpath)
	CreateAssetsFolder(Newpath)
}

func CreateCssFolder(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "css")
	err := os.MkdirAll(Newpath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	createfrontendfile.CreateCssFile(Newpath)
}

func CreateJsFolder(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "js")
	err := os.MkdirAll(Newpath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("js directory created successfully!")
	createfrontendfile.CreateJSFile(Newpath)
}

func CreateAssetsFolder(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "asset")
	err := os.MkdirAll(Newpath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	CreateImageFolder(Newpath)
	CreateFontFolder(Newpath)
}

func CreateImageFolder(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "image")
	err := os.MkdirAll(Newpath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CreateFontFolder(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "font")
	err := os.MkdirAll(Newpath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("font directory created successfully!")
}
