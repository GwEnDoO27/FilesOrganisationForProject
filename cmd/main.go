package main

import (
	"FilesOrganisationForProject/configs"
	creatingfolders "FilesOrganisationForProject/pkg/CreatingFolders"
	prompt "FilesOrganisationForProject/pkg/Prompt"
	"fmt"
)

func main() {
	var absolutepath = "/Users/gwendal/Desktop/"

	configs.Projectname = prompt.PromptProjectName("Enter the name of the project :")
	fmt.Printf("The name of your project %s\n", configs.Projectname)
	ProjectPath := creatingfolders.CreateMainFolder(configs.Projectname, absolutepath)
	creatingfolders.CreateBackendFolder(ProjectPath)
	creatingfolders.CreateFrontendFolder(ProjectPath)
}
