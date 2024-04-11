package backoff_test

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jasonhancock/go-backoff"
)

type myLogger struct{}

func (l *myLogger) Warn(msg any, keyvals ...any) {
	fmt.Println(msg)
}

func ExampleNew_minimum() {
	boff := backoff.New(backoff.WithLogger(&myLogger{}))
	i := 0
	err := boff.Backoff(func() error {
		i++
		if i < 4 {
			return errors.New("some error")
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleNew_customized() {
	boff := backoff.New(
		backoff.WithLogger(&myLogger{}),
		backoff.InitBackoff(10*time.Second),
		backoff.MaxCalls(2),
	)
	i := 0
	err := boff.Backoff(func() error {
		i++
		if i < 4 {
			return errors.New("some error")
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
