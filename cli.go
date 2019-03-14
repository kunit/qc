package qc

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/kunit/qc/parser"
	"io"
	"os"
	"reflect"
	"strings"
)

const (
	// ExitOK for exit code
	ExitOK int = 0

	// ExitErr for exit code
	ExitErr int = 1
)

type cli struct {
	env     Env
	command string
	args    []string
	All     bool `long:"all" short:"a" description:"show all queries"`
	Help    bool `long:"help" short:"h" description:"show this help message and exit"`
	Version bool `long:"version" short:"v" description:"show the version number"`
}

// Env 実行時の環境
type Env struct {
	Out, Err io.Writer
	Args     []string
	Version  string
}

// OutputJSON 出力するjson
type OutputJSON struct {
	Timestamp string      `json:"ts"`
	Status    string      `json:"status"`
	SQL       string      `json:"sql"`
	Values    interface{} `json:"values"`
}

// RunCLI runs as cli
func RunCLI(env Env) int {
	cli := &cli{env: env, All: false}
	return cli.run()
}

func (c *cli) buildHelp(names []string) []string {
	var help []string
	t := reflect.TypeOf(cli{})

	for _, name := range names {
		f, ok := t.FieldByName(name)
		if !ok {
			continue
		}

		tag := f.Tag
		if tag == "" {
			continue
		}

		var o, a string
		if a = tag.Get("arg"); a != "" {
			a = fmt.Sprintf("=%s", a)
		}
		if s := tag.Get("short"); s != "" {
			o = fmt.Sprintf("-%s, --%s%s", tag.Get("short"), tag.Get("long"), a)
		} else {
			o = fmt.Sprintf("--%s%s", tag.Get("long"), a)
		}

		desc := tag.Get("description")
		if i := strings.Index(desc, "\n"); i >= 0 {
			var buf bytes.Buffer
			buf.WriteString(desc[:i+1])
			desc = desc[i+1:]
			const indent = "                        "
			for {
				if i = strings.Index(desc, "\n"); i >= 0 {
					buf.WriteString(indent)
					buf.WriteString(desc[:i+1])
					desc = desc[i+1:]
					continue
				}
				break
			}
			if len(desc) > 0 {
				buf.WriteString(indent)
				buf.WriteString(desc)
			}
			desc = buf.String()
		}
		help = append(help, fmt.Sprintf("  %-40s %s", o, desc))
	}

	return help
}

func (c *cli) showHelp() {
	opts := strings.Join(c.buildHelp([]string{
		"All",
	}), "\n")

	help := `
Usage: qc [--version] [--help] <options>
Options:
%s
`
	fmt.Fprintf(c.env.Out, help, opts)
}

func (c *cli) run() int {
	p := flags.NewParser(c, flags.PassDoubleDash)
	_, err := p.ParseArgs(c.env.Args)
	if err != nil || c.Help {
		c.showHelp()
		return ExitErr
	}

	if c.Version {
		fmt.Fprintf(c.env.Err, "qc version %s\n", c.env.Version)
		return ExitOK
	}

	var psql, connID string

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		item, err := parser.ParseJSONLogItem(s.Bytes())
		if err != nil {
			fmt.Fprintf(c.env.Err, "json parse error = %v\n", err.Error())
		}

		var sql string
		var values []interface{}

		if item.Query != "" {
			sql = item.Query
		} else if item.StmtPrepareQuery != "" {
			psql = item.StmtPrepareQuery
			connID = item.ConnID
		} else if item.StmtExecuteValues != nil && connID == item.ConnID {
			sql = psql
			values = item.StmtExecuteValues
			psql = ""
			connID = ""
		}

		if sql != "" {
			q, err := parser.ParseSQL(sql)
			if q != nil && err == nil {
				status := "WARNING"
				if q.IsPrepared {
					status = "OK"
				}

				if c.All || !q.IsPrepared {
					json, err := json.Marshal(&OutputJSON{
						Timestamp: item.Timestamp,
						Status:    status,
						SQL:       sql,
						Values:    values,
					})
					if err == nil {
						fmt.Fprintf(c.env.Out, "%s\n", json)
					} else {
						fmt.Fprintf(c.env.Err, "json marshal error = %v\n", err.Error())
					}
				}
			}
		}
	}

	return ExitOK
}
