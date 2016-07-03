source /etc/prtstuff/config

complete -c prtprint -x -a "(ls $portdir)"
complete -c prtprint -f -o h -l help -d 'Print help and exit'
