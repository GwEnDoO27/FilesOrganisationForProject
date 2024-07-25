package main

import (
	"FilesOrganisationForProject/pkg"
	"fmt"
)

func main() {
	var absolutepath = "/Users/gwendal/Desktop/"

	Projectname := pkg.PromptProjectName("Enter the name of the project :")
	fmt.Printf("The name of your project %s\n", Projectname)
	ProjectPath := pkg.CreateMainFolder(Projectname, absolutepath)
	pkg.CreatebackendFolder(ProjectPath)
}
