syntax match typ_func /#\?\w\+(\@=/
highlight typ_func guifg=#4b69c6

syntax region typ_heading start=/^= / end=/$/
highlight typ_heading guifg=#19181f gui=bold

syntax match typ_num /\d\+%\?/
highlight typ_num guifg=#b60157

syntax region typ_str start=/"/ end=/"/
highlight typ_str guifg=#298e0d
