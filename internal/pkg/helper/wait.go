package helper

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type WaitFunc func() (finished bool, err error)

func WaitForCommand(f WaitFunc) error {
	waitTimeout := 1200
	if viper.GetInt("command_wait_timeout_seconds") > 0 {
		waitTimeout = viper.GetInt("command_wait_timeout_seconds")
	}
	sleepTimeout := 5
	if viper.GetInt("command_wait_sleep_seconds") > 0 {
		sleepTimeout = viper.GetInt("command_wait_sleep_seconds")
	}

	timeStart := time.Now()

	for {
		if time.Since(timeStart).Seconds() > float64(waitTimeout) {
			return errors.New("Timed out waiting for command")
		}

		finished, err := f()
		if err != nil {
			return fmt.Errorf("Error waiting for command: %s", err)
		}
		if finished {
			break
		}

		time.Sleep(time.Duration(sleepTimeout) * time.Second)
	}

	return nil
}
