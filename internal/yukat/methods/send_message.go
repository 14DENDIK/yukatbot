package methods

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/14DENDIK/yukatbot/api/telegram"
)

// SendMessage ...
func (m *Method) SendMessage(reqBody *telegram.SendMessage) error {
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	res, err := http.Post("https://api.telegram.org/bot"+m.token+"/sendMessage", "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("Unexpected status " + res.Status)
	}

	return nil
}
