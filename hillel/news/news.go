package news

var (
	chCollect chan string
	chResult  chan []Topic
	chRequest chan string
)

func init() {
	chCollect = make(chan string)
	chResult = make(chan []Topic)
	chRequest = make(chan string)
}

// Collect Do collection
func Collect(category string) {
	chCollect <- category
}

// Result Return result
func Result(category string) []Topic {
	chRequest <- category
	return <-chResult
}

// Serve Serve
func (a Archive) Serve() {
	for {
		select {
		case category := <-chCollect:
			a.collect(category)
		case category := <-chRequest:
			chResult <- a.result(category)

		}
	}

}
