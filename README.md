# OpenFaas Go Template with Workspaces

This provides three "example functions" that show how to use the experimental Golang Classic template from my fork.

This template using Go 1.18's workspaces instead of trying to manipulate teh `go.mod` file.

## Quickstart

```sh
faas-cli template pull https://github.com/LucasRoesler/templates#feat-upgrade-golang-to-118-and-support-mod
faas-cli build
```
