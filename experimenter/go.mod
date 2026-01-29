module go.cfdata.org/crypto/dome/changelog

go 1.23.4

replace go.cfdata.org/crypto/dome/collector => ../collector-core

// replace go.cfdata.org/crypto/dome/otis => /Users/innoobi/Documents/dev/other/dome-main/otislang/otis

replace code.cfops.it/crypto/dome/changelog => ../experimenter

require (
	github.com/common-nighthawk/go-figure v0.0.0-20210622060536-734e95fb86be
	github.com/sirupsen/logrus v1.9.3
	go.cfdata.org/crypto/dome/collector v0.0.0-00010101000000-000000000000
// go.cfdata.org/crypto/dome/otis v0.0.0-00010101000000-000000000000
)

require (
	github.com/robfig/cron/v3 v3.0.1 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
