package news

// Topic Топик
type Topic struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Author      string `json:"author"`
	URL         string `json:"url"`
}

//

// Archive Архив топиков
type Archive map[string][]Topic

// New Возвращает архив топиков
func New() Archive {
	return make(map[string][]Topic)
}

func (a Archive) collect(category string) int {
	sources := getSources(category)
	topics := getTopics(sources)
	a[category] = topics
	return len(topics)
}

func (a Archive) result(category string) []Topic {
	return a[category]
}
