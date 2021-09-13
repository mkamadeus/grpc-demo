package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/mkamadeus/grpc-demo/server/schemas"
	"google.golang.org/grpc"
)

const (
	address        = "localhost:1337"
	defaultFaculty = "FMIPA"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := schemas.NewStudentClient(conn)

	// Contact the server and print out its response.
	faculty := defaultFaculty
	if len(os.Args) > 1 {
		faculty = os.Args[1]
	}
	r, err := c.GetStudentsByFaculty(context.Background(), &schemas.StudentByFacultyRequest{Faculty: faculty})
	if err != nil {
		log.Fatalf("Could not get students: %v", err)
	}

	// Read from stream
	done := make(chan bool)
	go func() {
		for {
			resp, err := r.Recv()
			if err == io.EOF {
				done <- true //means stream is finished
				return
			}
			if err != nil {
				log.Fatalf("Cannot receive %v", err)
			}
			log.Printf("Student response received: %s", resp)
		}
	}()
	<-done
}
