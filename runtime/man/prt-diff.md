# prt-diff 8 "2017-01-02" prt "General Commands Manual"

## NAME

prt-diff - list outdated packages


## SYNOPSIS

prt diff [arguments]


## DESCRIPTION

prt diff lists outdated packages, it does does by comparing the version and release of the installed package
against the available version.


## OPTIONS

`-n` disable aliasing. Without this option ports get aliased using values found in `prt.toml(5)`. Same as `--no-alias`

`-v` print with version information. Same as `--version`


## EXAMPLES

List outdated packages with version information:

```
$ prt diff -v
libpng 1.6.26-1 -> 1.6.27-1
fish git-4 -> git-5
```


## AUTHORS

Camille Scholtz


## SEE ALSO

`cdp(1)`, `prt(8)`, `prt.toml(5)`, `prt-depends(8)`, `prt-info(8)`, `prt-install(8)`, `prt-list(8)`, 
`prt-loc(8)`, `prt-prov(8)`, `prt-pull(8)`, `prt-sysup(8)`, `prt-uninstall(8)`, `prt-get(8)`,
`pkgmk(8)`, `pkgrm(8)`, `pkgadd(8)`, `ports(8)`, `pkginfo(8)`, `prt-utils(1)`
