package config

import {
	"log"
	"github.com/BurntSushi/toml"
}

type config struct{
	Server string
	Database string
}
func (c *Config) Read(){
	if err :=toml.DecodeFile("config.toml", &c); err !=nil{
		log.Fatal(err)
	}
}