
## notes for windows:

To add c:\go\bin to PATH
setx PATH "%PATH%;c\go\bin"

To add GOPATH environment variable do:

> setx GOPATH c:\go-samples

restart command line to see changes:

> echo %GOPATH%
c:\go-samples

> go install src\hello
this will build the program and create a hello.exe under the %GOPATH%\bin\

To add the %GOPATH%\bin to PATH do the following:
setx PATH "%PATH%;c\go-samples\bin"