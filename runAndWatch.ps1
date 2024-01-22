Param([string]$FileDir)

if ([string]::IsNullOrEmpty($FileDir)) {
  Write-Host "Usage: .\runAndWatch.ps1 [-File] <file or directory>"
  exit 1
}

if (-not (Test-Path $FileDir)) {
  Write-Host "$FileDir no such file or directory"
  exit 1
}

if (-not (Get-Command nodemon -ErrorAction SilentlyContinue)) {
  Write-Host "nodemon command does not exist"
  exit 1
}

$BuildTemp = Join-Path $env:TEMP "build"

if (-not (Test-Path $BuildTemp)) {
  [void](New-Item -ItemType Directory $BuildTemp)
}

$FileDirPath = Get-ChildItem $FileDir

if (Test-Path $FileDir -PathType Leaf) {
  $BinaryOutputName = Join-Path $BuildTemp ($FileDirPath.BaseName + ".exe")
  $InputFileDir = $FileDirPath.FullName
} else {
  $BinaryOutputName = Join-Path $BuildTemp "main.exe"
  $InputFileDir = $FileDirPath.Directory.FullName + "\..."
}

$Command = "nodemon -e go,html,css,js,json -w {0} --exec 'go build -o {1} {2} && {1}'" -f $FileDirPath.Directory.FullName, $BinaryOutputName, $InputFileDir

$Command | Invoke-Expression