package servers

import (
	"context"

	"github.com/mkamadeus/grpc-demo/server/models"
	"github.com/mkamadeus/grpc-demo/server/schemas"
)

type StudentServer struct {
	StudentCtl models.StudentController
	schemas.UnimplementedStudentServer
}

func (server *StudentServer) GetStudentByNIM(ctx context.Context, req *schemas.StudentByNIMRequest) (*schemas.StudentResponse, error) {
	return server.StudentCtl.GetByNIM(req.GetNim())
}

func (server *StudentServer) GetStudentsByFaculty(req *schemas.StudentByFacultyRequest, srv schemas.Student_GetStudentsByFacultyServer) error {
	data, err := server.StudentCtl.GetByFaculty(req.GetFaculty())
	if err != nil {
		return err
	}
	for i := 0; i < 10; i++ {
		srv.Send(data[i])
	}
	return nil
}
