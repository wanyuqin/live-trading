# Live Trading
---

**📖 English | [中文](README_zh.md)**

[![Go Version](https://img.shields.io/badge/Go-1.20+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/wanyuqin/live-trading)](https://goreportcard.com/report/github.com/wanyuqin/live-trading)

A real-time stock trading monitoring tool developed in Go, providing an elegant Terminal User Interface (TUI) to monitor stock price changes and market conditions.

## ✨ Features

- 🚀 **Real-time Monitoring**: Support real-time price monitoring for custom stock code lists
- 📊 **Market Overview**: Real-time display of major indices (Shanghai, Shenzhen, ChiNext) market conditions
- 🎯 **TUI Interface**: Modern terminal-based user interface with keyboard shortcuts support
- ⚙️ **Flexible Configuration**: Support custom configuration files with runtime dynamic updates
- 🔄 **Auto Restart**: Support graceful restart and recovery of monitoring services
- 🛡️ **Safe & Reliable**: Comprehensive error handling and goroutine management

## 🚀 Quick Start

### Prerequisites

- Go 1.20 or higher
- Network connection (for fetching stock data)

### Installation

1. Clone the repository
```bash
git clone https://github.com/wanyuqin/live-trading.git
cd live-trading
```

2. Build the project
```bash
go build -o live-trading main.go
```

3. Run the program
```bash
./live-trading
```

### Using Custom Configuration

```bash
./live-trading -c /path/to/your/config.yaml
```

## 📖 Usage Guide

### Basic Operations

- **Add Stock**: Press `a` key, input 6-digit stock code (e.g., 000001)
- **Delete Stock**: Press `x` key to delete the currently selected stock
- **Exit Program**: Press `Ctrl+C`

### Stock Code Format

- **Shanghai Stock Exchange**: 6-digit numbers starting with `6` or `5` (e.g., 600000)
- **Shenzhen Stock Exchange**: 6-digit numbers starting with `0`, `1`, or `3` (e.g., 000001)
- **Indices**: Support major market indices (000001, 399001, 399006)

### Configuration

Default configuration file location: `~/.live_trading/trading.yaml`

```yaml
watchList:
  stock:
    - "000001"  # Shanghai Composite Index
    - "600000"  # SPD Bank
    - "000002"  # Vanke A
```

## 🏗️ Project Architecture

```
live-trading/
├── cmd/                    # Command line entry
├── internal/              # Internal packages
│   ├── configs/          # Configuration management
│   ├── domain/           # Domain layer (entities, services, repositories)
│   ├── infrastructure/   # Infrastructure layer (API integration)
│   ├── views/            # View layer (UI components)
│   └── watcher/          # Monitors
├── tools/                 # Utility packages
└── main.go               # Program entry point
```

### Tech Stack

- **Language**: Go 1.20+
- **UI Framework**: [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- **CLI Framework**: [Cobra](https://github.com/spf13/cobra)
- **Configuration**: YAML
- **Data Source**: East Money API

## 🔧 Development

### Local Development Environment

1. Install dependencies
```bash
go mod download
```

2. Run tests
```bash
go test ./...
```

3. Build development version
```bash
go build -o live-trading-dev main.go
```

### Project Structure Explanation

- **Domain Layer**: Contains business logic, entity definitions, and service interfaces
- **Infrastructure Layer**: Handles external API calls and data persistence
- **Views Layer**: Manages user interface and interaction logic
- **Watcher Layer**: Responsible for data monitoring and updates

## 📊 Data Source

This project uses East Money (eastmoney.com) public API as the data source:

- **Real-time Quotes**: Obtained through SSE (Server-Sent Events)
- **Load Balancing**: Support for multiple API server random selection
- **Data Fields**: Contains key information such as price, change percentage, volume, etc.

## 🤝 Contributing

Issues and Pull Requests are welcome!

### Contribution Guide

1. Fork this repository
2. Create a feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 Changelog

### [Unreleased]

- [ ] Optimize error handling mechanism
- [ ] Implement smooth refresh when adding/deleting stocks
- [ ] Add more technical indicators
- [ ] Support historical data queries

### [Current Version]

- [x] Delete stock code functionality
- [x] Market monitoring
- [x] Restart monitoring via context
- [x] Custom configuration file support
- [x] Basic TUI interface

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## 🙏 Acknowledgments

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Excellent TUI framework
- [Cobra](https://github.com/spf13/cobra) - Powerful CLI framework
- [East Money](http://www.eastmoney.com/) - Provides stock data API

## 📞 Contact

- Project Homepage: [https://github.com/wanyuqin/live-trading](https://github.com/wanyuqin/live-trading)
- Issue Reports: [Issues](https://github.com/wanyuqin/live-trading/issues)

---

⭐ If this project helps you, please give it a star!


