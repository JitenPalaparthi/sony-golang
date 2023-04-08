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

package,func

## builtin

println


## go tools

    - run go application : compile, link, build(generates bin) and then run
  
        ```go run hello.go```
    
    - build : compile, link, build(generates bin; only if main exists)

        ```go build hello.go```
    
    - build: specific output file name
  
        ```go build -o app hello.go```

    - to create a module

        ```go mod init <name of the module>```