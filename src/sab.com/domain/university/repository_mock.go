package university

import "github.com/stretchr/testify/mock"

type UniversityRepositoryMock struct {
	mock.Mock
}

func (m *UniversityRepositoryMock) Save(univ *University, countryCode string) error {
	args := m.Called(univ, countryCode)
	return args.Error(0)
}

func (m *UniversityRepositoryMock) GetById(id int64, countryCode string) (University, error) {
	args := m.Called(id, countryCode)
	return args.Get(0).(University), args.Error(1)
}

func (m *UniversityRepositoryMock) GetAll(countryCode string) ([]University, error) {
	args := m.Called(countryCode)
	return args.Get(0).([]University), args.Error(1)
}
