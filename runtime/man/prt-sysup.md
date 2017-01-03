# prt-sysup 8 "2017-01-02" prt "General Commands Manual"

## NAME

prt-sysup - update outdated packages


## SYNOPSIS

prt sysup [arguments] [ports to skip]


## DESCRIPTION

prt sysup builds and updates outdated packages, these are the same ports as reported
by `prt-diff(8)`. prt sysup also runs pre- and post-install scripts, and reports if the port as a README.


## OPTIONS

`-v` enable verbose output. Without this option stdout and stderr of `pkgmk(8)` and `pkgadd(8)` will
be redirected to `/dev/null`. Same as `--verbose`


## EXAMPLES

Update all outdated packages, but not `opt/libpng`:

```
$ prt diff
libpng
zlib

# prt sysup opt/libpng
Updating package 1/1, core/zlib.
-  Downloading sources
-  Unpacking sources
-  Building package
-  Updating package
```


## AUTHORS

Camille Scholtz


## SEE ALSO

`cdp(1)`, `prt(8)`, `prt.toml(5)`, `prt-depends(8)`, `prt-diff(8)`, `prt-info(8)`, `prt-install(8)`, 
`prt-list(8)`, `prt-loc(8)`, `prt-prov(8)`, `prt-pull(8)`, `prt-uninstall(8), prt-get(8)`,
`pkgmk(8)`, `pkgrm(8)`, `pkgadd(8)`, `ports(8)`, `pkginfo(8)`, `prt-utils(1)`
