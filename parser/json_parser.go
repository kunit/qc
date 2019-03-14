package parser

import (
	"encoding/json"
	"fmt"
	"github.com/kunit/qc/tcpdp"
)

// ParseJSONLogItem tcpdp が出力するログの1行をパースする
func ParseJSONLogItem(b []byte) (*tcpdp.LogItem, error) {
	var item tcpdp.LogItem
	if err := json.Unmarshal(b, &item); err != nil {
		return nil, fmt.Errorf("json parse error")
	}

	return &item, nil
}
