# test-agent-orchestrator

用于测试 [agent-orchestrator](https://github.com/ComposioHQ/agent-orchestrator) 的实验项目。

## 配置

`agent-orchestrator.yaml`：

```yaml
repo: yinzhezq/test-agent-orchestrator
defaultBranch: main

runtime: tmux
agent: claude-code
workspace: worktree
```

- **runtime**：`tmux`，每个 agent 在独立的 tmux 会话中运行
- **agent**：`claude-code`，使用 Claude Code 作为执行 agent
- **workspace**：`worktree`，每个任务在独立的 git worktree 中隔离执行

## 用途

验证 agent-orchestrator 在以下场景下的行为：

- 多 agent 并行调度
- worktree 隔离与合并
- tmux 会话管理
- Claude Code 与 orchestrator 的集成
