package typrails

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"

	sq "github.com/Masterminds/squirrel"
	"github.com/typical-go/typical-go/pkg/common"
	"go.uber.org/dig"
)

// Fetcher responsible to fetch entity
type Fetcher struct {
	dig.In
	*sql.DB
}

// InfoSchema is infomation schema from database
type InfoSchema struct {
	ColumnName string
	DataType   string
}

// Fetch entity from database based on table name
func (f *Fetcher) Fetch(pkg, table, entity string) (e *Entity, err error) {
	var infos []InfoSchema
	if infos, err = f.infoSchema(table); err != nil {
		return
	}
	if len(infos) < 1 {
		err = fmt.Errorf("No column in '%s'", table)
		return
	}
	var std common.StringDictionary
	std.Add("id", "int4")
	std.Add("updated_at", "timestamp")
	std.Add("created_at", "timestamp")
	fields := f.convert(infos)
	if err = f.validate(std, fields); err != nil {
		return
	}
	e = &Entity{
		Name:           entity,
		Table:          table,
		Type:           strcase.ToCamel(entity),
		Cache:          strings.ToUpper(table),
		ProjectPackage: pkg,
		Fields:         fields,
		Forms:          f.filter(std, fields),
	}
	return
}

func (f *Fetcher) filter(std common.StringDictionary, fields []Field) (filtered []Field) {
fields:
	for _, field := range fields {
		for _, ks := range std {
			if ks.Key == field.Column {
				continue fields
			}
		}
		filtered = append(filtered, field)
	}
	return
}

func (f *Fetcher) validate(std common.StringDictionary, fields []Field) (err error) {
	fieldMap := make(map[string]string)
	for _, field := range fields {
		fieldMap[field.Column] = field.Udt
	}
	var errs common.Errors
	for _, ks := range std {
		if udt, ok := fieldMap[ks.Key]; ok {
			if ks.Value == udt {
				continue
			}
		}
		errs.Append(fmt.Errorf("\"%s\" with underlying data type \"%s\" is missing",
			ks.Key, ks.Value))
	}

	err = errs.Unwrap()
	return
}

//
func (f *Fetcher) convert(infos []InfoSchema) (fields []Field) {
	for _, info := range infos {
		fields = append(fields, CreateField(info.ColumnName, info.DataType))
	}
	return
}

func (f *Fetcher) infoSchema(tableName string) (infos []InfoSchema, err error) {
	query := sq.Select("column_name", "udt_name").
		From("information_schema.COLUMNS").
		Where(sq.Eq{"table_name": tableName}).
		RunWith(f.DB).
		PlaceholderFormat(sq.Dollar)
	var rows *sql.Rows
	if rows, err = query.Query(); err != nil {
		return
	}
	for rows.Next() {
		var info InfoSchema
		if err = rows.Scan(&info.ColumnName, &info.DataType); err != nil {
			return
		}
		infos = append(infos, info)
	}
	return
}
