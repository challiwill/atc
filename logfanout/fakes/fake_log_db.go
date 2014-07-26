// This file was generated by counterfeiter
package fakes

import (
	"sync"
	"github.com/concourse/atc/logfanout"
)

type FakeLogDB struct {
	BuildLogStub        func(job string, build int) ([]byte, error)
	buildLogMutex       sync.RWMutex
	buildLogArgsForCall []struct {
		job   string
		build int
	}
	buildLogReturns struct {
		result1 []byte
		result2 error
	}
	AppendBuildLogStub        func(job string, build int, log []byte) error
	appendBuildLogMutex       sync.RWMutex
	appendBuildLogArgsForCall []struct {
		job   string
		build int
		log   []byte
	}
	appendBuildLogReturns struct {
		result1 error
	}
}

func (fake *FakeLogDB) BuildLog(job string, build int) ([]byte, error) {
	fake.buildLogMutex.Lock()
	defer fake.buildLogMutex.Unlock()
	fake.buildLogArgsForCall = append(fake.buildLogArgsForCall, struct {
		job   string
		build int
	}{job, build})
	if fake.BuildLogStub != nil {
		return fake.BuildLogStub(job, build)
	} else {
		return fake.buildLogReturns.result1, fake.buildLogReturns.result2
	}
}

func (fake *FakeLogDB) BuildLogCallCount() int {
	fake.buildLogMutex.RLock()
	defer fake.buildLogMutex.RUnlock()
	return len(fake.buildLogArgsForCall)
}

func (fake *FakeLogDB) BuildLogArgsForCall(i int) (string, int) {
	fake.buildLogMutex.RLock()
	defer fake.buildLogMutex.RUnlock()
	return fake.buildLogArgsForCall[i].job, fake.buildLogArgsForCall[i].build
}

func (fake *FakeLogDB) BuildLogReturns(result1 []byte, result2 error) {
	fake.BuildLogStub = nil
	fake.buildLogReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeLogDB) AppendBuildLog(job string, build int, log []byte) error {
	fake.appendBuildLogMutex.Lock()
	defer fake.appendBuildLogMutex.Unlock()
	fake.appendBuildLogArgsForCall = append(fake.appendBuildLogArgsForCall, struct {
		job   string
		build int
		log   []byte
	}{job, build, log})
	if fake.AppendBuildLogStub != nil {
		return fake.AppendBuildLogStub(job, build, log)
	} else {
		return fake.appendBuildLogReturns.result1
	}
}

func (fake *FakeLogDB) AppendBuildLogCallCount() int {
	fake.appendBuildLogMutex.RLock()
	defer fake.appendBuildLogMutex.RUnlock()
	return len(fake.appendBuildLogArgsForCall)
}

func (fake *FakeLogDB) AppendBuildLogArgsForCall(i int) (string, int, []byte) {
	fake.appendBuildLogMutex.RLock()
	defer fake.appendBuildLogMutex.RUnlock()
	return fake.appendBuildLogArgsForCall[i].job, fake.appendBuildLogArgsForCall[i].build, fake.appendBuildLogArgsForCall[i].log
}

func (fake *FakeLogDB) AppendBuildLogReturns(result1 error) {
	fake.AppendBuildLogStub = nil
	fake.appendBuildLogReturns = struct {
		result1 error
	}{result1}
}

var _ logfanout.LogDB = new(FakeLogDB)