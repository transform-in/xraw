package xraw

import (
	"errors"
	"log"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/radityaapratamaa/xraw/lib"

	"github.com/jmoiron/sqlx"
)

func (re *Engine) SetDB(db *sqlx.DB) {
	re.db = db
}

func (re *Engine) GetDB() *sqlx.DB {
	return re.db
}

func (re *Engine) GetPreparedValues() []interface{} {
	return re.preparedValue
}
func (re *Engine) GetMultiPreparedValues() [][]interface{} {
	return re.multiPreparedValue
}
func (re *Engine) Clear() {
	re.clearField()
}

func (re *Engine) SetIsMultiRows(state bool) {
	re.isMultiRows = state
}

// // New - init new RORM Engine
// func New(dbDriver, connectionURL string, tbPrefix ...string) *Engine {
// 	re := &Engine{}
// 	if err := re.Connect(dbDriver, connectionURL, tbPrefix...); err != nil {
// 		return nil
// 	}
// 	return re
// }

// New - init new RORM Engine
func New(cfg *DbConfig) (*Engine, error) {
	var err error
	re := &Engine{
		config:  cfg,
		options: &DbOptions{},
	}
	re.connectionString, err = generateConnectionString(cfg)
	if err != nil {
		return nil, err
	}
	if err := re.Connect(cfg.Driver, re.connectionString); err != nil {
		log.Println("Cannot Connect to DB: ", err.Error())
		return nil, err
	}
	// set default for table case format
	re.options.tbFormat = "snake"
	re.options.colFormat = "snake"
	re.syntaxQuote = "`"
	if cfg.Driver != "mysql" {
		re.syntaxQuote = "\""
	}
	return re, nil
}

func (re *Engine) GetConnectionString() string {
	return re.connectionString
}

func generateConnectionString(cfg *DbConfig) (connectionString string, err error) {
	if strings.TrimSpace(cfg.Driver) == "" || strings.TrimSpace(cfg.Host) == "" || strings.TrimSpace(cfg.Username) == "" || strings.TrimSpace(cfg.DbName) == "" {
		err = errors.New("Config is not set correctly")
		return
	}
	dbURL := &url.URL{
		Scheme:   cfg.Driver,
		User:     url.UserPassword(cfg.Username, cfg.Password),
		Host:     cfg.Host,
		RawQuery: cfg.Options,
	}

	query := url.Values{}
	dbURL.Path = cfg.DbName
	switch cfg.Driver {
	case "postgres":
		query.Add("sslmode", "disable")
		query.Add("search_path", "public")
		if cfg.DbScheme != "" {
			query.Set("search_path", cfg.DbScheme)
		}
		dbURL.Host += ":5432"
		if cfg.Port != "" {
			dbURL.Host = cfg.Host + ":" + cfg.Port
		}
	case "mysql":
		if cfg.Protocol == "" {
			cfg.Protocol = "tcp"
		}
		dbURL.Host += ":3306"
		if cfg.Port != "" {
			dbURL.Host = cfg.Host + ":" + cfg.Port
		}
		replacedStr := "@" + cfg.Protocol + "($1:$2)/"

		rgx := regexp.MustCompile(`@([a-zA-Z0-9]+):([0-9]+)/`)
		res := rgx.ReplaceAllString(dbURL.String(), replacedStr)
		res = res[8:]
		connectionString = res
		return
	case "sqlserver":
		dbURL.Path = ""
		if cfg.DbInstance != "" {
			dbURL.Path = cfg.DbInstance
		}
		query.Add("database", cfg.DbName)
	}
	dbURL.RawQuery = query.Encode()
	connectionString = dbURL.String()
	return
}

func (re *Engine) extractTableName(data interface{}) reflect.Value {
	dValue := reflect.ValueOf(data).Elem()

	sdValue := dValue
	if dValue.Kind() == reflect.Slice {
		re.isMultiRows = true
		re.multiPreparedValue = nil
		for i := 0; i < dValue.Len(); i++ {
			sdValue = dValue.Index(i)
			re.preparedValue = nil
			if i == dValue.Len()-1 {
				break
			}
			if sdValue.Kind() == reflect.Ptr {
				sdValue = sdValue.Elem()
			}
			for x := 0; x < sdValue.NumField(); x++ {
				if _, valid := re.getAndValidateTag(sdValue, x); !valid {
					continue
				}
				field := sdValue.Field(x)
				re.preparedValue = append(re.preparedValue, field.Interface())
			}
			re.multiPreparedValue = append(re.multiPreparedValue, re.preparedValue)
		}
	}
	tblName := ""
	switch sdValue.Kind() {
	case reflect.Ptr:
		sdValue = sdValue.Elem()
		tblName = sdValue.Type().Name()
	case reflect.Slice:
		sdType := dValue.Type()
		strName := sdType.String()
		tblName = strName[strings.Index(strName, ".")+1:]
	default:
		tblName = sdValue.Type().Name()
	}
	if re.options.tbFormat == "snake" {
		re.tableName = lib.CamelToSnakeCase(tblName)
	} else {
		re.tableName = lib.SnakeToCamelCase(tblName)
	}
	// re.tableName = re.syntaxQuote + re.tableName + re.syntaxQuote
	return sdValue
}

// Connect - connect to db Driver
func (re *Engine) Connect(dbDriver, connectionURL string) error {
	var err error
	re.db, err = sqlx.Open(dbDriver, connectionURL)
	return err
}

func (re *Engine) clearField() {
	re.condition = ""
	re.column = ""
	re.orderBy = ""
	re.tableName = ""
	re.limit = ""
	re.join = ""
	re.isRaw = false
	re.isBulk = false
	re.isMultiRows = false
	re.preparedValue = nil
	re.multiPreparedValue = nil
	re.counter = 0
	re.bulkCounter = 0
	re.updatedCol = nil
	if re.stmt != nil {
		re.stmt.Close()
	}
}

func (re *Engine) getAndValidateTag(field reflect.Value, keyIndex int) (string, bool) {

	fieldType := field.Type().Field(keyIndex)
	if fieldType.Type.Kind() == reflect.Ptr && field.Field(keyIndex).IsNil() {
		return "", false
	}
	fieldValue := field.Field(keyIndex)
	colNameTag := ""
	var valid bool
	if colNameTag, valid = re.checkStructTag(fieldType.Tag, fieldValue); !valid {

		return "", false
	}
	if colNameTag == "" {
		colNameTag = fieldType.Name
		if re.options.colFormat == "snake" {
			colNameTag = lib.CamelToSnakeCase(colNameTag)
		}

	}
	return colNameTag, true
}

func (re *Engine) checkStructTag(tagField reflect.StructTag, fieldVal reflect.Value) (string, bool) {
	colName := ""
	if tagField.Get("db") == "" {
		return colName, true
	}
	colName = strings.Split(tagField.Get("db"), " ")[0]
	identifierTagArr := strings.Split(tagField.Get("xraw"), " ")
	for _, val := range identifierTagArr {
		switch val {
		case "pk", "ai":
			return colName, false
		case "date":
			if fieldVal.String() == "" {
				return colName, false
			}
		}
	}
	return colName, true
}

func (re *Engine) StartBulkOptimized() {
	re.bulkOptimized = true
	re.bulkCounter = 0
}
func (re *Engine) StopBulkOptimized() {
	re.bulkOptimized = false
	re.clearField()
}

func (re *Engine) SetTableOptions(tbCaseFormat, tbPrefix string) {
	re.options.tbFormat = tbCaseFormat
	re.options.tbPrefix = tbPrefix
}

func (re *Engine) SetColumnOptions(columnFormat, columnPrefix string) {
	re.options.colFormat = columnFormat
}

func (re *Engine) adjustPreparedParam(old string) string {
	if strings.TrimSpace(re.config.Driver) != "mysql" {

		replacement := ""
		switch re.config.Driver {
		case "postgres":
			replacement = POSTGRES_PREPARED_PARAM
		case "sqlserver":
			replacement = MSSQL_PREPARED_PARAM
		default:
			replacement = ORACLE_PREPARED_PARAM
		}

		for {
			if strings.Index(old, "?") == -1 {
				break
			}
			re.counter++
			old = strings.Replace(old, "?", replacement+strconv.Itoa(re.counter), 1)
		}
	}
	return old
}