# codeSnippets
usefull code

## Install ZSH Completions for Prezto
```
curl -fLo ~/.zprezto/modules/completion/external/src/_docker \
  https://raw.github.com/felixr/docker-zsh-completion/master/_docker
exec zsh
```
## ssh kegen
`ssh-keygen -o -a 100 -t ed25519 -f <name>`

## vim config
```
" Name
" Colors {{{
syntax enable           " enable syntax processing
colorscheme molokai
" }}}
" Spaces & Tabs {{{
set tabstop=2           " 4 space tab
set expandtab           " use spaces for tabs
set softtabstop=2       " 4 space tab
set shiftwidth=2
set modelines=1
filetype indent on
filetype plugin on
set autoindent
" }}}
```

## Access the boot2docker vm on osx

`screen ~/Library/Containers/com.docker.docker/Data/com.docker.driver.amd64-linux/tty`

## Build Docker Image behind Proxy
`--build-arg http_proxy=<proxy-server> --build-arg https_proxy=<proxy-server>`
