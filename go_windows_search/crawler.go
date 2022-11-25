package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func getCrawling() {

	//get user input
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Directory Path:  ")

	dirPath, _ := reader.ReadString('\n')

	newPath := strings.TrimSpace(dirPath)

	os.Chdir(newPath)

	//newDir := newPath

	newPath, err := os.Getwd()

	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println(newPath)

		getContent, err := ioutil.ReadDir(newPath)

		if err != nil {
			fmt.Println("Reading Directory Error")
		} else {
			for _, val := range getContent {
				fmt.Println(val.Name(), " ...... ", val.Size())
				//fmt.Println(val.Sys())
			}
		}
	}

}
