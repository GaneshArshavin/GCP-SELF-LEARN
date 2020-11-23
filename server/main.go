package main

import (
	"flag"
	"log"
	"net"
	"time"

	pg "github.com/carousell/gcp-self-study/pg"
	pb "github.com/carousell/gcp-self-study/proto"
	"github.com/carousell/gcp-self-study/redis"
	"github.com/carousell/gcp-self-study/service"
	store "github.com/carousell/gcp-self-study/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 8080, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterUserLoginServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}

func newServer() service.UserLoginServer {
	s := &service.Svc{}
	pgMasterConfig := &pg.Config{
		Host:               "35.185.179.122",
		Port:               5432,
		User:               "rajagopalganesh",
		Password:           "Welcome@135",
		DBName:             "gcp-self-study_user",
		MaxIdleConnections: 10,
		MaxOpenConnections: 10,
	}
	pgSlaveConfig := &pg.Config{
		Host:               "35.185.179.122",
		Port:               5432,
		User:               "rajagopalganesh",
		DBName:             "gcp-self-study_user",
		Password:           "Welcome@135",
		MaxIdleConnections: 10,
		MaxOpenConnections: 10,
	}

	redisConfig := &redis.Config{
		//disabling it as GCP Is costing me will turn it on on request
		//
		//Host:        "10.103.214.219",
		Host:        "127.0.0.1",
		Port:        6379,
		DialTimeout: time.Duration(1000) * time.Millisecond,
	}
	var err error
	s.Storage, err = store.NewClient(pgMasterConfig, pgSlaveConfig, redisConfig)
	if err != nil {
		panic(err)
	}
	return s
}
