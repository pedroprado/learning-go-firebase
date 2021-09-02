package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"leaning.go.firebase/src/apis"
	"leaning.go.firebase/src/infra"
	"leaning.go.firebase/src/repository"
	"log"
)

var (
	dbProject = "xxx";
)

func main(){

	ctx := context.Background()
	db := getFirestoreDB(ctx)

	userRepo := repository.NewUserRepository(db)

	server := infra.NewServerHttp()
	group := server.GetGinRouterGroup("/go-fire")

	apis.RegisterUserApi(group, userRepo)

	err := server.StartServer("7800")
	if err != nil {
		log.Fatal(err)
	}
}

func getFirestoreDB(ctx context.Context) *firestore.Client{
	client, err := firestore.NewClient(ctx, dbProject)
	if err != nil {
		fmt.Println("Fatal db not connected")
	}

	return client
}