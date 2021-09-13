package models

import (
	"github.com/mkamadeus/grpc-demo/server/schemas"
)

type StudentController interface {
	GetByNIM(NIM string) (*schemas.StudentResponse, error)
	GetByFaculty(faculty string) ([]*schemas.StudentResponse, error)
}
