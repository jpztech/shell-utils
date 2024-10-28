package ugit

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Site struct {
	Name			string `yaml:"name"`
	BaseURL			string `yaml:"base_url"`
	DefaultBranch	string `yaml:"default_branch"`
}
type Config struct {
	CurrentSite	Site	`yaml:"current_site,omitempty"`
	Sites		[]Site	`yaml:"sites"`
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
}

func ApplySite(name string) error {
	for _, site := range config.Sites {
		if site.Name == name {
			config.CurrentSite = site
			return nil
		}
	}
	return fmt.Errorf("Site '%s' not found", name)
}

func GetURL(repo string) string {
	return fmt.Sprintf("%s/%s", config.CurrentSite.BaseURL, "repo")
}

func CurrentSite() string {
	return config.CurrentSite.Name
}