package main

import (
	"fmt"
	"randoquote/server"
)

func main() {
	fmt.Println("Hello, playground")
    
	q := server.Quote{
		Quote: "Hello World",
		Author: "Me",
	}

	fmt.Println(q.PostQuote())
	
}
