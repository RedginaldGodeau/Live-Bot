---

# 📡 LIVE BOT
### *Un bot Discord pour partager des vidéos et des images en temps réel*

---

## ⚡ Description
**LIVE BOT** est un bot Discord qui permet aux utilisateurs d'envoyer des vidéos, des images et des messages directement sur une page web dédiée, en temps réel.

Le backend est développé en **Golang** avec **Ent** et **Echo**, tandis que le frontend utilise **HTML**, **JavaScript**, et **TailwindCSS**.

---

## 🛠️ Fonctionnalités
- 🔄 Partage instantané de médias (images/vidéos) sur une page web.
- 💬 Personnalisation des messages envoyés.
- 📡 Commandes simples pour interagir avec le bot.

---

## 📋 Commandes Discord
```
$ping
➡️ Vérifie si le bot est connecté.

$livebot [message] [Fichier attaché]
➡️ Envoie le message et le média sur la page web.
```  

---

## 🏗️ Installation

1. **Cloner le dépôt**
   ```bash
   git clone <lien_du_repo>
   cd live-bot
   ```

2. **Compiler le projet**
   ```bash
   make build
   ```  

---

## 🚀 Lancer le Bot
1. **Remplire le .env**
```env
DISCORD_APP_ID= # ID de votre application discord (BOT)
DISCORD_PUBLIC_KEY= # la clef public de votre application discord (BOT)
FRONT_END_URL= # URL du front-end (localhost:8080 par defaut)
```

2. **Lancer le bot**
```bash
make run
```  

---

## 🧩 Prochaines Améliorations
- 📦 Ajouter un **Docker** pour la production (et retirer Air).
- ⚙️ Créer un **fichier de configuration** pour :
    - Les délais d'affichage
    - Les permissions
    - Les types de fichiers autorisés
- 💾 Ajouter des commandes pour **sauvegarder les médias**.

---

## 👤 Auteur
**Redginald Godeau**

---
