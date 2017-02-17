[![Go Report Card](https://goreportcard.com/badge/github.com/onodera-punpun/prt)](https://goreportcard.com/report/github.com/onodera-punpun/prt)

*This project is currently in early alpha stage!*

prt - CRUX port utility written in Go, aiming to replace prt-get, ports, and some pkgutils (on my machine)


## SYNOPSIS

prt command [arguments]


## DESCRIPTION

prt is like `prt-get(8)` a port/package management utility which provides additional functionality to the CRUX pkgutils. It works with the local ports tree and is therefore fully compatible with `ports(8)`, `pkgmk(8)`, `pkgadd(8)` and of course `prt-get(8)`. It offers the following features:

* listing dependencies of ports recursively, with an optional flag to print using tree view
* listing outdated package, by comparing port versions with the installed version
* easily printing port information such as the maintainer, version, release, et cetera
* install ports and their dependencies with a single command
* list ports, with optional flags to also only list installed ports, print with repo information or to print with additional version information
* print the location of a port
* searching through files ports provide, with an optional flag to only search through installed ports
* pull in ports using git(1)
* update outdated packages
* uninstall installed packages

like `prt-get(8)`, prt is basically a wrapper around `pkgmk(8)`/`pkgadd(8)` and provides some nice functionality such as listing and installing dependencies, getting the location of a port, aliasing ports (for example `core/openssl` to `6c37-dropin/libressl`), and ordering ports with the same name depending on how "important" the repo is the port resides in.

There are a few differences though, for example, unlike `prt-get(8)` you need to  be in the port's directory for most commands to work, like how `pkgmk(8)` works. This has a few advantages, for example you can quickly download a port
anywhere on the filesystem, and install it and its dependencies using `prt install`. Because `prt-get depinst` needs a port name, you can *only* install ports that are located in a predefined `prtdir`.

Another difference with `prt-get(8)` is that prt does not use a cache file, while still being nearly as fast or faster in some cases.

Aliasing is also handeled a bit different. `prt-get(8)` aliases ports based on name, but prt on name and repo. This makes it possible to alias `foo/bar` to `baz/bar`.


## COMMANDS

The prt syntax is inspired by `prt-get(8)`, `git(8)` and `go(8)`, and thus uses so called commands which always have to be the first non-option argument passed. The commands are:

`depends`   list dependencies recursively,

`diff`      list outdated packages

`info`      print port information

`install`   build and install ports and their dependencies

`list`      list porst and packages

`loc`       print port locations

`prov`      search ports for files

`pull`      pull in ports

`sysup`     update outdated packages

`uninstall` uninstall packages

`help`      print help and exit


## INSTALLATION

https://github.com/onodera-punpun/crux-ports/blob/master/prt/Pkgfile

Make sure to check `/etc/prt/config.toml` after installation and edit values to fit your needs and setup.

If you use `fish` a `cd` wrapper for `prt loc` will also be installed, and some handy completions.


## TODO

- [x] Implement `depends` command.
- [x] Implement `diff` command.
- [x] Implement `info` command.
- [x] Implement `install` command.
- [x] Implement `list` command.
- [x] Implement `loc` command.
- [ ] Implement `patch` command.
- [x] Implement `prov` command.
- [x] Implement `pull` command.
- [x] Implement `sysup` command. *(needs some improvements when it comes to installation order)*
- [x] Implement `uninstall` command.

---

- [x] Convert `pkgmk` `get_filename` function to Go. *(So uhh, pkgmk does something with "absolute paths", do I need this as well?)*
- [x] Convert `pkgmk` `get_basename` function to Go.
- [ ] Convert `pkgmk` `check_pkgfile` function to Go.
- [ ] Convert `pkgmk` `check_directory` function to Go.
- [ ] Convert `pkgmk` `check_file` function to Go.
- [x] Convert `pkgmk` `download_file` function to Go.
- [x] Convert `pkgmk` `download_source` function to Go.
- [x] Convert `pkgmk` `unpack_source` function to Go.
- [x] Convert `pkgmk` `make_md5sum` function to Go.
- [ ] Convert `pkgmk` `make_footprint` function to Go.
- [x] Convert `pkgmk` `check_md5sum` function to Go.
- [ ] Convert `pkgmk` `check_signature` function to Go.
- [ ] Convert `pkgmk` `make_signature` function to Go.
- [ ] Convert `pkgmk` `strip_files` function to Go.
- [ ] Convert `pkgmk` `compress_manpages` function to Go.
- [ ] Convert `pkgmk` `check_footprint` function to Go.
- [x] Convert `pkgmk` `make_work_dir` function to Go.
- [x] Convert `pkgmk` `remove_work_dir` function to Go.
- [ ] Convert `pkgmk` `install_package` function to Go.
- [x] Convert `pkgmk` `clean` function to Go. *(not going to implement, there is `rm`)*

---

- [x] Write fish `cdp` function.
- [x] Write bash `cdp` function. *(need to actually test this)*
- [x] Write fish completions.
- [ ] Write bash completions.

---

- [ ] Write tests.
- [x] Write README and man pages. *(needs some updates with changes)*


## AUTHORS

Camille Scholtz


## NOTES

Since this is my first Go project I'm probably making some mistakes, feedback is highly appreciated!
