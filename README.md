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
