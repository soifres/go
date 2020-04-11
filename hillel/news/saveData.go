package news

import (
	"os"
)

func saveToFile(category string, body []byte) {
	rootFolder := "json/"

}

func getFile(name string, max int) *File {
	ext := ".json"
	cname := name + ext
	file, err := os.Open("cname")
}
