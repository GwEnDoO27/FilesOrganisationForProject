package createbackendfile

import (
	"FilesOrganisationForProject/configs"
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

	var routes = []string{
		"package models\n", "//Contient la configuration des routes de l'application\n", "import(", fmt.Sprintf("\t\"%s/handler\"", configs.Projectname), "\t\"net/http\"", ")\n", "func SetupRoutes() {", "\thttp.HandleFunc(\"/\", handler.helloHandler)", "}",
	}

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

func CreateUtilsFile(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "utils.go")
	file, err := os.Create(Newpath)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	for _, v := range utils {
		fmt.Fprintln(file, v)
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

var utils = []string{
	"package utils\n", "Contient des fonctions utilitaires, souvent utilisées dans plusieurs parties de l'application.",
}

func CreateDatabaseFile(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "Db.go")
	file, err := os.Create(Newpath)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	for _, v := range Db {
		fmt.Fprintln(file, v)
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

var Db = []string{
	"package utils\n", "import(", "\t\"log\"", "\t\"database/sql\"", "\t\"fmt\"", "\t_ \"github.com/mattn/go-sqlite3\"", ")\n", "func CreateDatabase() {", "\tdb, err := sql.Open(\"sqlite3\",\"./backend/Configs/db.sqlite\")", "\tif err != nil {",
	"\t\tfmt.Println(err)", "\t}\n", "\tdefer db.Close()", "\tr := `", "\tCREATE TABLE IF NOT EXISTS User (", "\t\tID INTEGER PRIMARY KEY AUTOINCREMENT", "\t\tUsername VARCAHR(20) NOT NULL REFERENCEs \"Auth\"(\"username\"),",
	"\t\tEmail VARCHAR(100) NOT NULL REFERENCES \"Auth\"(\"email\"),", "\t\tName VARCHAR(20) NOT NULL,", "\t\tPassword VARCHAR(50) NOT NULL,", "\t\tInscription VARCHAR(10) NOT NULL,", "\t\tBirthday VARCHAR(10) NOT NULL,",
	"\t\tLastname VARCHAR(20) NOT NULL", "\t);", "\tCREATE TABLE IF NOT EXISTS Auth (", "\t\tID INTEGER PRIMARY KEY,", "\t\tID INTEGER PRIMARY KEY,", "\t\tUsername  VARCAHR(20) NOT NULL,", "\t\tPassword VARCAHR(50) NOT NULL,",
	"\t\temail VARCAHR(100) NOT NULL", "\t);", "\t`\n", "\t_, err = db.Exec(r)", "\tif err != nil {", "\t\tlog.Println(\"CREATE ERROR\")", "\t\tfmt.Println(err)", "\t\t}", "\t}",
}

func CreateDb(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "Db.sqlite")
	file, err := os.Create(Newpath)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
}

func CreateMainFile(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "main.go")
	file, err := os.Create(Newpath)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	for _, v := range main {
		fmt.Fprintln(file, v)
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

var main = []string{
	"package main\n", "import(", "\t\"fmt\"", "\t\"log\"", "\t\"net/http\"", fmt.Sprintf("\t\"%s/routes\"", configs.Projectname), fmt.Sprintf("\t\"%s/utils\"", configs.Projectname), ")\n", "func init() {", "\tutils.CreateDatabase()",
	"}\n", "func main() {", "\tfmt.Print(\"\\033[37m\")", "\tfmt.Println(\"Server Started: https://localhost:8080/\")",
	"\tfmt.Print(\"\\033[37m\")", "\thttp.Handle(\"/static/\", http.StripPrefix(\"/static/\", http.FileServer(http.Dir(\"Frontend/static/\"))))", "\troutes.SetupRoutes()",
	"\terr := http.ListenAndServe(\":8080\",nil)", "\tif err != nil {", "\t\tlog.Fatal(err)", "\t}", "}",
}
