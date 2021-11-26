package main

import (
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/mock"
)

type MockedClient struct {
	mock.Mock
}

func (m *MockedClient) SetValue(key string, value interface{}, expiration time.Duration) error {
	return m.Called(key, value, expiration).Get(0).(*redis.StatusCmd).Err()
}

func TestSet(t *testing.T) {
	mockClient := new(MockedClient)

	mockClient.On("SetValue", "demo", "test", time.Duration(0)*time.Second).Return(&redis.StatusCmd{})

	err := mockClient.SetValue("demo", "test", time.Duration(0)*time.Second)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSetValue(t *testing.T) {
	db := new(MockedClient)
	k := "demo"
	v := "test"

	db.On("SetValue", k, v, time.Duration(0)*time.Second).Return(&redis.StatusCmd{})

	err := db.SetValue(k, v, time.Duration(0)*time.Second)
	if err != nil {
		t.Error(err)
	}
}
