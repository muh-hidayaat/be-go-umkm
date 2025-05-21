package helpers

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateTrackingNumber generates a tracking number in the format "DSI ddmmyy XXXXX"
func GenerateTrackingNumber() string {
	currentDate := time.Now().Format("020106") // Format date as ddmmyy
	randomNumber := rand.Intn(100000)          // Generate a random 5-digit number
	return fmt.Sprintf("DSI %s %05d", currentDate, randomNumber)
}
