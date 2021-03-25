package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Hello World !!!!")
	fmt.Println("This is your second commit")
	fmt.Println(strings.ToUpper("Hello from me to your world, really loud"))

	
}

// go run filename :- Compiles and runs the .go file
// gofmt filename:- formats the file properly and outputs on the command line
// gofmt -w filename :- formats and saves the changes
// go build filename :- creates executable .exe file but saves in the same directory
// set GOPATH=C:\Users\arjun\Desktop\GO tutorial 
// go install //on command prompt generates a .exe file in the bin folder of workspace 
// ToUppercase function

/*
	GOPATH	
		/src
			/github.com
				/username
					/repo
			/bitbucket.com
				/username
					/repo
			/gitlab.com
				/username
					/repo   
							*/ // this is how github works
