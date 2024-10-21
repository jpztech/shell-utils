package ugit

import (
	"errors"
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

func FromFile(file string) *Config {
	c := &Config{}
	// if the file does not exists, return an empty map
	if _, err := os.Stat(file); os.IsNotExist(err) {
		fmt.Printf("File not exists")
	}
	// read the yaml file and deserialize it into a map
	yamlFile, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Error reading YAML file %s: %v\n", file, err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		fmt.Printf("Error parsing YAML file %s: %v\n", file, err)
	}
	return c
}

func (c *Config) ApplySite(name string) (*Config, error) {
	for _, site := range c.Sites {
		if site.Name == name {
			c.CurrentSite = site
			return c, nil
		}
	}
	return nil, errors.New("Site not found")
}

func (c *Config) GetURL(repo string) string {
	return fmt.Sprintf("%s/%s", c.CurrentSite.BaseURL, "repo")
}