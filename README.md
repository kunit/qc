# qc - Query Checker for tcpdp [![Build Status](https://travis-ci.org/kunit/qc.svg?branch=master)](https://travis-ci.org/kunit/qc) [![codecov](https://codecov.io/gh/kunit/qc/branch/master/graph/badge.svg)](https://codecov.io/gh/kunit/qc)

qc is a tool that analyzes the log output by [tcpdp](https://github.com/k1LoW/tcpdp) and checks if prepared statements are used.

## Usage

```
$ cat /path/to/tcpdp-log-dir/dump.log
{"ts":"2019-03-11T05:09:22.115Z","src_addr":"192.168.33.1:57372","dst_addr":"192.168.33.10:3306","query":"SELECT * FROM members WHERE id = 1","seq_num":0,"command_id":3,"interface":"eth1","probe_target_addr":"3306","conn_id":"bi30tiiab4ht0rb8uong","mss":1460,"character_set":"latin1","username":"test","database":"test"}
{"ts":"2019-03-11T05:10:12.837Z","src_addr":"192.168.33.1:54785","dst_addr":"192.168.33.10:3306","stmt_prepare_query":"SELECT * FROM members WHERE name = ?","seq_num":0,"command_id":22,"interface":"eth1","probe_target_addr":"3306","conn_id":"bi2urs2ab4ht1rvrtpdg","mss":1460,"character_set":"utf8","username":"test","database":"test"}
{"ts":"2019-03-11T05:10:15.838Z","src_addr":"192.168.33.1:54785","dst_addr":"192.168.33.10:3306","stmt_id":1,"stmt_execute_values":["kunit"],"seq_num":0,"command_id":23,"interface":"eth1","probe_target_addr":"3306","conn_id":"bi2urs2ab4ht1rvrtpdg","mss":1460,"character_set":"utf8","username":"test","database":"test"}
{"ts":"2019-03-11T05:11:00.050Z","src_addr":"192.168.33.1:54824","dst_addr":"192.168.33.10:3306","stmt_prepare_query":"SELECT * FROM members WHERE name = 'kunit'","seq_num":0,"command_id":22,"interface":"eth1","probe_target_addr":"3306","conn_id":"bi2uu52ab4ht0rb8uoj0","mss":1460,"character_set":"utf8","username":"test","database":"test"}
{"ts":"2019-03-11T05:16:04.053Z","src_addr":"192.168.33.1:54824","dst_addr":"192.168.33.10:3306","stmt_id":1,"stmt_execute_values":[],"seq_num":0,"command_id":23,"interface":"eth1","probe_target_addr":"3306","conn_id":"bi2uu52ab4ht0rb8uoj0","mss":1460,"character_set":"utf8","username":"test","database":"test"}

$ qc --help
Usage: qc [--version] [--help] <options>
Options:
  -a, --all                                show all queries

$ echo /path/to/tcpdp-log-dir/dump.log | qc
{"ts":"2019-03-11T05:09:22.115Z","status":"WARNING","sql":"SELECT * FROM members WHERE id = 1","values":[]}
{"ts":"2019-03-11T05:11:41.050Z","status":"WARNING","sql":"SELECT * FROM members WHERE name = 'kunit'","values":[]}

$ echo /path/to/tcpdp-log-dir/dump.log/dump.log | qc -a
{"ts":"2019-03-11T05:09:22.115Z","status":"WARNING","sql":"SELECT * FROM members WHERE id = 1","values":[]}
{"ts":"2019-03-11T05:10:12.838Z","status":"OK","sql":"SELECT * FROM members WHERE name = ?","values":["kunit"]}
{"ts":"2019-03-11T05:11:41.053Z","status":"WARNING","sql":"SELECT * FROM members WHERE name = 'kunit'","values":[]}
```

## References

- https://github.com/k1LoW/tcpdp
