package main

import (
	"fmt"
	"hsite"
	_ "mysql"
)

func main() {
	fmt.Println("STARTING...")
	var app = &hsite.App{}
	app.Run()
	fmt.Println("EXITING...")
}
