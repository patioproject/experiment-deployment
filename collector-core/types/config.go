package types

import (
	"errors"
	"os"
	"strconv"
	"time"
)

/*---------- OsEnv ---------------*/
type OsEnv struct{}

func (OsEnv) Getenv(key string) string {
	return os.Getenv(key)
}

type HasEnv interface {
	Getenv(key string) string
}

func ParseIntValue(val string, fallback int) int {
	if len(val) > 0 {
		parsedVal, parseErr := strconv.Atoi(val)
		if parseErr == nil && parsedVal >= 0 {
			return parsedVal
		}
	}
	return fallback
}

func ParseIntOrDurationValue(val string, fallback time.Duration) time.Duration {
	if len(val) > 0 {
		parsedVal, parsedErr := strconv.Atoi(val)
		if parsedErr == nil && parsedVal >= 0 {
			return time.Duration(parsedVal) * time.Second
		}
	}

	duration, durationErr := time.ParseDuration(val)
	if durationErr != nil {
		return fallback
	}
	return duration
}

func ParseIntOrDuration(val string) (int, error) {
	i, err := strconv.ParseInt(val, 10, 0)
	if err == nil {
		return int(i), nil
	}

	if err != nil && errors.Is(err, strconv.ErrRange) {
		return int(i), err
	}

	d, err := time.ParseDuration(val)
	if err != nil {
		return 0, err
	}

	return int(d.Seconds()), nil
}

func ParseBoolValue(val string, fallback bool) bool {
	if len(val) > 0 {
		return val == "true"
	}
	return fallback
}

func ParseString(val string, fallback string) string {
	if len(val) > 0 {
		return val
	}
	return fallback
}

func GetSecrets(val string, fallback string) string {
	if len(val) > 0 {
		return val
	}
	return fallback
}

type Config struct {
	TCPPort             int // public port for API
	Endpoint            string
	ReadTimeout         time.Duration
	WriteTimeout        time.Duration
	MaxIdleConns        int
	MaxIdleConnsPerHost int
	Secrets             Secrets
}

type Handlers struct{}
