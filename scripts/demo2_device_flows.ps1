# RNDIS / 存储 / 键盘 处置演示（占位：当前展示状态查询）
Invoke-WebRequest -UseBasicParsing http://127.0.0.1:7777/status | Select-Object -ExpandProperty Content
Write-Host "[DEMO] 设备事件与审计将在实现设备订阅后展示。"
