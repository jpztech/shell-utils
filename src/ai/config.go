package ai

import (
	"bytes"
	"fmt"
	"os"

	"text/template"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Model       string            `yaml:"model"`
	UrlTemplate string            `yaml:"url_template"`
	Properties  map[string]string `yaml:"properties"`
	Proxy	   	string            `yaml:"proxy"`
	URL		 	string            
}

var config *Config

func ReadConfigFromFile(file string) {
	config = &Config{}
	// if the file does not exists, return an empty map
	if _, err := os.Stat(file); os.IsNotExist(err) {
		fmt.Printf("File not exists")
	}
	// read the yaml file and deserialize it into a map
	yamlFile, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Error reading YAML file %s: %v\n", file, err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Printf("Error parsing YAML file %s: %v\n", file, err)
	}

	t := template.Must(template.New("url").Parse(config.UrlTemplate))
	var urlBuffer bytes.Buffer
	err = t.Execute(&urlBuffer, config.Properties)
	if err != nil {
		fmt.Printf("Error executing template: %v\n", err)
	}
	config.URL = urlBuffer.String()
	InitAIClient(config)
}