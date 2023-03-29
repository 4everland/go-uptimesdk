package message

import (
	"encoding/json"
	"fmt"
)

type Quota struct {
	Name  string
	Value int
}

type Message struct {
	Value  int
	Values []Quota
}

type Header struct {
	ServerName   string `json:"serverName"`
	InstanceName string `json:"instanceName"`
}

type Package struct {
	Header
	ResourceMap string `json:"resourceMapJsonStr"`

	Extra string `json:"extra"`
}

func (q Quota) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"name":"%s", "value":"%d"}`, q.Name, q.Value)), nil
}

func (m Message) MarshalJSON() ([]byte, error) {
	if len(m.Values) > 0 {
		values, _ := json.Marshal(m.Values)
		return []byte(fmt.Sprintf(`{"value":"%d", "values":%s}`, m.Value, string(values))), nil
	}
	return []byte(fmt.Sprintf(`{"value":"%d"}`, m.Value)), nil
}
