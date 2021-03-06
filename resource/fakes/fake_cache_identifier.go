// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc/resource"
	"github.com/concourse/atc/worker"
	"github.com/pivotal-golang/lager"
)

type FakeCacheIdentifier struct {
	FindOnStub        func(lager.Logger, worker.Client) (worker.Volume, bool, error)
	findOnMutex       sync.RWMutex
	findOnArgsForCall []struct {
		arg1 lager.Logger
		arg2 worker.Client
	}
	findOnReturns struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	CreateOnStub        func(lager.Logger, worker.Client) (worker.Volume, error)
	createOnMutex       sync.RWMutex
	createOnArgsForCall []struct {
		arg1 lager.Logger
		arg2 worker.Client
	}
	createOnReturns struct {
		result1 worker.Volume
		result2 error
	}
	VolumeIdentifierStub        func() worker.VolumeIdentifier
	volumeIdentifierMutex       sync.RWMutex
	volumeIdentifierArgsForCall []struct{}
	volumeIdentifierReturns     struct {
		result1 worker.VolumeIdentifier
	}
}

func (fake *FakeCacheIdentifier) FindOn(arg1 lager.Logger, arg2 worker.Client) (worker.Volume, bool, error) {
	fake.findOnMutex.Lock()
	fake.findOnArgsForCall = append(fake.findOnArgsForCall, struct {
		arg1 lager.Logger
		arg2 worker.Client
	}{arg1, arg2})
	fake.findOnMutex.Unlock()
	if fake.FindOnStub != nil {
		return fake.FindOnStub(arg1, arg2)
	} else {
		return fake.findOnReturns.result1, fake.findOnReturns.result2, fake.findOnReturns.result3
	}
}

func (fake *FakeCacheIdentifier) FindOnCallCount() int {
	fake.findOnMutex.RLock()
	defer fake.findOnMutex.RUnlock()
	return len(fake.findOnArgsForCall)
}

func (fake *FakeCacheIdentifier) FindOnArgsForCall(i int) (lager.Logger, worker.Client) {
	fake.findOnMutex.RLock()
	defer fake.findOnMutex.RUnlock()
	return fake.findOnArgsForCall[i].arg1, fake.findOnArgsForCall[i].arg2
}

func (fake *FakeCacheIdentifier) FindOnReturns(result1 worker.Volume, result2 bool, result3 error) {
	fake.FindOnStub = nil
	fake.findOnReturns = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCacheIdentifier) CreateOn(arg1 lager.Logger, arg2 worker.Client) (worker.Volume, error) {
	fake.createOnMutex.Lock()
	fake.createOnArgsForCall = append(fake.createOnArgsForCall, struct {
		arg1 lager.Logger
		arg2 worker.Client
	}{arg1, arg2})
	fake.createOnMutex.Unlock()
	if fake.CreateOnStub != nil {
		return fake.CreateOnStub(arg1, arg2)
	} else {
		return fake.createOnReturns.result1, fake.createOnReturns.result2
	}
}

func (fake *FakeCacheIdentifier) CreateOnCallCount() int {
	fake.createOnMutex.RLock()
	defer fake.createOnMutex.RUnlock()
	return len(fake.createOnArgsForCall)
}

func (fake *FakeCacheIdentifier) CreateOnArgsForCall(i int) (lager.Logger, worker.Client) {
	fake.createOnMutex.RLock()
	defer fake.createOnMutex.RUnlock()
	return fake.createOnArgsForCall[i].arg1, fake.createOnArgsForCall[i].arg2
}

func (fake *FakeCacheIdentifier) CreateOnReturns(result1 worker.Volume, result2 error) {
	fake.CreateOnStub = nil
	fake.createOnReturns = struct {
		result1 worker.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeCacheIdentifier) VolumeIdentifier() worker.VolumeIdentifier {
	fake.volumeIdentifierMutex.Lock()
	fake.volumeIdentifierArgsForCall = append(fake.volumeIdentifierArgsForCall, struct{}{})
	fake.volumeIdentifierMutex.Unlock()
	if fake.VolumeIdentifierStub != nil {
		return fake.VolumeIdentifierStub()
	} else {
		return fake.volumeIdentifierReturns.result1
	}
}

func (fake *FakeCacheIdentifier) VolumeIdentifierCallCount() int {
	fake.volumeIdentifierMutex.RLock()
	defer fake.volumeIdentifierMutex.RUnlock()
	return len(fake.volumeIdentifierArgsForCall)
}

func (fake *FakeCacheIdentifier) VolumeIdentifierReturns(result1 worker.VolumeIdentifier) {
	fake.VolumeIdentifierStub = nil
	fake.volumeIdentifierReturns = struct {
		result1 worker.VolumeIdentifier
	}{result1}
}

var _ resource.CacheIdentifier = new(FakeCacheIdentifier)
