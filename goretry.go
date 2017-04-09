package goretry

import (
	"fmt"
	"log"
	"time"
)

func Retry(attempts int, sleep time.Duration, execute func() error) (err error) {
	for i := 0; ; i++ {
		err = execute()
		if err == nil {
			return
		}

		if i >= (attempts - 1) {
			break
		}

		time.Sleep(sleep)

		log.Println("retrying after error:", err)
	}
	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}

func RetryDuring(duration time.Duration, sleep time.Duration, execute func() error) (err error) {
	t0 := time.Now()
	i := 0
	for {
		i++

		err = execute()
		if err == nil {
			return
		}

		delta := time.Now().Sub(t0)
		if delta > duration {
			return fmt.Errorf("after %d attempts (during %s), last error: %s", i, delta, err)
		}

		time.Sleep(sleep)

		log.Println("retrying after error:", err)
	}
}
