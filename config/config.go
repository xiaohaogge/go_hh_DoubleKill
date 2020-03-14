package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"

)

var Setting *Conf

type Conf struct {
	Server *serverModel `yaml:"server"`
}

type serverModel struct {
	Mode string `yaml:mode`
	Host string `yaml:host`
	Port string `yaml:port`
}

//LoadConfigInformation load config information for application
func InitConfig() {
	filepath,_ := os.Getwd()
	fileFullPath := path.Join(filepath ,"\\go_sugared\\config\\dev_config.yaml")
	fmt.Println("fileFullPath: ",fileFullPath)
	configData, err := ioutil.ReadFile(fileFullPath)
	if err != nil {
		fmt.Printf(" config file read failed: %s", err)
		os.Exit(-1)
	}
	var configApplication *Conf

	err = yaml.Unmarshal(configData, &configApplication)
	if err != nil {
		fmt.Printf(" config parse failed: %s", err)
		os.Exit(-1)
	}
	Setting = configApplication
}
