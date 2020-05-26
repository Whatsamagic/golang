package main

import(
	"fmt"
	"time"
	"math/rand"
)

func howLong(seed int64)int{ //return a random value between 0-31
	rand.Seed(seed)
	x:=rand.Intn(31)
	return x
	}

func isItWeekend(day int)bool{ //check if weekday(int) is weekend
	if day == 5|| day ==6{
	return true}
	if day == 0{
	fmt.Println("..the Muses remind you that Sundays are for rest and fear of Mondays.." )
	return false
}
	return false
 }


	
func main(){
	now := time.Now()
	random:=howLong(now.UTC().UnixNano())	
	then := now.AddDate(0,0,random)
	thenday:= then.Weekday()

	fmt.Printf("\nToday, on %v %v Anno %v..\nYou have approached the Muses of Go for the permission of consumption of a brewed beverage %v days from now..\n", now.Month(), now.Day(), now.Year(),random )
	fmt.Printf("..that day being %v, %v %v Anno %v..\n",then.Weekday(),then.Month(),then.Day(),then.Year())

	if random <7{
		fmt.Printf("..which is less than a week from now..\n")
	}
	if random ==0{
		fmt.Printf("..which happens to be today..")
	}
	        
	defer fmt.Println("You will be grateful of this attention and shall not approach the Muses of Go again before the day has passed.")
	if isItWeekend(int(thenday))==true {
		fmt.Printf("\nThe muses have decided to allow you to consume a beer!\n Think kindly of them during this weekend!\n\n")
	}  else {
	fmt.Printf("During these days of work, you will be allowed to consume only mineral water.\n")}
	
}
