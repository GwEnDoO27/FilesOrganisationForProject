package main

import (
	creatingfolders "FilesOrganisationForProject/pkg/CreatingFolders"
	prompt "FilesOrganisationForProject/pkg/Prompt"
	"fmt"
)

func main() {
	var absolutepath = "/Users/gwendal/Desktop/"

	Projectname := prompt.PromptProjectName("Enter the name of the project :")
	fmt.Printf("The name of your project %s\n", Projectname)
	ProjectPath := creatingfolders.CreateMainFolder(Projectname, absolutepath)
	creatingfolders.CreatebackendFolder(ProjectPath)
}
