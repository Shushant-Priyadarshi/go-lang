package main

/*
Format Constants
	1) Date => 02/01/2006
	2) Time => 15:04:05 
	3) Time 12 Hour =>  3:04 PM
	4) Day => Monday
*/

import (
	"fmt"
	"time"
)

func main(){
	curentTime := time.Now()
	fmt.Println("Current time: ",curentTime)

	fomattedTime := curentTime.Format("02/01/2006, 15:04:05, Monday")
	fmt.Println("Formatted Time: ",fomattedTime) //Formatted Time:  15/06/2025, 19:10:31, Sunday


	fomattedTime12Hour := curentTime.Format("2006/01/02, 3:04 PM, Monday")
	fmt.Println("Formatted Time: ",fomattedTime12Hour) //Formatted Time:  2025/06/15, 7:10 PM, Sunday

	date_str := "24/09/2006"
	layout_str := "02/01/2006"
	date_str_now_is_a_date,_ := time.Parse(layout_str,date_str)
	fmt.Println("Date Str: ",date_str_now_is_a_date)

	tomorrow_date_Time := curentTime.Add(24 * time.Hour)
	tomorrow_date_Time_format := tomorrow_date_Time.Format("02/01/2006, 3:04 PM, Monday")
	fmt.Println("Kal: ",tomorrow_date_Time_format)
}