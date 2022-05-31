package observer

import (
	"bufio"
	"encoding/json"
	"github.com/ledgerwatch/erigon/eth/protocols/eth"
	"io"
	"strconv"
	"strings"
)

type SentryCandidatesLog struct {
	scanner *bufio.Scanner
}

type SentryCandidatesLogEvent struct {
	Message      string   `json:"msg"`
	PeerIDHex    string   `json:"peer,omitempty"`
	NodeURL      string   `json:"nodeURL,omitempty"`
	ClientID     string   `json:"clientID,omitempty"`
	Capabilities []string `json:"capabilities,omitempty"`
}

func NewSentryCandidatesLog(logReader io.Reader) *SentryCandidatesLog {
	scanner := bufio.NewScanner(logReader)
	return &SentryCandidatesLog{scanner}
}

func (log *SentryCandidatesLog) Read() (*SentryCandidatesLogEvent, error) {
	for log.scanner.Scan() {
		lineData := log.scanner.Bytes()
		var event SentryCandidatesLogEvent
		if err := json.Unmarshal(lineData, &event); err != nil {
			return nil, err
		}
		if event.Message == "Sentry peer did Connect" {
			return &event, nil
		}
	}
	return nil, log.scanner.Err()
}

func (event *SentryCandidatesLogEvent) EthVersion() uint {
	var maxVersion uint64
	for _, capability := range event.Capabilities {
		if !strings.HasPrefix(capability, eth.ProtocolName) {
			continue
		}
		versionStr := capability[len(eth.ProtocolName)+1:]
		version, _ := strconv.ParseUint(versionStr, 10, 32)
		if version > maxVersion {
			maxVersion = version
		}
	}
	return uint(maxVersion)
}
