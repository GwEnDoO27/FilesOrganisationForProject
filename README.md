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
cd FilesOrganisationForProject
cd cmd
go run .
```

2. Alternative :
- Telechargz l'un des executable selon votre systeme