# Codelo — hướng dẫn phát triển

Dự án Go. Fork từ Reasonix, Việt hóa và mở rộng.
MIT-licensed. Go ≥1.25 required.

## Công nghệ

- **Ngôn ngữ** — Go 1.25+
- **CLI** — Bubble Tea v2 + Lipgloss v2 (TUI)
- **Web UI** — Vanilla HTML/CSS/JS SPA, nhúng Go binary
- **Build** — Makefile, CGO_ENABLED=0

## Cấu trúc

| Đường dẫn | Mô tả |
|---|---|
| `cmd/codelo/` | Entry point CLI |
| `internal/cli/` | CLI commands + Bubble Tea TUI |
| `internal/tool/` | Tool definitions |
| `internal/provider/` | Model providers (DeepSeek, Anthropic, OpenAI) |
| `internal/plugin/` | Plugin system |
| `internal/hook/` | Lifecycle hooks |
| `internal/agent/` | Multi-agent system |
| `internal/config/` | Config loader (TOML) |
| `internal/skill/` | Skills (built-in + custom) |
| `internal/serve/` | HTTP + SSE server + Web SPA |
| `internal/i18n/` | Đa ngôn ngữ (en, zh, vi) |
| `internal/lsp/` | LSP tools |

## Lệnh

```sh
make build       # → bin/codelo
make vet         # go vet ./...
make fmt         # gofmt -w .
make test        # go test ./...
make cross       # cross-compile → dist/
```

## Thêm ngôn ngữ mới

1. Copy `internal/i18n/messages_en.go` thành `messages_XX.go`
2. Dịch tất cả field
3. Thêm case trong `setLanguage()` ở `i18n.go`
4. Thêm detection trong `normalize()` ở `i18n.go`