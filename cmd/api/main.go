package main

import (
	"flag"
	"fmt"

	metadataService "github.com/paisit04/building-modern-cli/services/metadata"
)

func main() {
	var port int
	flag.IntVar(&port, "p", 8000, "Port to metadata service")
	flag.Parse()
	fmt.Printf("Starting API at http://localhost:%d\n", port)
	metadataService.Run(port)
}
