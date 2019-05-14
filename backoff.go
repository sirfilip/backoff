package backoff

import "time"

func Retry(f func() error, times int, duration time.Duration) error {
	var err error
	var durations []time.Duration

	for i := 0; i < times; i++ {
		durations = append(durations, duration)
		duration = time.Duration(duration + duration)
	}

	for _, duration := range durations {
		err = f()
		if err == nil {
			break
		}
		time.Sleep(duration)
	}
	return err
}
