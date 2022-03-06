# bearings

A fast, clean, super-customisable shell prompt.

- Supports zsh, bash, fish, and more.
- Easily write your own modules using any language.
- Simple configuration with YAML - works out of the box with a nice default configuration too. 
- Smart-colouring handles powerline character colour inverting intelligently.

## Examples Gallery

<table>
    <tr><td align="center"><img src="_examples/default/screenshot.png" /></td><td align="center"><a href="_examples/default/config.yml">Default</a></td></tr>
    <tr><td align="center"><img src="_examples/halflife/screenshot.png" /></td><td align="center"><a href="_examples/halflife/config.yml">Half Life</a></td></tr>
    <tr><td align="center"><img src="_examples/traditional/screenshot.png" /></td><td align="center"><a href="_examples/traditional/config.yml">Traditional</a></td></tr>
</table>

## Installation

You can download the latest binaries [here](https://github.com/liamg/bearings/releases/latest). Make sure you `chmod +x`  the binary and place it somewhere in your `PATH`. Then follow the instructions for your shell below.

It is recommended to install font(s) which include powerline characters, especially [nerd-fonts](https://github.com/ryanoasis/nerd-fonts).

## Configuration

### Automatic

You can attempt to automatically configure your shell by running `bearings install`. This will modify your shell configuration files in order to set bearings as your PS1 generator. For advanced configurations, you should use the manual methods below.

### ZSH

```zsh
#bearings-auto:start
function configure_bearings() {
    PROMPT="$(bearings prompt -s zsh -e $?)"
}
[ ! "$TERM" = "linux" ] && precmd_functions+=(configure_bearings)
#bearings-auto:end
```

### Bash

```bash
#bearings-auto:start
bearings_prompt() { export PS1=$(bearings prompt -s bash -e $?); }
[[ ! "$TERM" = "linux" ]] && export PROMPT_COMMAND=bearings_prompt
#bearings-auto:end
```

### Fish

```fish
#bearings-auto:start
function fish_prompt
    bearings prompt -s fish -e $status
end
#bearings-auto:end
```

## Customisation

The config file is read from `~/.config/bearings/config.yml`. You can create a default config file by running `bearings prompt` for the first time.

You can find example configurations with screenshots for each in the [examples directory](_examples).

| Property | Default | Description |
|----------|---------|-------------|
| padding  | 1       | Number of spaces before and after each module. Can be overriden on a per-module basis.
| end      |  (powerline character) | The string to render at the end (right) of the prompt.
| divider  |  (powerline character) | The string to render between modules. Can be overriden on a per-module basis.
| fg       | white   | Default foreground colour for all modules. Can be overridden on a per-module basis.
| bg       | black   | Default background colour for all modules. Can be overridden on a per-module basis.
| lines_above | 1    | Number of blank lines to render above the prompt.
| modules  | exitcode, cwd, git | A list of modules and their configurations.

Colours can be specified in hexidecimal, e.g. `#ffffff`. You can also refer to your terminal colour scheme colours using `default` (for default fg/bg), `red`, `green`, `yellow`, `blue`, `magenta`, `cyan`, `white`, `black`.

All modules support the following options:

| Property       | Default | Description |
|----------------|---------|-------------|
| label          | %s      | Text to render alongside the module output. Use %s as the placeholder for the module content. 
| fg             | _inherits from top-level fg_ | 
| bg             | _inherits from top-level bg_ |
| padding_before | _inherits from top-level padding_ |
| padding_after  | _inherits from top-level padding_ |
| divider        | _inherits from top-level divider_ |

## Available Modules

### Current Working Directory (`cwd`)

Show the current working directory.

![cwd](_assets/cwd.png)

| Property       | Default | Description |
|----------------|---------|-------------|
| max_depth      | 0       | The maximum number of directories to render in the path. If this number is exceeded, the output will be truncated to show `...` followed by the lowest `max_depth` number of directories.

### Exit Code (`exitcode`)

Show the exit code of the previous command. By default will only show when the command fails, but can also show a success icon/message.

![exit code](_assets/exitcode.png)

| Property       | Default | Description |
|----------------|---------|-------------|
| show_success   | false   | Show the module when the previous command succeeded (exit code zero).
| success_bg     | _inherits from bg, top-level bg_ | Background colour for the module when the previous command succeeded.
| failure_bg     | _inherits from bg, top-level bg_ | Background colour for the module when the previous command failed.
| success_fg     | green   | Foreaground colour for the module when the previous command succeeded.
| failure_fg     | red     | Foreground colour for the module when the previous command failed.
| success_output |        | Output for the module when the previous command succeeded.
| failure_output |        | Output for the module when the previous command failed.

### Git Overview (`git`)

Show an overview of the current git status. Displays the branch name, a set of possible icons, and the number of commits ahead/behind of the base branch.

| Property       | Default | Description |
|----------------|---------|-------------|
| icon_stashed   | S       | The icon/text to display when stashed changes are available.
| icon_untracked | ?       | The icon/text to display when untracked files are present.
| icon_modified  | M       | The icon/text to display when tracked files are modified.
| icon_staged    | A       | The icon/text to display when changes are staged.
| icon_conflicts | !       | The icon/text to display when conflicts are present.

### Command (`command`)

Run a shell command and use the combined output streams as the module output.

| Property       | Default | Description |
|----------------|---------|-------------|
| command        | _none_  | The shell command to run.

### Hostname (`hostname`)

Show the current hostname.

### New line (`newline`)

Output a single new line. Before/after padding values default to `0` for convenience.

### Text (`text`)

Output the specified text.

| Property       | Default | Description |
|----------------|---------|-------------|
| text           | _none_  | The text to output.

### Username (`username`)

Show the current username.

