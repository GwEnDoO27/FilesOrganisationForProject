package createscript

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateInitGoModFile(ProjectPath string) {
	Newpath := filepath.Join(ProjectPath, "Backend", "init_go_mod.sh")
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
	"#!/bin/bash", "set -e", "MODULE_NAME=\"$1\"", "TARGET_DIR=\"$2\"", "if [ ! -d \"$TARGET_DIR\" ]; then", "\techo \"Le répertoire cible n'existe pas: $TARGET_DIR\"", "\texit 1",
	"fi", "cd \"$TARGET_DIR\"", "echo \"Initialisation du module Go pour le module $MODULE_NAME...\"", "go mod init $MODULE_NAME",
	"echo \"Téléchargement des dépendances...\"", "go mod tidy", "echo \"Fichiers go.mod et go.sum créés et mis à jour avec succès!\"",
}
