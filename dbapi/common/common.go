package common

import (
	lr "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

var Logr = &lr.Logger{
	Out:   os.Stdout,
	Level: lr.TraceLevel,
	Formatter: &prefixed.TextFormatter{
		DisableColors:   false,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceFormatting: true,
	},
}

func ScanDir(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			Logr.Fatal(err)
		}
		*files = append(*files, path)
		return nil
	}
}

type Conf struct {
	ApiPath string `yaml:"apiPath"`
}

func (c *Conf) GetConf() *Conf {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		Logr.Printf("GetConfig: can't ipen config.yaml file (%v)\n", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		Logr.Fatalf("GetConfig: config.yaml unmarshal error (%v)\n", err)
	}
	return c
}
