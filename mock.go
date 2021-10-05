package json2keys

import "github.com/stretchr/testify/mock"

type Json2KeysMock struct {
	mock.Mock
}

func (m Json2KeysMock) Parse(jsonStr string) (keys map[string]interface{}, err error) {
	args := m.Called(jsonStr)
	return args.Get(0).(map[string]interface{}), args.Error(1)
}
