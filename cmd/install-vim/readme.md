# Install Vim

## Syntax highlighting only

I only care about the highlighting, nothing else. So I got this file:

https://github.com/fatih/vim-go/blob/master/syntax/go.vim

and copy to my local Vim install. However I get these errors:

~~~
Unknown function: go#config#FoldEnable
~~~

can you please configure this package such that people can just drop in the
`go.vim` only and be done with it?

## switch `nil` and `iota` away from booleans

(sorry for new issue, but previous issue was quickly closed, signalling that
further discussion is not welcome)

currently `nil` and `iota` are defined as `goPredefinedIdentifiers`:

https://github.com/fatih/vim-go/blob/3d16fa33/syntax/go.vim#L53

which is then linked to `goBoolean`:

https://github.com/fatih/vim-go/blob/3d16fa33/syntax/go.vim#L57

and then highlighted as `Boolean`:

https://github.com/fatih/vim-go/blob/3d16fa33/syntax/go.vim#L56

in some other code, I have seen it declared like this:

~~~vim
syn keyword     goConstants         iota true false nil
hi def link     goConstants         Keyword
~~~

I think this is a sensible option, as `nil` and `iota` are not booleans.
