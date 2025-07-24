package main

import (
	"countreez/configs"
)

func main() {
	configs.GetConnection()
	defer configs.DB.Close()
}