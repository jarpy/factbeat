$workdir = Split-Path $MyInvocation.MyCommand.Path

if (Get-Service factbeat -ErrorAction SilentlyContinue) {
  $service = Get-WmiObject -Class Win32_Service -Filter "name='factbeat'"
  $service.StopService()
  Start-Sleep -s 1
  $service.delete()
}

New-Service -name factbeat `
  -displayName factbeat `
  -binaryPathName "`"$workdir\\factbeat.exe`" -c `"$workdir\\factbeat.yml`""
