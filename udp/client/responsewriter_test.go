package client

import (
	"bytes"
	"context"
	"github.com/plgd-dev/go-coap/v2/message"
	"github.com/plgd-dev/go-coap/v2/message/codes"
	"github.com/plgd-dev/go-coap/v2/udp/message/pool"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestResponseWriter_SetResponse(t *testing.T) {
	msg := pool.AcquireMessage(context.Background())
	w := NewResponseWriter(msg, nil, message.Options{})

	err := w.SetResponse(codes.Valid, message.AppCBOR, bytes.NewReader([]byte{1}))
	require.NoError(t, err)

	w.SendReset()
}
