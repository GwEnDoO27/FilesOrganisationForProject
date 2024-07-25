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
	fmt.Println("Frontend directory created successfully!")
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
	fmt.Println("static directory created successfully!")
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
	fmt.Println("css directory created successfully!")

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
	fmt.Println("asset directory created successfully!")
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
	fmt.Println("image directory created successfully!")
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
