// credits: https://github.com/sadlil/go-trigger

package discord

import (
	"errors"
	"reflect"
	"sync"
)

type eventHandler struct {
	functionMap map[string]interface{}
	Mux         sync.Mutex
}

// Trigger ...
type Trigger interface {
	On(event string, task interface{}) error
	Register(event string, params ...interface{}) ([]reflect.Value, error)
}

// Event ...
func Event() Trigger {
	return &eventHandler{
		functionMap: make(map[string]interface{}),
	}
}

var event = Event()

// On listens for the event and executes function once event is called.
func On(eventName string, task interface{}) error {
	return event.On(eventName, task)
}

// Register a new event.
func Register(eventName string, Args ...interface{}) ([]reflect.Value, error) {
	return event.Register(eventName, Args...)
}

func (e *eventHandler) Register(eventName string, Args ...interface{}) (result []reflect.Value, err error) {
	if f, in, err := e.Read(eventName, Args...); err == nil {
		result := f.Call(in)
		return result, nil
	}
	return nil, err
}

func (e *eventHandler) On(eventName string, task interface{}) error {
	e.Mux.Lock()
	defer e.Mux.Unlock()
	if _, ok := e.functionMap[eventName]; ok {
		return errors.New("event already defined")
	}
	if reflect.ValueOf(task).Type().Kind() != reflect.Func {
		return errors.New("task is not a function")
	}
	e.functionMap[eventName] = task
	return nil
}

func (e *eventHandler) Read(eventName string, Args ...interface{}) (reflect.Value, []reflect.Value, error) {
	e.Mux.Lock()
	task, ok := e.functionMap[eventName]
	e.Mux.Unlock()
	if !ok {
		return reflect.Value{}, nil, errors.New("no task found for event")
	}
	f := reflect.ValueOf(task)
	if len(Args) != f.Type().NumIn() {
		return reflect.Value{}, nil, errors.New("parameter mismatched")
	}
	in := make([]reflect.Value, len(Args))
	for k, param := range Args {
		in[k] = reflect.ValueOf(param)
	}
	return f, in, nil
}
