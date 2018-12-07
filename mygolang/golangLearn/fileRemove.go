package main

import (
	"fmt"
	"os"
)

func RemoveFile(){
	path := "./file/test2"
	err := os.Remove(path)
	if err !=nil{
		fmt.Println(err)
	}

}

func RemoveDir(){
	path :="./testDir"
	err := os.RemoveAll(path)
	if err != nil{
		fmt.Println(err)
	}

}

func CreateFile(){
	path := "./file/test3"
	f,err := os.Create(path)
	if err != nil{
		fmt.Println(err)
		fmt.Println(f)
	}
}

func CreateDir(){
	path := "./testDir"
	err :=os.MkdirAll(path,os.ModePerm)
	if err !=nil{
		fmt.Println(err)
	}


}

func main(){
	//CreateDir()
	CreateFile()
	//RemoveFile()
	//RemoveDir()
}

