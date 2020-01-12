## Build
```
cd src/solution
go build -o ../../bin/solution
```

## Run tests
```
cd src/solution
go test ./...
```

## Usage example

### 1. First create files and place them in same directory with binary
```
$ cat a.txt
1
2
3

$ cat b.txt
2
3
4

$ cat c.txt
3
4
5
```

### 2. Run binary with expression as an argument

    ./solution [ SUM [ DIF a.txt b.txt c.txt ] [ INT b.txt c.txt ] ]

Expected output:
```
1
3
4
```