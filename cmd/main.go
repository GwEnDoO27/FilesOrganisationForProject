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
	//var absolutepath = "/Users/gwendal/Desktop/"

	configs.Projectname = prompt.PromptProjectName("Enter the name of the project :")
	fmt.Printf("The name of your project %s\n", configs.Projectname)
	configs.ProjectPath = prompt.PromptProjecPath("Enter the path where your project will be located in this format(/somefolders/anotherfolder) :")
	ProjectPath := creatingbackendfolders.CreateMainFolder(configs.Projectname, configs.ProjectPath)
	creatingbackendfolders.CreateBackFolder(ProjectPath)
	creatingfrontendfolders.CreateFrontendFolder(ProjectPath)
	creatingbackendfolders.CreateDatabaseFolder(ProjectPath)
	createscript.CreateInitGoModFile(ProjectPath)
	fmt.Println("All directories created")

	scriptPath := filepath.Join(ProjectPath, "Backend", "init_go_mod.sh")

	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		fmt.Printf("Le script n'existe pas à l'emplacement spécifié: %s\n", scriptPath)
		return
	}

	if err := os.Chmod(scriptPath, 0755); err != nil {
		fmt.Printf("Error making script executable: %s\n", err)
		return
	}

	targetDir := filepath.Join(ProjectPath, "Backend")

	cmd := exec.Command("bash", scriptPath, configs.Projectname, targetDir)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running script: %s\n", err)
		return
	}

	// Afficher la sortie du script
	fmt.Printf("Script output:\n%s\n", string(output))

	// Supprimer le script après exécution
	if err := os.Remove(scriptPath); err != nil {
		fmt.Printf("Error removing script: %s\n", err)
		return
	}
	fmt.Println("Script removed successfully.")
}
