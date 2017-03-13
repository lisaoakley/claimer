// This file was generated by counterfeiter
package lockerfakes

import (
	"sync"
)

type FakeFs struct {
	LsStub        func(dir string) ([]string, error)
	lsMutex       sync.RWMutex
	lsArgsForCall []struct {
		dir string
	}
	lsReturns struct {
		result1 []string
		result2 error
	}
	lsReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	MvStub        func(src, dst string) error
	mvMutex       sync.RWMutex
	mvArgsForCall []struct {
		src string
		dst string
	}
	mvReturns struct {
		result1 error
	}
	mvReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFs) Ls(dir string) ([]string, error) {
	fake.lsMutex.Lock()
	ret, specificReturn := fake.lsReturnsOnCall[len(fake.lsArgsForCall)]
	fake.lsArgsForCall = append(fake.lsArgsForCall, struct {
		dir string
	}{dir})
	fake.recordInvocation("Ls", []interface{}{dir})
	fake.lsMutex.Unlock()
	if fake.LsStub != nil {
		return fake.LsStub(dir)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.lsReturns.result1, fake.lsReturns.result2
}

func (fake *FakeFs) LsCallCount() int {
	fake.lsMutex.RLock()
	defer fake.lsMutex.RUnlock()
	return len(fake.lsArgsForCall)
}

func (fake *FakeFs) LsArgsForCall(i int) string {
	fake.lsMutex.RLock()
	defer fake.lsMutex.RUnlock()
	return fake.lsArgsForCall[i].dir
}

func (fake *FakeFs) LsReturns(result1 []string, result2 error) {
	fake.LsStub = nil
	fake.lsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeFs) LsReturnsOnCall(i int, result1 []string, result2 error) {
	fake.LsStub = nil
	if fake.lsReturnsOnCall == nil {
		fake.lsReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.lsReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeFs) Mv(src string, dst string) error {
	fake.mvMutex.Lock()
	ret, specificReturn := fake.mvReturnsOnCall[len(fake.mvArgsForCall)]
	fake.mvArgsForCall = append(fake.mvArgsForCall, struct {
		src string
		dst string
	}{src, dst})
	fake.recordInvocation("Mv", []interface{}{src, dst})
	fake.mvMutex.Unlock()
	if fake.MvStub != nil {
		return fake.MvStub(src, dst)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.mvReturns.result1
}

func (fake *FakeFs) MvCallCount() int {
	fake.mvMutex.RLock()
	defer fake.mvMutex.RUnlock()
	return len(fake.mvArgsForCall)
}

func (fake *FakeFs) MvArgsForCall(i int) (string, string) {
	fake.mvMutex.RLock()
	defer fake.mvMutex.RUnlock()
	return fake.mvArgsForCall[i].src, fake.mvArgsForCall[i].dst
}

func (fake *FakeFs) MvReturns(result1 error) {
	fake.MvStub = nil
	fake.mvReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFs) MvReturnsOnCall(i int, result1 error) {
	fake.MvStub = nil
	if fake.mvReturnsOnCall == nil {
		fake.mvReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.mvReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFs) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.lsMutex.RLock()
	defer fake.lsMutex.RUnlock()
	fake.mvMutex.RLock()
	defer fake.mvMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeFs) recordInvocation(key string, args []interface{}) {
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