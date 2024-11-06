/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"github.com/tyagnii/gw-exchanger/config"
	"github.com/tyagnii/gw-exchanger/gen/exchanger/v1"
	"github.com/tyagnii/gw-exchanger/internal/db"
	"github.com/tyagnii/gw-exchanger/internal/logger"
	"github.com/tyagnii/gw-exchanger/internal/server"
	"google.golang.org/grpc"
	"net"
	"time"

	"os"

	"github.com/spf13/cobra"
)

var configFile = "config.env"

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
		var ctx context.Context = context.Background()
		var startUpTimeout time.Duration
		var DBConn db.DBConnector

		// Fetch startUpTimeout
		startUpTimeout, err := time.ParseDuration(os.Getenv("START_UP_TIMEOUT"))
		if err != nil {
			startUpTimeout = 30 * time.Second
		}

		// Init Logger
		sLogger, err := logger.NewSugaredLogger()
		if err != nil {
			panic(err)
		}

		// Read configuration from env file to os environment variables
		if err := config.ReadConfig(configFile); err != nil {
			sLogger.DPanicf("Configuration file not found: %s", configFile)
			os.Exit(1)
		}

		// Fetch server address string
		addr := os.Getenv("EXCHANGE_SEVER_ADDRESS_STRING")

		// Create dbconnection
		ctx, cancel := context.WithTimeout(ctx, startUpTimeout)
		defer cancel()
		connString := config.BuildConnString()
		DBConn, err = db.NewPGConnector(ctx, connString, sLogger)
		if err != nil {
			sLogger.DPanicf("Database connection error: %s", err.Error())
			os.Exit(1)
		}

		// Create Exchanger Server instance
		exchangeServer := server.NewExchangeServer(DBConn, addr, sLogger)
		sLogger.Debugf("Created exchange server")

		// Listener configuration for gRPC connection
		tcpListen, err := net.Listen("tcp", addr)
		if err != nil {
			sLogger.DPanicf("Error listening on %s: %v", addr, err)
			panic(err)
		}

		// Create new gRPC server to handle services
		grpcServer := grpc.NewServer()

		// Register service on gRPC server
		exchanger.RegisterExchangeServiceServer(grpcServer, exchangeServer)

		// Run gRPC server
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
