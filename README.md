# pulumi-stack-switcher

A lightweight CLI tool that makes it easy to switch between Pulumi stacks with shell autocompletion.

## Features

- Automatically detects available stacks by scanning for `Pulumi.<stack-name>.yaml` files
- Shell autocompletion for quick stack switching (tab-completion)
- Cross-platform support (Linux, macOS, Windows, FreeBSD)

## Installation

### From Release

Download the latest release for your platform from the [Releases](https://github.com/trondhindenes/pulumi-stack-switcher/releases) page.

### From Source

```bash
go install github.com/trondhindenes/pulumi-stack-switcher@latest
```

Or clone and build:

```bash
git clone https://github.com/trondhindenes/pulumi-stack-switcher.git
cd pulumi-stack-switcher
make install
```

## Usage

```bash
# List available stacks
pulumi-stack-switcher

# Switch to a specific stack
pulumi-stack-switcher dev
pulumi-stack-switcher production
pulumi-stack-switcher pr<tab> --> lets you switch to a stack matching "pr"
pulumi-stack-switcher <tab> --> lets you switch to any stack
```

## Shell Completion Setup

### Bash

Add to your `~/.bashrc`:

```bash
source <(pulumi-stack-switcher completion bash)
```

Or install system-wide:

```bash
# Linux
pulumi-stack-switcher completion bash > /etc/bash_completion.d/pulumi-stack-switcher

# macOS with Homebrew
pulumi-stack-switcher completion bash > $(brew --prefix)/etc/bash_completion.d/pulumi-stack-switcher
```

### Zsh

Add to your `~/.zshrc`:

```zsh
source <(pulumi-stack-switcher completion zsh)
```

Or install to your fpath:

```bash
# Linux
pulumi-stack-switcher completion zsh > "${fpath[1]}/_pulumi-stack-switcher"

# macOS with Homebrew
pulumi-stack-switcher completion zsh > $(brew --prefix)/share/zsh/site-functions/_pulumi-stack-switcher
```

If shell completion is not already enabled in your zsh environment, add this to `~/.zshrc`:

```zsh
autoload -U compinit; compinit
```

### Fish

```bash
pulumi-stack-switcher completion fish > ~/.config/fish/completions/pulumi-stack-switcher.fish
```

### PowerShell

Add to your PowerShell profile:

```powershell
pulumi-stack-switcher completion powershell | Out-String | Invoke-Expression
```

To find your profile path:

```powershell
echo $PROFILE
```

## How It Works

The tool scans the current directory for files matching the pattern `Pulumi.<stack-name>.yaml` and uses these to determine available stacks. When you select a stack, it runs `pulumi stack select <stack-name>` under the hood.

## Development

```bash
# Run tests
make test

# Build locally
make build

# Format code
make fmt

# Run linter
make lint
```

## License

MIT
