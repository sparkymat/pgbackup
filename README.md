# pgsnap
pgsnap is a script/utility for managing postgresql backups using the psql command under the hood.

## Installation

`go get -u github.com/sparkymat/pgsnap`

N.B: You will need to have .pgbackup.toml to specify the name of the repository (or specify the config file each time) in the current working directory. Also, pgsnap assumes that you have administrator privileges with the local postgtresql instance.

## Usage

`pgsnap help` will print the available commands and a quick description of what they do
