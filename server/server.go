package main

import (
	"errors"
	"log"
	"net"

	firebase "firebase.google.com/go"
	"github.com/rajveermalviya/grpc-browser/server/pb"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

// Some global vars
var globalCtx = context.Background()

// get the Firebase Creds file
var opt = option.WithCredentialsFile("creds.json")

// Instantiate the Firebase admin SDK
var app, appErr = firebase.NewApp(globalCtx, nil, opt)

// Instantiate Firestore
var client, dbErr = app.Firestore(globalCtx)

// MyServer is a starting point to define all my services.
type MyServer struct{}

// Check is a gRPC end point to check if an user exists in the database
func (s *MyServer) Check(ctx context.Context, username *pb.CheckUsernameRequest) (*pb.CheckUsernameResponse, error) {
	// get username from the incoming request.
	uname := username.GetUsername()

	if uname != "" {
		// do logging in different goroutine
		go func() { log.Println("Service : Check :", uname) }()

		// get the firestore document
		doc, _ := client.Collection("users").Where("username", "==", uname).Documents(globalCtx).Next()

		// if document is nil, no username exists, username is valid
		if doc == nil {
			return &pb.CheckUsernameResponse{IsValid: true}, nil
		}

		// else return false
		return &pb.CheckUsernameResponse{IsValid: false}, nil
	}

	return nil, errors.New("inavlid username recieved")
}

// GetUser is an API to get userdata given the username.
func (s *MyServer) GetUser(ctx context.Context, username *pb.GetUserDetailsRequest) (*pb.GetUserDetailsResponse, error) {
	// get username from the incoming request.
	uname := username.GetUsername()

	if uname != "" {
		// do logging in different goroutine
		go func() { log.Println("Service : GetUser :", uname) }()

		// get the firestore document
		doc, _ := client.Collection("users").Where("username", "==", uname).Documents(globalCtx).Next()

		// if doc is nil no user exists, so return  an error.
		if doc == nil {
			return nil, errors.New("can't get userdata, no user exists with username : " + uname)
		}

		// else fetch the data
		docData := doc.Data()
		// and return it.
		return &pb.GetUserDetailsResponse{
			Name:     docData["name"].(string),
			Username: docData["username"].(string),
			ImageUrl: docData["img"].(string),
			About:    docData["about"].(string),
		}, nil
	}

	return nil, errors.New("inavlid username recieved")
}

func main() {
	// close the firestore connection at the end of main func
	defer client.Close()

	if appErr != nil {
		log.Fatalf("error initializing app: %v\n", appErr)
	}

	if dbErr != nil {
		log.Fatalf("error initializing firestore: %v\n", dbErr)
	}

	// create a tcp socket connection on port 9090
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterHeuristGrpcServer(grpcServer, &MyServer{})

	log.Println("Starting server on 9090")
	// start the gRPC server
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("error occured in starting gRPC server: %v", err)
	}
}
