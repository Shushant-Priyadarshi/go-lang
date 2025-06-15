package main

import (
	"fmt"
	"strings"
)

func main(){
	data := "Manjistha,Kavya,Shushant,Manjistha,Manjistha"
	data_split_commas := strings.Split(data,",")
	fmt.Println(data_split_commas)

	data_count := strings.Count(data,"Manjistha")
	fmt.Println(data_count)

	new_data := "    ....Helooo World ...      "
	new_data_trim := strings.TrimSpace(new_data)
	fmt.Println(new_data_trim)

	rel := "Loves"
	joining_string := strings.Join([]string{"Shushant",rel,"Manjistha"}," ")
	fmt.Println(joining_string)

}