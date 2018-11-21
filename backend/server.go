package main

import (
	"errors"
	"log"
	"net"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/rajveermalviya/grpc-browser/backend/pb"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

// HeuristGRPC is a starting point to define all my services.
type HeuristGRPC struct{ db *firestore.Client }

// Check is a gRPC end point to check if an user exists in the database
func (s *HeuristGRPC) Check(ctx context.Context, username *pb.CheckUsernameRequest) (*pb.CheckUsernameResponse, error) {
	// get username from the incoming request.
	uname := username.GetUsername()

	if uname != "" {
		log.Println("Service : Check :", uname)

		// get the firestore document
		doc, _ := s.db.Collection("users").Where("username", "==", uname).Documents(ctx).Next()

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
func (s *HeuristGRPC) GetUser(ctx context.Context, username *pb.GetUserDetailsRequest) (*pb.GetUserDetailsResponse, error) {
	// get username from the incoming request.
	uname := username.GetUsername()

	if uname != "" {
		log.Println("Service : GetUser :", uname)

		// get the firestore document
		doc, _ := s.db.Collection("users").Where("username", "==", uname).Documents(ctx).Next()

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

	return nil, errors.New("invalid username recieved")
}

func main() {
	ctx := context.Background()

	// get the Firebase Creds env var.
	opt := option.WithCredentialsJSON([]byte(os.Getenv("FIREBASE_CREDENTIALS")))

	// Instantiate the Firebase admin SDK
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	// Instantiate Firestore
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("error initializing firestore: %v\n", err)
	}

	// close the firestore connection at the end of main func
	defer client.Close()

	// create a tcp socket connection on port 9090
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalln("failed to listen:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterHeuristGrpcServer(grpcServer, &HeuristGRPC{db: client})

	log.Println("Starting server on 9090")

	// start the gRPC server
	log.Fatalln(grpcServer.Serve(lis))
}
