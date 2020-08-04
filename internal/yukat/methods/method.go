package methods

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/14DENDIK/yukatbot/api/telegram"
)

// Method ...
type Method struct {
	token string
}

// New ...
func New(token string) *Method {
	return &Method{
		token: token,
	}
}

// RunPostMethod ...
func (m *Method) RunPostMethod(method string, reqBody interface{}) error {
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

// GetChatMember ...
func (m *Method) GetChatMember(method string, reqBody *telegram.GetChatMember) (*telegram.ChatMember, error) {
	response := &telegram.ChatMember{}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	res, err := http.Post("https://api.telegram.org/bot"+m.token+"/"+method, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("Unexpected status " + res.Status)
	}

	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}

	return response, nil
}
