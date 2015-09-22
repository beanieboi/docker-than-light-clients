package main

import (
	"fmt"
	"math/rand"
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
			fmt.Printf("Unable to scan: %s\n", err.Error())
		} else {
			index := 0
			if len(sectors) > 1 {
				index = rand.Intn(len(sectors) - 1)
			}
			if len(sectors) > 0 {
				escapeSector := sectors[index]
				fmt.Printf("Selected escape sector %s\n", escapeSector.Name)
				if len(ships) > 0 {
					fmt.Println("ALERT! Detected Ship, evac!!!")
					evac(ship, escapeSector)
				}
			}
		}
		time.Sleep(10 * time.Second)
	}
}

func evac(ship *dtl.Ship, escapeSector *dtl.Sector) error {
	for {
		if ship.CanTravel() {
			fmt.Printf("Evaccing to %s\n", escapeSector.Name)
			return ship.Travel(escapeSector)
		}
		time.Sleep(1 * time.Second)
	}
}

func scan(ship *dtl.Ship) ([]*dtl.Ship, []*dtl.Sector, error) {
	for {
		if ship.CanScan() {
			return ship.ScanSector()
		}
		time.Sleep(1 * time.Second)
	}
}
