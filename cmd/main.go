package main

import (
	"FilesOrganisationForProject/configs"
	createscript "FilesOrganisationForProject/pkg/CreateScript"
	creatingbackendfolders "FilesOrganisationForProject/pkg/CreatingBackendFolders"
	creatingfrontendfolders "FilesOrganisationForProject/pkg/CreatingFrontendFolders"
	prompt "FilesOrganisationForProject/pkg/Prompt"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	var absolutepath = "/Users/gwendal/Desktop/"

	configs.Projectname = prompt.PromptProjectName("Enter the name of the project :")
	fmt.Printf("The name of your project %s\n", configs.Projectname)
	ProjectPath := creatingbackendfolders.CreateMainFolder(configs.Projectname, absolutepath)
	creatingbackendfolders.CreateBackFolder(ProjectPath)
	creatingfrontendfolders.CreateFrontendFolder(ProjectPath)
	creatingbackendfolders.CreateDatabaseFolder(ProjectPath)
	createscript.CreateInitGoModFile(ProjectPath)
	fmt.Println("All directories created")

	scriptPath := filepath.Join(ProjectPath, "init_go_mod.sh")

	if err := os.Chmod(scriptPath, 0755); err != nil {
		fmt.Printf("Error making script executable: %s\n", err)
		return
	}

	cmd := exec.Command("bash", scriptPath, configs.Projectname)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running script: %s\n", err)
		return
	}

	// Afficher la sortie du script
	fmt.Printf("Script output:\n%s\n", string(output))
}
