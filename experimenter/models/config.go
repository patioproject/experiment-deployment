package models

import (
	"time"

	core "go.cfdata.org/crypto/dome/collector/types"
)

const (
	defaultReadTimeout  = 10 * time.Second
	defaultMaxIdleConns = 1024
)

type ReadConfig struct {
}

//R2.Config should mvoe out of collect and here as a map of configs

func (ReadConfig) Read(hasEnv core.HasEnv) (*core.Config, error) {
	cfg := &core.Config{
		Secrets: core.Secrets{
			User:     hasEnv.Getenv("CF_USER"),
			Password: hasEnv.Getenv("CF_PASS"),
			Token:    hasEnv.Getenv("CF_TOKEN"),
			R2: core.R2Config{
				BucketName:      hasEnv.Getenv("R2_BUCKETNAME"),
				AccountId:       hasEnv.Getenv("R2_ACCOUNTID"),
				AccessKeyId:     hasEnv.Getenv("R2_ACCESSKEYID"),
				AccessKeySecret: hasEnv.Getenv("R2_ACCESSKEYSECRET"),
			},
		},
		ReadTimeout:  core.ParseIntOrDurationValue(hasEnv.Getenv("READ_TIMEOUT"), time.Second*10),
		WriteTimeout: core.ParseIntOrDurationValue(hasEnv.Getenv("WRITE_TIMEOUT"), time.Second*10),
	}

	cfg.MaxIdleConns = defaultMaxIdleConns
	cfg.MaxIdleConnsPerHost = defaultMaxIdleConns
	return cfg, nil
}
