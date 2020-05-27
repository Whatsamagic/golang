package main

import  (
	"log"
	"os"
//	"fmt"
	)

func main(){
	fileinfo,err := os.Stat("testfile")
	if os.IsNotExist(err) {
		log.Fatal("File does not exist.")
	}
	
	log.Println("type is: ",fileinfo)
}
