
# Building from source

### Requirements
* go.15+

### Steps
1. Run `git clone https://github.com/BigPapaChas/ow-bot.git`
2. Navigate to the `/cmd/bot/` directory within a terminal window or IDE.
3. Run `go build`. A successful build should have created an executable within `/cmd/bot`.


# Running the server
1. Run the executable that was created when building the project.
2. Navigate to <localhost:8080> to check the status of the server.

# Debugging the server

In vscode, add the following launch configuration then select "runserver" (or "F5") in the debug menu.

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

- [LeoV's Embed Visualizer](https://leovoel.github.io/embed-visualizer/) - Preview how some given embed (JSON payload) might appear on Discord, but it also serves as a reference for some limits. This vizualization tool does not render all types of embeds. Check out the About page for caveats or consult the [Embed docs](https://discord.com/developers/docs/resources/channel#embed-object).

<!-- links -->

[discordjs]: <https://discordjs.guide/preparations/adding-your-bot-to-servers.html#bot-invite-links> "Javascript SDK for Discord - great docs"

[bot-token]: <https://github.com/andersfylling/disgord/wiki/Get-bot-token-and-add-it-to-a-server> "Get bot token and add to server"

[dc-dev]: <https://discord.com/developers> "Discord developer portal"

[discordgo ]: <github.com/bwmarrin/discordgo> "Go bindings for Discord"
[dgrouter]: <github.com/Necroforger/dgrouter> "router for discordgo"
[gocilint]: <https://golangci-lint.run/usage/quick-start/> "linting tool to use locally & in remote builds"