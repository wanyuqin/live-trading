# Live Trading

---

**📖 [English](README.md) | 中文**

[![Go Version](https://img.shields.io/badge/Go-1.20+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/wanyuqin/live-trading)](https://goreportcard.com/report/github.com/wanyuqin/live-trading)

一个基于Go语言开发的实时股票交易监控工具，提供优雅的终端用户界面（TUI）来监控股票价格变化和市场行情。

## ✨ 特性

- 🚀 **实时监控**: 支持自定义股票代码列表的实时价格监控
- 📊 **市场行情**: 实时显示主要指数（上证、深证、创业板）行情
- 🎯 **TUI界面**: 基于终端的现代化用户界面，支持键盘快捷键
- ⚙️ **灵活配置**: 支持自定义配置文件，运行时动态更新
- 🔄 **自动重启**: 支持监控服务的优雅重启和恢复
- 🛡️ **安全可靠**: 完善的错误处理和goroutine管理

## 🚀 快速开始

### 前置要求

- Go 1.20 或更高版本
- 网络连接（用于获取股票数据）

### 安装

1. 克隆仓库
```bash
git clone https://github.com/wanyuqin/live-trading.git
cd live-trading
```

2. 构建项目
```bash
go build -o live-trading main.go
```

3. 运行程序
```bash
./live-trading
```

### 使用自定义配置

```bash
./live-trading -c /path/to/your/config.yaml
```

## 📖 使用说明

### 基本操作

- **添加股票**: 按 `a` 键，输入6位股票代码（如：000001）
- **删除股票**: 按 `x` 键删除当前选中的股票
- **退出程序**: 按 `Ctrl+C`

### 股票代码格式

- **上海证券交易所**: 以 `6` 或 `5` 开头的6位数字（如：600000）
- **深圳证券交易所**: 以 `0`、`1`、`3` 开头的6位数字（如：000001）
- **指数**: 支持主要市场指数（000001、399001、399006）

### 配置说明

配置文件默认位置：`~/.live_trading/trading.yaml`

```yaml
watchList:
  stock:
    - "000001"  # 上证指数
    - "600000"  # 浦发银行
    - "000002"  # 万科A
```

## 🏗️ 项目架构

```
live-trading/
├── cmd/                    # 命令行入口
├── internal/              # 内部包
│   ├── configs/          # 配置管理
│   ├── domain/           # 领域层（实体、服务、仓储）
│   ├── infrastructure/   # 基础设施层（API集成）
│   ├── views/            # 视图层（UI组件）
│   └── watcher/          # 监控器
├── tools/                 # 工具包
└── main.go               # 程序入口
```

### 技术栈

- **语言**: Go 1.20+
- **UI框架**: [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- **命令行**: [Cobra](https://github.com/spf13/cobra)
- **配置**: YAML
- **数据源**: 东方财富API

## 🔧 开发

### 本地开发环境

1. 安装依赖
```bash
go mod download
```

2. 运行测试
```bash
go test ./...
```

3. 构建开发版本
```bash
go build -o live-trading-dev main.go
```

### 项目结构说明

- **Domain Layer**: 包含业务逻辑、实体定义和服务接口
- **Infrastructure Layer**: 处理外部API调用和数据持久化
- **Views Layer**: 管理用户界面和交互逻辑
- **Watcher Layer**: 负责数据监控和更新

## 📊 数据源

本项目使用东方财富（eastmoney.com）的公开API作为数据源：

- **实时行情**: 通过SSE（Server-Sent Events）获取
- **负载均衡**: 支持多个API服务器随机选择
- **数据字段**: 包含价格、涨跌幅、成交量等关键信息

## 🤝 贡献

欢迎提交Issue和Pull Request！

### 贡献指南

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

## 📝 更新日志

### [未发布]

- [ ] 优化错误处理机制
- [ ] 实现股票添加/删除时的平滑刷新
- [ ] 增加更多技术指标
- [ ] 支持历史数据查询

### [当前版本]

- [x] 删除选股代码功能
- [x] 市场行情监控
- [x] 通过context重启监控
- [x] 自定义配置文件支持
- [x] 基础TUI界面

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情

## 🙏 致谢

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - 优秀的TUI框架
- [Cobra](https://github.com/spf13/cobra) - 强大的命令行框架
- [东方财富](http://www.eastmoney.com/) - 提供股票数据API

## 📞 联系方式

- 项目主页: [https://github.com/wanyuqin/live-trading](https://github.com/wanyuqin/live-trading)
- 问题反馈: [Issues](https://github.com/wanyuqin/live-trading/issues)

---

⭐ 如果这个项目对你有帮助，请给它一个星标！
