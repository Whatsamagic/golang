package main

import ("fmt"
	"strings"
	"time"
	"math/rand"
)

func eSlice()[]string{
	return []string{"E","F","f#/gb","G","g#/ab","A","a#/bb","B","C","c#/db","D","d#/eb","E"}
}
func aSlice() []string{
	return []string{"A","a#/bb","B","C","c#/db","D","d#/eb","E","F","f#/gb","G","g#/ab","A"}
}
func dSlice() []string{
	return []string{"D","d#/eb","E","F","f#/gb","G","g#/ab","A","a#/bb","B","C","c#/db","D"}
}
func gSlice() []string{
	return []string{"G","g#/ab","A","a#/bb","B","C","c#/db","D","d#/eb","E","F","f#/gb","G"}
}
func bSlice() []string{
	return []string{"B","C","c#/db","D","d#/eb","E","F","f#/gb","G","g#/ab","A","a#/bb","B"}
}



func fretboard(){ //display simplified guitar fretboard with fret numbers
	fmt.Println("0  1 2 3 4 5 6 7 8 9 10 11 12")

	for i := 0 ;i<6;i++{
		if i==2||i==3{
		 fmt.Printf(" || | | | |.| | | |.| | |:||\n")
		}else{
                fmt.Printf(" || | | | | | | | | | | | ||\n")
         	}
		}
	} 

func  showNotes(input string,onestring []string){ //print all requested full tone on fretboard
	tmpSlice := make([]string,13)
	for i:=0; i<13; i++{
		if onestring[i]==input {
		tmpSlice[i]=input
		}else{
		tmpSlice[i]=" "
		}}
	fmt.Println(strings.Join(tmpSlice,"|"),"|")
	}

func hideNotes(input string, onestring []string){ //print only single mystery note
	tmpSlice := make([]string,13)
	if len(onestring) < 12 {
		for i:=0; i<13; i++ {
		tmpSlice[i]=" "
		}}else{
			for i:=0; i<13; i++ {
				if onestring[i]==input {
				tmpSlice[i]="*"
				}else{
				tmpSlice[i]=" "
			}
	}}
	fmt.Println(strings.Join(tmpSlice,"|"),"|")
	}

func guessNote(x,y int){
	var answer string =""
	fmt.Println("0 1 2 3 4 5 6 7 8 9 10 11 12")

	switch y{
	case 0:
		answer=eSlice()[x]
		hideNotes(eSlice()[x],eSlice())
                hideNotes("",make([]string,1))
                hideNotes("",make([]string,1))
                hideNotes("",make([]string,1))
                hideNotes("",make([]string,1))
                hideNotes("",make([]string,1))
	case 1:
		answer=bSlice()[x]
                hideNotes("",make([]string,1))
		hideNotes(bSlice()[x],bSlice())
                hideNotes("",make([]string,1))
                hideNotes("",make([]string,1))
                hideNotes("",make([]string,1))
                hideNotes("",make([]string,1))
	case 2:
		answer=gSlice()[x]
                hideNotes("",make([]string,1))
                hideNotes("",make([]string,1))
		hideNotes(gSlice()[x],gSlice())
                hideNotes("",make([]string,1))
                hideNotes("",make([]string,1))
                hideNotes("",make([]string,1))
	case 3:
		answer=dSlice()[x]
                hideNotes("",make([]string,1))
                hideNotes("",make([]string,1))
                hideNotes("",make([]string,1))
                hideNotes(dSlice()[x],dSlice())
                hideNotes("",make([]string,1))
                hideNotes("",make([]string,1))
	case 4:
		answer=aSlice()[x]
                hideNotes("",make([]string,1))
                hideNotes("",make([]string,1))
                hideNotes("",make([]string,1))
                hideNotes("",make([]string,1))
 		hideNotes(aSlice()[x],aSlice())
                hideNotes("",make([]string,1))
	case 5:
		answer=eSlice()[x]
                hideNotes("",make([]string,1))
                hideNotes("",make([]string,1))
                hideNotes("",make([]string,1))
                hideNotes("",make([]string,1))
                hideNotes("",make([]string,1))
		hideNotes(eSlice()[x],eSlice())

	}
	fmt.Println("Press enter for answer.")
	fmt.Scanln() // wait for user input
	fmt.Println("Answer is:",answer)
	time.Sleep(2*time.Second)
	main()
}

func whichNote(x string){	// show all selected notes on fretboard
	fmt.Println("0 1 2 3 4 5 6 7 8 9 10 11 12")

	showNotes(x,eSlice())
	showNotes(x,bSlice())
	showNotes(x,gSlice())
	showNotes(x,dSlice())
	showNotes(x,aSlice())
	showNotes(x,eSlice())
	time.Sleep(2*time.Second)
	main()
}

func cheat(){ //show all notes on each guitarstring
	fmt.Printf("\nNotes per each string:\n%v\n%v\n%v\n%v\n%v\n%v\n",strings.Join(eSlice(),"|"),strings.Join(bSlice(),"|"), strings.Join(gSlice(),"|"),strings.Join(dSlice(),"|"),strings.Join(aSlice(),"|"), strings.Join(eSlice(),"|"))
	time.Sleep(2*time.Second)
	main()
}

func main(){

	rand.Seed(time.Now().UnixNano())
	x:=rand.Intn(12)
	y:=rand.Intn(6)


	fmt.Printf("\n\nThe guitar fretboard, 12 frets:\n")
	fretboard()

	fmt.Println("\nWhat do You want to do:  1)Guess the note, 2)Show note on fretboard, 3)Show all notes, 4)Quit")
	var show string 
	var selection int
     	fmt.Scanln(&selection)                  // Take input from user

	switch selection {
	case 1:
		guessNote(x,y)

	case 2:
		fmt.Println("Which note you want to be displayed? C,D,E,F,G,A or B?")
     		fmt.Scanln(&show)                  // Take input from user
		switch strings.ToUpper(show){
		case "C":
			whichNote("C")
		case "D":
			whichNote("D")
		case "E":
			whichNote("E")
		case "F":
			whichNote("F")
		case "G":
			whichNote("G")
		case "A":
			whichNote("A")
		case "B":
			whichNote("B")
		}
	case 3:
		cheat()
	default:
		fmt.Println("Goodbye!")
	}
}



