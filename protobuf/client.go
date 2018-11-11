package main

import (
	"fmt"
	"os"

	api "github.com/KarelKubat/go-demo/protobuf/pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) == 4 && os.Args[1] == "add" {
		add(os.Args[2], os.Args[3])
	} else if len(os.Args) == 2 && os.Args[1] == "list" {
		list()
	} else {
		panic(`
Usage: go run client.go add FIRST LAST - add a person to the list
   or: go run client.go list           - list known persons
`)
	}
}

func getClient() api.HRClient {
	conn, err := grpc.Dial(":1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return api.NewHRClient(conn)
}

func add(first, last string) {
	client := getClient()
	resp, err := client.Add(context.Background(), &api.Person{
		FirstName: first,
		LastName:  last,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d persons added so far\n", resp.GetCount())
}

func list() {
	client := getClient()
	resp, err := client.List(context.Background(), &api.ListRequest{})
	if err != nil {
		panic(err)
	}
	for _, p := range resp.GetPersons() {
		fmt.Printf("%s %s\n", p.GetFirstName(), p.GetLastName())
	}
}
