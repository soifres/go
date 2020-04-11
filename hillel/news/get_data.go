package news

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/soifres/go_lessons/filework/utils"
)

type source struct {
	ID       string `json:"id"`
	Category string `json:"category"`
}

type sourcesAPI struct {
	Sources []source `json:"sources"`
}

type topicsAPI struct {
	Aricles []Topic `json:"articles"`
}

func getSources(category string) []source {
	body := getData(sourceURL(category))
	var sourceAPI sourcesAPI
	json.Unmarshal(body, &sourceAPI)

	writeInFile("sources", "sources_of_"+category, body)

	return sourceAPI.Sources
}

func getTopics(sources []source) []Topic {
	var topics []Topic

	for _, src := range sources {
		body := getData(topicURL(src.ID))

		var topicAPI topicsAPI
		json.Unmarshal(body, &topicAPI)

		writeInFile("topics", "topics_of_"+src.ID, body)

		topics = append(topics, topicAPI.Aricles...)

	}
	return topics
}

func writeInFile(dirname string, name string, body []byte) {
	nfile := utils.GetNewNFile()
	nfile.Name = name
	nfile.Extension = "json"
	nfile.Directory = "./json/" + dirname + "/" //"./json/"
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
	file.WriteString(string(body))
}

// q

// qInTitle

// sources
// ddd, aaa, ggg

// domains
// ddd.org, aaa.ru, ggg.com

// excludeDomains

// from
// 2020-04-08
// to
// 2020-04-08

// sortBy
//// relevancy статьи, более тесно связанные с qпервым
//// popularity статьи из популярных источников и издателей на первом месте
//// publishedAt  новые статьи на первом месте

// language
// ar de en es fr he it nl no pt ru se ud zh

func topicURL(id string) string {
	return fmt.Sprintf("https://newsapi.org/v2/everything?sources=%s&apiKey=57cdad5e23744cb8be009055886fad29", id)
	// return fmt.Sprintf("https://newsapi.org/v2/sources?language=en&id=%s&apiKey=57cdad5e23744cb8be009055886fad29", id)
}

// category
// business entertainment general health science sports technology

// language
// ru en

// country
// ru en us

//57cdad5e23744cb8be009055886fad29

func sourceURL(category string) string {
	return fmt.Sprintf("https://newsapi.org/v2/sources?&category=%s&apiKey=57cdad5e23744cb8be009055886fad29", category)
}

func getData(url string) []byte {

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body) // was "res", now "res.Body"
	if err != nil {
		panic(err)
	}
	// res.Close() // was absent
	return body

}
