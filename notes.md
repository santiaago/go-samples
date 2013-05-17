
## notes for windows:

To add GOPATH environment variable do:
> setx GOPATH c:\go-samples

restart command line to see changes:

> echo %GOPATH%
c:\go-samples

To add c:\go\bin and c:\go-samples\bin and add the %GOPATH%\bin to PATH do the following:
> setx PATH "c\go\bin;c:\go-samples\bin"

To  build the program and create a hello.exe under the %GOPATH%\bin\
> go install src\hello

To build a package do:
> go build package name.

To testd do an "import package "testing"" and run 
> go test pkgname

To get a remote package:
> go get code.google.com/p/go.example/hello

## notes for mac:

[TODO]

## go notes:

Capital letter in functions of a package manage private public.