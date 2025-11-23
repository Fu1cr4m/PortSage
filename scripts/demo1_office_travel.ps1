# 办公室 → 锁屏 → 解锁 → 旅途 场景流（控制面占位：/status 可用）
Start-Process -FilePath .\bin\ps-svc.exe -ArgumentList '-policy assets/policy/sample_office_travel.yaml' -NoNewWindow
Start-Sleep -s 2
Invoke-WebRequest -UseBasicParsing http://127.0.0.1:7777/status | Select-Object -ExpandProperty Content
Write-Host "[DEMO] 切换场景请在后续版本通过 /context API 演示。"
