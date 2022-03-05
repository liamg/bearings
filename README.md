# bearings

A fast, clean, super-customisable shell prompt.

- Supports zsh, bash, fish, and more.
- Easily write your own modules using any language.
- Simple configuration with YAML - works out of the box with a nice default configuration too. 
- Smart-colouring handles powerline character colour inverting intelligently.

## Gallery

// TODO

## Installation

You can download the latest binaries [here](https://github.com/liamg/bearings/releases/latest). Make sure you `chmod +x`  the binary and place it somewhere in your `PATH`. Then follow the instructions for your shell below.

## Configuration

### Automatic

You can attempt to automatically configure your shell by running `bearings install`. This will modify your shell configuration files in order to set bearings as your PS1 generator. For advanced configurations, you should use the manual methods below.

### ZSH

```zsh
#bearings-auto:start
function configure_bearings() {
    PROMPT="$(bearings prompt -e $?)"
}
[ ! "$TERM" = "linux" ] && precmd_functions+=(configure_bearings)
#bearings-auto:end
```

### Bash

```bash
#bearings-auto:start
bearings_prompt() { export PROMPT=$(bearings prompt -e $?); }
[[ ! "$TERM" = "linux" ]] && export PROMPT_COMMAND=bearings_prompt
#bearings-auto:end
```

### Fish

```fish
#bearings-auto:start
function fish_prompt
    source bearings prompt -e $status
end
#bearings-auto:end
```

## Customisation

The config file is read from `~/.config/bearings/config.yml`. You can create a default config file by running `bearings` for the first time.

You can find example configurations with screenshots for each in the [examples directory](_examples).

## Available Modules

### Working Directory (`workdir`)
### Exit Code (`exitcode`)
### Git (`git`)
### Custom (`custom`)

## TODO

- [ ] customise depth of workdir
- [ ] add examples and screenshots
- [ ] add module documentation
- [ ] test bash
- [ ] test fish
