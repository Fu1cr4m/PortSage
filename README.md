# PortSage v2 🛡️

<p align="center">
  <img src="https://img.shields.io/badge/Platform-Windows-blue?style=for-the-badge&logo=windows" alt="Platform">
  <img src="https://img.shields.io/badge/Language-Go-00ADD8?style=for-the-badge&logo=go" alt="Go">
  <img src="https://img.shields.io/badge/UI-Fyne-purple?style=for-the-badge" alt="Fyne">
</p>

> **PortSage** 是一款现代化的 Windows USB 设备安全管理工具。V2 版本采用 Fyne 构建了原生 GUI 界面，支持对每一个 USB 设备进行独立的启用、禁用和安全管理。

## ✨ 新版特性 (v2.0)

*   **🖥️ 图形化界面 (GUI)**: 抛弃繁琐的命令行，提供直观的鼠标交互体验。
*   **🎮 精细化控制 (Per-Device Control)**: 支持**单独禁用**某一个 U 盘或摄像头，而不影响其他设备（基于 Windows PnP 管理）。
*   **📝 审计日志 (Audit Logs)**: 实时记录所有禁用、启用和扫描操作，便于安全审计。
*   **⚠️ 风险可视化**: 通过图标颜色（红/橙/绿）直观展示设备风险等级。

## 🚀 下载与运行

### 1. 下载
前往 [Releases](https://github.com/Fu1cr4m/PortSage/releases) 页面下载最新的 `portsage_v2.exe`。

### 2. 运行
⚠️ **注意**: 由于涉及底层硬件控制，**必须以管理员身份运行**。
*   右键点击 `portsage_v2.exe` -> 选择 **"以管理员身份运行"**。

## 🛠️ 编译指南

如果你想自己编译源码：

1.  安装 **Go** (1.20+) 和 **TDM-GCC (64-bit)**。
2.  克隆仓库并编译：
    ```powershell
    git clone https://github.com/Fu1cr4m/PortSage.git
    cd PortSage
    go mod tidy
    # 编译并隐藏控制台窗口
    go build -ldflags "-s -w -H=windowsgui" -o portsage_v2.exe ./cmd/app/
    ```

## ⚠️ 免责声明
本工具修改系统 PnP 设备状态，仅供安全研究和授权管理使用。
