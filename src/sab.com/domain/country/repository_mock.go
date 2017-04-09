package country

import (
	"github.com/stretchr/testify/mock"
)

type CountryRepositoryMock struct {
	mock.Mock
}

func (m CountryRepositoryMock) Save(ctr *Country) error {
	args := m.Called(ctr)
	return args.Error(0)
}

func (m CountryRepositoryMock) GetByCode(code string) (Country, error) {
	args := m.Called(code)
	return args.Get(0).(Country), args.Error(1)
}

func (m CountryRepositoryMock) GetAll() ([]Country, error) {
	args := m.Called()
	return args.Get(0).([]Country), args.Error(1)
}

func (m CountryRepositoryMock) HasCountryWithCode(code string) (bool, error) {
	args := m.Called(code)
	return args.Bool(0), args.Error(1)
}
