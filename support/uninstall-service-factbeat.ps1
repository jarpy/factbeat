if (Get-Service factbeat -ErrorAction SilentlyContinue) {
  $service = Get-WmiObject -Class Win32_Service -Filter "name='factbeat'"
  $service.delete()
}
