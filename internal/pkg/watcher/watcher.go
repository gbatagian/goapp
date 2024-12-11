package watcher

import (
	"sync"

	"github.com/google/uuid"
)

type Watcher struct {
	id          string         // Watcher ID.
	inCh        chan string    // Input channel.
	outCh       chan *Counter  // Updates to counter will notify this channel.
	counter     *Counter       // The counter.
	counterLock *sync.RWMutex  // Lock for counter.
	quitChannel chan struct{}  // Quit.
	running     sync.WaitGroup // Run, Amy, Run!
}

func New() *Watcher {
	w := Watcher{}
	w.id = uuid.NewString()
	w.inCh = make(chan string, 1)
	w.outCh = make(chan *Counter, 1)
	w.counter = &Counter{Iteration: 0}
	w.counterLock = &sync.RWMutex{}
	w.quitChannel = make(chan struct{})
	w.running = sync.WaitGroup{}
	return &w
}

// Start watcher in another Go routine, Stop() must be called at the end.
func (w *Watcher) Start() error {
	w.running.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case <-w.inCh:
				w.counter.Iteration += 1
				select {
				case w.outCh <- w.counter:
				case <-w.quitChannel:
					return
				}
			case <-w.quitChannel:
				return
			}
		}
	}(&w.running)

	return nil
}

func (w *Watcher) Stop() {
	w.counterLock.Lock()
	defer w.counterLock.Unlock()

	close(w.quitChannel)
	w.running.Wait()
}

func (w *Watcher) GetWatcherId() string { return w.id }

func (w *Watcher) Send(str string) { w.inCh <- str }

func (w *Watcher) Recv() <-chan *Counter { return w.outCh }

func (w *Watcher) ResetCounter() {
	w.counterLock.Lock()
	defer w.counterLock.Unlock()

	w.counter.Iteration = 0

	select {
	case w.outCh <- w.counter:
	case <-w.quitChannel:
		return
	}
}
