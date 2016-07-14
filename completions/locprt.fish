complete -c locprt -x -a "(lsprt | cut -d '/' -f 2)"
complete -c locprt -f -o d -l duplicate -d 'List duplicate ports as well'
complete -c locprt -f -o n -l no-alias -d 'Disable aliasing'
complete -c locprt -f -o h -l help -d 'Print help and exit'
