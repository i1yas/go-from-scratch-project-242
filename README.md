# hexlet-path-size

Utility for calculating file and directory sizes

[![Actions Status](https://github.com/i1yas/go-from-scratch-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/i1yas/go-from-scratch-project-242/actions)

## Usage in Go

```go
size, err := code.GetPathSize(path, isRecursive, isHumanReadable, includeHidden)
```

## Usage in CLI

```bash
make build # build
./bin/hexlet-path-size my-file # run binary from bin
```

### Human-readable format

By default show size in bytes. When enabled automatically picks appropriate unit (B, KB, MB, GB, ...).

```bash
./bin/hexlet-path-size -H my-file
./bin/hexlet-path-size --human my-file
```

### Recursive

By default counts only first level of directory.

```bash
./bin/hexlet-path-size -r my-dir
./bin/hexlet-path-size --recursive my-dir
```

### Include hidden files

By default does not count hidden files (starting with `.`)

```bash
./bin/hexlet-path-size -a my-dir
./bin/hexlet-path-size --all my-dir
```

## File types

The reported size depends of the file type:

- __Regular file__: file's own size
- __Symlink__: only size of symlink itself, _not_ the target file
- __Directory__: sum of the sizes of it's children files and subdirectories (depending on `recursive` flag)

All other files are not supported.

- If program's target is an unsupported file type, it returns an error.
- If program's target is directory that contains unsupported files then those files will be skipped

### Example

[![asciicast](https://asciinema.org/a/ceVnx26DHQAauxkv.svg)](https://asciinema.org/a/ceVnx26DHQAauxkv)
