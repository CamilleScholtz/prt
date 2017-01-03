# prt-loc 8 "2017-01-02" prt "General Commands Manual"

## NAME

prt-loc - print port locations


## SYNOPSIS

prt loc [arguments] [ports]


## DESCRIPTION

prt loc prints the location of a ports, if there are multiple matches (for example 
`foo/bar` and `baz/bar`) it will print the most "important" ports using the order
value in `prt.toml(5)`.


## OPTIONS

`-d` list duplicate as well. Without this flag only the most "important" port will be listed. Same as `--duplicate`

`-n` disable aliasing. Without this option ports get aliased using values found in `prt.toml(5)`. Same as `--no-alias`


## EXAMPLES

Get location of `mpv`:

```
$ prt loc mpv
6c37-git/mpv
```

Get location of `mpv` and `fish`, list duplicate ports as well:

```
$ prt loc -d mpv fish
6c37-git/mpv
-  contrib/mpv
6c37-git/fish
-  6c37/fish
```


## AUTHORS

Camille Scholtz


## SEE ALSO

`cdp(1)`, `prt(8)`, `prt.toml(5)`, `prt-depends(8)`, `prt-diff(8)`, `prt-info(8)`, `prt-install(8)`, 
`prt-list(8)`, `prt-prov(8)`, `prt-pull(8)`, `prt-sysup(8)`, `prt-uninstall(8), prt-get(8)`,
`pkgmk(8)`, `pkgrm(8)`, `pkgadd(8)`, `ports(8)`, `pkginfo(8)`, `prt-utils(1)`
