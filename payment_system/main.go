package main

import (
	"log"
	"net"

	"github.com/acblacktea/payment_system/payment_system/client"
	"github.com/acblacktea/payment_system/payment_system/dal/repo"
	payment_system "github.com/acblacktea/payment_system/payment_system/kitex_gen/acblacktea/payment_system/payment_system/paymentsystem"
	payment_system_service "github.com/acblacktea/payment_system/payment_system/service/payment_system"
	"github.com/acblacktea/payment_system/payment_system/service/wallet"
	"github.com/cloudwego/kitex/server"
)

func main() {
	db := client.NewDBClient()
	walletRepo := repo.CreateWalletRepo(db)
	walletService := wallet.CreateWalletService(walletRepo)
	paymentService := payment_system_service.CreatePaymentSystemService(walletService)

	addr, _ := net.ResolveTCPAddr("tcp", ":8889")
	svr := payment_system.NewServer(&PaymentSystemImpl{
		PaymentSystemService: paymentService,
	}, server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
