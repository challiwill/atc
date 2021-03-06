// This file was generated by counterfeiter
package fakes

import (
	"sync"
	"time"

	"github.com/concourse/atc/db"
	"github.com/concourse/atc/worker"
)

type FakeWorkerDB struct {
	WorkersStub        func() ([]db.SavedWorker, error)
	workersMutex       sync.RWMutex
	workersArgsForCall []struct{}
	workersReturns     struct {
		result1 []db.SavedWorker
		result2 error
	}
	GetWorkerStub        func(string) (db.SavedWorker, bool, error)
	getWorkerMutex       sync.RWMutex
	getWorkerArgsForCall []struct {
		arg1 string
	}
	getWorkerReturns struct {
		result1 db.SavedWorker
		result2 bool
		result3 error
	}
	CreateContainerStub        func(db.Container, time.Duration, time.Duration) (db.SavedContainer, error)
	createContainerMutex       sync.RWMutex
	createContainerArgsForCall []struct {
		arg1 db.Container
		arg2 time.Duration
		arg3 time.Duration
	}
	createContainerReturns struct {
		result1 db.SavedContainer
		result2 error
	}
	GetContainerStub        func(string) (db.SavedContainer, bool, error)
	getContainerMutex       sync.RWMutex
	getContainerArgsForCall []struct {
		arg1 string
	}
	getContainerReturns struct {
		result1 db.SavedContainer
		result2 bool
		result3 error
	}
	FindContainerByIdentifierStub        func(db.ContainerIdentifier) (db.SavedContainer, bool, error)
	findContainerByIdentifierMutex       sync.RWMutex
	findContainerByIdentifierArgsForCall []struct {
		arg1 db.ContainerIdentifier
	}
	findContainerByIdentifierReturns struct {
		result1 db.SavedContainer
		result2 bool
		result3 error
	}
	UpdateExpiresAtOnContainerStub        func(handle string, ttl time.Duration) error
	updateExpiresAtOnContainerMutex       sync.RWMutex
	updateExpiresAtOnContainerArgsForCall []struct {
		handle string
		ttl    time.Duration
	}
	updateExpiresAtOnContainerReturns struct {
		result1 error
	}
	ReapContainerStub        func(handle string) error
	reapContainerMutex       sync.RWMutex
	reapContainerArgsForCall []struct {
		handle string
	}
	reapContainerReturns struct {
		result1 error
	}
	InsertVolumeStub        func(db.Volume) error
	insertVolumeMutex       sync.RWMutex
	insertVolumeArgsForCall []struct {
		arg1 db.Volume
	}
	insertVolumeReturns struct {
		result1 error
	}
	GetVolumesByIdentifierStub        func(db.VolumeIdentifier) ([]db.SavedVolume, error)
	getVolumesByIdentifierMutex       sync.RWMutex
	getVolumesByIdentifierArgsForCall []struct {
		arg1 db.VolumeIdentifier
	}
	getVolumesByIdentifierReturns struct {
		result1 []db.SavedVolume
		result2 error
	}
	GetVolumeTTLStub        func(volumeHandle string) (time.Duration, bool, error)
	getVolumeTTLMutex       sync.RWMutex
	getVolumeTTLArgsForCall []struct {
		volumeHandle string
	}
	getVolumeTTLReturns struct {
		result1 time.Duration
		result2 bool
		result3 error
	}
	ReapVolumeStub        func(handle string) error
	reapVolumeMutex       sync.RWMutex
	reapVolumeArgsForCall []struct {
		handle string
	}
	reapVolumeReturns struct {
		result1 error
	}
	SetVolumeTTLStub        func(string, time.Duration) error
	setVolumeTTLMutex       sync.RWMutex
	setVolumeTTLArgsForCall []struct {
		arg1 string
		arg2 time.Duration
	}
	setVolumeTTLReturns struct {
		result1 error
	}
	SetVolumeSizeStub        func(string, uint) error
	setVolumeSizeMutex       sync.RWMutex
	setVolumeSizeArgsForCall []struct {
		arg1 string
		arg2 uint
	}
	setVolumeSizeReturns struct {
		result1 error
	}
}

func (fake *FakeWorkerDB) Workers() ([]db.SavedWorker, error) {
	fake.workersMutex.Lock()
	fake.workersArgsForCall = append(fake.workersArgsForCall, struct{}{})
	fake.workersMutex.Unlock()
	if fake.WorkersStub != nil {
		return fake.WorkersStub()
	} else {
		return fake.workersReturns.result1, fake.workersReturns.result2
	}
}

func (fake *FakeWorkerDB) WorkersCallCount() int {
	fake.workersMutex.RLock()
	defer fake.workersMutex.RUnlock()
	return len(fake.workersArgsForCall)
}

func (fake *FakeWorkerDB) WorkersReturns(result1 []db.SavedWorker, result2 error) {
	fake.WorkersStub = nil
	fake.workersReturns = struct {
		result1 []db.SavedWorker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerDB) GetWorker(arg1 string) (db.SavedWorker, bool, error) {
	fake.getWorkerMutex.Lock()
	fake.getWorkerArgsForCall = append(fake.getWorkerArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.getWorkerMutex.Unlock()
	if fake.GetWorkerStub != nil {
		return fake.GetWorkerStub(arg1)
	} else {
		return fake.getWorkerReturns.result1, fake.getWorkerReturns.result2, fake.getWorkerReturns.result3
	}
}

func (fake *FakeWorkerDB) GetWorkerCallCount() int {
	fake.getWorkerMutex.RLock()
	defer fake.getWorkerMutex.RUnlock()
	return len(fake.getWorkerArgsForCall)
}

func (fake *FakeWorkerDB) GetWorkerArgsForCall(i int) string {
	fake.getWorkerMutex.RLock()
	defer fake.getWorkerMutex.RUnlock()
	return fake.getWorkerArgsForCall[i].arg1
}

func (fake *FakeWorkerDB) GetWorkerReturns(result1 db.SavedWorker, result2 bool, result3 error) {
	fake.GetWorkerStub = nil
	fake.getWorkerReturns = struct {
		result1 db.SavedWorker
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorkerDB) CreateContainer(arg1 db.Container, arg2 time.Duration, arg3 time.Duration) (db.SavedContainer, error) {
	fake.createContainerMutex.Lock()
	fake.createContainerArgsForCall = append(fake.createContainerArgsForCall, struct {
		arg1 db.Container
		arg2 time.Duration
		arg3 time.Duration
	}{arg1, arg2, arg3})
	fake.createContainerMutex.Unlock()
	if fake.CreateContainerStub != nil {
		return fake.CreateContainerStub(arg1, arg2, arg3)
	} else {
		return fake.createContainerReturns.result1, fake.createContainerReturns.result2
	}
}

func (fake *FakeWorkerDB) CreateContainerCallCount() int {
	fake.createContainerMutex.RLock()
	defer fake.createContainerMutex.RUnlock()
	return len(fake.createContainerArgsForCall)
}

func (fake *FakeWorkerDB) CreateContainerArgsForCall(i int) (db.Container, time.Duration, time.Duration) {
	fake.createContainerMutex.RLock()
	defer fake.createContainerMutex.RUnlock()
	return fake.createContainerArgsForCall[i].arg1, fake.createContainerArgsForCall[i].arg2, fake.createContainerArgsForCall[i].arg3
}

func (fake *FakeWorkerDB) CreateContainerReturns(result1 db.SavedContainer, result2 error) {
	fake.CreateContainerStub = nil
	fake.createContainerReturns = struct {
		result1 db.SavedContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerDB) GetContainer(arg1 string) (db.SavedContainer, bool, error) {
	fake.getContainerMutex.Lock()
	fake.getContainerArgsForCall = append(fake.getContainerArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.getContainerMutex.Unlock()
	if fake.GetContainerStub != nil {
		return fake.GetContainerStub(arg1)
	} else {
		return fake.getContainerReturns.result1, fake.getContainerReturns.result2, fake.getContainerReturns.result3
	}
}

func (fake *FakeWorkerDB) GetContainerCallCount() int {
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	return len(fake.getContainerArgsForCall)
}

func (fake *FakeWorkerDB) GetContainerArgsForCall(i int) string {
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	return fake.getContainerArgsForCall[i].arg1
}

func (fake *FakeWorkerDB) GetContainerReturns(result1 db.SavedContainer, result2 bool, result3 error) {
	fake.GetContainerStub = nil
	fake.getContainerReturns = struct {
		result1 db.SavedContainer
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorkerDB) FindContainerByIdentifier(arg1 db.ContainerIdentifier) (db.SavedContainer, bool, error) {
	fake.findContainerByIdentifierMutex.Lock()
	fake.findContainerByIdentifierArgsForCall = append(fake.findContainerByIdentifierArgsForCall, struct {
		arg1 db.ContainerIdentifier
	}{arg1})
	fake.findContainerByIdentifierMutex.Unlock()
	if fake.FindContainerByIdentifierStub != nil {
		return fake.FindContainerByIdentifierStub(arg1)
	} else {
		return fake.findContainerByIdentifierReturns.result1, fake.findContainerByIdentifierReturns.result2, fake.findContainerByIdentifierReturns.result3
	}
}

func (fake *FakeWorkerDB) FindContainerByIdentifierCallCount() int {
	fake.findContainerByIdentifierMutex.RLock()
	defer fake.findContainerByIdentifierMutex.RUnlock()
	return len(fake.findContainerByIdentifierArgsForCall)
}

func (fake *FakeWorkerDB) FindContainerByIdentifierArgsForCall(i int) db.ContainerIdentifier {
	fake.findContainerByIdentifierMutex.RLock()
	defer fake.findContainerByIdentifierMutex.RUnlock()
	return fake.findContainerByIdentifierArgsForCall[i].arg1
}

func (fake *FakeWorkerDB) FindContainerByIdentifierReturns(result1 db.SavedContainer, result2 bool, result3 error) {
	fake.FindContainerByIdentifierStub = nil
	fake.findContainerByIdentifierReturns = struct {
		result1 db.SavedContainer
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorkerDB) UpdateExpiresAtOnContainer(handle string, ttl time.Duration) error {
	fake.updateExpiresAtOnContainerMutex.Lock()
	fake.updateExpiresAtOnContainerArgsForCall = append(fake.updateExpiresAtOnContainerArgsForCall, struct {
		handle string
		ttl    time.Duration
	}{handle, ttl})
	fake.updateExpiresAtOnContainerMutex.Unlock()
	if fake.UpdateExpiresAtOnContainerStub != nil {
		return fake.UpdateExpiresAtOnContainerStub(handle, ttl)
	} else {
		return fake.updateExpiresAtOnContainerReturns.result1
	}
}

func (fake *FakeWorkerDB) UpdateExpiresAtOnContainerCallCount() int {
	fake.updateExpiresAtOnContainerMutex.RLock()
	defer fake.updateExpiresAtOnContainerMutex.RUnlock()
	return len(fake.updateExpiresAtOnContainerArgsForCall)
}

func (fake *FakeWorkerDB) UpdateExpiresAtOnContainerArgsForCall(i int) (string, time.Duration) {
	fake.updateExpiresAtOnContainerMutex.RLock()
	defer fake.updateExpiresAtOnContainerMutex.RUnlock()
	return fake.updateExpiresAtOnContainerArgsForCall[i].handle, fake.updateExpiresAtOnContainerArgsForCall[i].ttl
}

func (fake *FakeWorkerDB) UpdateExpiresAtOnContainerReturns(result1 error) {
	fake.UpdateExpiresAtOnContainerStub = nil
	fake.updateExpiresAtOnContainerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorkerDB) ReapContainer(handle string) error {
	fake.reapContainerMutex.Lock()
	fake.reapContainerArgsForCall = append(fake.reapContainerArgsForCall, struct {
		handle string
	}{handle})
	fake.reapContainerMutex.Unlock()
	if fake.ReapContainerStub != nil {
		return fake.ReapContainerStub(handle)
	} else {
		return fake.reapContainerReturns.result1
	}
}

func (fake *FakeWorkerDB) ReapContainerCallCount() int {
	fake.reapContainerMutex.RLock()
	defer fake.reapContainerMutex.RUnlock()
	return len(fake.reapContainerArgsForCall)
}

func (fake *FakeWorkerDB) ReapContainerArgsForCall(i int) string {
	fake.reapContainerMutex.RLock()
	defer fake.reapContainerMutex.RUnlock()
	return fake.reapContainerArgsForCall[i].handle
}

func (fake *FakeWorkerDB) ReapContainerReturns(result1 error) {
	fake.ReapContainerStub = nil
	fake.reapContainerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorkerDB) InsertVolume(arg1 db.Volume) error {
	fake.insertVolumeMutex.Lock()
	fake.insertVolumeArgsForCall = append(fake.insertVolumeArgsForCall, struct {
		arg1 db.Volume
	}{arg1})
	fake.insertVolumeMutex.Unlock()
	if fake.InsertVolumeStub != nil {
		return fake.InsertVolumeStub(arg1)
	} else {
		return fake.insertVolumeReturns.result1
	}
}

func (fake *FakeWorkerDB) InsertVolumeCallCount() int {
	fake.insertVolumeMutex.RLock()
	defer fake.insertVolumeMutex.RUnlock()
	return len(fake.insertVolumeArgsForCall)
}

func (fake *FakeWorkerDB) InsertVolumeArgsForCall(i int) db.Volume {
	fake.insertVolumeMutex.RLock()
	defer fake.insertVolumeMutex.RUnlock()
	return fake.insertVolumeArgsForCall[i].arg1
}

func (fake *FakeWorkerDB) InsertVolumeReturns(result1 error) {
	fake.InsertVolumeStub = nil
	fake.insertVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorkerDB) GetVolumesByIdentifier(arg1 db.VolumeIdentifier) ([]db.SavedVolume, error) {
	fake.getVolumesByIdentifierMutex.Lock()
	fake.getVolumesByIdentifierArgsForCall = append(fake.getVolumesByIdentifierArgsForCall, struct {
		arg1 db.VolumeIdentifier
	}{arg1})
	fake.getVolumesByIdentifierMutex.Unlock()
	if fake.GetVolumesByIdentifierStub != nil {
		return fake.GetVolumesByIdentifierStub(arg1)
	} else {
		return fake.getVolumesByIdentifierReturns.result1, fake.getVolumesByIdentifierReturns.result2
	}
}

func (fake *FakeWorkerDB) GetVolumesByIdentifierCallCount() int {
	fake.getVolumesByIdentifierMutex.RLock()
	defer fake.getVolumesByIdentifierMutex.RUnlock()
	return len(fake.getVolumesByIdentifierArgsForCall)
}

func (fake *FakeWorkerDB) GetVolumesByIdentifierArgsForCall(i int) db.VolumeIdentifier {
	fake.getVolumesByIdentifierMutex.RLock()
	defer fake.getVolumesByIdentifierMutex.RUnlock()
	return fake.getVolumesByIdentifierArgsForCall[i].arg1
}

func (fake *FakeWorkerDB) GetVolumesByIdentifierReturns(result1 []db.SavedVolume, result2 error) {
	fake.GetVolumesByIdentifierStub = nil
	fake.getVolumesByIdentifierReturns = struct {
		result1 []db.SavedVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerDB) GetVolumeTTL(volumeHandle string) (time.Duration, bool, error) {
	fake.getVolumeTTLMutex.Lock()
	fake.getVolumeTTLArgsForCall = append(fake.getVolumeTTLArgsForCall, struct {
		volumeHandle string
	}{volumeHandle})
	fake.getVolumeTTLMutex.Unlock()
	if fake.GetVolumeTTLStub != nil {
		return fake.GetVolumeTTLStub(volumeHandle)
	} else {
		return fake.getVolumeTTLReturns.result1, fake.getVolumeTTLReturns.result2, fake.getVolumeTTLReturns.result3
	}
}

func (fake *FakeWorkerDB) GetVolumeTTLCallCount() int {
	fake.getVolumeTTLMutex.RLock()
	defer fake.getVolumeTTLMutex.RUnlock()
	return len(fake.getVolumeTTLArgsForCall)
}

func (fake *FakeWorkerDB) GetVolumeTTLArgsForCall(i int) string {
	fake.getVolumeTTLMutex.RLock()
	defer fake.getVolumeTTLMutex.RUnlock()
	return fake.getVolumeTTLArgsForCall[i].volumeHandle
}

func (fake *FakeWorkerDB) GetVolumeTTLReturns(result1 time.Duration, result2 bool, result3 error) {
	fake.GetVolumeTTLStub = nil
	fake.getVolumeTTLReturns = struct {
		result1 time.Duration
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorkerDB) ReapVolume(handle string) error {
	fake.reapVolumeMutex.Lock()
	fake.reapVolumeArgsForCall = append(fake.reapVolumeArgsForCall, struct {
		handle string
	}{handle})
	fake.reapVolumeMutex.Unlock()
	if fake.ReapVolumeStub != nil {
		return fake.ReapVolumeStub(handle)
	} else {
		return fake.reapVolumeReturns.result1
	}
}

func (fake *FakeWorkerDB) ReapVolumeCallCount() int {
	fake.reapVolumeMutex.RLock()
	defer fake.reapVolumeMutex.RUnlock()
	return len(fake.reapVolumeArgsForCall)
}

func (fake *FakeWorkerDB) ReapVolumeArgsForCall(i int) string {
	fake.reapVolumeMutex.RLock()
	defer fake.reapVolumeMutex.RUnlock()
	return fake.reapVolumeArgsForCall[i].handle
}

func (fake *FakeWorkerDB) ReapVolumeReturns(result1 error) {
	fake.ReapVolumeStub = nil
	fake.reapVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorkerDB) SetVolumeTTL(arg1 string, arg2 time.Duration) error {
	fake.setVolumeTTLMutex.Lock()
	fake.setVolumeTTLArgsForCall = append(fake.setVolumeTTLArgsForCall, struct {
		arg1 string
		arg2 time.Duration
	}{arg1, arg2})
	fake.setVolumeTTLMutex.Unlock()
	if fake.SetVolumeTTLStub != nil {
		return fake.SetVolumeTTLStub(arg1, arg2)
	} else {
		return fake.setVolumeTTLReturns.result1
	}
}

func (fake *FakeWorkerDB) SetVolumeTTLCallCount() int {
	fake.setVolumeTTLMutex.RLock()
	defer fake.setVolumeTTLMutex.RUnlock()
	return len(fake.setVolumeTTLArgsForCall)
}

func (fake *FakeWorkerDB) SetVolumeTTLArgsForCall(i int) (string, time.Duration) {
	fake.setVolumeTTLMutex.RLock()
	defer fake.setVolumeTTLMutex.RUnlock()
	return fake.setVolumeTTLArgsForCall[i].arg1, fake.setVolumeTTLArgsForCall[i].arg2
}

func (fake *FakeWorkerDB) SetVolumeTTLReturns(result1 error) {
	fake.SetVolumeTTLStub = nil
	fake.setVolumeTTLReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorkerDB) SetVolumeSize(arg1 string, arg2 uint) error {
	fake.setVolumeSizeMutex.Lock()
	fake.setVolumeSizeArgsForCall = append(fake.setVolumeSizeArgsForCall, struct {
		arg1 string
		arg2 uint
	}{arg1, arg2})
	fake.setVolumeSizeMutex.Unlock()
	if fake.SetVolumeSizeStub != nil {
		return fake.SetVolumeSizeStub(arg1, arg2)
	} else {
		return fake.setVolumeSizeReturns.result1
	}
}

func (fake *FakeWorkerDB) SetVolumeSizeCallCount() int {
	fake.setVolumeSizeMutex.RLock()
	defer fake.setVolumeSizeMutex.RUnlock()
	return len(fake.setVolumeSizeArgsForCall)
}

func (fake *FakeWorkerDB) SetVolumeSizeArgsForCall(i int) (string, uint) {
	fake.setVolumeSizeMutex.RLock()
	defer fake.setVolumeSizeMutex.RUnlock()
	return fake.setVolumeSizeArgsForCall[i].arg1, fake.setVolumeSizeArgsForCall[i].arg2
}

func (fake *FakeWorkerDB) SetVolumeSizeReturns(result1 error) {
	fake.SetVolumeSizeStub = nil
	fake.setVolumeSizeReturns = struct {
		result1 error
	}{result1}
}

var _ worker.WorkerDB = new(FakeWorkerDB)
