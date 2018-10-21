package main

import (
	"log"
	"net"

	firebase "firebase.google.com/go"
	"github.com/rajveermalviya/grpc-browser/server/pb"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

var globalCtx = context.Background()
var opt = option.WithCredentialsFile("creds.json")
var app, appErr = firebase.NewApp(globalCtx, nil, opt)
var client, dbErr = app.Firestore(globalCtx)

type myServer struct{}

func (s *myServer) Check(ctx context.Context, username *pb.CheckUsernameRequest) (*pb.CheckUsernameResponse, error) {
	uname := username.GetUsername()
	log.Println(uname)

	doc, _ := client.Collection("users").Where("username", "==", uname).Documents(globalCtx).Next()

	isValid := false
	if doc == nil {
		isValid = true
	}

	return &pb.CheckUsernameResponse{IsValid: isValid}, nil
}

func (s *myServer) GetUser(ctx context.Context, username *pb.GetUserDetailsRequest) (*pb.GetUserDetailsResponse, error) {
	uname := username.GetUsername()
	log.Println(uname)

	doc, _ := client.Collection("users").Where("username", "==", uname).Documents(globalCtx).Next()

	if doc == nil {
		return &pb.GetUserDetailsResponse{
			Name:     "No User",
			Username: "No User",
			ImageUrl: "No User",
			About:    "No User",
		}, nil
	}

	docData := doc.Data()

	userDetails := &pb.GetUserDetailsResponse{
		Name:     docData["name"].(string),
		Username: docData["username"].(string),
		ImageUrl: docData["img"].(string),
		About:    docData["about"].(string),
	}

	return userDetails, nil
}

func main() {

	defer client.Close()

	if appErr != nil {
		log.Fatalf("error initializing app: %v\n", appErr)
	}

	if dbErr != nil {
		log.Fatalf("error initializing firestore: %v\n", dbErr)
	}

	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterHeuristGrpcServer(grpcServer, &myServer{})

	log.Println("Starting server on 9090")
	grpcServer.Serve(lis)
}
