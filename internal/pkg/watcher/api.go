package watcher

type Counter struct {
	Iteration int    `json:"iteration"`
	Value     string `json:"value"`
}

type CounterReset struct{}
