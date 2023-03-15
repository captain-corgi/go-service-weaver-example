package reverser

import (
	"context"

	"github.com/ServiceWeaver/weaver"
)

//go:generate weaver generate .

type (
	Reverser interface {
		Reverse(context.Context, string) (string, error)
	}
	reverser struct {
		weaver.Implements[Reverser]
	}
)

func (r *reverser) Reverse(ctx context.Context, s string) (string, error) {
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}

	return string(runes), nil
}
