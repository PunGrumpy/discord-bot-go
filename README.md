# Discord Bot in Go

This project is a Discord bot written in Go. It uses the discordgo package and includes a variety of features, such as command handling, environment variable configuration, Docker deployment, and a server for status monitoring and alerts.

## Features

- Command handling: The bot is capable of handling various commands and responds accordingly.
- Environment variable configuration: Both `.env` file and system environment variables are supported for configuration.
- Docker deployment: The bot is containerized using Docker, which makes it easy to deploy anywhere.
- Server for status monitoring and alerts: A server written in Go monitors the status of the bot and sends alerts when the bot is down or encounters issues.

## Setup and Running

1. Clone this repository.

```bash
git clone https://github.com/PunGrumpy/discord-bot-go.git
cd discord-bot-go
```

2. Create a `.env` file in the root of the project and fill in your Discord bot token and guild id:

```bash
BOT_TOKEN=your_bot_token
GUILD_ID=your_guild_id
```

3. Build the Docker image and start the services:

```bash
docker-compose up -d
# or
docker compose up -d
```

The bot should now be running in your Discord server, and the status monitoring server should be available at `http://localhost:8080`.

# License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
