package cfg

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Serv struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Db struct {
		Host     string `yaml:"host"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Dbname   string `yaml:"name"`
		Port     string `yaml:"port"`
		Sslmode  string `yaml:"sslmode"`
	} `yaml:"db"`
}

func (config *Config) ReadYaml(filepath string) error {
	yamlFile, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("error Reading Config file with path: %s. %s", filepath, err.Error())
	}
	if err = yaml.Unmarshal(yamlFile, config); err != nil {
		return err
	}
	return nil
}
func (dbc *Config) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sswlmode=%s",
		dbc.Db.Host, dbc.Db.User, dbc.Db.Password, dbc.Db.Dbname, dbc.Db.Port, dbc.Db.Sslmode)
}
