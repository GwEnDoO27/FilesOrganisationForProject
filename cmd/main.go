package main

import (
	"FilesOrganisationForProject/configs"
	creatingbackendfolders "FilesOrganisationForProject/pkg/CreatingBackendFolders"
	creatingfrontendfolders "FilesOrganisationForProject/pkg/CreatingFrontendFolders"
	prompt "FilesOrganisationForProject/pkg/Prompt"
	"fmt"
)

func main() {
	var absolutepath = "/Users/gwendal/Desktop/"

	configs.Projectname = prompt.PromptProjectName("Enter the name of the project :")
	fmt.Printf("The name of your project %s\n", configs.Projectname)
	ProjectPath := creatingbackendfolders.CreateMainFolder(configs.Projectname, absolutepath)
	creatingbackendfolders.CreateBackFolder(ProjectPath)
	creatingfrontendfolders.CreateFrontendFolder(ProjectPath)
	creatingbackendfolders.CreateDatabaseFolder(ProjectPath)
}
