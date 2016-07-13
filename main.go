package main

import (
	"fmt"
	"os"
	"go-test/cli"
)


func main() {

	fmt.Println("Hello")

	commands := []*cli.Command{
		&cli.Command{
			Name:        "Text Command",
			Cmd:         "text",
			Callback:    testHandler,
			Description: "This is the test command",
		},
		&cli.Command{
			Name: "About Command",
			Cmd:  "about",
			Callback: aboutHandler,
			Description: "This is the about command",
		},
	}


	for _, cmd := range commands {
		fmt.Printf("Name: %s\n", cmd.Name);
	}


	if ok := HandleArgs(os.Args[1:]); !ok {
		fmt.Println("No Valid arguments passed")
		os.Exit(1);
	}
	
}


func HandleArgs(args []string) bool {

	if len(args) <= 1 {
		return false
	}


	return true
}


func testHandler(ctx cli.Context) {
	fmt.Printf("The name from the context is %v\n", ctx.Name)
}

func aboutHandler(ctx cli.Context) {
	fmt.Printf("This is the about handler\n");
}