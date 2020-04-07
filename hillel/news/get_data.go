package news

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type source struct {
	ID string `json:"id"`
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

	return sourceAPI.Sources
}

func getTopics(sources []source) []Topic {
	var topics []Topic

	for _, soure := range sources {
		body := getData(topicURL(source.ID))

		var topicAPI topicsAPI
		json.Unmarshal(topics, &topicAPI)

		topics = append(topics, topicAPI.Aricles...)

	}
	return topics
}

func topicURL(id string) string {
	return fmt.Sprintf()
}

//57cdad5e23744cb8be009055886fad29

func sourceURL(category string) string {
	return fmt.Sprintf("https://newsapi.org/v2/sources?language=en&category=%s&apiKey=57cdad5e23744cb8be009055886fad29", category)
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
	res.Close() // was absent
	return body

}
