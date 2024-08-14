package util

import (
	"time"
)

/*
Wrapper function to handle retry
ref: https://rednafi.com/go/retry_function/
*/
func RetryFunc(fn func() error, maxRetry int, startBackoff time.Duration, maxBackoff time.Duration) error {
	var err error

	for attempt := 0; ; attempt++ {
		if err = fn(); err == nil {
			return nil
		}

		if attempt == maxRetry-1 {
			return err
		}

		time.Sleep(startBackoff)
		if startBackoff < maxBackoff {
			startBackoff *= 2
		}
	}
}
