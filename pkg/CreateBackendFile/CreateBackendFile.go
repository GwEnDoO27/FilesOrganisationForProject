package createbackendfile

import (
	"FilesOrganisationForProject/configs"
	"fmt"
	"os"
	"path/filepath"
)

func CreateHandlerFile(ProjectPath string) {
	var handler = []string{
		"package handler\n", "import(", "\t\"net/http\"", "\t\"path/filepath\"", fmt.Sprintf("\t\"%s/utils\"", configs.Projectname), ")", "func HelloHandler(w http.ResponseWriter, r *http.Request){", "\tpath := filepath.Join(\"../Frontend\", \"index.html\")", "\tutils.Rendertemp(w,path)", "}",
	}

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
		"package routes\n", "//Contient la configuration des routes de l'application\n", "import(", fmt.Sprintf("\t\"%s/handler\"", configs.Projectname), "\t\"net/http\"", ")\n", "func SetupRoutes() {", "\thttp.HandleFunc(\"/\", handler.HelloHandler)", "}\n",
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
	"package utils\n", "//Contient des fonctions utilitaires, souvent utilisées dans plusieurs parties de l'application.",
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
	"package utils\n", "import(", "\t\"log\"", "\t\"database/sql\"", "\t\"fmt\"", "\t_ \"github.com/mattn/go-sqlite3\"", ")\n", "func CreateDatabase() {", "\tdb, err := sql.Open(\"sqlite3\",\"../Database/Db.sqlite\")", "\tif err != nil {",
	"\t\tfmt.Println(err)", "\t}\n", "\tdefer db.Close()", "\tr := `", "\tCREATE TABLE IF NOT EXISTS Auth (", "\t\tID INTEGER PRIMARY KEY AUTOINCREMENT,", "\t\tUsername VARCHAR(20) NOT NULL UNIQUE,",
	"\t\tPassword VARCHAR(50) NOT NULL,", "\t\tEmail VARCHAR(100) NOT NULL UNIQUE", "\t);", "\tCREATE TABLE IF NOT EXISTS User (", "\t\tID INTEGER PRIMARY KEY AUTOINCREMENT,", "\t\tUsername VARCHAR(20) NOT NULL,",
	"\t\tEmail VARCHAR(100) NOT NULL,", "\t\tName VARCHAR(20) NOT NULL,", "\t\tPassword VARCHAR(50) NOT NULL,", "\t\tInscription VARCHAR(10) NOT NULL,", "\t\tBirthday VARCHAR(10) NOT NULL,",
	"\t\tLastname VARCHAR(20) NOT NULL,", "\t\tFOREIGN KEY (Username, Email) REFERENCES Auth (Username, Email)", "\t);", "\t`\n", "\t_, err = db.Exec(r)", "\tif err != nil {", "\t\tlog.Println(\"CREATE ERROR\")", "\t\tfmt.Println(err)", "\t\t}", "\t}",
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
	var main = []string{
		"package main\n", "import(", "\t\"fmt\"", "\t\"log\"", "\t\"net/http\"", "\t\"path/filepath\"", fmt.Sprintf("\t\"%s/routes\"", configs.Projectname), fmt.Sprintf("\t\"%s/utils\"", configs.Projectname), ")\n", "func init() {", "\tutils.CreateDatabase()",
		"}\n", "func main() {", "\tfmt.Print(\"\\033[37m\")", "\tfmt.Println(\"Server Started: https://localhost:8080/\")",
		"\tfmt.Print(\"\\033[37m\")", "\tstaticPath := filepath.Join(\"..\", \"Frontend\", \"static\")", "\thttp.Handle(\"/static/\", http.StripPrefix(\"/static/\", http.FileServer(http.Dir(staticPath))))", "\troutes.SetupRoutes()",
		"\terr := http.ListenAndServe(\":8080\",nil)", "\tif err != nil {", "\t\tlog.Fatal(err)", "\t}", "}",
	}

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

func CreateTemplateFile(ProjectPath string) {

	Newpath := filepath.Join(ProjectPath, "Templates.go")
	file, err := os.Create(Newpath)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	for _, v := range Templates {
		fmt.Fprintln(file, v)
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

var Templates = []string{
	"package utils\n", "import(", "\t\"fmt\"", "\t\"net/http\"", "\t\"text/template\"", ")\n", "func Rendertemp(w http.ResponseWriter, path string) {", "\tt, err := template.ParseFiles(path)",
	"\tif err != nil {", "\t\tfmt.Println(fmt.Printf(\"%s Error Parsing Template\", err))", "\t}", "\terr = t.Execute(w, nil)", "\tif err != nil {", "\t\tfmt.Println(fmt.Printf(\"%s Error Execute Template\", err))",
	"\t}", "}",
}

func CreateGitignoreFile(ProjectPath string) {

	Newpath := filepath.Join(ProjectPath, ".gitignore")
	file, err := os.Create(Newpath)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	for _, v := range gitignore {
		fmt.Fprintln(file, v)
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

var gitignore = []string{"# Logs", "logs", "*.log", "npm-debug.log*", "yarn-debug.log*", "yarn-error.log*\n", "# Runtime data", "pids", "*.pid", "*.seed", "*.pid.lock\n",
	"# Dependency directories", "node_modules/", "jspm_packages/\n", "# TypeScript v1 declaration files", "typings/\n", "# Optional npm cache directory", ".npm\n", "# IDEs and editors", "# Eclipse", ".metadata", ".classpath", ".project",
	".c9/", "*.launch", ".settings/", "*.sublime-workspace\n", "# IntelliJ IDEA", "*.iml", "*.ipr", "*.iws", ".idea/", "out/\n", "# Visual Studio Code", ".vscode/", ".history/\n", "# Sublime Text", "*.sublime-workspace\n", "# Vim", "*.swp", "*.swo\n",
	"# MacOS", ".DS_Store", ".AppleDouble", ".LSOverride\n", "# Windows", "Thumbs.db", "ehthumbs.db", "Desktop.ini", "$RECYCLE.BIN/\n", "# Python", "__pycache__/", "*.py[cod]", "*$py.class\n", "# Jupyter Notebook", ".ipynb_checkpoints\n",
	"# Java", "*.class", "*.jar", "*.war", "*.ear", "*.class\n", "# C/C++", "*.obj", "*.o", "*.a", "*.lib", "*.so", "*.dylib", "*.exe", "*.dll", "*.class\n", "# Go", "*.exe", "*.test", "*.out\n", "# Rust", "/target/\n", "# Swift", ".build/",
	"*.xcodeproj/\n", "# Ruby", "*.gem", "*.rbc", ".rvm/", ".bundle/", "vendor/bundle/", "lib/bundler/man/\n", "# PHP", "vendor/\n", "# Composer", "/vendor/\n", "# Laravel", "/vendor/", ".phpunit.result.cache", ".env", "Homestead.yaml", "Homestead.json\n",
	"# CakePHP", "tmp/", "logs/", "bin/cake\n", "# Symfony", "/app/bootstrap.php.cache", "/app/cache/*", "/app/logs/*", "/app/sessions/*", "/vendor/*", "/web/bundles/", "/bin/*", "/build/\n", "# Django", "*.log", "*.pot", "*.pyc", "__pycache__/", "local_settings.py", "db.sqlite3", "media\n",
	"# Flask", "instance/", ".webassets-cache\n", "# CakePHP", "tmp/cache/persistent/*", "tmp/cache/models/*", "tmp/cache/views/*", "tmp/logs/*", "tmp/sessions/*", "tmp/test_logs/*", "tmp/tests/*\n", "# WordPress", "wp-content/uploads/", "wp-content/upgrade/\n",
	"# Magento", "/media/", "var/\n", "# Joomla", "/logs/", "tmp/\n", "# Drupal", "sites/all/modules/", "sites/all/libraries/", "sites/all/themes/", "sites/default/files/\n", "# TYPO3", "typo3temp/", "typo3conf/l10n/", "uploads/\n",
}
