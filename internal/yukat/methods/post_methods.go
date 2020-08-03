package methods

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

// RunMethod ...
func (m *Method) RunMethod(method string, reqBody interface{}) error {
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	res, err := http.Post("https://api.telegram.org/bot"+m.token+"/"+method, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("Unexpected status " + res.Status)
	}

	return nil
}
