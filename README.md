# 🪟 FODHelper UAC Bypass

Simple UAC bypass using FODHelper registry hijacking.

## How it works

Modifies `HKLM\SOFTWARE\Classes\ms-settings\Shell\Open\command` to execute `cmd.exe` when `fodhelper.exe` runs.

## Build

```bash
go build -o fodhelper.exe fodhelper.go
```
##Usage
```
fodhelper.exe
```
