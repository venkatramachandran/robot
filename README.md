Robot
============
This is the solution to a toy robot problem described [here](Problem.MD).
The solution is done in Golang. To get the code to work, follow the steps below:

# How To Install Go
Please refer to instructions [here](https://golang.org/doc/install#install).

# Setting up GOPATH
Please refer to instructions [here](https://golang.org/doc/code.html#GOPATH)

# Getting the source code
Once Go is installed and a $GOPATH is configured, run 
```
go get -u github.com/venkatramachandran/robot
```
This will download the code and put it in `$GOPATH/src/github.com/venkatramachandran/robot`

# Building and running the code
To build, run
```
go build
```
from `$GOPATH/src/github.com/venkatramachandran/robot`

This will compile the code and create an executable in the same directory.

The executable can then be run with:
```
./robot -input <filename>
```
See `example.txt` for a sample input file.

# Running Tests
To run the tests, run
```go test```
from `$GOPATH/src/github.com/venkatramachandran/robot`
