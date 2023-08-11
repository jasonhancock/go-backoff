# go-backup

[![Go Reference](https://pkg.go.dev/badge/github.com/jasonhancock/go-backoff.svg)](https://pkg.go.dev/github.com/jasonhancock/go-backoff)

This is a fork of the backup from [gmkit](https://github.com/graymeta/gmkit) with external dependencies removed.

### Minimal configuration

```go
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
```

### Customized config

```go
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
```
