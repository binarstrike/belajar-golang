<#
  Sebuah script powershell simpel yang digunakan untuk build kode program
  
  contoh penggunaan:
  .\build.ps1 -output foo.exe
#>
param([string]$output=".\main.exe")

$mainPackagePath = "$PSScriptRoot\..."
$version = $(Get-Content $PSScriptRoot\version.txt)

$buildCommand = "go build -o $output -ldflags '-X main.versionString=$version' $mainPackagePath"

Write-Host "building..."
Write-Host $buildCommand

Invoke-Expression $buildCommand
