source /etc/prtstuff/config

complete -c prtpull -x -a "(ls $portdir)"
complete -c prtpull -f -o h -l help -d 'Print help and exit'
