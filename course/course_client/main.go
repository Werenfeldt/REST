package main

import (
	"context"
	"log"

	// "os"
	pd "REST/course/course"
	"time"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:8080"
	defaultName = "course"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pd.NewCoursesClient(conn)

	// Contact the server and print out its response.
	// Id := defaultName
	// if len(os.Args) > 1 {
	// 	name = os.Args[1]
	// }
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AddCourse(ctx, &pd.CourseRequest{Id: "6666"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetId())
}