# Conteye (Container Eye) 🐳👁️

A simple CLI tool to monitor Docker container statistics, focusing on container restart counts and health information.

## Features

- Monitor container restart counts
- View container creation times and IDs
- Simple and straightforward CLI interface

## Installation

### Using Go

```bash
go install github.com/zwsq/conteye@latest
```

### Binary Downloads

Download the latest binary for your platform from the [releases page](https://github.com/zwsq/conteye/releases).

## Usage

```bash
conteye restart    # Show container restart counts
conteye help      # Show help message
```

## Building from Source

```bash
git clone https://github.com/zwsq/conteye.git
cd conteye
go build
```

## License

MIT License 
