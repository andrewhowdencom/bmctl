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

## Supported Commands

`bmctl` currently supports a subset of commands from the Blackmagic API:
- `bmctl audio`: Control camera audio components including inputs (`input`, `description`, `supported-inputs`), levels (`level`), `phantom-power`, `padding`, and `low-cut-filter`.
- `bmctl video`: Set ISO and white balance.
- `bmctl lens`: Control lens metadata and features.

To view detailed help on any command, use the `--help` flag:
```bash
$ bmctl audio level set --help
```

[Bingo]: https://github.com/bwplotka/bingo


### Task Runner

The task runner is [Taskfile]. See the available tasks with:


```bash
$ TASK list
```

[Taskfile]: https://taskfile.dev/