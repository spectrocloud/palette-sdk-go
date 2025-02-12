package retry

import (
	"fmt"
	"math/rand"
	"time"
)

func Retry(retryMsg string, attempts int, sleep time.Duration, f func() error) error {
	return retryOp(retryMsg, attempts, sleep, f)
}

func retryOp(retryMsg string, attempts int, sleep time.Duration, f func() error) error {
	var err error

	r := rand.New(rand.NewSource(time.Now().UnixNano())) // #nosec G404 - non-cryptographic RNG is fine for jitter
	t := attempts

	for ; attempts >= 0; attempts-- {
		if t > attempts {
			fmt.Printf("retrying (%d/%d): %s ", t-attempts, t, retryMsg)
		}
		err = f()
		if err != nil {
			if sleep > 0 {
				time.Sleep(sleep)
				jitter := time.Duration(r.Int63n(int64(sleep)))
				sleep = (2 * sleep) + jitter/2 // exponential sleep with jitter
			}
		} else {
			return nil
		}
	}
	return err
}
