package xraw

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type (
	// XrawEngine - Engine of Raw Query ORM Library
	XrawEngine interface {
		SetTableOptions(tbCaseFormat, tbPrefix string)
		SetColumnOptions(columnFormat, columnPrefix string)
		Cols(col string, otherCols ...string) XrawEngine
		SetIsMultiRows(state bool)
		SetDB(db *sqlx.DB)
		GetDB() *sqlx.DB
		GetPreparedValues() []interface{}
		GetMultiPreparedValues() [][]interface{}

		PrepareData(ctx context.Context, command string, data interface{}) error
		ExecuteCUDQuery(ctx context.Context, preparedValue ...interface{}) (int64, error)
		Clear()
		GenerateSelectQuery()
		GenerateRawCUDQuery(command string, data interface{})

		GetLastQuery() string
		// GetResults() []map[string]string
		// GetSingleResult() map[string]string

		Select(col ...string) XrawEngine
		SelectSum(col string, colAlias ...string) XrawEngine
		SelectAverage(col string, colAlias ...string) XrawEngine
		SelectMax(col string, colAlias ...string) XrawEngine
		SelectMin(col string, colAlias ...string) XrawEngine
		SelectCount(col string, colAlias ...string) XrawEngine

		Where(col string, value interface{}, opt ...string) XrawEngine
		WhereRaw(args string, value ...interface{}) XrawEngine

		WhereIn(col string, listOfValues ...interface{}) XrawEngine
		WhereNotIn(col string, listOfValues ...interface{}) XrawEngine
		WhereLike(col, value string) XrawEngine
		WhereBetween(col string, val1, val2 interface{})
		WhereNotBetween(col string, val1, val2 interface{})

		Or(col string, value interface{}, opt ...string) XrawEngine
		OrIn(col string, listOfValues ...interface{}) XrawEngine
		OrNotIn(col string, listOfValues ...interface{}) XrawEngine
		OrLike(col, value string) XrawEngine
		OrBetween(col string, val1, val2 interface{})
		OrNotBetween(col string, val1, val2 interface{})

		OrderBy(col, value string) XrawEngine
		Asc(col string) XrawEngine
		Desc(col string) XrawEngine

		Limit(limit int, offset ...int) XrawEngine
		From(tableName string) XrawEngine

		SQLRaw(rawQuery string, values ...interface{}) XrawEngine
		Get(pointerStruct interface{}) error

		Insert(data interface{}) (int64, error)
		Update(data interface{}) (int64, error)
		Delete(data interface{}) (int64, error)
	}
	// XrawTransaction - Transaction Structure
	XrawTransaction struct {
	}
)
