# OpenFaas Go Template with Workspaces

This provides three "example functions" that show how to use the experimental Golang Classic template from my fork.

This template using Go 1.18's workspaces instead of trying to manipulate teh `go.mod` file.

## Quickstart

```sh
faas-cli template pull https://github.com/LucasRoesler/templates#feat-upgrade-golang-to-118-and-support-mod
faas-cli build
```

### Supported modes

1. Normal Go modules with external packages: see [`tester`](./tester/), this example uses the external package `logrus`
2. Named sub-packages i.e. routes, imported into the handler: see [`tester`](./tester/), this example also contains sub-packages `yo` and `pulp`
3. Using a local replace on a package: see [`mitreplace`](./mitreplace/), this example has a replace of the `logrus`
4. Function usage vendoring for dependencies: see [`mitvendor`](./mitvendor/)
   - this is the pattern you need to use if you have a private module.
   - note that it requires the following build args
     ```yaml
     GO111MODULE: off
     GOFLAGS: "-mod=vendor"
     ```

### Unsupported modes

- Normal Go modules on, with public vendored packages?

Over all this use case is covered in (4) above. However, if you try to enable `GO111MODULE` and enable vendoring, it will produce an error like this

```sh
go: -mod may only be set to readonly when in workspace mode, but it is set to "vendor"
	Remove the -mod flag to use the default readonly value,
	or set GOWORK=off to disable workspace mode.
executor failed running [/bin/sh -c GOOS=${TARGETOS} GOARCH=${TARGETARCH} CGO_ENABLED=${CGO_ENABLED} go test ./... -cover]: exit code: 1
------
```

Because the template requires workspaces, we can not set `GOWORK=off`. Hence, to use vendoring, you must use the combination

```yaml
GO111MODULE: off
GOFLAGS: "-mod=vendor"
```
