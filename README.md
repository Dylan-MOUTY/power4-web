# power4-web

🟡 Puissance 4 – Version Web (Go)

Un jeu Puissance 4 entièrement développé en Go, jouable depuis un
navigateur web grâce à un mini-serveur intégré.
Le projet inclut une logique de jeu complète, un système de détection de
victoire et une interface web simple mais efficace.

🚀 Fonctionnement du projet

Le jeu fonctionne comme une application web :
- un serveur HTTP en Go lance automatiquement votre navigateur
- l’interface web affiche la grille, les boutons, les assets
- toutes les actions (jouer un coup, reset, détecter gagnant, etc.)
utilisent des routes HTTP

🛠️ Technologies utilisées

-   Go (Golang)
-   HTML / CSS
-   Templates Go
-   Assets media (images, vidéos, polices)


▶️ Lancer le jeu

1.  Cloner le projet : git clone cd power4-web

2.  Lancer le serveur : go run main.go

3.  Ouvrir dans le navigateur : http://localhost:7070/

🎮 Fonctionnalités principales

-   Logique complète du Puissance 4
-   Détection de victoire (horizontal, vertical, diagonales)
-   Détection d’égalité
-   Interface web dynamique
-   Serveur web intégré