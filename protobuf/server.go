package main

import (
	"fmt"
	"net"
	"sync"

	api "github.com/KarelKubat/go-demo/protobuf/pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type firstlast struct {
	first, last string
}

type server struct {
	mutex   sync.Mutex
	persons []firstlast
}

func (s *server) Add(ctx context.Context, in *api.Person) (*api.AddResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, p := range s.persons {
		if p.first == in.GetFirstName() && p.last == in.GetLastName() {
			return nil, fmt.Errorf("person already known")
		}
	}

	s.persons = append(s.persons, firstlast{
		first: in.GetFirstName(),
		last:  in.GetLastName(),
	})
	return &api.AddResponse{
		Count: int64(len(s.persons)),
	}, nil
}

func (s *server) List(ctx context.Context, in *api.ListRequest) (*api.ListResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	resp := &api.ListResponse{}
	for _, p := range s.persons {
		resp.Persons = append(resp.Persons, &api.Person{
			FirstName: p.first,
			LastName:  p.last,
		})
	}
	return resp, nil
}

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	s := &server{}
	grpcServer := grpc.NewServer()
	api.RegisterHRServer(grpcServer, s)

	fmt.Printf("server starting on localhost:1234\n")
	if err = grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
