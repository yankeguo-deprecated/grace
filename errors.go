package grace

import (
	"strings"
	"sync"
)

type Errors []error

func (errs Errors) Error() string {
	sb := &strings.Builder{}
	for i, err := range errs {
		if i > 0 {
			sb.WriteRune(';')
		}
		sb.WriteString(err.Error())
	}
	return sb.String()
}

type ErrorGroup struct {
	errs []error
	lock *sync.RWMutex
}

func NewErrorGroup() *ErrorGroup {
	return &ErrorGroup{
		lock: &sync.RWMutex{},
	}
}

func (eg *ErrorGroup) Add(err error) {
	if err == nil {
		return
	}

	eg.lock.Lock()
	defer eg.lock.Unlock()

	eg.errs = append(eg.errs, err)
}

func (eg *ErrorGroup) Unwrap() error {
	eg.lock.RLock()
	defer eg.lock.RUnlock()

	if len(eg.errs) == 0 {
		return nil
	} else if len(eg.errs) == 1 {
		return eg.errs[0]
	} else {
		return Errors(eg.errs)
	}
}
