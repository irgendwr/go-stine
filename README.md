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
curl -O -L https://github.com/irgendwr/go-stine/releases/latest/download/stine_Linux_x86_64.tar.gz
tar -xvzf stine_Linux_x86_64.tar.gz
```

Create a file called `.stine.yaml` inside this folder (e.g. using `nano .stine.yaml`) and edit it to fit your needs.
See [config](#config) section for examples.

### Config

**Note: Do not use Tabs! Indent config with spaces instead.**

Example:

```yaml
Username: baw1234
Password: your-password-here
```

## Usage

`./stine help`

Export a schedule: `./stine scheduler export Y2020M07 -o ./2020_07.ics`

## Build

Run `make`.
