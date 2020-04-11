package main

import (
	"fmt"

	"github.com/soifres/go_lessons/filework/utils"
)

var dirJSON string = "./json/"

func main() {
	utils.Tst(dirJSON)
	fmt.Println("Done...Tst completed")

	result, err := utils.HasFileInDirectory(dirJSON, "tst")
	if err != nil {
		fmt.Println("Done...TstString completed WITH ERROR")
	} else {
		fmt.Println(fmt.Sprintf("Done...TstString completed with %t", result))
	}

	file := utils.CreateNewFile(dirJSON, "job.health")
	if file == nil {
		fmt.Println("Error in main")
		return
	}
	defer file.Close()
	file.WriteString("{\"status\": \"error\",\"code\": \"apiKeyInvalid\"," +
		"\"message\": \"Your API key is invalid or incorrect. Check your key, or go to https://newsapi.org to create a free API key.\"}")

}
