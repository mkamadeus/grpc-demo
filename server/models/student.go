package models

import (
	"github.com/mkamadeus/grpc-demo/server/schemas"
)

type StudentController interface {
	GetByNIM(NIM string) (*schemas.StudentResponse, error)
	GetByFaculty(faculty string, stream schemas.Student_GetStudentsByFacultyServer) error
}
