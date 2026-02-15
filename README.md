# Oh My Gosh

A sleek, customizable theme for the [GoSH](https://github.com/gosh-terminal/gosh) terminal, inspired by the popular Oh My Zsh framework.

![Oh My Gosh Screenshot](https://via.placeholder.com/800x450)

## Features

- 🎨 **Beautiful Themes**: Choose from a variety of carefully crafted color schemes
- ⚡ **Performance Focused**: Minimal impact on terminal startup time
- 🔌 **Plugin System**: Extend functionality with community plugins
- 💻 **Cross-Platform**: Works on macOS, Linux, and Windows
- 🚀 **Smart Autocompletion**: Context-aware suggestions for commands and arguments
- 📊 **Informative Prompt**: Shows git status, command execution time, and more

## Things To Do!!!

- Add "viper" for shell configuration.
- Add "golangci-lint" support.
- Add install developement dependencies scripts.
- Add some keyboard controls to shell in "Raw Terminal" mode. (x/term)
- Add double "linked-list" data structure in "context" for supporting navigate history with arrow keys. (x/term)
- Add left & right move controls in prompts. (x/term)
- Add color configuration using "viper".
- Add manual page using "cobra CLI" framework or something else.
- Using "buffers", "bufio" and string builders for better performance.
- Doing some refactors on "code architecture" for easier maintenance and easy to write tests.
- Doing some enhancements on "database connection lifecycle".
- Using "syscall" if needed
- Using some features of low-level programming ("unsafe") for better performance and control.
- Adding "encryption" and "security" aspects on shell sqlite database.
- Database location must be in /var path (for linux) or somewhere in other operating systems.
-

## Installation

### Prerequisites

- [GoSH](https://github.com/gosh-terminal/gosh) terminal installed
- Git

### Quick Install

```bash
go install github.com/itsjustabavaar/oh-my-gosh@latest
```

Or clone the repository manually:

```bash
git clone https://github.com/itsjustabavaar/oh-my-gosh.git ~/.oh-my-gosh
cd ~/.oh-my-gosh
./install.sh
```

## Configuration

Add the following to your `~/.goshrc` file:

```go
import "github.com/itsjustabavaar/oh-my-gosh"

func init() {
    // Choose your theme
    ohmygosh.SetTheme("monokai")

    // Enable plugins
    ohmygosh.LoadPlugin("git")
    ohmygosh.LoadPlugin("golang")
}
```

## Available Themes

- `monokai` - Dark theme with vibrant colors
- `nord` - Cool, bluish dark theme
- `solarized-dark` - Dark theme with distinct colors
- `solarized-light` - Light theme with distinct colors
- `tokyo-night` - Dark blue theme inspired by Tokyo at night

To preview themes:

```bash
oh-my-gosh theme preview
```

## Plugins

### Core Plugins

- **git** - Git integration and status indicators
- **golang** - Go development environment enhancements
- **docker** - Docker container and image management shortcuts
- **node** - NodeJS version and package information
- **python** - Python virtual environment and package tools

### Installing Custom Plugins

Place custom plugins in `~/.oh-my-gosh/plugins/`:

```bash
cp my-custom-plugin.go ~/.oh-my-gosh/plugins/
```

Then load them in your configuration:

```go
ohmygosh.LoadPlugin("my-custom-plugin")
```

## Customization

### Custom Prompt Segments

Create your own prompt segments by implementing the `PromptSegment` interface:

```go
type MyCustomSegment struct{}

func (s *MyCustomSegment) Content() string {
    return "⚡"
}

func (s *MyCustomSegment) Foreground() string {
    return "#ffff00"
}

func (s *MyCustomSegment) Background() string {
    return "#000000"
}

// Add to your prompt
ohmygosh.AddSegment(&MyCustomSegment{})
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by [Oh My Zsh](https://github.com/ohmyzsh/ohmyzsh)
- Thanks to the [GoSH](https://github.com/gosh-terminal/gosh) team for the amazing terminal
- All contributors who have helped shape this project
