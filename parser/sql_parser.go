package parser

import (
	"github.com/knocknote/vitess-sqlparser/sqlparser"
)

// Query クエリの情報
type Query struct {
	SQL        string
	IsPrepared bool
}

// ParseSQL 引数で渡されたsqlをパース
func ParseSQL(sql string) (*Query, error) {
	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		return nil, err
	}

	query := &Query{SQL: sql, IsPrepared: true}
	WalkSelect(stmt, query)

	return query, nil
}

// WalkSelect Select 文を再帰的にパース
func WalkSelect(node sqlparser.SQLNode, query *Query) {
	switch node.(type) {
	case *sqlparser.Select:
		for _, expr := range node.(*sqlparser.Select).SelectExprs {
			switch expr.(type) {
			case *sqlparser.AliasedExpr:
				switch expr.(*sqlparser.AliasedExpr).Expr.(type) {
				case *sqlparser.Subquery:
					WalkSelect(expr.(*sqlparser.AliasedExpr).Expr.(*sqlparser.Subquery).Select, query)
				}
			}
		}
		for _, expr := range node.(*sqlparser.Select).From {
			WalkTableExpr(expr, query)
		}
		if node.(*sqlparser.Select).Where != nil {
			WalkExpr(node.(*sqlparser.Select).Where.Expr, query)
		}
		if node.(*sqlparser.Select).Having != nil {
			WalkExpr(node.(*sqlparser.Select).Having.Expr, query)
		}
	}
}

// WalkTableExpr Table Expr を再帰的にパース
func WalkTableExpr(expr sqlparser.TableExpr, query *Query) {
	switch expr.(type) {
	case *sqlparser.AliasedTableExpr:
		switch expr.(*sqlparser.AliasedTableExpr).Expr.(type) {
		case *sqlparser.Subquery:
			WalkSelect(expr.(*sqlparser.AliasedTableExpr).Expr.(*sqlparser.Subquery).Select, query)
		}
	case *sqlparser.ParenTableExpr:
		for _, expr2 := range expr.(*sqlparser.ParenTableExpr).Exprs {
			WalkTableExpr(expr2, query)
		}
	case *sqlparser.JoinTableExpr:
		WalkTableExpr(expr.(*sqlparser.JoinTableExpr).LeftExpr, query)
		WalkTableExpr(expr.(*sqlparser.JoinTableExpr).RightExpr, query)
	}
}

// WalkExpr Expr をタイプに応じて再帰的にパース
func WalkExpr(expr interface{}, query *Query) {
	switch expr.(type) {
	case *sqlparser.Subquery:
		WalkSelect(expr.(*sqlparser.Subquery).Select, query)
	case *sqlparser.AndExpr:
		WalkExpr(expr.(*sqlparser.AndExpr).Left, query)
		WalkExpr(expr.(*sqlparser.AndExpr).Right, query)
	case *sqlparser.OrExpr:
		WalkExpr(expr.(*sqlparser.OrExpr).Left, query)
		WalkExpr(expr.(*sqlparser.OrExpr).Right, query)
	case *sqlparser.NotExpr:
		WalkExpr(expr.(*sqlparser.NotExpr).Expr, query)
	case *sqlparser.ParenExpr:
		WalkExpr(expr.(*sqlparser.ParenExpr).Expr, query)
	case *sqlparser.ComparisonExpr:
		WalkExpr(expr.(*sqlparser.ComparisonExpr).Left, query)
		WalkExpr(expr.(*sqlparser.ComparisonExpr).Right, query)
	case *sqlparser.RangeCond:
		WalkExpr(expr.(*sqlparser.RangeCond).Left, query)
		WalkExpr(expr.(*sqlparser.RangeCond).From, query)
		WalkExpr(expr.(*sqlparser.RangeCond).To, query)
	case *sqlparser.IsExpr:
		WalkExpr(expr.(*sqlparser.IsExpr).Expr, query)
	case *sqlparser.ExistsExpr:
		WalkSelect(expr.(*sqlparser.ExistsExpr).Subquery.Select, query)
	case sqlparser.ValTuple:
		for _, expr2 := range expr.(sqlparser.ValTuple) {
			WalkExpr(expr2, query)
		}
	case *sqlparser.SQLVal:
		if expr.(*sqlparser.SQLVal).Type != 5 {
			query.IsPrepared = false
		}
	}
}
