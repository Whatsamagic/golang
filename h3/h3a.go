package main

import ("fmt"
	"flag"
	"time"
	"io/ioutil"
	"log"
	"os"
)
func  logFile(file string, dur int){
	content := []byte("Start of logfile\n")
	_ ,err := os.Stat(file)
	if os.IsNotExist(err) {
		err :=ioutil.WriteFile(file, content, 0644)
		if err != nil{
			fmt.Println("File error:", err)
			log.Panic(err)
		}	
	}
	
	appendfile, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil{
		log.Println(err)
	}
	defer appendfile.Close()

	logger := log.New(appendfile, "", log.LstdFlags)

	for i:=dur; i>0;i--{
		time.Sleep(time.Second)			
		//if _, err := appendfile.WriteString("additional line\n"); err !=nil{
		//log.Fatal(err)
		logger.Println("näin vaan mennään vielä",i,"kertaa")
	}
}


func main(){
	var filename string
	var duration int	
	
	fmt.Println("Program will create a log file based on provided flags")
	flag.StringVar(&filename, "filename", "testfile","Name of log file")
	flag.IntVar(&duration, "duration", 10, "duration of logging")
	flag.Parse()
	
	log.Printf("Starting writing file: %v for %v s...\n",filename,duration)
	logFile(filename,duration)	
	//fmt.Println("file: %v duration %v", filename, duration)
	log.Println("Completed logging!")
}
