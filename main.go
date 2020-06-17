package main

import (
	"test/cmd"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	cmd.Execute()
}
