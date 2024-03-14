package rmq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRecieveMessage(t *testing.T) {
	_, err := RecieveMessage("amqps://rigqkizo:CBNChsj9lZoMSSzHXKB84-0glFjLZsT8@hawk.rmq.cloudamqp.com/rigqkizo")
	require.NoError(t, err)

}
