package rmq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendMessage(t *testing.T) {
	err := SendMessage("amqps://rigqkizo:CBNChsj9lZoMSSzHXKB84-0glFjLZsT8@hawk.rmq.cloudamqp.com/rigqkizo", "test5")
	require.NoError(t, err)
}
