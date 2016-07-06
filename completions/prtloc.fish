complete -c prtloc -x -a "(prtls | cut -d '/' -f 2)"
complete -c prtloc -f -o d -l duplicate -d 'list duplicate ports as well'
complete -c prtloc -f -o h -l help -d 'Print help and exit'
