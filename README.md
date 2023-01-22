# go_discbot
A simple POC discord bot written is Go using the discordgo library.

This code is based on this fantastic tutorial: https://kavitmht.medium.com/discord-bot-in-go-74eabf4090a7

Please note that since the tutorial was written, the discord UI, and the bot scopes have changed.

## Setup

### Create a discord bot

1. Go to https://discord.com/developers/applications
2. Click on "New Application"
3. Give it a name and click "Create"
4. Click on "Bot" on the left menu
5. Click on "Add Bot"
6. Make sure to copy the token, you will need it later (you can regenerate it later if you need).

### Create a discord server

1. Go to the discord UI and click on the "+" icon on the left menu
2. Click on "Create my own" -> "For me and my friends"
3. Give it a name and click "Create"

### Invite the bot to the server

1. On the application page for your bot, click on "OAuth2" on the left menu
2. Check the "bot" scope
3. Check the following permissions:
    - Read Messages/View Channels
    - Send Messages
4. Copy the URL and open it in a browser
5. Select the server you created in the previous step

### Run the bot

1. Clone the repo
2. Run 'go get "github.com/bwmarrin/discordgo"'
3. Switch to g_discbot/cmd and edit the config.json file to replace the token with the one you copied earlier
4. Run 'go run main.go'

## Usage

The bot will respond to direct messages and to messages starting with the prefix "!".

The following commands are available:
- !ping: The bot will respond with "pong"
- !tell: The bot will say "I'm a bot"

