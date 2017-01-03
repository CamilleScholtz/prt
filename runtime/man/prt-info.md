# prt-info 8 "2017-01-02" prt "General Commands Manual"

## NAME

prt-info - print port information


## SYNOPSIS

prt info [arguments]


## DESCRIPTION

prt info prinst port information, basically an alias for `grep something Pkgfile`.


## OPTIONS

`-d` print description. Same as `--description`

`-u` print url. Same as `--url`

`-m` print maintainer. Same as `--maintainer`

`-e` print dependencies. Same as `--depends`

`-o` print optional dependencies. Same as `--optional`

`-v` print version. Same as `--version`

`-r` print release. Same as `--release`


## EXAMPLES

Print information:

```
$ prt info
Description: CRUX port utils written in Go.
URL: https://github.com/onodera-punpun/prt
Maintainer: onodera, https://github.com/onodera-punpun/crux-ports/issues
Depends on: go
Nice to have: fish
Version: git
Release: 1
```

Only print maintainer, version and release:

```
prt info -r -v -m
Maintainer: onodera, https://github.com/onodera-punpun/crux-ports/issues
Version: git
Release: 1
```


## AUTHORS

Camille Scholtz


## SEE ALSO

`cdp(1)`, `prt(8)`, `prt.toml(5)`, `prt-depends(8)`, `prt-diff(8)`, `prt-install(8)`, `prt-list(8)`, 
`prt-loc(8)`, `prt-prov(8)`, `prt-pull(8)`, `prt-sysup(8)`, `prt-uninstall(8)`, `prt-get(8)`,
`pkgmk(8)`, `pkgrm(8)`, `pkgadd(8)`, `ports(8)`, `pkginfo(8)`, `prt-utils(1)`
