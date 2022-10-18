package main

import (
	"github.com/high-performance-payment-gateway/balance-service/balance"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("devops/.env")
	balanceModule := balance.NewModule()
	balanceModule.Start()
}
