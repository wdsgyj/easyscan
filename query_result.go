// query_result
package easyscan

import (
	"bytes"
	"database/sql"
	"fmt"
)

type QueryResult struct {
	rows    *sql.Rows
	buf     []interface{}
	values  map[string]*interface{}
	columns []string
}

func NewQueryResult(r *sql.Rows) (*QueryResult, error) {
	rs := new(QueryResult)
	rs.rows = r
	err := rs.buildBuf()
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func (self *QueryResult) buildBuf() error {
	if self.values == nil {
		var err error
		self.columns, err = self.rows.Columns()
		if err != nil {
			return err
		}

		size := len(self.columns)
		self.values = make(map[string]*interface{}, size)
		self.buf = make([]interface{}, size)
		raw := make([]interface{}, size)

		for i, column := range self.columns {
			self.values[column] = &raw[i]
			self.buf[i] = &raw[i]
		}
	}

	return nil
}

func (self *QueryResult) Next() bool {
	rs := false
	if self.rows.Next() {
		rs = true
		// fill data to self.values
		self.rows.Scan(self.buf...)
	}
	return rs
}

func (self *QueryResult) Close() error {
	rs := self.rows.Close()
	for i := range self.buf {
		self.buf[i] = nil
	}
	for k := range self.values {
		self.values[k] = nil
	}
	return rs
}

func (self *QueryResult) Dump() string {
	buf := bytes.NewBuffer(nil)
	first := true

	str := func(i interface{}) interface{} {
		switch va := i.(type) {
		case []byte:
			return string(va)
		default:
			return va
		}
	}

	for _, v := range self.columns {
		if first {
			first = false
		} else {
			fmt.Fprintln(buf)
		}
		fmt.Fprintf(buf, "%s (%T): %v", v, *self.values[v], str(*self.values[v]))
	}
	return buf.String()
}

func (self *QueryResult) ColumnSize() int {
	return len(self.columns)
}

func (self *QueryResult) ColumnNames() []string {
	rs := make([]string, self.ColumnSize())
	copy(rs, self.columns)
	return rs
}

func (self *QueryResult) get(i int) interface{} {
	return *(self.buf[i].(*interface{}))
}

func (self *QueryResult) getColumn(column string) interface{} {
	return *(self.values[column])
}

func (self *QueryResult) ValueOf(i int) interface{} {
	return self.get(i)
}

func (self *QueryResult) ValueOfColumn(column string) interface{} {
	return self.getColumn(column)
}

func (self *QueryResult) StringOf(i int) (string, bool) {
	v := self.get(i)
	if v == nil {
		return "", false
	}
	return any2String(v)
}

func (self *QueryResult) StringOfColumn(column string) (string, bool) {
	v := self.getColumn(column)
	if v == nil {
		return "", false
	}
	return any2String(v)
}

func (self *QueryResult) IntOf(i int) (int64, bool) {
	v := self.get(i)
	if v == nil {
		return 0, false
	}
	return any2Int(v)
}

func (self *QueryResult) IntOfColumn(column string) (int64, bool) {
	v := self.getColumn(column)
	if v == nil {
		return 0, false
	}
	return any2Int(v)
}

func (self *QueryResult) FloatOf(i int) (float64, bool) {
	v := self.get(i)
	if v == nil {
		return 0, false
	}
	return any2Float(v)
}

func (self *QueryResult) FloatOfColumn(column string) (float64, bool) {
	v := self.getColumn(column)
	if v == nil {
		return 0, false
	}
	return any2Float(v)
}

func (self *QueryResult) BytesOf(i int) ([]byte, bool) {
	v := self.get(i)
	if v == nil {
		return nil, false
	}
	return any2Bytes(v)
}

func (self *QueryResult) BytesOfColumn(column string) ([]byte, bool) {
	v := self.getColumn(column)
	if v == nil {
		return nil, false
	}
	return any2Bytes(v)
}

func (self *QueryResult) BoolOf(i int) (bool, bool) {
	v := self.get(i)
	if v == nil {
		return false, false
	}
	return any2Bool(v)
}

func (self *QueryResult) BoolOfColumn(column string) (bool, bool) {
	v := self.getColumn(column)
	if v == nil {
		return false, false
	}
	return any2Bool(v)
}
