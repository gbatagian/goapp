package httpsrv

import (
	"context"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"goapp/internal/pkg/watcher"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	strChan      <-chan string               // String channel.
	server       *http.Server                // Gorilla HTTP server.
	watchers     map[string]*watcher.Watcher // Counter watchers (k: counterId).
	watchersLock *sync.RWMutex               // Counter lock.
	sessionStats []sessionStats              // Session stats.
	quitChannel  chan struct{}               // Quit channel.
	running      sync.WaitGroup              // Running goroutines.
}

func New(strChan <-chan string) *Server {
	s := Server{}
	s.strChan = strChan
	s.server = nil // Set below.
	s.watchers = make(map[string]*watcher.Watcher)
	s.watchersLock = &sync.RWMutex{}
	s.sessionStats = []sessionStats{}
	s.quitChannel = make(chan struct{})
	s.running = sync.WaitGroup{}
	return &s
}

func (s *Server) Start() error {
	// Create router.
	r := mux.NewRouter()

	// Register routes.
	for _, route := range s.myRoutes() {
		if route.Method == "ANY" {
			r.Handle(route.Pattern, route.HFunc)
		} else {
			r.Handle(route.Pattern, route.HFunc).Methods(route.Method)
			if route.Queries != nil {
				r.Handle(route.Pattern, route.HFunc).Methods(route.Method).Queries(route.Queries...)
			}
		}
	}

	// Create HTTP server.
	s.server = &http.Server{
		Addr:         "localhost:8080",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  10 * time.Second,
		Handler:      handlers.CombinedLoggingHandler(os.Stdout, r),
	}

	// Start HTTP server.
	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				panic(err)
			}
		}
	}()

	s.running.Add(1)
	go s.mainLoop()

	return nil
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		log.Printf("error: %v\n", err)
	}

	close(s.quitChannel)
	s.running.Wait()
}

func (s *Server) mainLoop() {
	defer s.running.Done()

	for {
		select {
		case str := <-s.strChan:
			s.notifyWatchers(str)
		case <-s.quitChannel:
			return
		}
	}
}
