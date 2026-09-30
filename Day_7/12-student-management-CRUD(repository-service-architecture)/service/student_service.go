package service
import (
	"student_app/model"
	"student_app/repository"
	"student_app/utility"
)

type StudentService interface {
	Save(student model.Student) error
	FindByID(id int) (model.Student, error)
	FindAll() ([]model.Student, error)
	Update(student model.Student) error
	Delete(id int) error
}

type StudentServiceImpl struct {
	repository repository.StudentRepository
}

func NewStudentService(
	repository repository.StudentRepository,
) StudentService {

	return &StudentServiceImpl{
		repository: repository,
	}
}

func (s *StudentServiceImpl) Save(
	student model.Student,
) error {

	err := utility.ValidateStudent(student)

	if err != nil {
		return err
	}

	err = utility.ValidateEmail(student.Email)

	if err != nil {
		return err
	}

	return s.repository.Save(student)
}

func (s *StudentServiceImpl) FindByID(
	id int,
) (model.Student, error) {

	err := utility.ValidateID(id)

	if err != nil {
		return model.Student{}, err
	}

	return s.repository.FindByID(id)
}

func (s *StudentServiceImpl) FindAll() (
	[]model.Student,
	error,
) {

	return s.repository.FindAll()
}

func (s *StudentServiceImpl) Update(
	student model.Student,
) error {

	err := utility.ValidateID(student.ID)

	if err != nil {
		return err
	}

	err = utility.ValidateStudent(student)

	if err != nil {
		return err
	}

	err = utility.ValidateEmail(student.Email)

	if err != nil {
		return err
	}

	return s.repository.Update(student)
}

func (s *StudentServiceImpl) Delete(
	id int,
) error {

	err := utility.ValidateID(id)

	if err != nil {
		return err
	}

	return s.repository.Delete(id)
}
