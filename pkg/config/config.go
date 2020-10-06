package config

import (
	"encoding/json"
	"io"
)

type Config struct {
	// Subsystem is a metric prefix to add to all exported metrics.  This can be used to avoid name collisions on multiple exporters.
	Subsystem   string             `json:"subsystem,omitempty"`
	CloudSQL    *CloudSQLConfig    `json:"cloudsql,omitempty"`
	Compute     *ComputeConfig     `json:"compute,omitempty`
	Filestore   *FilestoreConfig   `json:"filestore,omitempty`
	GCS         *GCSConfig         `json:"gcs,omitempty"`
	Stackdriver *StackdriverConfig `json:"stackdriver,omitempty"`
}

func (c *Config) Marshal(w io.Writer) error {
	data, err := json.MarshalIndent(c, " ", "\t")
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	if err != nil {
		return err
	}
	return nil
}

type CloudSQLConfig struct {
}

type ComputeConfig struct {
}

type FilestoreConfig struct {
}

type GCSConfig struct {
}

type StackdriverConfig struct {
}
