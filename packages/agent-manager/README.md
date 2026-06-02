# Codelo Agent Manager

Quản lý nhiều agent session song song với git worktree isolation.

## Cấu trúc

```
.kilo/worktrees/
├── feature-a/
├── bugfix-b/
├── agent-manager.json
└── run-script
```

## Sử dụng

```bash
codelo worktree create feature-login "triển khai trang đăng nhập"
codelo worktree list
codelo worktree apply feature-login
codelo worktree pr feature-login --title "Thêm trang đăng nhập"
```