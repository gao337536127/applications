package main

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/gao337536127/configuration"
	"github.com/gao337536127/status"
)

//go:embed test.ini
var defaultConfig []byte

func init() {
	configuration.AppendDefaultConfigurationBytes(defaultConfig)
}

func main() {
	iniConfig := configuration.IniConfig{}
	fmt.Println(iniConfig.GetConfigWithEnvironment("fff", "common", "master", "c"))

	status.DefaultStatus.SetStatus("time", func() (string, error) {
		return fmt.Sprintf("%d", time.Now().Unix()), nil
	})
	for i := 0; i < 10; i++ {
		s, _ := status.DefaultStatus.GetStatus("time")
		fmt.Printf("s: %v\n", s)
		time.Sleep(time.Second)
	}
}
