// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package event

import (
	"context"
	"sync"
)

// Ensure, that RecorderMock does implement Recorder.
// If this is not the case, regenerate this file with moq.
var _ Recorder = &RecorderMock{}

// RecorderMock is a mock implementation of Recorder.
//
//	func TestSomethingThatUsesRecorder(t *testing.T) {
//
//		// make and configure a mocked Recorder
//		mockedRecorder := &RecorderMock{
//			RecordEventFunc: func(contextMoqParam context.Context, event Event)  {
//				panic("mock out the RecordEvent method")
//			},
//		}
//
//		// use mockedRecorder in code that requires Recorder
//		// and then make assertions.
//
//	}
type RecorderMock struct {
	// RecordEventFunc mocks the RecordEvent method.
	RecordEventFunc func(contextMoqParam context.Context, event Event)

	// calls tracks calls to the methods.
	calls struct {
		// RecordEvent holds details about calls to the RecordEvent method.
		RecordEvent []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// Event is the event argument value.
			Event Event
		}
	}
	lockRecordEvent sync.RWMutex
}

// RecordEvent calls RecordEventFunc.
func (mock *RecorderMock) RecordEvent(contextMoqParam context.Context, event Event) {
	if mock.RecordEventFunc == nil {
		panic("RecorderMock.RecordEventFunc: method is nil but Recorder.RecordEvent was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		Event           Event
	}{
		ContextMoqParam: contextMoqParam,
		Event:           event,
	}
	mock.lockRecordEvent.Lock()
	mock.calls.RecordEvent = append(mock.calls.RecordEvent, callInfo)
	mock.lockRecordEvent.Unlock()
	mock.RecordEventFunc(contextMoqParam, event)
}

// RecordEventCalls gets all the calls that were made to RecordEvent.
// Check the length with:
//
//	len(mockedRecorder.RecordEventCalls())
func (mock *RecorderMock) RecordEventCalls() []struct {
	ContextMoqParam context.Context
	Event           Event
} {
	var calls []struct {
		ContextMoqParam context.Context
		Event           Event
	}
	mock.lockRecordEvent.RLock()
	calls = mock.calls.RecordEvent
	mock.lockRecordEvent.RUnlock()
	return calls
}

// Ensure, that EventMock does implement Event.
// If this is not the case, regenerate this file with moq.
var _ Event = &EventMock{}

// EventMock is a mock implementation of Event.
//
//	func TestSomethingThatUsesEvent(t *testing.T) {
//
//		// make and configure a mocked Event
//		mockedEvent := &EventMock{
//			GetNameFunc: func() Name {
//				panic("mock out the GetName method")
//			},
//			isEventFromThisPackageFunc: func()  {
//				panic("mock out the isEventFromThisPackage method")
//			},
//		}
//
//		// use mockedEvent in code that requires Event
//		// and then make assertions.
//
//	}
type EventMock struct {
	// GetNameFunc mocks the GetName method.
	GetNameFunc func() Name

	// isEventFromThisPackageFunc mocks the isEventFromThisPackage method.
	isEventFromThisPackageFunc func()

	// calls tracks calls to the methods.
	calls struct {
		// GetName holds details about calls to the GetName method.
		GetName []struct {
		}
		// isEventFromThisPackage holds details about calls to the isEventFromThisPackage method.
		isEventFromThisPackage []struct {
		}
	}
	lockGetName                sync.RWMutex
	lockisEventFromThisPackage sync.RWMutex
}

// GetName calls GetNameFunc.
func (mock *EventMock) GetName() Name {
	if mock.GetNameFunc == nil {
		panic("EventMock.GetNameFunc: method is nil but Event.GetName was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetName.Lock()
	mock.calls.GetName = append(mock.calls.GetName, callInfo)
	mock.lockGetName.Unlock()
	return mock.GetNameFunc()
}

// GetNameCalls gets all the calls that were made to GetName.
// Check the length with:
//
//	len(mockedEvent.GetNameCalls())
func (mock *EventMock) GetNameCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetName.RLock()
	calls = mock.calls.GetName
	mock.lockGetName.RUnlock()
	return calls
}

// isEventFromThisPackage calls isEventFromThisPackageFunc.
func (mock *EventMock) isEventFromThisPackage() {
	if mock.isEventFromThisPackageFunc == nil {
		panic("EventMock.isEventFromThisPackageFunc: method is nil but Event.isEventFromThisPackage was just called")
	}
	callInfo := struct {
	}{}
	mock.lockisEventFromThisPackage.Lock()
	mock.calls.isEventFromThisPackage = append(mock.calls.isEventFromThisPackage, callInfo)
	mock.lockisEventFromThisPackage.Unlock()
	mock.isEventFromThisPackageFunc()
}

// isEventFromThisPackageCalls gets all the calls that were made to isEventFromThisPackage.
// Check the length with:
//
//	len(mockedEvent.isEventFromThisPackageCalls())
func (mock *EventMock) isEventFromThisPackageCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockisEventFromThisPackage.RLock()
	calls = mock.calls.isEventFromThisPackage
	mock.lockisEventFromThisPackage.RUnlock()
	return calls
}
