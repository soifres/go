package main

import (
	"fmt"
	"io/ioutil"

	"github.com/soifres/go_lessons/filework/utils"
)

var dirJSON string = "./json/"

func main() {
	result, err := utils.IsFileInDirectory(dirJSON, "tst")
	if err != nil {
		fmt.Println("Done...TstString completed WITH ERROR")
	} else {
		fmt.Println(fmt.Sprintf("Done...TstString completed with %t", result))
	}

	nfile := utils.GetNewNFile()
	nfile.Name = "job"
	nfile.Extension = "json"
	nfile.Directory = "./json/"
	nfile.MaxCount = 1000
	file, err := nfile.Create()
	if err != nil {
		fmt.Println("Error in NFile.Create():\n" + err.Error())
		return
	}
	if file == nil {
		fmt.Println("File does not created, NFile.Create() return \"nil\"")
		return
	}
	defer file.Close()
	file.WriteString("{\"status\": \"error\",\"code\": \"apiKeyInvalid\"," +
		"\"message\": \"Your API key is invalid or incorrect. Check your key, or go to https://newsapi.org to create a free API key.\"}")

	data, err := ioutil.ReadFile("./json/job.json")
	if err != nil {
		fmt.Println("Error")
		return
	}

	fmt.Println(string(data))

}
