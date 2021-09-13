package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/mkamadeus/grpc-demo/server/schemas"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

const (
	address        = "localhost:1337"
	defaultFaculty = "FMIPA"
)

var rootCmd = &cobra.Command{}

func Execute(c schemas.StudentClient) {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "nim",
		Short: "Command to get student by NIM",
		Run: func(cmd *cobra.Command, args []string) {
			GetByNIM(c, args[0])
		},
	})
	rootCmd.AddCommand(&cobra.Command{
		Use:   "faculty",
		Short: "Command to get students by faculty",
		Run: func(cmd *cobra.Command, args []string) {
			faculty := defaultFaculty
			if args[0] != "" {
				faculty = args[0]
			}
			GetByFacultyName(c, faculty)
		},
	})

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func GetByNIM(c schemas.StudentClient, NIM string) {
	res, err := c.GetStudentByNIM(context.Background(), &schemas.StudentByNIMRequest{
		Nim: NIM,
	})

	if err != nil {
		log.Fatalf("Error when calling GetByNIM: %s", err)
	}

	fmt.Println(res)
}

func GetByFacultyName(c schemas.StudentClient, faculty string) {

	// Contact the server and print out its response.
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

func main() {
	//Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := schemas.NewStudentClient(conn)

	// Execute command
	Execute(c)
}
