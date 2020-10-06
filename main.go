package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/drud/gcp-exporter/pkg/config"
)

func printConfigAndExit() {
	c := &config.Config{
		Subsystem:   "",
		CloudSQL:    &config.CloudSQLConfig{},
		Compute:     &config.ComputeConfig{},
		Filestore:   &config.FilestoreConfig{},
		GCS:         &config.GCSConfig{},
		Stackdriver: &config.StackdriverConfig{},
	}
	err := c.Marshal(os.Stdout)
	if err != nil {
		log.Fatalf("failed to generate config: %v", err)
	}
	os.Exit(0)
}

func main() {
	configFile := flag.String("config", "", "the location of the config file for this service.  Config files can be generated with -genConfig")
	genConfig := flag.Bool("genConfig", false, "generate an example config and exit")
	flag.Parse()

	if *genConfig {
		printConfigAndExit()
	}

	if *configFile == "" {
		log.Fatal("'config' flag must be specified")
	}

	configData, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatalf("Failed to read configuration file %s: %v", *configFile, err)
	}

	var exporterConfig config.Config
	err = json.Unmarshal(configData, &exporterConfig)
	if err != nil {
		log.Fatalf("Error marshalling configuration in file %s: %v", *configFile, err)
	}
	_ = exporterConfig
}
