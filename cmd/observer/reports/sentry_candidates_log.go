package reports

import (
	"bufio"
	"encoding/json"
	"io"
)

type SentryCandidatesLog struct {
	scanner *bufio.Scanner
}

type SentryCandidatesLogEvent struct {
	Message   string `json:"msg"`
	PeerIDHex string `json:"peer,omitempty"`
	NodeURL   string `json:"nodeURL,omitempty"`
	ClientID  string `json:"clientID,omitempty"`
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
