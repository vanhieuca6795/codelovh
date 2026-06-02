<p align="center">
  <img src="docs/logo.svg" alt="Codelo" width="640"/>
</p>

<p align="center">
  <strong>Tiếng Việt</strong>
  &nbsp;·&nbsp;
  <a href="./README.md">English</a>
  &nbsp;·&nbsp;
  <a href="./README.zh-CN.md">简体中文</a>
</p>

<h3 align="center">Codelo — Trợ lý lập trình AI cho terminal của bạn.</h3>
<p align="center">Fork từ Reasonix, Việt hóa toàn bộ và mở rộng với hệ sinh thái Kilo. Tối ưu DeepSeek prefix-cache để chi phí token luôn thấp.</p>

## Cài đặt

Yêu cầu Go ≥ 1.25. Hỗ trợ macOS · Linux · Windows.

```bash
git clone https://github.com/vanhieuca6795/codelovh.git
cd codelovh
make build
./bin/codelo chat
```

## Bắt đầu nhanh

```bash
./bin/codelo chat                           # Phiên tương tác
./bin/codelo run "viết hàm kiểm tra SNT"     # Chạy tác vụ
./bin/codelo setup                          # Thiết lập ban đầu
./bin/codelo serve                          # Mở web SPA
CODELO_LANG=vi ./bin/codelo chat            # Giao diện tiếng Việt
```

## Cấu hình

Tệp cấu hình: `./codelo.toml` hoặc `~/.config/codelo/config.toml`.

```toml
default_model = "deepseek"
language = "vi"

[[providers]]
name = "deepseek"
kind = "openai"
base_url = "https://api.deepseek.com"
models = ["deepseek-v4-flash", "deepseek-v4-pro"]
api_key_env = "DEEPSEEK_API_KEY"

[agent]
system_prompt = "Bạn là Codelo, trợ lý lập trình AI."
```

## Tính năng

| Tính năng | Codelo |
|-----------|--------|
| Giao diện tiếng Việt | Toàn bộ CLI + Web SPA |
| Đa mô hình | DeepSeek + Anthropic + OpenAI |
| Plugin system | Stdio + HTTP transports |
| Hooks | PreToolUse, PostToolUse, lifecycle |
| Multi-agent | Coordinator + task orchestration |
| Web SPA | Zero-dependency, nhúng Go binary |
| Cache-first loop | 99%+ cache hit với DeepSeek |
| LSP tools | Go, Rust, Python, TypeScript |

## Các lệnh

| Lệnh | Mô tả |
|---|---|
| `/compact` | Nén ngữ cảnh |
| `/new` | Tạo phiên mới |
| `/resume` | Tiếp tục phiên đã lưu |
| `/model` | Đổi mô hình |
| `/mcp` | Quản lý MCP server |
| `/hooks` | Quản lý hook |
| `/skill` | Quản lý skill |
| `/init` | Khởi tạo AGENTS.md |
| `/quit` | Thoát |

## Giấy phép

MIT — xem [LICENSE](./LICENSE). Fork từ [Reasonix](https://github.com/esengine/DeepSeek-Reasonix).