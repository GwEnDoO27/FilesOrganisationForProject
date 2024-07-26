#!/bin/bash

URL="https://raw.githubusercontent.com/GwEnDoO27/FilesOrganisationForProject/main/FilesOrganisationForProject-macos"

# Téléchargement du binaire
echo "Téléchargement du binaire depuis $URL"
curl -L $URL -o /usr/local/bin/FilesOrganisationForProject

# Vérification si le téléchargement a réussi
if [ $? -ne 0 ]; then
    echo "Échec du téléchargement du binaire."
    exit 1
fi

# Rendre le binaire exécutable
sudo chmod +x /usr/local/bin/FilesOrganisationForProject

# Vérifier le shell et définir le fichier de profil approprié
if [ "$SHELL" = "/bin/zsh" ] || [ "$SHELL" = "/usr/bin/zsh" ]; then
    SHELL_PROFILE="$HOME/.zshrc"
elif [ "$SHELL" = "/bin/bash" ] || [ "$SHELL" = "/usr/bin/bash" ]; then
    SHELL_PROFILE="$HOME/.bashrc"
else
    SHELL_PROFILE="$HOME/.profile"
fi
# Ajouter un alias dans le fichier de profil
if grep -Fxq "alias Newproject='/usr/local/bin/FilesOrganisationForProject'" $SHELL_PROFILE; then
    echo "Alias 'Newproject' déjà présent dans $SHELL_PROFILE"
else
    echo "alias Newproject='/usr/local/bin/FilesOrganisationForProject'" >> $SHELL_PROFILE
    echo "Alias 'Newproject' ajouté dans $SHELL_PROFILE"
fi

# Recharger le fichier de profil
source $SHELL_PROFILE

echo "Installation terminée. Vous pouvez utiliser le programme en tapant 'Newproject'."

