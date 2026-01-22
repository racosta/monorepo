# Javascript support

Bazel JavaScript dependencies use the [pnpm package manager](https://pnpm.io).

## Update pnpm lock file

From repo root:

```bash
bazel run @pnpm install -- --dir $PWD
```

## Add dependencies

From repo root:

```bash
bazel run @pnpm add -- --dir $PWD <package>
```
