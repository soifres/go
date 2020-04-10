package news

var (
	chCollect      chan string
	chRequest      chan string
	chResultTopics chan []Topic
	chResultCount  chan int
)

func init() {
	chCollect = make(chan string)
	chResultTopics = make(chan []Topic)
	chRequest = make(chan string)
}

// Collect Do collection
func Collect(category string) {
	chCollect <- category
	//return <-chResultCount
}

// Result Return result
func Result(category string) []Topic {
	chRequest <- category
	return <-chResultTopics
}

// Serve Находится в бесконечном цикле ожидания,
// когда какой-нибудь канал получит значение.
func (a Archive) Serve() {
	for {
		select {
		case category := <-chCollect:
			a.collect(category)
			// chResultCount <- a.collect(category)
		case category := <-chRequest:
			chResultTopics <- a.result(category)

		}
	}

}
