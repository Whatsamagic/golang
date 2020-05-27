package main

import (
"fmt"
"flag")

func main(){
	var userinput string
	flag.StringVar(&userinput, "feature", "", "Desired feature")
	flag.Parse()
	if userinput==""{
		fmt.Println ("Please provide required feature as parameter!")
	} else {
		fmt.Println ("feature", userinput, "is not available")
	}
}
