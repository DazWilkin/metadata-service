package main

import (
	"log"

	"cloud.google.com/go/compute/metadata"
)

func main() {
	instanceID, err := metadata.InstanceID()
	if err != nil {
		log.Println("Error getting instance ID:", err)
		instanceID = "unknown"
	}
	zone, err := metadata.Zone()
	if err != nil {
		log.Println("Error getting zone:", err)
		zone = "unknown"
	}
	projectID, err := metadata.ProjectID()
	if err != nil {
		log.Println("Error getting Project ID: ", err)
		projectID = "unknown"
	}

	log.Printf("[%s] %s/%s", projectID, zone, instanceID)
}
