complete -c prt -n '__fish_use_subcommand' -a depends -d 'list dependencies recursivly'
complete -c prt -n '__fish_seen_subcommand_from depends' -o a -l all -d 'also list installed dependencies'
complete -c prt -n '__fish_seen_subcommand_from depends' -o n -l no-alias -d 'disable aliasing'
complete -c prt -n '__fish_seen_subcommand_from depends' -o t -l tree -d 'list using tree view'
complete -c prt -n '__fish_seen_subcommand_from depends' -o h -l help -d 'print help and exit'

complete -c prt -n '__fish_use_subcommand' -a diff -d 'list oudated packages'
complete -c prt -n '__fish_seen_subcommand_from diff' -o v -l version -d 'list installed and availabe version'
complete -c prt -n '__fish_seen_subcommand_from diff' -o h -l help -d 'print help and exit'

complete -c prt -n '__fish_use_subcommand' -a build -d 'build and install ports'
complete -c prt -n '__fish_seen_subcommand_from build' -o s -l script -d 'toggle execution of pre- and post-install'
complete -c prt -n '__fish_seen_subcommand_from build' -o v -l verbose -d 'enable verbose output'
complete -c prt -n '__fish_seen_subcommand_from build' -o h -l help -d 'print help and exit'

complete -c prt -n '__fish_use_subcommand' -a info -d 'print port information'
complete -c prt -n '__fish_seen_subcommand_from info' -o d -l description -d 'print description'
complete -c prt -n '__fish_seen_subcommand_from info' -o u -l url -d 'print url'
complete -c prt -n '__fish_seen_subcommand_from info' -o m -l maintainer -d 'print maintainer'
complete -c prt -n '__fish_seen_subcommand_from info' -o e -l depends -d 'print dependencies'
complete -c prt -n '__fish_seen_subcommand_from info' -o o -l optional -d 'print optional dependencies'
complete -c prt -n '__fish_seen_subcommand_from info' -o v -l version -d 'print version'
complete -c prt -n '__fish_seen_subcommand_from info' -o r -l release -d 'print release'
complete -c prt -n '__fish_seen_subcommand_from info' -o h -l help -d 'print help and exit'

complete -c prt -n '__fish_use_subcommand' -a list -d 'list repos and ports'
complete -c prt -n '__fish_seen_subcommand_from list' -o i -l installed -d 'list installed ports'
complete -c prt -n '__fish_seen_subcommand_from list' -o r -l repos -d 'list repos'
complete -c prt -n '__fish_seen_subcommand_from list' -o h -l help -d 'print help and exit'

complete -c prt -n '__fish_use_subcommand' -a location -d 'prints port locations'
complete -c prt -n '__fish_seen_subcommand_from location' -x -a "(prt list | cut -d '/' -f 2)"
complete -c prt -n '__fish_seen_subcommand_from location' -o d -l duplicate -d 'list duplicate ports as well'
complete -c prt -n '__fish_seen_subcommand_from location' -o n -l no-alias -d 'disable aliasing'
complete -c prt -n '__fish_seen_subcommand_from location' -o h -l help -d 'print help and exit'

complete -c prt -n '__fish_use_subcommand' -a patch -d 'patch ports'
complete -c prt -n '__fish_seen_subcommand_from patch' -o h -l help -d 'print help and exit'

complete -c prt -n '__fish_use_subcommand' -a provide -d 'search ports for files'
complete -c prt -n '__fish_seen_subcommand_from provide' -o h -l help -d 'print help and exit'

complete -c prt -n '__fish_use_subcommand' -a remove -d 'remove installed ports'

complete -c prt -n '__fish_use_subcommand' -a pull -d 'pull in ports'
complete -c prt -n '__fish_seen_subcommand_from location' -x -a "(prt list -r)"
complete -c prt -n '__fish_seen_subcommand_from pull' -o h -l help -d 'print help and exit'

complete -c prt -n '__fish_use_subcommand' -a sysup -d 'update outdated packages'
complete -c prt -n '__fish_seen_subcommand_from sysup' -o s -l script -d 'toggle execution of pre- and post-install'
complete -c prt -n '__fish_seen_subcommand_from sysup' -o v -l verbose -d 'enable verbose output'
complete -c prt -n '__fish_seen_subcommand_from sysup' -o h -l help -d 'print help and exit'

complete -c prt -n '__fish_use_subcommand' -a help -d 'print help and exit'
