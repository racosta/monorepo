# Golang support

Bazel Go dependencies are based on the `go.work` file at the repo root. That will include
references to all Go modules in the repository.

```bash
bazel run @rules_go//go work use <project-path>
```

## Module-level

Use `go mod tidy` to maintain a `go.mod` for each individual project. Use the `Go` executable from
`@rules_go` by

```bash
bazel run @rules_go//go mod tidy
```

## Gazelle

Use `bazel run //:gazelle` to (re-)generate BUILD files.
