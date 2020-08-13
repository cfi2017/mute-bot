# Mute Bot

This discord bot mutes any voice user on a server wide basis and unmutes them depending on the mute toggle state.

The bot looks for either a json, yaml or toml config file called mute-bot in the following locations:
- /etc/discord/
- $HOME/.discord/
- ./config/
- ./

For example, this is a valid config file path:
- /etc/discord/mute-bot.toml

Check the config/ directory for an example config file.

The bot also accepts environment variables in the form `DISCORD_<FLAG>` where flag is a valid config line, upper case.
