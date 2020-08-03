# go-stine

[![Build status](https://github.com/irgendwr/go-stine/workflows/build/badge.svg)](https://github.com/irgendwr/go-stine/actions?query=workflow%3Abuild)
[![Release status](https://github.com/irgendwr/go-stine/workflows/release/badge.svg)](https://github.com/irgendwr/go-stine/actions?query=workflow%3Arelease)
[![Go Report Card](https://goreportcard.com/badge/github.com/irgendwr/go-stine)](https://goreportcard.com/report/github.com/irgendwr/go-stine)
[![GitHub Release](https://img.shields.io/github/release/irgendwr/go-stine.svg)](https://github.com/irgendwr/go-stine/releases)

[STiNE](https://www.stine.uni-hamburg.de) CLI/library written in [Go](https://golang.org/).

You can find more details about STiNE here:

- https://www2.informatik.uni-hamburg.de/fachschaft/wiki/index.php/STiNE
- https://www2.informatik.uni-hamburg.de/fachschaft/wiki/index.php/STiNE-Interna

## Demo

[![asciicast](https://asciinema.org/a/351399.svg)](https://asciinema.org/a/351399)

## Installation

### Linux

Download and unpack the latest release:
```bash
# download
curl -O -L https://github.com/irgendwr/go-stine/releases/latest/download/stine_Linux_x86_64.tar.gz
# unpack
tar -xvzf stine_Linux_x86_64.tar.gz
# copy to folder in $PATH
sudo cp ./stine /usr/bin/stine
```

Create a file called `.stine.yaml` (either inside your home folder or the folder containing the program) (e.g. using `nano ~/.stine.yaml`) and edit it to fit your needs.
See [config](#config) section for examples.

## Config

If no config file is specified using the `-c`/`--config` flag, the program looks for a file called `.stine.yaml` in the following paths:

1. Program directory (path the `stine` binary is in)
2. CWD (current working directory)
3. Home folder

**Note: Do not use Tabs! Indent config with spaces instead.**

Example:

```yaml
username: baw1234
password: your-password-here
```

## Usage

List of commands and flags: `stine help`

### Examples

List exams: `stine exams`

List all exam results: `stine examresults -a`

Show schedule of a given day: `stine schedule 06.07 --day`

Show schedule of the current week: `stine schedule --week`

Show schedule of next week: `stine schedule --next --week`

Export a schedule: `stine schedule export Y2020M07 -o ./2020_07.ics`

## Build

Run `make`.

## Related Projects

- [felsenhower/stine_calendar_bot](https://github.com/felsenhower/stine_calendar_bot) - Automatic calendar download bot for STiNE
- [Flova/ModernStineTheme](https://github.com/Flova/ModernStineTheme) - A better STiNE CSS theme

## Contributing

Contributions are welcome! Feel free to open an issue, or even better: submit a pull-request.
