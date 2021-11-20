# mk-inject

`mk-inject` allows you to inject the input of a command into your markdown files.

## Example

The following example shows how we inject the date into a README.md file.

### Input

```markdown
# Title

Date below:
<!-- mk-inject:start:dateHere -->
<!-- mk-inject:end:dateHere -->

Some more content.
```

```shell
date | mk-inject --label dateHere README.md
```

### Result

```
# Title

Date below:
<!-- mk-inject:start:dateHere -->
Fri Nov 19 15:41:01 CET 2021
<!-- mk-inject:end:dateHere -->

Some more content.
```


## Installation

```shell
# homebrew
brew install stenic/tap/mk-inject

# gofish
gofish rig add https://github.com/stenic/fish-food
gofish install github.com/stenic/fish-food/mk-inject

# scoop
scoop bucket add mk-inject https://github.com/stenic/scoop-bucket.git
scoop install mk-inject

# go
go install github.com/stenic/mk-inject@latest

# docker 
docker pull ghcr.io/stenic/mk-inject:latest

# dockerfile
COPY --from=ghcr.io/stenic/mk-inject:latest /mk-inject /usr/local/bin/
```

> For even more options, check the [releases page](https://github.com/stenic/mk-inject/releases).


## Running

```shell
# docker
docker run -ti -v $(pwd):/workspace ghcr.io/stenic/mk-inject:latest sh -c "cat tests/assets/single-line.txt | mk-inject --label single-line tests/basic.in.md"
```

## Documentation

<!-- mk-inject:start:help prefix="```shell" suffix="```" -->
```shell
mk-inject

Usage:
  mk-inject [--label labelName] file [flags]

Flags:
  -h, --help           help for mk-inject
  -i, --inplace        Edit in place
  -l, --label string   Replace a specific tag
  -v, --version        version for mk-inject
```
<!-- mk-inject:end:help -->
