# hexlet-path-size

CLI utility for calculating file and directory sizes

[![Actions Status](https://github.com/i1yas/go-from-scratch-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/i1yas/go-from-scratch-project-242/actions)

## Usage

```bash
./bin/hexlet-path-size my-file
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

By default does not count hidden files starting with `.`

```bash
./bin/hexlet-path-size -a my-dir
./bin/hexlet-path-size --all my-dir
```

### Example
[![asciicast](https://asciinema.org/a/ceVnx26DHQAauxkv.svg)](https://asciinema.org/a/ceVnx26DHQAauxkv)
