# mk-inject

`mk-inject` allow you to inject the input of a command into your markdown files.

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
