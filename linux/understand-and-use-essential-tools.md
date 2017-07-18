# Understand and use essential tools
## Access a shell prompt and issue commands with correct syntax
## Use input-output redirection (>, >>, |, 2>, etc.)
  ```
  1>file #stdout to file
  1>>file # append stdout to file
  2>file, 2>>file # stderr to file
  &>file # stdout & stderr to file
  2>&1 # stderr to stdout
  ```
## Use grep and regular expressions to analyze text
  `grep <search-string> <file/stdin>`
  `grep 'hello\.gif' file` # escape the . (wildcard)
  
## Log to journal
`logger <my-msg>`
## Access remote systems using ssh
## Log in and switch users in multiuser targets
## Archive, compress, unpack, and uncompress files using tar, star, gzip, and bzip2
## Create and edit text files
## Create, delete, copy, and move files and directories
## Create hard and soft links
## List, set, and change standard ugo/rwx permissions
## Locate, read, and use system documentation including man, info, and files in /usr/share/doc

