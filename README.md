# FilesOrganisationForProject

Ce projet en Go génère une architecture de dossier pour une application web. Il facilite la configuration initiale des dossiers et des fichiers nécessaires pour démarrer rapidement un projet web.

## Fonctionnalités

- Génération automatique des dossiers de base pour une application web.
- Création des fichiers essentiels pour une application Go.
- Creation du go.mod automatiquement
- Le sever peut direcetemetnt être lancé

## Architecture des dossiers

Voici un exemple d'architecture générée par cet outil :

Your-app
├── Backend
│   ├── go.mod
│   ├── go.sum
│   ├── handler
│   │   └── handler.go
│   ├── main.go
│   ├── middleware
│   │   └── middleware.go
│   ├── models
│   │   └── models.go
│   ├── routes
│   │   └── routes.go
│   └── utils
│       ├── Db.go
│       ├── Templates.go
│       └── utils.go
├── Database
│   └── Db.sqlite
└── Frontend
    ├── index.html
    └── static
        ├── asset
        │   ├── font
        │   └── image
        ├── css
        │   └── style.css
        └── js
            └── app.js



## Installation

1. Clonez le dépôt :

```sh
git clone https://github.com/GwEnDoO27/FilesOrganisationForProject.git
```

2. Alternative :
- Telechargz l'un des executable selon votre systeme dans le dossier Apps

3. Autres Alternatives : 
- Telechargz l'un des executable selon votre systeme dans le dossier Apps
- Deplacer le dans le dossier que vous souhaitiez et retenez le chemin d'acces


## Utilisation 

Methode 1 : Exécuter via Go
- Naviguez vers le répertoire du projet et exécutez le programme: 
```sh
cd FilesOrganisationForProject
cd cmd
go run .
```
Methode 2 : Utiliser l'exécutable
- Dans le terminale tapez 
```sh
sudo bash FilesOrganisationForProject-linux
```

Methode 3 : Ajouter un alias pour un accès rapide
1. Dans le terminale tapez 
```sh
nano ~/.zshrc
```
2. Ajoutez l'alias suivant à la fin du fichier :
```sh
alias Nom_d_appel='/votre chemin d'acces/FilesOrganisationForProject'
```
3. Sauvegardez et fermez l'éditeur (^X puis y).
4. Maintenant, vous pouvez appeler et lancer le fichier n'importe où en utilisant l'alias
```sh
Nom_d_appel
```

## Licence

Ce projet est sous licence MIT. Voir le fichier LICENSE pour plus de détails.