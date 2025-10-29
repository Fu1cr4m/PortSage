# PortSage (端智) – Windows USB Risk-Aware Orchestration

一套以 **运行时可观测（ETW/RawInput）+ 环境上下文（网络/会话/时间）+ 自适应策略** 为核心的 **Windows USB 安全系统**。

## 目录
- 服务：`cmd/ps-svc`（Windows Service/也可控制台运行）  
- CLI：`cmd/psctl`（策略管理、场景切换、临时放行）  
- 策略：`assets/policy/*.yaml`（YAML DSL）  
- 脚本：`scripts/*.ps1`（演示与初始化）

## 快速开始（本地运行）
1. 安装 Go 1.22+（Windows）。
2. 打开 **PowerShell**（管理员），执行：
   ```powershell
   go build -o bin\ps-svc.exe ./cmd/ps-svc
   go build -o bin\psctl.exe ./cmd/psctl
   ./bin/ps-svc.exe -policy assets/policy/sample_office_travel.yaml
   ```
   另起窗口执行 `./bin/psctl.exe -cmd status` 查看状态。
3. 安装为服务（可选）：见 `scripts/setup_service.ps1`。

## GitHub 同步（Git Bash）
参见 `scripts/git_init_push.sh`。

## 免责声明
本仓库为用户态研究与工程样例，不包含内核驱动。请在测试环境验证后再用于生产。
