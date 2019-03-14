package parser

import (
	"testing"
)

func TestParseSQL(t *testing.T) {
	tests := []struct {
		name string
		sql  string
		want bool
	}{
		{
			name: "simple query",
			sql:  `SELECT id FROM foo WHERE id = ?`,
			want: true,
		},
		{
			name: "sub query",
			sql:  `SELECT id FROM foo WHERE id IN (SELECT id FROM bar WHERE age > ?)`,
			want: true,
		},
		{
			name: "and expr",
			sql:  `SELECT id FROM foo WHERE id = ? AND age > ?`,
			want: true,
		},
		{
			name: "or expr",
			sql:  `SELECT id FROM foo WHERE id = ? OR age > ?`,
			want: true,
		},
		{
			name: "and/or expr",
			sql:  `SELECT id FROM foo WHERE id = ? AND age > ? OR age < ?`,
			want: true,
		},
		{
			name: "not expr",
			sql:  `SELECT id FROM foo WHERE NOT age > ?`,
			want: true,
		},
		{
			name: "paren expr",
			sql:  `SELECT id FROM foo WHERE NOT (age > ? OR age < ?)`,
			want: true,
		},
		{
			name: "range cond expr",
			sql:  `SELECT id FROM foo WHERE age BETWEEN ? AND ? `,
			want: true,
		},
		{
			name: "is expr",
			sql:  `SELECT id FROM foo WHERE age IS NULL`,
			want: true,
		},
		{
			name: "exists expr",
			sql:  `SELECT id FROM foo WHERE EXISTS (SELECT id FROM bar WHERE age > ?)`,
			want: true,
		},
		{
			name: "tuple value expr",
			sql:  `SELECT id FROM foo WHERE id IN (?, ?, ?)`,
			want: true,
		},
		{
			name: "sub query in select exprs",
			sql:  `SELECT id, (SELECT name FROM bar WHERE age > ?) AS name FROM foo WHERE id = ?`,
			want: true,
		},
		{
			name: "aliased table expr",
			sql:  `SELECT id FROM foo, bar AS b WHERE id = ?`,
			want: true,
		},
		{
			name: "aliased table expr with sub query",
			sql:  `SELECT id FROM foo, (SELECT name bar WHERE age > ?) AS bar WHERE id = ?`,
			want: true,
		},
		{
			name: "paren table expr",
			sql:  `SELECT id FROM (foo, bar) WHERE id = ?`,
			want: true,
		},
		{
			name: "paren table expr with sub query",
			sql:  `SELECT id FROM (foo, (SELECT name bar WHERE age > ?) AS bar) WHERE id = ?`,
			want: true,
		},
		{
			name: "join table expr",
			sql:  `SELECT id FROM foo JOIN bar ON foo.id = bar.id WHERE foo.id = ?`,
			want: true,
		},
		{
			name: "join table expr with sub query",
			sql:  `SELECT id FROM foo JOIN (SELECT * FROM bar WHERE age > ?) AS bar ON foo.id = bar.id WHERE foo.id = ?`,
			want: true,
		},
		{
			name: "simple query",
			sql:  `SELECT id FROM foo WHERE id = 1`,
			want: false,
		},
		{
			name: "sub query",
			sql:  `SELECT id FROM foo WHERE id IN (SELECT id FROM bar WHERE age > 20)`,
			want: false,
		},
		{
			name: "and expr",
			sql:  `SELECT id FROM foo WHERE id = 1 AND age > ?`,
			want: false,
		},
		{
			name: "or expr",
			sql:  `SELECT id FROM foo WHERE id = 1 OR age > ?`,
			want: false,
		},
		{
			name: "and/or expr",
			sql:  `SELECT id FROM foo WHERE id = 1 AND age > ? OR age < ?`,
			want: false,
		},
		{
			name: "not expr",
			sql:  `SELECT id FROM foo WHERE NOT age > 20`,
			want: false,
		},
		{
			name: "paren expr",
			sql:  `SELECT id FROM foo WHERE NOT (age > 20 OR age < ?)`,
			want: false,
		},
		{
			name: "range cond expr",
			sql:  `SELECT id FROM foo WHERE age BETWEEN 20 AND ? `,
			want: false,
		},
		{
			name: "exists expr",
			sql:  `SELECT id FROM foo WHERE EXISTS (SELECT id FROM bar WHERE age > 20)`,
			want: false,
		},
		{
			name: "tuple value expr",
			sql:  `SELECT id FROM foo WHERE id IN (20, ?, ?)`,
			want: false,
		},
		{
			name: "sub query in select exprs",
			sql:  `SELECT id, (SELECT name FROM bar WHERE age > 20) AS name FROM foo WHERE id = ?`,
			want: false,
		},
		{
			name: "aliased table expr",
			sql:  `SELECT id FROM foo, bar AS b WHERE id = 1`,
			want: false,
		},
		{
			name: "aliased table expr with sub query",
			sql:  `SELECT id FROM foo, (SELECT name bar WHERE age > 20) AS bar WHERE id = ?`,
			want: false,
		},
		{
			name: "paren table expr",
			sql:  `SELECT id FROM (foo, bar) WHERE id = 1`,
			want: false,
		},
		{
			name: "paren table expr with sub query",
			sql:  `SELECT id FROM (foo, (SELECT name bar WHERE age > 20) AS bar) WHERE id = ?`,
			want: false,
		},
		{
			name: "join table expr",
			sql:  `SELECT id FROM foo JOIN bar ON foo.id = bar.id WHERE foo.id = 1`,
			want: false,
		},
		{
			name: "join table expr with sub query",
			sql:  `SELECT id FROM foo JOIN (SELECT * FROM bar WHERE age > 20) AS bar ON foo.id = bar.id WHERE foo.id = ?`,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseSQL(tt.sql)
			if err != nil {
				t.Errorf("ParseSQL() error = %#v", err)
			}
			if got.IsPrepared != tt.want {
				t.Errorf("ParseSQL() got = %#v, want = %#v", got, tt.want)
			}
		})
	}
}
