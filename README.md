# power4-web

# Puissance 4 – Application Web en Go

Ce projet est une application web qui recrée le jeu classique **Puissance 4**.  
Deux joueurs peuvent jouer l’un après l’autre directement depuis un navigateur, sur un plateau affiché en HTML/CSS, avec une logique de jeu gérée côté serveur en **Go**.

---

## 🎯 Objectif du projet

L’objectif de ce projet est :

- de **reproduire le jeu Puissance 4** en version numérique,
- de **gérer automatiquement les règles du jeu** (tours, placement des pions, victoire),
- de **proposer une interface web simple** et accessible via un navigateur,
- de structurer une petite application **client–serveur** en Go.

---

## 🧩 Fonctionnalités principales

- Jeu de **Puissance 4 à deux joueurs**.
- **Interface web** : le joueur interagit via des pages HTML.
- **Choix de la difficulté** avec plusieurs tailles de grilles :
  - facile
  - normal
  - difficile
- Gestion de la **chute des pions** dans les colonnes (effet de gravité).
- Détection de la **victoire** :
  - alignement horizontal,
  - alignement vertical,
  - alignement diagonal.
- Ouverture automatique du **navigateur** au lancement du serveur (selon le système).

---

## 🛠 Technologies utilisées

- **Langage** : Go (Golang)
- **Serveur web** : `net/http`
- **Templates** : HTML (rendu côté serveur)
- **Style / interface** : HTML / CSS
- **Exécution** : application locale (serveur HTTP + navigateur)

---

## 📁 Organisation du projet

L’arborescence générale du projet est la suivante (simplifiée) :

```text
.
├── main.go            # Point d'entrée : serveur, routes, handlers
├── go.mod             # Module Go
├── game/
│   └── game.go        # Logique du jeu : structure Game, grille, victoire
├── templates/
│   ├── Page1.html     # Page de choix de la difficulté
│   └── index.html     # Page principale du plateau de jeu
└── assets/
    ├── images/...     # Images, éventuel logo, fond
    └── ...            # Autres ressources (polices, etc.)
