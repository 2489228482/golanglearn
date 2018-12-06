package main

import (
	"fmt"

)


func main(){
	//type1
	for i:=0;i<=10;i++{
		fmt.Printf("type1-%d\n",i)

	}

	//type 2
	i:=0
	for i<=10{
		i++
		fmt.Printf("type2-%d\n",i)
	}


	//type 3
	s := []int{1,2,3,4,5,6,7,8,9,10}
	for v := range s{
		fmt.Print("type3-%d\n",v)

	}


	//type 4
	i =0
	for {
		fmt.Printf("type4-%d\n",i)
		i++
	}



}