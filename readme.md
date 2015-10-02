##This project is for testing the capabilities of the go programming language 

### Installation of the go lang

**Note:** In order to properly set the environment please pay attention to the os variable setting.

$ export GOPATH=$WORKSPACE
$ export GOROOT=only set it if the golang is not install in the default location (/usr/local/go)

**Read more here**  

    https://golang.org/doc/code.html
    https://github.com/golang/go/wiki/InstallTroubleshooting


##### Tasks

1. Create a simple REST example with gorilla mux GO library

    How to run:

    $ go get github.com/gorilla/mux

    $ go run src/code.1.1/code.1.1.go &

    $ curl http://localhost:8000

    Expected result: {"ID":1,"FirstName":"Antonel","LastName":"Pazargic"}

#### Git stuff

1. For changing the message of the last commit

    $ git commit --amend

#### IntelliJ Idea stuff

1. Start the intelliJ with the following script commands

    export GOPATH=$PATH_TO_PROJECT_ROOT

    PATH_TO_IDEA/bin/idea.sh


