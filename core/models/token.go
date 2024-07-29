package models

import "fmt"

type Token struct {
	Username string `json:"username"`
	Value    string `json:"value"`
}

func (t Token) Key() string {
	return fmt.Sprintf("token:%s", t.Value)
}
