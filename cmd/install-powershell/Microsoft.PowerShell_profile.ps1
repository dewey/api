# 2022-04-12
$env:RIPGREP_CONFIG_PATH = 'C:\Users\Steven\Documents\ripgrep.txt'

# 2023-05-10
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

# 2023-05-10
Set-PSReadLineKeyHandler Ctrl+UpArrow {
   Set-Location ..
   [Microsoft.PowerShell.PSConsoleReadLine]::InvokePrompt()
}

# 2023-05-10
# `git diff` unicode
[Console]::OutputEncoding = [System.Text.UTF8Encoding]::new()

# 2023-05-10
# git commit -v
$env:EDITOR = 'gvim'

# 2023-03-24
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

# 2023-03-19
Set-PSReadLineOption -AddToHistoryHandler $null

# 2023-03-06
Get-Alias | Remove-Alias -Force

# 2023-02-04
$env:CGO_ENABLED = 1
