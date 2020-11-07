## INSTALLING

#### Installing from go

First, use go get to install the latest version. This command will install the `tasks` and its dependencies:

`go get -u github.com/mrinjamul/tasks`

#### Installing from Binaries

Download for your platform

[Download](https://github.com/mrinjamul/tasks/releases)

For Linux,

```sh
wget ...
tar ...
chmod +x tasks
sudo mv tasks /usr/local/bin
```

or you can put the executable file into your env variables `$PATH`

For Termux,

You need to have `tar wget`. To install simply type `pkg install tar wget`

```sh
cd ~
wget ...
tar ...
chmod +x tasks
mv tasks ../usr/bin
```

## Usage

    tasks will help you get more done in less time.
    It's designed to be as simple as possible to help you
    accomplish your goals.

    Usage:
    tasks [command]

    Available Commands:
    help        Help about any command
    version     Prints version

    Flags:
        --config string   config file (default is $HOME/.tasks.yaml)
    -h, --help            help for tasks

    Use "tasks [command] --help" for more information about a command.

### version

    Prints version
