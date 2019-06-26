// Code generated by counterfeiter. DO NOT EDIT.
package k8sfakes

import (
	io "io"
	sync "sync"

	lager "code.cloudfoundry.org/lager"
	k8s "github.com/concourse/concourse/atc/k8s"
)

type FakeBlobstoreIO struct {
	InputBlobReaderStub        func(lager.Logger, string, string) (io.Reader, error)
	inputBlobReaderMutex       sync.RWMutex
	inputBlobReaderArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
	}
	inputBlobReaderReturns struct {
		result1 io.Reader
		result2 error
	}
	inputBlobReaderReturnsOnCall map[int]struct {
		result1 io.Reader
		result2 error
	}
	OutputBlobWriterStub        func(lager.Logger, string, string) (io.Writer, error)
	outputBlobWriterMutex       sync.RWMutex
	outputBlobWriterArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
	}
	outputBlobWriterReturns struct {
		result1 io.Writer
		result2 error
	}
	outputBlobWriterReturnsOnCall map[int]struct {
		result1 io.Writer
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBlobstoreIO) InputBlobReader(arg1 lager.Logger, arg2 string, arg3 string) (io.Reader, error) {
	fake.inputBlobReaderMutex.Lock()
	ret, specificReturn := fake.inputBlobReaderReturnsOnCall[len(fake.inputBlobReaderArgsForCall)]
	fake.inputBlobReaderArgsForCall = append(fake.inputBlobReaderArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("InputBlobReader", []interface{}{arg1, arg2, arg3})
	fake.inputBlobReaderMutex.Unlock()
	if fake.InputBlobReaderStub != nil {
		return fake.InputBlobReaderStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.inputBlobReaderReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBlobstoreIO) InputBlobReaderCallCount() int {
	fake.inputBlobReaderMutex.RLock()
	defer fake.inputBlobReaderMutex.RUnlock()
	return len(fake.inputBlobReaderArgsForCall)
}

func (fake *FakeBlobstoreIO) InputBlobReaderArgsForCall(i int) (lager.Logger, string, string) {
	fake.inputBlobReaderMutex.RLock()
	defer fake.inputBlobReaderMutex.RUnlock()
	argsForCall := fake.inputBlobReaderArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeBlobstoreIO) InputBlobReaderReturns(result1 io.Reader, result2 error) {
	fake.InputBlobReaderStub = nil
	fake.inputBlobReaderReturns = struct {
		result1 io.Reader
		result2 error
	}{result1, result2}
}

func (fake *FakeBlobstoreIO) InputBlobReaderReturnsOnCall(i int, result1 io.Reader, result2 error) {
	fake.InputBlobReaderStub = nil
	if fake.inputBlobReaderReturnsOnCall == nil {
		fake.inputBlobReaderReturnsOnCall = make(map[int]struct {
			result1 io.Reader
			result2 error
		})
	}
	fake.inputBlobReaderReturnsOnCall[i] = struct {
		result1 io.Reader
		result2 error
	}{result1, result2}
}

func (fake *FakeBlobstoreIO) OutputBlobWriter(arg1 lager.Logger, arg2 string, arg3 string) (io.Writer, error) {
	fake.outputBlobWriterMutex.Lock()
	ret, specificReturn := fake.outputBlobWriterReturnsOnCall[len(fake.outputBlobWriterArgsForCall)]
	fake.outputBlobWriterArgsForCall = append(fake.outputBlobWriterArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("OutputBlobWriter", []interface{}{arg1, arg2, arg3})
	fake.outputBlobWriterMutex.Unlock()
	if fake.OutputBlobWriterStub != nil {
		return fake.OutputBlobWriterStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.outputBlobWriterReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBlobstoreIO) OutputBlobWriterCallCount() int {
	fake.outputBlobWriterMutex.RLock()
	defer fake.outputBlobWriterMutex.RUnlock()
	return len(fake.outputBlobWriterArgsForCall)
}

func (fake *FakeBlobstoreIO) OutputBlobWriterArgsForCall(i int) (lager.Logger, string, string) {
	fake.outputBlobWriterMutex.RLock()
	defer fake.outputBlobWriterMutex.RUnlock()
	argsForCall := fake.outputBlobWriterArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeBlobstoreIO) OutputBlobWriterReturns(result1 io.Writer, result2 error) {
	fake.OutputBlobWriterStub = nil
	fake.outputBlobWriterReturns = struct {
		result1 io.Writer
		result2 error
	}{result1, result2}
}

func (fake *FakeBlobstoreIO) OutputBlobWriterReturnsOnCall(i int, result1 io.Writer, result2 error) {
	fake.OutputBlobWriterStub = nil
	if fake.outputBlobWriterReturnsOnCall == nil {
		fake.outputBlobWriterReturnsOnCall = make(map[int]struct {
			result1 io.Writer
			result2 error
		})
	}
	fake.outputBlobWriterReturnsOnCall[i] = struct {
		result1 io.Writer
		result2 error
	}{result1, result2}
}

func (fake *FakeBlobstoreIO) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.inputBlobReaderMutex.RLock()
	defer fake.inputBlobReaderMutex.RUnlock()
	fake.outputBlobWriterMutex.RLock()
	defer fake.outputBlobWriterMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBlobstoreIO) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ k8s.BlobstoreIO = new(FakeBlobstoreIO)