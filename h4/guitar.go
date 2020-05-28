package main

import ("fmt"
	"strings"
	)


func fretboard(){
	fmt.Println("0  1 2 3 4 5 6 7 8 9 10 11 12")

	for i := 0 ;i<6;i++{
		if i==2||i==3{
		 fmt.Printf(" || | | | |.| | | |.| | |:||\n")
		}else{
                fmt.Printf(" || | | | | | | | | | | | ||\n")
         	}
		}

	} 

func main(){


	eSlice:= []string{"E","F","F#","G","G#","A","A#","B","C","C#","D","D#"}
	aSlice:= []string{"A","A#","B","C","C#","D","D#","E","F","F#","G","G#"}
	dSlice:= []string{"D","D#","E","F","F#","G","G#","A","A#","B","C","C#"}
	gSlice:= []string{"G","G#","A","A#","B","C","C#","D","D#","E","F","F#"}
	bSlice:= []string{"B","C","C#","D","D#","E","F","F#","F","G","G#","A"}

	fmt.Println("The guitar fretboard, 12 frets:")
	fretboard()
	fmt.Printf("\nNotes per each string:\n%v\n%v\n%v\n%v\n%v\n%v\n",strings.Join(eSlice,"|"),strings.Join(bSlice,"|"), strings.Join(gSlice,"|"),strings.Join(dSlice,"|"),strings.Join(aSlice,"|"), strings.Join(eSlice,"|"))



}
