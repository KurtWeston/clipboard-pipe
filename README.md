# clipboard-pipe

Pipe command output to clipboard or paste clipboard content as stdin - seamless terminal-to-GUI bridge

## Features

- Read from stdin and write to system clipboard (copy mode)
- Read from clipboard and write to stdout (paste mode)
- Auto-detect mode: if stdin is available, copy; otherwise paste
- Cross-platform support for Linux (X11/Wayland), macOS, and Windows
- Preserve newlines and special characters in clipboard content
- Silent operation by default with optional verbose flag
- Exit codes: 0 for success, 1 for clipboard access errors
- Handle empty clipboard gracefully
- Support piping in shell command chains seamlessly
- Minimal binary size with no external runtime dependencies

## How to Use

Use this project when you need to:

- Quickly solve problems related to clipboard-pipe
- Integrate go functionality into your workflow
- Learn how go handles common patterns

## Installation

```bash
# Clone the repository
git clone https://github.com/KurtWeston/clipboard-pipe.git
cd clipboard-pipe

# Install dependencies
go build
```

## Usage

```bash
./main
```

## Built With

- go

## Dependencies

- `github.com/atotto/clipboard`

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
