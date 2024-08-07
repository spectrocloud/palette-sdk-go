package retry

import (
	"fmt"
	"math/rand"
	"time"
)

type RetryOption struct {
	retryMsg  string
	attempts  int
	sleep     time.Duration
	retryFlag bool
}

func NewRetryOption(retryMsg string, attempts int, sleep time.Duration, retryFlag bool) *RetryOption {
	return &RetryOption{retryMsg: retryMsg, attempts: attempts, sleep: sleep, retryFlag: retryFlag}
}

func (retryOption *RetryOption) Retry(f func() error) error {
	return retryOp(retryOption.retryMsg, retryOption.attempts, retryOption.sleep, f, retryOption.retryFlag)
}

func Retry(retryMsg string, attempts int, sleep time.Duration, f func() error) error {
	return retryOp(retryMsg, attempts, sleep, f, true)
}

func RetryWithErrRetryCond(retryMsg string, attempts int, sleep time.Duration, f func() error) error {
	return retryOp(retryMsg, attempts, sleep, f, false)
}

func retryOp(retryMsg string, attempts int, sleep time.Duration, f func() error, retryFlag bool) error {

	rand.Seed(time.Now().UnixNano())
	var err error
	t := attempts
	for ; attempts >= 0; attempts-- {
		if t > attempts {
			fmt.Printf("retrying (%d/%d): %s ", t-attempts, t, retryMsg)
		}
		err = f()
		if err != nil {
			if sleep > 0 {
				time.Sleep(sleep)
				jitter := time.Duration(rand.Int63n(int64(sleep)))
				sleep = (2 * sleep) + jitter/2 //exponential sleep with jitter
			}
		} else {
			return nil
		}
	}
	return err
}
