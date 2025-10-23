# Git Archiver

Git Archiver is a command-line tool that periodically pulls Git repositories according to a specified schedule defined in a YAML configuration file.

## Features

- Schedule Git repository pulls at specified intervals
- Maintain and update multiple repositories simultaneously
- Configure repositories individually with custom settings
- Cross-platform support (Linux, Windows)

## Installation

### Pre-compiled Binaries

Download the latest release for your platform from the Releases page.

### Build from Source

#### Prerequisites

- Go 1.24.1 or higher
- Git

#### Steps

```bash
# Clone the repository
git clone https://github.com/RatDotRaw/git-archiver.git # Make sure to get the correct repo URL
cd git-archiver

# Build directly
go build -o git-archiver cmd/git-archiver/main.go

# Or use the build scripts
./scripts/build-linux.sh    # For Linux
./scripts/build-windows.sh  # For Windows
```

## Usage

1. Create a `configs/config.yaml` file with your repository configurations (or modify the example provided)
2. Run the executable:

```bash
# Run directly with Go
go run cmd/git-archiver/main.go

# Or run the built executable
./git-archiver
```

## Configuration

Create a `configs/config.yaml` file based on the example provided. Here's a sample configuration:

```yaml
repositories:
  - url: https://github.com/username/repo1.git
    path: /path/to/local/storage/repo1
    interval: 1d
    
  - url: https://github.com/username/repo2.git
    path: ./storage/repo2
    interval: 24h
    
  - url: ssh://github.com/username/repo3.git
    path: C:\GitArchives\repo3
    interval: 30m
```

### Configuration Options

| Option   | Description                                  | Format                        |
|----------|----------------------------------------------|-------------------------------|
| url      | Git repository URL                           | String                        |
| path     | Local path where repository should be stored | String                        |
| interval | Time between pulls                           | Duration (e.g., 30m, 1h, 24h) |

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

Don't know what to contribute, search for TODO: in the code (if there are any)!

---

Yes, all of this could be done in a simple cron task.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
