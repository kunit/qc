package parser

import (
	"github.com/kunit/qc/tcpdp"
	"reflect"
	"testing"
)

func TestParseJsonLog(t *testing.T) {
	tests := []struct {
		name string
		json string
		want *tcpdp.LogItem
	}{
		{
			name: "query log",
			json: `{"ts": "2019-03-12T12:00:00.000Z","src_addr": "192.168.33.1:50000","dst_addr": "192.168.33.10:3306","query": "SELECT * FROM members WHERE id = 1","seq_num": 0,"command_id": 1,"interface": "eth1","probe_target_addr": "3306","conn_id": "xyz12345678901234567",
  "mss": 1460,"character_set": "utf8","username": "testuser","database": "testdb"}`,
			want: &tcpdp.LogItem{
				Timestamp:       "2019-03-12T12:00:00.000Z",
				SrcAddr:         "192.168.33.1:50000",
				DstAddr:         "192.168.33.10:3306",
				SeqNum:          0,
				CommandID:       1,
				Interface:       "eth1",
				ProbeTargetAddr: "3306",
				ConnID:          "xyz12345678901234567",
				Mss:             1460,
				CharacterSet:    "utf8",
				Username:        "testuser",
				Database:        "testdb",
				Query:           "SELECT * FROM members WHERE id = 1",
			},
		},
		{
			name: "stmt_prepare_query log",
			json: `{"ts": "2019-03-12T12:00:00.000Z","src_addr": "192.168.33.1:50000","dst_addr": "192.168.33.10:3306","stmt_prepare_query": "SELECT * FROM members WHERE id = ?","seq_num": 0,"command_id": 1,"interface": "eth1","probe_target_addr": "3306","conn_id": "xyz12345678901234567",
  "mss": 1460,"character_set": "utf8","username": "testuser","database": "testdb"}`,
			want: &tcpdp.LogItem{
				Timestamp:        "2019-03-12T12:00:00.000Z",
				SrcAddr:          "192.168.33.1:50000",
				DstAddr:          "192.168.33.10:3306",
				SeqNum:           0,
				CommandID:        1,
				Interface:        "eth1",
				ProbeTargetAddr:  "3306",
				ConnID:           "xyz12345678901234567",
				Mss:              1460,
				CharacterSet:     "utf8",
				Username:         "testuser",
				Database:         "testdb",
				StmtPrepareQuery: "SELECT * FROM members WHERE id = ?",
			},
		},
		{
			name: "stmt_execute_values log",
			json: `{"ts": "2019-03-12T12:00:00.000Z","src_addr": "192.168.33.1:50000","dst_addr": "192.168.33.10:3306","stmt_id":1,"stmt_execute_values": ["1"],"seq_num": 0,"command_id": 1,"interface": "eth1","probe_target_addr": "3306","conn_id": "xyz12345678901234567",
  "mss": 1460,"character_set": "utf8","username": "testuser","database": "testdb"}`,
			want: &tcpdp.LogItem{
				Timestamp:         "2019-03-12T12:00:00.000Z",
				SrcAddr:           "192.168.33.1:50000",
				DstAddr:           "192.168.33.10:3306",
				SeqNum:            0,
				CommandID:         1,
				Interface:         "eth1",
				ProbeTargetAddr:   "3306",
				ConnID:            "xyz12345678901234567",
				Mss:               1460,
				CharacterSet:      "utf8",
				Username:          "testuser",
				Database:          "testdb",
				StmtID:            1,
				StmtExecuteValues: []interface{}{"1"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseJSONLogItem([]byte(tt.json))
			if err != nil {
				t.Errorf("ParseJSONLogItem() error = %#v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseJSONLogItem() got = %#v, want = %#v", got, tt.want)
			}
		})
	}

}
