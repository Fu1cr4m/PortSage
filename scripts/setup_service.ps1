param(
  [string]$BinPath = "$(Resolve-Path .\bin\ps-svc.exe)",
  [string]$Policy = "assets\policy\sample_office_travel.yaml",
  [string]$Name = "PortSage"
)
$cmd = "`"$BinPath`" -policy `"$Policy`""
sc.exe create $Name binPath= $cmd start= auto
sc.exe description $Name "PortSage USB risk-aware service"
sc.exe start $Name
sc.exe query $Name
