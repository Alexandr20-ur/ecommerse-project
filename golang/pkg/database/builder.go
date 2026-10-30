package database

import (
	"errors"
	"fmt"
	"strings"
)

type Builder struct {
	selectParts []string
	from        string
	joins       []join
	wheres      []where
	orderBy     []string
	args        []any
	err         error
}

type join struct {
	kind  string
	table string
	on    string
}

type where struct {
	column string
	op     string
	value  any
}

var allowedOps = map[string]struct{}{
	"=": {}, "<>": {}, ">": {}, ">=": {}, "<": {}, "<=": {},
}

func Query() *Builder {
	return &Builder{}
}

func (b *Builder) Select(selectParts ...string) *Builder {
	b.selectParts = append(b.selectParts, selectParts...)
	return b
}

func (b *Builder) From(from string) *Builder {
	b.from = from
	return b
}

func (b *Builder) Join(kind, table, on string) *Builder {
	b.joins = append(b.joins, join{kind: kind, table: table, on: on})
	return b
}

func (b *Builder) Where(column, op string, value any) *Builder {
	if !isAllowedOp(op) {
		b.err = errors.New("invalid operator: " + op)
		return b
	}
	b.wheres = append(b.wheres, where{column: column, op: op, value: value})
	b.args = append(b.args, value)
	return b
}

func (b *Builder) OrderBy(column string) *Builder {
	b.orderBy = append(b.orderBy, column)
	return b
}

func (b *Builder) ToSql() (string, []any, error) {
	if b.err != nil {
		return "", nil, b.err
	}
	if b.from == "" {
		return "", nil, fmt.Errorf("from must not be empty")
	}

	var sb strings.Builder

	sb.WriteString("SELECT ")
	if len(b.selectParts) == 0 {
		sb.WriteString("*")
	} else {
		sb.WriteString(strings.Join(b.selectParts, ","))
	}

	sb.WriteString(" FROM ")
	sb.WriteString(b.from)

	// JOIN
	for _, j := range b.joins {
		sb.WriteString(" ")
		sb.WriteString(strings.ToUpper(j.kind))
		sb.WriteString(" JOIN ")
		sb.WriteString(j.table)
		sb.WriteString(" ON ")
		sb.WriteString(j.on)
	}

	// WHERE
	if len(b.wheres) > 0 {
		sb.WriteString(" WHERE ")
		for i, w := range b.wheres {
			if i > 0 {
				sb.WriteString(" AND ")
			}
			sb.WriteString(w.column)
			sb.WriteString(" ")
			sb.WriteString(w.op)
			sb.WriteString(" ")
			// $1, $2 ... для postgres; для mysql — ?
			sb.WriteString(fmt.Sprintf("$%d", i+1))
		}
	}

	// ORDER BY
	if len(b.orderBy) > 0 {
		sb.WriteString(" ORDER BY ")
		sb.WriteString(strings.Join(b.orderBy, ", "))
	}

	return sb.String(), b.args, nil
}

func isAllowedOp(op string) bool {
	_, ok := allowedOps[op]
	return ok
}
