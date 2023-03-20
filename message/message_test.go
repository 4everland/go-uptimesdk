package message

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMarshalJSON(t *testing.T) {
	m := Message{
		Value: 10,
		Values: []Quota{
			{Name: "abc", Value: 50},
		},
	}
	b, err := json.Marshal(m)
	fmt.Println(err)
	fmt.Println(string(b))
}
