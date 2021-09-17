package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Hey %s you've just executed random code from the internet. Did you expect that?\n", os.Getenv("USER"))
}
