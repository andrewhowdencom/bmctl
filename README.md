# BMCTL

(Or, "Black Magic Control"). CLI tool & library for interacting with the Black Magic camera series, based on their
OpenAPI definition.

## Development
### Tools

The Go-related tools for this project are managed by [Bingo]. To install the required tools, run:

```bash
$ bingo get
```

There is a useful environment file to convert the versions into something that can be executed:

```bash
. .bingo/variables.env
```

After which the applications can be run via a variable:

```bash
$ $YQ
Usage:
  yq [flags]
  yq [command]

```

[Bingo]: https://github.com/bwplotka/bingo


### Task Runner

The task runner is [Taskfile]. See the available tasks with:


```bash
$ TASK list
```

[Taskfile]: https://taskfile.dev/