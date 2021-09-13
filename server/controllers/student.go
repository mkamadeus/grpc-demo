package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/mkamadeus/grpc-demo/server/models"
	"github.com/mkamadeus/grpc-demo/server/schemas"
	"github.com/mkamadeus/grpc-demo/server/util"
)

func NewStudentController() (models.StudentController, error) {
	facultyMap, err := util.LoadMapping("json/kode_fakultas.json")
	if err != nil {
		return nil, err
	}

	majorMap, err := util.LoadMapping("json/kode_jurusan.json")
	if err != nil {
		return nil, err
	}

	return &StudentController{facultyMap, majorMap}, nil
}

type StudentController struct {
	facultyMap map[string]string
	majorMap   map[string]string
}

func (controller *StudentController) GetByNIM(NIM string) (*schemas.StudentResponse, error) {
	response, err := http.Get("https://cdn.jsdelivr.net/gh/mkamadeus/nim-finder-v2@main/src/json/data_13_21.json")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var students [][]string
	if err := json.Unmarshal(body, &students); err != nil {
		return nil, err
	}

	var studentRes *schemas.StudentResponse
	for _, student := range students {
		var majorNIM string
		if len(student) == 3 {
			majorNIM = student[2]
		}
		if student[1] == NIM || majorNIM == NIM {
			var major string
			if majorNIM != "" {
				major = controller.majorMap[majorNIM[:3]]
			}
			studentRes = &schemas.StudentResponse{
				Name:       student[0],
				FacultyNIM: student[1],
				MajorNIM:   majorNIM,
				Faculty:    controller.facultyMap[student[1][:3]],
				Major:      major,
			}
			break
		}
	}

	return studentRes, nil
}

func (controller *StudentController) GetByFaculty(faculty string, stream schemas.Student_GetStudentsByFacultyServer) (err error) {
	response, err := http.Get("https://cdn.jsdelivr.net/gh/mkamadeus/nim-finder-v2@main/src/json/data_13_21.json")
	if err != nil {
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	var students [][]string
	if err = json.Unmarshal(body, &students); err != nil {
		return
	}

	for _, student := range students {
		var majorNIM string
		if len(student) == 3 {
			majorNIM = student[2]
		}
		if controller.facultyMap[student[1][:3]] == faculty {
			var major string
			if majorNIM != "" {
				major = controller.majorMap[majorNIM[:3]]
			}
			stream.Send(&schemas.StudentResponse{
				Name:       student[0],
				FacultyNIM: student[1],
				MajorNIM:   majorNIM,
				Faculty:    faculty,
				Major:      major,
			})
			time.Sleep(time.Duration(1) * time.Second) // simulate delay
		}
	}

	return nil
}
