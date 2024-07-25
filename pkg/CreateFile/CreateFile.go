package createfile

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateHandlerFile(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "handler.go")
	file, err := os.Create(Newpath)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	for _, v := range handler {
		fmt.Fprintln(file, v)
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

var handler = []string{
	"package handler\n", "import(", "\t\"fmt\"", "\t\"net/http\"", ")", "func helloHandler(w http.ResponseWriter, r *http.Request){", "\tfmt.Fprintf(w, \"Hello, World!\")", "}",
}

func CreateMiddlewaresFile(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "middleware.go")
	file, err := os.Create(Newpath)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	for _, v := range middleware {
		fmt.Fprintln(file, v)
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

var middleware = []string{
	"package middleware\n", "//Contient les middlewares, qui sont des fonctions exécutées pendant le traitement des requêtes",
}

func CreateModelsFile(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "models.go")
	file, err := os.Create(Newpath)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	for _, v := range models {
		fmt.Fprintln(file, v)
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

var models = []string{
	"package models\n", "//Contient les modèles de données, qui définissent les structures des données utilisées dans l'application\n", " type Auth struct {", "\tUsername string", "\tPassword string", "}",
}

func CreateRoutesFile(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "routes.go")
	file, err := os.Create(Newpath)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	for _, v := range routes {
		fmt.Fprintln(file, v)
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

var routes = []string{
	"package models\n", "//Contient la configuration des routes de l'application\n", "import(", "\t\"\"NameofTheProject/handler\"", "\t\"net/http\"", ")\n", "func SetupRoutes() {", "\thttp.HandleFunc(\"/\", handlers.helloHandler)", "}",
}