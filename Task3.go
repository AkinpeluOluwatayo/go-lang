package assignment

import "fmt"

func GasMileage() {
	var miles, gallons int
	var totalMiles, totalGallons int

	for {
		fmt.Print("Enter miles driven (or -1 to quit): ")
		fmt.Scan(&miles)
		if miles == -1 {
			break
		}

		fmt.Print("Enter gallons used: ")
		fmt.Scan(&gallons)

		milesPerGallon := float64(miles) / float64(gallons)
		fmt.Printf("MPG for this trip: %.2f\n", milesPerGallon)

		totalMiles += miles
		totalGallons += gallons

		totalMPG := float64(totalMiles) / float64(totalGallons)
		fmt.Printf("Combined MPG so far: %.2f\n\n", totalMPG)
	}
}
