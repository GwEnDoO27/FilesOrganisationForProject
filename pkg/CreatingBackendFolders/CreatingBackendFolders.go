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
	return Newpath
}

func CreateBackFolder(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "Backend")
	err := os.MkdirAll(Newpath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	createbackendfile.CreateMainFile(Newpath)
	createbackendfile.CreateGitignoreFile(Newpath)
	CreateHandlerFolder(Newpath)
	CreateMiddlewareFolder(Newpath)
	CreateModelsFolder(Newpath)
	CreateRoutesFolder(Newpath)
	CreateUtilsFolder(Newpath)
}

func CreateDatabaseFolder(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "Database")
	err := os.MkdirAll(Newpath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	createbackendfile.CreateDb(Newpath)
}

func CreateHandlerFolder(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "handler")
	err := os.MkdirAll(Newpath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	createbackendfile.CreateHandlerFile(Newpath)
}

func CreateMiddlewareFolder(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "middleware")
	err := os.MkdirAll(Newpath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	createbackendfile.CreateMiddlewaresFile(Newpath)
}

func CreateModelsFolder(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "models")
	err := os.MkdirAll(Newpath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	createbackendfile.CreateModelsFile(Newpath)
}

func CreateRoutesFolder(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "routes")
	err := os.MkdirAll(Newpath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	createbackendfile.CreateRoutesFile(Newpath)
}
func CreateUtilsFolder(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "utils")
	err := os.MkdirAll(Newpath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	createbackendfile.CreateUtilsFile(Newpath)
	createbackendfile.CreateDatabaseFile(Newpath)
	createbackendfile.CreateTemplateFile(Newpath)
}
