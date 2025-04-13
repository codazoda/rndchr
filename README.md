# rndchr

A command line tool that generates a random set of characters.


## Syntax

```
rndchr [characters] [--lower]
```


## Examples

```
rndchr
O3IGFFQE

rndchr 16
QTDGLLW9CBCSJG7L

rndchr 32
KA2OEDN6RJBQ9Y2Y46S8QL79XE5SQ90P

rndchr 32 --lower
zd31w2kpi4ob6z0xybjzvurxyl8jkpfh

rndchr --lower
59itqusf
```


## Compiling

Clone the repo and run the following command line to generate a binary for your platform.

```
go build
```


## Installing

You can also install the binary into the default go bin directory with the following command.

```
go install
```
