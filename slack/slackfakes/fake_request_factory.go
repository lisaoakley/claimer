// This file was generated by counterfeiter
package slackfakes

import (
	"sync"

	"github.com/mdelillo/claimer/slack"
)

type FakeRequestFactory struct {
	NewUsernameRequestStub        func(userId string) slack.UsernameRequest
	newUsernameRequestMutex       sync.RWMutex
	newUsernameRequestArgsForCall []struct {
		userId string
	}
	newUsernameRequestReturns struct {
		result1 slack.UsernameRequest
	}
	newUsernameRequestReturnsOnCall map[int]struct {
		result1 slack.UsernameRequest
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRequestFactory) NewUsernameRequest(userId string) slack.UsernameRequest {
	fake.newUsernameRequestMutex.Lock()
	ret, specificReturn := fake.newUsernameRequestReturnsOnCall[len(fake.newUsernameRequestArgsForCall)]
	fake.newUsernameRequestArgsForCall = append(fake.newUsernameRequestArgsForCall, struct {
		userId string
	}{userId})
	fake.recordInvocation("NewUsernameRequest", []interface{}{userId})
	fake.newUsernameRequestMutex.Unlock()
	if fake.NewUsernameRequestStub != nil {
		return fake.NewUsernameRequestStub(userId)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.newUsernameRequestReturns.result1
}

func (fake *FakeRequestFactory) NewUsernameRequestCallCount() int {
	fake.newUsernameRequestMutex.RLock()
	defer fake.newUsernameRequestMutex.RUnlock()
	return len(fake.newUsernameRequestArgsForCall)
}

func (fake *FakeRequestFactory) NewUsernameRequestArgsForCall(i int) string {
	fake.newUsernameRequestMutex.RLock()
	defer fake.newUsernameRequestMutex.RUnlock()
	return fake.newUsernameRequestArgsForCall[i].userId
}

func (fake *FakeRequestFactory) NewUsernameRequestReturns(result1 slack.UsernameRequest) {
	fake.NewUsernameRequestStub = nil
	fake.newUsernameRequestReturns = struct {
		result1 slack.UsernameRequest
	}{result1}
}

func (fake *FakeRequestFactory) NewUsernameRequestReturnsOnCall(i int, result1 slack.UsernameRequest) {
	fake.NewUsernameRequestStub = nil
	if fake.newUsernameRequestReturnsOnCall == nil {
		fake.newUsernameRequestReturnsOnCall = make(map[int]struct {
			result1 slack.UsernameRequest
		})
	}
	fake.newUsernameRequestReturnsOnCall[i] = struct {
		result1 slack.UsernameRequest
	}{result1}
}

func (fake *FakeRequestFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newUsernameRequestMutex.RLock()
	defer fake.newUsernameRequestMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeRequestFactory) recordInvocation(key string, args []interface{}) {
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

var _ slack.RequestFactory = new(FakeRequestFactory)