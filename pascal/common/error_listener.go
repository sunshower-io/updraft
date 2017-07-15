package common

import "gitlab.com/sunshower.io/updraft/common/observer"

type ErrorListener struct {
	observer.EventListener
}

func (e *ErrorListener) Id() string {
	return "lang::error-listener"
}

func (e *ErrorListener) OnMessage(m observer.Message) {

}

func (e *ErrorListener) ListensFor(m observer.Message) bool {
	return false

}
