// Code generated by mockery v1.0.0. DO NOT EDIT.
package kafkaclient

import mock "github.com/stretchr/testify/mock"
import sarama "github.com/Shopify/sarama"

// MockAsyncProducer is an autogenerated mock type for the MockAsyncProducer type
type MockAsyncProducer struct {
	mock.Mock
}

// AsyncClose provides a mock function with given fields:
func (_m *MockAsyncProducer) AsyncClose() {
	_m.Called()
}

// Close provides a mock function with given fields:
func (_m *MockAsyncProducer) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Errors provides a mock function with given fields:
func (_m *MockAsyncProducer) Errors() <-chan *sarama.ProducerError {
	ret := _m.Called()

	var r0 <-chan *sarama.ProducerError
	if rf, ok := ret.Get(0).(func() <-chan *sarama.ProducerError); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *sarama.ProducerError)
		}
	}

	return r0
}

// Input provides a mock function with given fields:
func (_m *MockAsyncProducer) Input() chan<- *sarama.ProducerMessage {
	ret := _m.Called()

	var r0 chan<- *sarama.ProducerMessage
	if rf, ok := ret.Get(0).(func() chan<- *sarama.ProducerMessage); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan<- *sarama.ProducerMessage)
		}
	}

	return r0
}

// Successes provides a mock function with given fields:
func (_m *MockAsyncProducer) Successes() <-chan *sarama.ProducerMessage {
	ret := _m.Called()

	var r0 <-chan *sarama.ProducerMessage
	if rf, ok := ret.Get(0).(func() <-chan *sarama.ProducerMessage); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *sarama.ProducerMessage)
		}
	}

	return r0
}
