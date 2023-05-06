# GO

## features of Go

    - Go is an opensource general purpose programming lanaguage
    - Go is a garbage collected programming language
    - Fast compilation
    - Statically typed
    - Concurrency is the most important feature of golang
    - Go is fast
    - Cross platform compilation
    - Go runtime is included with in the binary

## go(golang) code structure
 GOROOT
    bin
    pkg
    src

## what are different types of packages

    - standard packages         --> created by go lang folks
    - user defined packages     --> created by users
    - thirdparty packages       --> develped and made available by other people/orgs

GOROOT is used to compile user defined and thirdparty packages

## Keywords


break, case, ,chan,continue, const, default, else,fallthrough, for , func, go,goto, if, import, map,package,range,return,select, struct,switch,type,var 

## builtin

append,cap,close,complex,copy,delete,len,make,new,panic,println,print


## go tools

    - run go application : compile, link, build(generates bin) and then run
  
        ```go run hello.go```
    
    - build : compile, link, build(generates bin; only if main exists)

        ```go build hello.go```
    
    - build: specific output file name
  
        ```go build -o app hello.go```

    - to create a module

        ```go mod init <name of the module>```

## cross compile go appliations

    ### get the list of arch and os that go supports

        ```go tool dist list```

 ```GOOS=windows && GOARCH=amd64 go build -o app.exe hello.go```

 ## stripe off the binary while building

 ```go build -ldflags="-s -w" -o app hello.go```

 ## use tool compile

 ```go tool compile hello.go```

 ## use tool link

 ```go tool link -v -o hello hello.o```

## go install

```go install github.com/cweill/gotests/gotests@v1.6.0```
