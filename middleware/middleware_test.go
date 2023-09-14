package middleware

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	tele "github.com/mikhailche/telebot"
)

var b, _ = tele.NewBot(tele.Settings{Offline: true})

func TestRecover(t *testing.T) {
	onError := func(err error) {
		require.Error(t, err, "recover test")
	}

	h := func(ctx context.Context, c tele.Context) error {
		panic("recover test")
	}

	assert.Panics(t, func() {
		h(context.Background(), nil)
	})

	assert.NotPanics(t, func() {
		Recover(onError)(h)(context.Background(), nil)
	})
}
