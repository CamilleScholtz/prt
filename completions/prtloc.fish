complete -c prtloc -x -a (ports -l | cut -d '/' -f 2)
complete -c prtloc -f -o h -l help -d 'Print help and exit'
