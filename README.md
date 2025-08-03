# Minecraft Query

A simple CLI app to query Minecraft server information.

Currently only supports querying of online players.

Planned Features:
- [x] Query player status
- [ ] Get server ping
- [ ] More server info (resource packs, modloader)

## Usage

```command
go build .
```

```command
./mc-query <ip|domain-name> [<port>]
./mc-query --help

# For debug output
LOG_LEVEL= ./mc-query ...
```
