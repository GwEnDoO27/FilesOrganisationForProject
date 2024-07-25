package creatingbackendfolders

import (
	createbackendfile "FilesOrganisationForProject/pkg/CreateBackendFile"
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

func CreateBackFolder(ProjectPath string) {
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

func CreateFrontendFolder(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "Frontend")
	err := os.MkdirAll(Newpath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Frontend directory created successfully!")
	CreateCssFolder(Newpath)
	CreateJsFolder(Newpath)
	CreateAssetsFolder(Newpath)
}

func CreateDatabaseFolder(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "Database")
	err := os.MkdirAll(Newpath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Frontend directory created successfully!")
	createbackendfile.CreateDb(Newpath)
}

func CreateHandlerFolder(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "handler")
	err := os.MkdirAll(Newpath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("handler directory created successfully!")

	createbackendfile.CreateHandlerFile(Newpath)
}

func CreateMiddlewareFolder(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "middleware")
	err := os.MkdirAll(Newpath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("middleware directory created successfully!")
	createbackendfile.CreateMiddlewaresFile(Newpath)
}

func CreateModelsFolder(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "models")
	err := os.MkdirAll(Newpath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("models directory created successfully!")
	createbackendfile.CreateModelsFile(Newpath)
}

func CreateRoutesFolder(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "routes")
	err := os.MkdirAll(Newpath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("routes directory created successfully!")
	createbackendfile.CreateRoutesFile(Newpath)
}
func CreateUtilsFolder(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "utils")
	err := os.MkdirAll(Newpath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("utils directory created successfully!")
	createbackendfile.CreateUtilsFile(Newpath)
	createbackendfile.CreateDatabaseFile(Newpath)
}

func CreateCssFolder(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "css")
	err := os.MkdirAll(Newpath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("css directory created successfully!")
}

func CreateJsFolder(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "js")
	err := os.MkdirAll(Newpath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("js directory created successfully!")
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