package config

import (
  "fmt"
  "io/ioutil"
  yaml "gopkg.in/yaml.v2"
)

type Config struct {
  Symbols []string `yaml:"symbols"`
  LiveFrequency int `yaml:"liveFrequency"`
  HistFrequency string `yaml:"histFrequency"`
  HistRange string `yaml:"histRange"`
  Threshold float64 `yaml:"threshold"`
  CarryoverWeight float64 `yaml:"carryoverWeight"`
}

func Parse(pathToConfig string) (conf Config, err error) {
  conf = Config{}

  bytes, err := ioutil.ReadFile(pathToConfig)
  if err != nil {
    fmt.Println(err)
    return conf, err
  }

  err = yaml.Unmarshal(bytes, &conf)
  if err != nil {
    fmt.Println(err)
    return conf, err
  }

  return conf, nil
}
