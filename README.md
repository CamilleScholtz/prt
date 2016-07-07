prtstuff
========

Consitent CRUX port utilities writtin in fish, aiming to replace prt-get, ports, and pkgutils.


----


depls
=====

List dependencies recursively.


Help
----

```
Usage: depls [options]

options:
  -a,   --all             also list installed dependencies
  -t,   --tree            list using tree view
  -h,   --help            print help and exit
```


depmk
=====

Install dependencies recursivly.


Help
----

```
Usage: depmk [options]

options:
  -s,   --script          toggle execution of pre- and post-install
  -r,   --readme          toggle opening of readmes
  -h,   --help            print help and exit
```


prtloc
======

Prints port location.


Help
----

```
Usage: prtloc [options] [ports]

options:
  -d,   --duplicate       list duplicate ports as well
  -h,   --help            print help and exit
```


prtls
======

List repos and ports.


Help
----

```
Usage: prtls [options]

options:
  -r,   --repos           list repos
  -i,   --installed       list installed ports
  -h,   --help            print help and exit
```


prtpatch
========

Patches ports.


Help
----

```
Usage: prtpatch [ports]

options:
  -h,   --help            print help and exit
```


prtprint
========

Prints port information.


Help
----

```
Usage: prtprint [options]

options:
  -d,   --description     print description
  -u,   --url             print url
  -m,   --maintainer      print maintainer
  -v,   --version         print version
  -r,   --release         print release
  -h,   --help            print help and exit
```


prtprovide
==========

Search ports for files they provide.


Help
----

```
Usage: prtprovide [options] [query]

options:
  -h,   --help            print help and exit
```


prtpull
=======

Pull in ports using git.


Help
----

```
Usage: prtpull [options] [repos]

options:
  -h,   --help            print help and exit
```
