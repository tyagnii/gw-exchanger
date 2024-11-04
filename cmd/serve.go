/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/tyagnii/gw-exchanger/config"
	"github.com/tyagnii/gw-exchanger/gen/exchanger/v1"
	"github.com/tyagnii/gw-exchanger/internal/db"
	"github.com/tyagnii/gw-exchanger/internal/server"
	"google.golang.org/grpc"
	"net"

	"os"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")

		if err := config.ReadConfig("config.env"); err != nil {
			panic(err)
		}

		// fetch server address string
		addr := os.Getenv("EXCHANGE_SEVER_ADDRESS_STRING")

		// create db connection
		dbconn, err := db.NewPGConnector(context.Background(), "")
		if err != nil {
			panic(err)
		}
		exchangeServer := server.NewExchangeServer(dbconn, addr)

		// Listener configuration for gRPC connection
		tcpListen, err := net.Listen("tcp", addr)
		if err != nil {
			panic(err)
		}

		// Create new gRPC server to handle services
		grpcServer := grpc.NewServer()

		// Register service on gRPC server
		exchanger.RegisterExchangeServiceServer(grpcServer, exchangeServer)

		grpcServer.Serve(tcpListen)

	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
