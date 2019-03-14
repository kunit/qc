package tcpdp

// LogItem tcpdp のログの一行
type LogItem struct {
	Timestamp         string        `json:"ts"`
	SrcAddr           string        `json:"src_addr"`
	DstAddr           string        `json:"dst_addr"`
	SeqNum            uint          `json:"seq_num"`
	CommandID         uint          `json:"command_id"`
	Interface         string        `json:"interface"`
	ProbeTargetAddr   string        `json:"probe_target_addr"`
	ConnID            string        `json:"conn_id"`
	Mss               uint          `json:"mss"`
	CharacterSet      string        `json:"character_set"`
	Username          string        `json:"username"`
	Database          string        `json:"database"`
	Query             string        `json:"query"`
	StmtPrepareQuery  string        `json:"stmt_prepare_query"`
	StmtID            uint          `json:"stmt_id"`
	StmtExecuteValues []interface{} `json:"stmt_execute_values"`
}
