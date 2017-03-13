// This file was generated by counterfeiter
package claimerfakes

import (
	"sync"
)

type FakeLocker struct {
	ClaimLockStub        func(pool string) error
	claimLockMutex       sync.RWMutex
	claimLockArgsForCall []struct {
		pool string
	}
	claimLockReturns struct {
		result1 error
	}
	claimLockReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLocker) ClaimLock(pool string) error {
	fake.claimLockMutex.Lock()
	ret, specificReturn := fake.claimLockReturnsOnCall[len(fake.claimLockArgsForCall)]
	fake.claimLockArgsForCall = append(fake.claimLockArgsForCall, struct {
		pool string
	}{pool})
	fake.recordInvocation("ClaimLock", []interface{}{pool})
	fake.claimLockMutex.Unlock()
	if fake.ClaimLockStub != nil {
		return fake.ClaimLockStub(pool)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.claimLockReturns.result1
}

func (fake *FakeLocker) ClaimLockCallCount() int {
	fake.claimLockMutex.RLock()
	defer fake.claimLockMutex.RUnlock()
	return len(fake.claimLockArgsForCall)
}

func (fake *FakeLocker) ClaimLockArgsForCall(i int) string {
	fake.claimLockMutex.RLock()
	defer fake.claimLockMutex.RUnlock()
	return fake.claimLockArgsForCall[i].pool
}

func (fake *FakeLocker) ClaimLockReturns(result1 error) {
	fake.ClaimLockStub = nil
	fake.claimLockReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLocker) ClaimLockReturnsOnCall(i int, result1 error) {
	fake.ClaimLockStub = nil
	if fake.claimLockReturnsOnCall == nil {
		fake.claimLockReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.claimLockReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeLocker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.claimLockMutex.RLock()
	defer fake.claimLockMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeLocker) recordInvocation(key string, args []interface{}) {
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
