#!/bin/bash
VERSION="v1.0.0"
URL="https://github.com/votre-utilisateur/votre-repo/releases/download/$VERSION/monprogramme-linux"

# Téléchargement du binaire
curl -L $URL -o /usr/local/bin/monprogramme

# Rendre le binaire exécutable
chmod +x /usr/local/bin/monprogramme

# Ajouter un alias dans le fichier .bashrc ou .zshrc
if [ -n "$ZSH_VERSION" ]; then
    SHELL_PROFILE="$HOME/.zshrc"
else
    SHELL_PROFILE="$HOME/.bashrc"
fi

echo "alias monprog='/usr/local/bin/monprogramme'" >> $SHELL_PROFILE
source $SHELL_PROFILE

echo "Installation terminée. Vous pouvez utiliser le programme en tapant 'monprog'."
