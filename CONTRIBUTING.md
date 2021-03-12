
# Build and Run

### Requirements
* `go.15+` - [install here](https://golang.org/doc/install)
* `git clone https://github.com/BigPapaChas/ow-bot.git`

### Steps

```bash
# 1. Run `make build` from within project directory.
$ make build
go build -o ./build/runserver ./cmd/bot
# 2. You may set the bot token as an environment variable or provide it inline.
$ export BOT_TOKEN=<value>
# 3. Run the newly created executable.
$ ./build/runserver --token $BOT_TOKEN --prefix !
INFO[0000] Bot is now running.  Press CTRL-C to exit.    version=v0.1.0
# 4. CTRL-C to exit
Gracefully exiting...                                    version=v0.1.0
```

# Debugging

In [VSCode](https://code.visualstudio.com/), add the following launch configuration then select "runserver" in the debug menu. Click ► Play in the debug menu (or F5) to start debugging.

```json
{
    "name": "runserver",
    "type": "go",
    "request": "launch",
    "mode": "debug",
    "program": "${workspaceFolder}${pathSeparator}cmd${pathSeparator}bot${pathSeparator}main.go",
    "args": ["-t", "${env:BOT_TOKEN}"]
}
```

# Resources

## Documentation

- [Discord Developer Portal][dc-dev] - navigate left sidebar for API info and docs
- [JavaScript SDK Docs][discordjs] - extremely clean and discriptive docs to supplement in any knowledge gaps left by the dev portal.
- [How to add your bot to a discord server server][bot-token]
    - Note: The [Developer Portal][dc-dev] did not specify how to build the Auth URL; but the [JavaScript SDK Docs][discordjs] had a good explanation.
    - (detailed instructions needed, i.e. pics)

## Go packages

- [discordgo][discordgo] - provides go bindings for the Discord API
- [dgrouter][dgrouter] - command router for discordgo
- [golangci-lint][gocilint] - linter used in remote builds, run locally before pushing

## Other tools

- [LeoV's Embed Visualizer](https://leovoel.github.io/embed-visualizer/) - Preview how some given embed (JSON payload) might appear on Discord, but it also serves as a reference for some limits. This visualization tool does not render all types of embeds. Check out the About page for caveats or consult the [Embed docs](https://discord.com/developers/docs/resources/channel#embed-object).

<!-- links -->

[discordjs]: <https://discordjs.guide/preparations/adding-your-bot-to-servers.html#bot-invite-links> "Javascript SDK for Discord - great docs"

[bot-token]: <https://github.com/andersfylling/disgord/wiki/Get-bot-token-and-add-it-to-a-server> "Get bot token and add to server"

[dc-dev]: <https://discord.com/developers> "Discord developer portal"

[discordgo ]: <github.com/bwmarrin/discordgo> "Go bindings for Discord"
[dgrouter]: <github.com/Necroforger/dgrouter> "router for discordgo"
[gocilint]: <https://golangci-lint.run/usage/quick-start/> "linting tool to use locally & in remote builds"