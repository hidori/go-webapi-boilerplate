package xtest

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

// Case is test case.
type Case[TContext any, TFields any, TArgs any, TWant any] struct {
	Name      string
	Context   TContext
	Fields    TFields
	Args      TArgs
	Up        func(tt Case[TContext, TFields, TArgs, TWant]) error
	Down      func(tt Case[TContext, TFields, TArgs, TWant]) error
	Want      TWant
	WantError bool
	Error     string
	WantPanic bool
	More      func(t *testing.T, tt Case[TContext, TFields, TArgs, TWant], got interface{})
}

// Run runs specified test case.
func Run[TContext any, TFields any, TArgs any, TWant any](t *testing.T, tt Case[TContext, TFields, TArgs, TWant], fc func(t *testing.T, tt Case[TContext, TFields, TArgs, TWant])) {
	t.Helper()

	t.Run(tt.Name, func(t *testing.T) {
		if tt.Up != nil {
			err := tt.Up(tt)
			if err != nil {
				t.Errorf("fail to tt.Up(): err=%v, tt=%v", err, tt)
				return
			}
		}

		if !assert.NotPanics(t, func() { fc(t, tt) }) {
			t.Errorf("panics fc(): tt=%v", tt)
			return
		}

		if tt.Down != nil {
			err := tt.Down(tt)
			if err != nil {
				t.Errorf("fail to tt.Down(): err=%v, tt=%v", err, tt)
				return
			}
		}
	})
}

// Equal compares want and got.
func Equal(t *testing.T, want interface{}, got interface{}, opts ...cmp.Option) bool {
	t.Helper()

	diff := cmp.Diff(want, got, opts...)
	if diff != "" {
		t.Errorf(diff)
		return false
	}

	return true
}
