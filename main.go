package main

import (
    "fmt"
    "time"
)


func main() {

	var costPerMessage float32 = 0.2
	var messageLimit int = 1000000
	var hasPermission bool
	var totalCost float32 = 0.0
	var username string

	const premiumPlanName = "premium plan"




	fmt.Println("Simple messaging library and platform specific")
	fmt.Printf("Cost per message: %.2f USD\n", costPerMessage)
	fmt.Printf("Message limit: %d\n", messageLimit)
	fmt.Println("\n--- Simulating messages....\n", time.Now())

	fmt.Println("Your subscription plan is :",premiumPlanName)
	
	fmt.Sprintf("\n --- %f, %d, %v, %d, %s",  costPerMessage, messageLimit, hasPermission, totalCost, username)
}
