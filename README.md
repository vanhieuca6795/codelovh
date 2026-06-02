<p align="center">
  <img src="docs/logo.svg" alt="Codelo" width="640"/>
</p>

<p align="center">
  <strong>English</strong>
  &nbsp;·&nbsp;
  <a href="./README.vi.md">Tiếng Việt</a>
  &nbsp;·&nbsp;
  <a href="./README.zh-CN.md">简体中文</a>
</p>

<h3 align="center">Codelo — A Vietnamese AI coding agent for your terminal.</h3>
<p align="center">Fork of Reasonix, fully Vietnamese-localized. Engineered around DeepSeek prefix-cache stability — so token costs stay low across long sessions.</p>

## Install

Requires Go ≥ 1.25. macOS · Linux · Windows.

```bash
git clone https://github.com/vanhieuca6795/codelovh.git
cd codelovh
make build
./bin/codelo chat
```

## Quick Start

```bash
./bin/codelo chat                           # Interactive session
./bin/codelo run "implement the TODOs"       # One-shot task
./bin/codelo setup                          # Config wizard
./bin/codelo serve                          # Web SPA (localhost)
CODELO_LANG=vi ./bin/codelo chat            # Vietnamese interface
```

## Configuration

Config file: `./codelo.toml` or `~/.config/codelo/config.toml`.

```toml
default_model = "deepseek"
language = "vi"                  # UI language: vi, en, zh

[[providers]]
name = "deepseek"
kind = "openai"
base_url = "https://api.deepseek.com"
models = ["deepseek-v4-flash", "deepseek-v4-pro"]
api_key_env = "DEEPSEEK_API_KEY"

[[providers]]
name = "claude"
kind = "anthropic"
model = "claude-opus-4-8"
api_key_env = "ANTHROPIC_API_KEY"
```

## What makes Codelo different

- **Vietnamese-first**: CLI + Web SPA fully translated (100+ i18n fields)
- **Zero-dependency Web UI**: 244-line vanilla HTML/CSS/JS embedded in Go binary
- **Cache-first loop**: 99%+ DeepSeek cache hit → ~$12/month
- **Multi-provider**: DeepSeek + Anthropic Claude + OpenAI
- **Plugin + Hook system**: Stdio/HTTP plugins, lifecycle hooks

## Features

| Feature | Status |
|---------|--------|
| Vietnamese CLI TUI | ✅ Complete |
| Vietnamese Web SPA | ✅ Complete (zero deps) |
| Multi-provider | ✅ DeepSeek + Anthropic + OpenAI |
| Plugin system | ✅ Stdio + HTTP |
| Hooks | ✅ PreToolUse, PostToolUse, lifecycle |
| Multi-agent | ✅ Coordinator + task orchestration |
| MCP server | ✅ stdio + SSE |
| LSP tools | ✅ Go, Rust, Python, TypeScript |
| VS Code Extension | 🚧 Skeleton |
| Agent Manager | 🚧 Skeleton |

## Slash Commands

| Command | Description |
|---------|------------|
| `/compact` | Compact context |
| `/new` | New session |
| `/resume` | Resume saved session |
| `/model` | Switch model |
| `/mcp` | Manage MCP servers |
| `/hooks` | Manage hooks |
| `/skill` | Manage skills |
| `/memory` | View memory files |
| `/init` | Initialize AGENTS.md |
| `/quit` | Exit |

## License

MIT — see [LICENSE](./LICENSE). Forked from [Reasonix](https://github.com/esengine/DeepSeek-Reasonix).

## Community

- [Reasonix Discord](https://discord.gg/XF78rEME2D)
- [GitHub Issues](https://github.com/vanhieuca6795/codelovh/issues)