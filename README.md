# Password Generator

Simple password generator

## Usage

```
$ pg [-c] [-C] [-d <digit>] [-s]
```

### Number only

```
$ pg
$ 60023148
```

### Digit

```
$ pg -d 10
$ 1452707794 // 10-digit password will be generated.
```

### Lowercase letters

```
$ pg -c
$ gu0n81y3
```

### Uppercase letters

```
$ pg -C
$ QCS529W6
```

### Symbol

```
$ pg -s
$ 9#936!&%
```

### Combine

```
$ pg -cCs -d 10
$ VBW6h$v%W!
```

## Install

```sh
$ go install github.com/kyamashiro/pg@latest
$ pg
```

```sh
$ git clone https://github.com/kyamashiro/pg.git
$ make build
$ ./pg
```
