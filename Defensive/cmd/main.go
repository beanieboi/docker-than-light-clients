package main

import (
	"fmt"
	"time"

	dtl "github.com/bfosberry/docker-than-light-client-go"
)

func main() {
	fmt.Println("Starting..")
	ship, err := dtl.New()
	if err != nil {
		fmt.Println("Failed to create ship")
		panic(err)
	}

	for {
		fmt.Println("Scanning...")
		ships, sectors, err := scan(ship)
		if err != nil {
			fmt.Println("Unable to scan")
		} else {
			escapeSector := sectors[0]
			fmt.Printf("Selected escape sector %s", escapeSector.Name)
			if len(ships) > 0 {
				fmt.Println("ALERT! Detected Ship, evac!!!")
				evac(ship, escapeSector)
			}
		}
	}
}

func evac(ship dtl.Ship, escapeSector dtl.Sector) error {
	for {
		if ship.CanTravel() {
			return ship.Travel(escapeSector)
		}
		time.Sleep(1 * time.Second)
	}
}

func scan(ship dtl.Ship) ([]dtl.Ship, []dtl.Sector, error) {
	for {
		if ship.CanScan() {
			return ship.ScanSector()
		}
		time.Sleep(1 * time.Second)
	}
}
