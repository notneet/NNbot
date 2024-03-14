package main

import (
	"fmt"

	"github.com/notneet/NNbot/pkg/common"
)

func main() {
	config, err := common.LoadConfig(common.DEFAULT_CONFIG_DIR)
	common.FatalIfError(err)

	fmt.Println("Hello World", config.Database.Password)
}
