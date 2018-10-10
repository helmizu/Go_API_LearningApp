package library

import (
	"log"

	"PW/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserCtr to export func
type UserCtr struct {
	Server   string
	Database string
}

var sess *mgo.Session
var db *mgo.Database

const (
	cGuru     = "guru"
	cStudents = "siswa"
)

// Connect use to connect MongoDB
func (m *UserCtr) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	sess = session
	db = session.DB(m.Database)
}

// FindAllGuru use to Find list of users
func (m *UserCtr) FindAllGuru() ([]models.GuruPrimary, error) {
	var guru []models.GuruPrimary
	err := db.C(cGuru).Find(bson.M{}).All(&guru)
	return guru, err
}

// FindGuruByID use to Find a guru by its id
func (m *UserCtr) FindGuruByID(id string) (models.GuruPrimary, error) {
	var guru models.GuruPrimary

	err := db.C(cGuru).FindId(bson.ObjectIdHex(id)).One(&guru)

	return guru, err
}

// SignupGuru a user into database
func (m *UserCtr) SignupGuru(guru models.GuruPrimary) error {
	err := db.C(cGuru).Insert(&guru)
	return err
}

// DeleteGuru an existing guru
func (m *UserCtr) DeleteGuru(guru models.GuruPrimary) error {
	err := db.C(cGuru).Remove(&guru)
	return err
}

// UpdateGuru an existing user
func (m *UserCtr) UpdateGuru(guru models.GuruPrimary) error {
	err := db.C(cGuru).UpdateId(guru.ID, &guru)
	return err
}

// FindAllStudents use to Find list of users
func (m *UserCtr) FindAllStudents() ([]models.SiswaPrimary, error) {
	var students []models.SiswaPrimary
	err := db.C(cStudents).Find(bson.M{}).All(&students)
	return students, err
}

// FindStudentByID use to Find a student by its id
func (m *UserCtr) FindStudentByID(id string) (models.SiswaPrimary, error) {
	var student models.SiswaPrimary
	err := db.C(cStudents).FindId(bson.ObjectIdHex(id)).One(&student)
	return student, err
}

// SignupStudent a user into database
func (m *UserCtr) SignupStudent(student models.SiswaPrimary) error {
	err := db.C(cStudents).Insert(&student)
	return err
}

// DeleteStudent an existing student
func (m *UserCtr) DeleteStudent(student models.SiswaPrimary) error {
	err := db.C(cStudents).Remove(&student)
	return err
}

// UpdateStudent an existing user
func (m *UserCtr) UpdateStudent(student models.SiswaPrimary) error {
	err := db.C(cStudents).UpdateId(student.ID, &student)
	return err
}
