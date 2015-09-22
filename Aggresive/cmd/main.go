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
		time.Sleep(5 * time.Second)
		fmt.Println("Scanning...")
		ships, sectors, err := scan(ship)
		if err != nil {
			fmt.Printf("Unable to scan: %s\n", err.Error())
		} else {
			index := 0
			if len(ships) > 1 {
				index = rand.Intn(len(ships) - 1)
			}
			if len(ships) > 0 {
				target := ships[index]
				fmt.Printf("Selected target ship %s\n", target.Name)
				err := fireLoop(ship, target)
				if err != nil {
					fmt.Printf("Failed to attack %s, %s\n", target.Name, err.Error())
				}
			} else {
				index := 0
				if len(sectors) > 1 {
					index = rand.Intn(len(sectors) - 1)
				}
				if len(sectors) > 0 {
					nextSector := sectors[index]
					if err := travel(ship, nextSector); err != nil {
						fmt.Printf("Failed to travel: %s\n", err.Error())
					}
				}

			}
		}
	}
}

func travel(ship *dtl.Ship, sector *dtl.Sector) error {
	for {
		if ship.CanTravel() {
			fmt.Printf("Travelling to %s\n", sector.Name)
			return ship.Travel(sector)
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

func fireLoop(ship, target *dtl.Ship) error {
	for {
		if ship.CanFire() {
			fmt.Printf("Fixing on %s\n", target.Name)
			if err := ship.Fire(target.Name); err != nil {
				fmt.Println("Fail!")
				return err
			} else {
				fmt.Println("HIT!")
			}
		}
		time.Sleep(1 * time.Second)
	}
}
