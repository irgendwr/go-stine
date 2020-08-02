# go-stine

![Build status](https://github.com/irgendwr/go-stine/workflows/build/badge.svg)
![Release status](https://github.com/irgendwr/go-stine/workflows/release/badge.svg)
[![GitHub Release](https://img.shields.io/github/release/irgendwr/go-stine.svg)](https://github.com/irgendwr/go-stine/releases)

STiNE CLI / library written in/for [golang](https://golang.org/).

You can find more details about STiNE here:

- https://www2.informatik.uni-hamburg.de/fachschaft/wiki/index.php/STiNE
- https://www2.informatik.uni-hamburg.de/fachschaft/wiki/index.php/STiNE-Interna

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

Export a schedule: `stine scheduler export Y2020M07 -o ./2020_07.ics`

List exams: `stine exams`

List all exam results: `stine examresults -a`

## Build

Run `make`.

## Contributing

Contributions are welcome! Feel free to open an issue, or even better: submit a pull-request.
