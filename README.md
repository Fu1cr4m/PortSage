# PortSage 🛡️

<p align="center">
  <img src="https://img.shields.io/badge/Platform-Windows-blue?style=for-the-badge&logo=windows" alt="Platform">
  <img src="https://img.shields.io/badge/Language-Go-00ADD8?style=for-the-badge&logo=go" alt="Go">
  <img src="https://img.shields.io/badge/License-MIT-green?style=for-the-badge" alt="License">
  <img src="https://img.shields.io/badge/UI-TUI-purple?style=for-the-badge" alt="TUI">
</p>

> **PortSage** 是一款专为 Windows 打造的轻量级 USB 端口安全审计与管理工具。它采用原生 Win32 API 构建，摒弃了繁重的第三方驱动依赖，提供黑客风格的终端交互界面（TUI），让 USB 安全管理变得直观且高效。

---

## ✨ 功能特性 (Features)

*   **🛡️ 零依赖 (Native Win32)**: 直接调用 Windows WMI 和 Registry API，无需安装 LibUSB 或 Zadig 驱动，开箱即用。
*   **🔍 智能风险评估 (Risk Heuristics)**: 内置启发式算法，实时识别伪装成键盘的 BadUSB 攻击设备及未经授权的存储介质。
*   **🔒 全局锁定 (Global Lockdown)**: 一键禁用 `USBSTOR` 服务，在物理层面上阻断数据通过 USB 存储设备泄露。
*   **🖥️ 沉浸式 UI (Cyberpunk TUI)**: 基于 `Bubbletea` 打造的现代化终端界面，支持全键盘操作。

## 🚀 快速开始 (Getting Started)

### 依赖环境
*   **OS**: Windows 10/11 (x64)
*   **权限**: 需要 **管理员权限 (Run as Administrator)** 以执行注册表锁定操作。

### 安装 (Installation)

#### 方式 1: 直接下载二进制文件
前往 [Releases](https://github.com/Fu1cr4m/PortSage/releases) 页面下载最新的 `portsage.exe`。

#### 方式 2: 源码编译
如果你安装了 Go (1.20+):

```powershell
git clone https://github.com/Fu1cr4m/PortSage.git
cd PortSage
go mod tidy
go build -o portsage.exe ./cmd/app/main.go

```

## 🎮 使用指南 (Usage)

右键以 **管理员身份** 运行 portsage.exe。

| 按键           | 功能描述                                          |
| -------------- | ------------------------------------------------- |
| **Space**      | **切换锁定状态**：开启或禁用 Windows USB 存储服务 |
| **R**          | **刷新**：重新扫描硬件总线上的设备变动            |
| **Q** / Ctrl+C | **退出**：关闭程序                                |

### 界面状态说明

- **🟢 SAFE**: 已验证的驱动或受信任设备。
- **🟡 LOW**: 普通存储设备（存在数据泄露风险）。
- **🔴 HIGH**: 疑似 BadUSB 注入设备或驱动异常的输入设备。

## 🛠️ 技术架构 (Architecture)

- **Core**: Golang w/ x/sys/windows & wmi
- **UI Framework**: [Bubbletea](https://www.google.com/url?sa=E&q=https%3A%2F%2Fgithub.com%2Fcharmbracelet%2Fbubbletea) (MVU pattern)
- **Security Logic**: 注册表键值控制 (SYSTEM\CurrentControlSet\Services\USBSTOR)

## ⚠️ 免责声明 (Disclaimer)

本工具仅供安全研究和授权的系统管理使用。修改 Windows 注册表存在一定风险，作者不对因使用本工具导致的任何系统故障或数据丢失承担责任。
=======
