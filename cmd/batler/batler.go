package main

import (
	"flag"
	"log"

	"github.com/vitorbaraujo/batler/configuration"
	"github.com/vitorbaraujo/batler/xcodebuild"
)

var config *configuration.Configuration

func init() {
	projectPath := flag.String("p", "", "Project path to run batler")
	flag.Parse()

	var err error
	config, err = configuration.FetchConfiguration(*projectPath)
	if err != nil {
		log.Fatalf("could not fetch configuration: %v", err)
	}
}

func main() {
	client, err := xcodebuild.NewClient(config,
		xcodebuild.WithClean(), xcodebuild.WithBuild(), xcodebuild.WithTest())
	if err != nil {
		log.Fatalf("creating xcodebuild client: %v", err)
	}

	if err := client.Run(); err != nil {
		log.Printf("could not run client: %v", err)
	}
}
