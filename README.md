# tcpdp-qc - Query Checker for tcpdp

tcpdp-qc は [tcpdp](https://github.com/k1LoW/tcpdp) が出力する json ログを解析し、実行されているSQLがプリペアドステートメントが使われているかをチェックするツールです。

## 使い方

```
$ cat dump.log
{"ts":"2019-03-11T05:10:12.837Z","src_addr":"192.168.33.1:54785","dst_addr":"192.168.33.10:3306","stmt_prepare_query":"SELECT * FROM members WHERE name = ?","seq_num":0,"command_id":22,"interface":"eth1","probe_target_addr":"3306","conn_id":"bi2urs2ab4ht1rvrtpdg","mss":1460,"character_set":"utf8","username":"test","database":"test"}
{"ts":"2019-03-11T05:10:15.838Z","src_addr":"192.168.33.1:54785","dst_addr":"192.168.33.10:3306","stmt_id":1,"stmt_execute_values":["kunit"],"seq_num":0,"command_id":23,"interface":"eth1","probe_target_addr":"3306","conn_id":"bi2urs2ab4ht1rvrtpdg","mss":1460,"character_set":"utf8","username":"test","database":"test"}
{"ts":"2019-03-11T05:11:00.050Z","src_addr":"192.168.33.1:54824","dst_addr":"192.168.33.10:3306","stmt_prepare_query":"SELECT * FROM members WHERE name = 'kunit'","seq_num":0,"command_id":22,"interface":"eth1","probe_target_addr":"3306","conn_id":"bi2uu52ab4ht0rb8uoj0","mss":1460,"character_set":"utf8","username":"test","database":"test"}
{"ts":"2019-03-11T05:16:04.053Z","src_addr":"192.168.33.1:54824","dst_addr":"192.168.33.10:3306","stmt_id":1,"stmt_execute_values":[],"seq_num":0,"command_id":23,"interface":"eth1","probe_target_addr":"3306","conn_id":"bi2uu52ab4ht0rb8uoj0","mss":1460,"character_set":"utf8","username":"test","database":"test"}

$ qc --help
Usage: qc [--version] [--help] <options>
Options:
  -a, --all                                show all queries

$ echo dump.log | qc
{"ts":"2019-03-11T05:11:41.050Z","status":"WARNING","sql":"SELECT * FROM members WHERE name = 'kunit'","values":[]}

$ echo dump.log | qc -a
{"ts":"2019-03-11T05:10:12.838Z","status":"OK","sql":"SELECT * FROM members WHERE name = ?","values":["kunit"]}
{"ts":"2019-03-11T05:11:41.053Z","status":"WARNING","sql":"SELECT * FROM members WHERE name = 'kunit'","values":[]}
```

## TODO

- [ ] 動作を確認できる docker-compose.yml を準備
- [ ] 特定のSQLを警告対象から除外する(whitelist)

## 参照

- https://github.com/k1LoW/tcpdp
