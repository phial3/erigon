package observer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSentryCandidatesLogEventEthVersion(t *testing.T) {
	event := SentryCandidatesLogEvent{}
	event.Capabilities = []string{"wit/0", "eth/64", "eth/65", "eth/66"}
	version := event.EthVersion()
	assert.Equal(t, uint(66), version)
}
