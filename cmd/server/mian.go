package main

import (
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"kobe/api"
	"kobe/pkg/config"
	"kobe/server"
	"log"
	"net"
)

func main() {
	config.InitConfig()
	address := fmt.Sprintf("%s:%d",
		viper.GetString("server.host"),
		viper.GetInt("server.port"))
	s, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	gs := grpc.NewServer()
	api.RegisterKobeApiServer(gs, &server.Kobe{})
	if err := gs.Serve(s); err != nil {
		log.Fatal(err)
	}
}