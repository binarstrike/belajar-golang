Param([String]$GoFileParam)

If ([String]::IsNullOrEmpty($GoFileParam)) {
  Write-Host "Usage: .\runAndWatch.ps1 <path to go file>"
  exit 1
}

If (-Not (Test-Path $GoFileParam)) {
  Write-Host "Can't find go file, $GoFileParam no such file or directory"
  exit 1
}

If (-Not (Get-Command nodemon -ErrorAction SilentlyContinue)) {
  Write-Host "nodemon command does not exist"
  exit 1
}

$ScriptLocation = $PSScriptRoot

$BuildDir = Join-Path $ScriptLocation "build"

If (-Not (Test-Path $BuildDir -ErrorAction SilentlyContinue)) {
  New-Item -ItemType Directory $BuildDir
}

$GoFile = Get-ChildItem $GoFileParam
$GoFileBasename = $GoFile | % {$_.BaseName} 
$GoFileFullName = $GoFile | % {$_.FullName}

$OutputBinary = Join-Path $BuildDir "$GoFileBasename.exe"
$NodemonConfigFilePath = Join-Path $ScriptLocation "nodemon.json"

# $Command = "nodemon --config $NodemonConfigFilePath --exec 'go run $GoFileFullName'"
$Command = "nodemon --config $NodemonConfigFilePath --exec 'go build -o $OutputBinary $GoFileFullName && $OutputBinary'"

Invoke-Expression $Command
