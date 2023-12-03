package services

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Fef0/go-5sim/fivesim"
)

func MakeClient() *fivesim.Client {
	key := os.Getenv("FiveSimKey")
	client := fivesim.NewClient(key)
	return client
}

func OrderActivation(client *fivesim.Client) (*fivesim.ActivationOrder, error) {
	fmt.Println("buying number")
	country := os.Getenv("Country")
	operator := os.Getenv("Operator")
	fmt.Println(country)
	fmt.Println(operator)
	order, err := client.BuyActivationNumber(country, operator, "google", "")
	if err != nil {		
		log.Fatalf("ERROR OCCURED WHILE BUYING A NUMBER : %v", err)
	}

	for order.Status != "RECEIVED" {
		time.Sleep(1 * time.Second)
		log.Printf("Current Status: %v", order.Status)
	}

	return order, err
}

func CancelOrder(client *fivesim.Client, orderID int) (string, error) {
	order, err := client.CancelOrder(orderID)
	if err != nil {
		log.Fatalf("ERROR OCCURED WHILE CANCELING ORDER: %v", err)
	}

	return order.Status, err
}
