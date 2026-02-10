package types

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	Changelog  Changelog  `yaml:"changelog"`
	Metadata   Metadata   `yaml:"metadata"`
	Clickhouse Clickhouse `yaml:"clickhouse"`
}

func (cf *Conf) GetMetaData() Metadata {
	return cf.Metadata
}

type Metadata struct {
	Name           string `yaml:"name"`
	Port           int    `yaml:"port"`
	Description    string `yaml:"description"`
	ExportDataPath string `yaml:"export_data_path"`
	ApiURL 	       string `yaml:"url"` // added this to make my api resource work
	ApiKey		   string `yaml:"apikey"`	// added this to make my api resource work
}
type Changelog struct {
	Name             string `yaml:"name"`
	Format           string `yaml:"format"`
	Endpoint         string `yaml:"endpoint"`
	Size             int    `yaml:"size"`
	Service          string `yaml:"service"`
	GreaterThanEqual string `yaml:"gte"`
	LessThan         string `yaml:"lt"`
}

type Clickhouse struct {
	Name             string `yaml:"name"`
	Columns          string `yaml:"columns"`
	Service          string `yaml:"service"`
	Endpoint         string `yaml:"endpoint"`
	GreaterThanEqual string `yaml:"gte"`
	Limit            string `yaml:"limit"`
}

func (c *Conf) GetConf(f string) *Conf {

	yamlFile, err := os.ReadFile(f)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
