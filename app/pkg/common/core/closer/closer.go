package closer

import (
	"io"
	"sync"

	errors "github.com/hashicorp/go-multierror"
)

type WoContext interface {
	Close()
}

type Closer = io.Closer
type NCloser = WoContext

type CloseFunc func() error
type NCloseFunc func()

var (
	defaultCloser = NewLifoCloser()
	once          sync.Once
)

func Add(closers ...Closer) {
	defaultCloser.Add(closers...)
}

func AddN(closers ...NCloser) {
	defaultCloser.AddN(closers...)
}

func Close() (err error) {
	once.Do(func() {
		err = defaultCloser.Close()
	})

	return
}

type LifoCloser struct {
	mu       sync.Mutex
	closers  []Closer
	nClosers []NCloser
}

func NewLifoCloser() *LifoCloser {
	return new(LifoCloser)
}

func (s *LifoCloser) Add(closers ...Closer) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.closers = append(s.closers, closers...)
}

func (s *LifoCloser) AddN(closers ...NCloser) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.nClosers = append(s.nClosers, closers...)
}

func (s *LifoCloser) Close() (errs error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := len(s.closers) - 1; i >= 0; i-- {
		closer := s.closers[i]
		if err := closer.Close(); err != nil {
			errs = errors.Append(errs, err)
		}
	}

	for i := len(s.nClosers) - 1; i >= 0; i-- {
		closer := s.nClosers[i]
		closer.Close()
	}

	return
}
