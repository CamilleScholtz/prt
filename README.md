[![Go Report Card](https://goreportcard.com/badge/github.com/onodera-punpun/prt)](https://goreportcard.com/report/github.com/onodera-punpun/prt)

# prt

CRUX port utility writtin in go, aiming to replace prt-get, ports, and some pkgutils (on my machine). 


## Difference from `prt-get`

`prt` differs from `prt get` in a few ways:

* It doesn't make use of a cache file, but still tries to be nearly as fast as `prt-get`.
  For example, `prt-get deptree` takes around `0.02` seconds on my machine, `prt depends -ta`
  takes around `0.04` seconds. `prt-get fsearch *.h` takes around `0.34` seconds, and `prt-get prov *.h`
  takes around `0.23` seconds.

* `prt-get` takes a port name for most commands, for example `prt-get depinst portname`, `prt` does it more
   like `pkgmk`, so first you `cd` to `portname`, and then you run `prt install`. This has a few advantages, for
   example you can quickly `httpup sync` a port anywhere in the filesystem, and install it *and* its dependencies
   using `prt install`. Since `prt-get depinst` needs a port name, you can *only* install ports that are located in a
   predefined `prtdir`.

* `prt` has some pretty colors you can customize using `config.toml`.

* Pretty much everything port related is defined in `config.toml`. The git repos you want to pull (if you choose to use
  `prt pull` over `ports -u`), repo ordering, aliasing, et cetera.

* I tried to keep `prt` pretty minimal. `prt-get` is around 7500 lines of C++, `prt` is around 2000 lines of go.

* `prt-get` aliases ports based on name, `prt` on name and repo. This makes it possible to alias `foo/bar` to `baz/bar`.


## Installation

https://github.com/onodera-punpun/crux-ports/blob/master/prt/Pkgfile

Make sure to check `/etc/prt/config.toml` after Installation and edit values to fir your needs and setup.

If you use `fish` a `cd` wrapper for `prt loc` will also be installed, and some handy completions.


## Todo

* Add `prt patch` command, this will patch ports using user created diff files. This removes the need
  to fork ports for minor changes.

* Add `prt remove` command, pretty self explanatory.

* Add `prt rebuild` command, again, pretty self explanatory.

* Add `prt maildiff`, what this is basically going to do is generate a diff with changes the user made to a port, 
  and semi-automatically mail it to the maintainer.

* Handle flags differently maybe.

* Handle commands differently maybe.

* The config gets loaded multiple times now I think (for example first by `commands.go`, but if `ports` get called also by `ports.go`).
  This isn't really a problem but it just annoys me.


## Notes

Since this is my first go project I'm probably making some mistakes, feedback is highly appreciated!
