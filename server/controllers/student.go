package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

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
		if student[1] == NIM || student[2] == NIM {
			studentRes = &schemas.StudentResponse{
				Name:       student[0],
				FacultyNIM: student[1],
				MajorNIM:   student[2],
				Faculty:    controller.facultyMap[student[1][:3]],
				Major:      controller.majorMap[student[2][:3]],
			}
			break
		}
	}

	return studentRes, nil
}

func (controller *StudentController) GetByFaculty(faculty string) ([]*schemas.StudentResponse, error) {
	return nil, errors.New("Not Implemented yet")
}
