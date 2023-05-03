package main

import (
	"echo-test-1/databases"
	"echo-test-1/routes"
)

func main() {
	databases.StartDB()
	routes.Routes()
}
