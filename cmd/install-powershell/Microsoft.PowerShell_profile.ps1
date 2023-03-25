## 2023-03-24
$env:PATH = @(
   'C:\Users\Steven\go\bin'
   'C:\python'
   'C:\python\Scripts'
   'D:\MinGit\mingw64\bin'
   'D:\bin'
   'D:\c\bin'
   'D:\go\bin'
   'D:\rust\bin'
   'D:\vim'
) -Join ';'

## 2023-03-19
Set-PSReadLineOption -AddToHistoryHandler $null

## 2023-03-06
Get-Alias | Remove-Alias -Force

## 2023-02-04
$env:CGO_ENABLED = 1

## 2022-04-18
$env:LESS = -join @(
   # Quit if entire file fits on first screen.
   'F'
   # Output "raw" control characters.
   'R'
   # Don't use termcap init/deinit strings.
   'X'
   # Ignore case in searches that do not contain uppercase.
   'i'
)

## 2022-04-17
Set-PSReadLineKeyHandler Ctrl+UpArrow {
   Set-Location ..
   [Microsoft.PowerShell.PSConsoleReadLine]::InvokePrompt()
}

## 2022-04-15
# color output
[Console]::OutputEncoding = [System.Text.UTF8Encoding]::new()

## 2022-04-12
$env:RIPGREP_CONFIG_PATH = 'C:\Users\Steven\_ripgrep'

## 2022-04-11
$env:EDITOR = 'gvim'
