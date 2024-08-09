package model

import (
	"bytes"
	"context"
	"database/sql"
	"github.com/orestonce/korm"
	"strconv"
	"strings"
	"time"
)

type OrmAll struct {
	db   *sql.DB // db, tx任选其一
	tx   *sql.Tx
	mode string // sqlite, mysql
}

func (this *OrmAll) ExecRawQuery(query string, args ...interface{}) (*sql.Rows, error) {
	if this.db != nil {
		return this.db.Query(query, args...)
	} else if this.tx != nil {
		return this.tx.Query(query, args...)
	}
	panic("ExecRawQuery: OrmAll must include db or tx")
}

func OrmAllNew(db *sql.DB, mode string) *OrmAll {
	return &OrmAll{
		db:   db,
		mode: mode,
	}
}

func (this *OrmAll) ExecRaw(query string, args ...interface{}) (sql.Result, error) {
	if this.db != nil {
		return this.db.Exec(query, args...)
	} else if this.tx != nil {
		return this.tx.Exec(query, args...)
	}
	panic("ExecRaw: OrmAll must include db or tx")
}

func (this *OrmAll) MustTxCallback(cb func(db *OrmAll)) {
	if this.tx != nil {
		panic("MustSingleThreadTxCallback repeat call")
	}
	tx, err := this.db.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	cb(&OrmAll{
		tx:   tx,
		mode: this.mode,
	})
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}

type KORM_MustNewDbMysqlReq struct {
	Addr        string
	UserName    string
	Password    string
	UseDatabase string
}

func KORM_MustNewDbMysql(req KORM_MustNewDbMysqlReq) (db *sql.DB) {
	var err error

	db, err = sql.Open("mysql", req.UserName+":"+req.Password+"@tcp("+req.Addr+")/")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + req.UseDatabase)
	if err != nil {
		panic(err)
	}
	_ = db.Close()
	db, err = sql.Open("mysql", req.UserName+":"+req.Password+"@tcp("+req.Addr+")/"+req.UseDatabase+"?charset=utf8")
	if err != nil {
		panic(err)
	}
	return db
}
func KORM_MustInitTableAll(db *sql.DB, mode string) {
	var err error
	err = korm.InitTable(korm.InitTableReq{
		Mode:      mode,
		TableName: "Work_D",
		FieldList: []korm.FieldSqlDefine{
			{
				Type:         korm.SqlTypeChar255,
				Name:         "ID",
				IsPrimaryKey: true,
			}, {
				Type:         korm.SqlTypeLongBlob,
				Name:         "Name",
				IsPrimaryKey: false,
			}, {
				Type:         korm.SqlTypeLongBlob,
				Name:         "Url",
				IsPrimaryKey: false,
			}, {
				Type:         korm.SqlTypeLongBlob,
				Name:         "SaveDir",
				IsPrimaryKey: false,
			}, {
				Type:         korm.SqlTypeBigInt,
				Name:         "State",
				IsPrimaryKey: false,
			}, {
				Type:         korm.SqlTypeLongBlob,
				Name:         "Info",
				IsPrimaryKey: false,
			}, {
				Type:         korm.SqlTypeChar255,
				Name:         "CreateTime",
				IsPrimaryKey: false,
			}, {
				Type:         korm.SqlTypeChar255,
				Name:         "UpdateTime",
				IsPrimaryKey: false,
			},
		},
		Db: db,
	})
	if err != nil {
		panic(err)
	}

}

type korm_scan_resp struct {
	argList    []interface{}
	argParseFn []func(v string) (err error)
}

func korm_Work_D_scan(joinNode *korm.KORM_leftJoinNode, info *Work_D) (resp korm_scan_resp) {
	for _, one := range joinNode.SelectFieldNameList {
		switch one {
		default:
			panic("GetScanInfoCode error" + strconv.Quote(one))
		case "ID":
			{
				resp.argList = append(resp.argList, new(sql.NullString))
				tmpFn := func(v string) (err error) {
					info.ID = v

					return nil
				}
				resp.argParseFn = append(resp.argParseFn, tmpFn)
			}
		case "Name":
			{
				resp.argList = append(resp.argList, new(sql.NullString))
				tmpFn := func(v string) (err error) {
					info.Name = v

					return nil
				}
				resp.argParseFn = append(resp.argParseFn, tmpFn)
			}
		case "Url":
			{
				resp.argList = append(resp.argList, new(sql.NullString))
				tmpFn := func(v string) (err error) {
					info.Url = v

					return nil
				}
				resp.argParseFn = append(resp.argParseFn, tmpFn)
			}
		case "SaveDir":
			{
				resp.argList = append(resp.argList, new(sql.NullString))
				tmpFn := func(v string) (err error) {
					info.SaveDir = v

					return nil
				}
				resp.argParseFn = append(resp.argParseFn, tmpFn)
			}
		case "State":
			{
				resp.argList = append(resp.argList, new(sql.NullString))
				tmpFn := func(v string) (err error) {
					vi, err := strconv.ParseInt(v, 10, 64)
					if err != nil {
						return err
					}
					info.State = int(vi)

					return nil
				}
				resp.argParseFn = append(resp.argParseFn, tmpFn)
			}
		case "Info":
			{
				resp.argList = append(resp.argList, new(sql.NullString))
				tmpFn := func(v string) (err error) {
					info.Info = v

					return nil
				}
				resp.argParseFn = append(resp.argParseFn, tmpFn)
			}
		case "CreateTime":
			{
				resp.argList = append(resp.argList, new(sql.NullString))
				tmpFn := func(v string) (err error) {
					vt, err := time.Parse(time.RFC3339Nano, v)
					if err != nil {
						return err
					}
					info.CreateTime = vt

					return nil
				}
				resp.argParseFn = append(resp.argParseFn, tmpFn)
			}
		case "UpdateTime":
			{
				resp.argList = append(resp.argList, new(sql.NullString))
				tmpFn := func(v string) (err error) {
					vt, err := time.Parse(time.RFC3339Nano, v)
					if err != nil {
						return err
					}
					info.UpdateTime = vt

					return nil
				}
				resp.argParseFn = append(resp.argParseFn, tmpFn)
			}
		}
	}
	for _, sub := range joinNode.ThisLevelJoinList {
		switch sub.FieldName {
		default:
			panic("Work_D")
		}
	}
	return resp
}

type KORM_Work_D struct {
	supper *OrmAll
}

func (this *OrmAll) Work_D() *KORM_Work_D {
	return &KORM_Work_D{supper: this}
}
func korm_fillSelectFieldNameList_Work_D(joinNode *korm.KORM_leftJoinNode) {
	joinNode.SelectFieldNameList = []string{"ID", "Name", "Url", "SaveDir", "State", "Info", "CreateTime", "UpdateTime"}
	for _, sub := range joinNode.ThisLevelJoinList {
		switch sub.FieldName {
		default:
			panic("korm_fillSelectFieldNameList_Work_D" + strconv.Quote(sub.FieldName))
		}
	}
}
func (this *KORM_Work_D) MustInsert(info Work_D) {
	var err error
	vCreateTime := info.CreateTime.UTC().Format(time.RFC3339Nano)
	vUpdateTime := info.UpdateTime.UTC().Format(time.RFC3339Nano)
	_, err = this.supper.ExecRaw("INSERT INTO `Work_D`(`ID` ,`Name` ,`Url` ,`SaveDir` ,`State` ,`Info` ,`CreateTime` ,`UpdateTime` ) VALUES(?,?,?,?,?,?,?,?)", info.ID, info.Name, info.Url, info.SaveDir, info.State, info.Info, vCreateTime, vUpdateTime)
	if err != nil {
		panic(err)
	}

	return
}
func (this *KORM_Work_D) MustSet(info Work_D) {
	var err error
	vCreateTime := info.CreateTime.UTC().Format(time.RFC3339Nano)
	vUpdateTime := info.UpdateTime.UTC().Format(time.RFC3339Nano)
	_, err = this.supper.ExecRaw("REPLACE INTO `Work_D`(`ID` ,`Name` ,`Url` ,`SaveDir` ,`State` ,`Info` ,`CreateTime` ,`UpdateTime` ) VALUES(?,?,?,?,?,?,?,?)", info.ID, info.Name, info.Url, info.SaveDir, info.State, info.Info, vCreateTime, vUpdateTime)
	if err != nil {
		panic(err)
	}

	return
}

// Select Work_D
type KORM_Work_D_SelectObj struct {
	supper      *OrmAll
	bufWhere    bytes.Buffer
	argsWhere   []interface{}
	orderBy     []string
	limit       int
	offset      int
	linkOpList  []string
	isLinkBegin bool
	joinNode    *korm.KORM_leftJoinNode
	rootInfo    korm.KORM_leftJoinRootInfo
}

func (this *KORM_Work_D) Select() *KORM_Work_D_SelectObj {
	one := &KORM_Work_D_SelectObj{supper: this.supper, linkOpList: []string{"AND"}, isLinkBegin: true}
	one.joinNode = &korm.KORM_leftJoinNode{TableName: "_0"}
	one.joinNode.Root = &one.rootInfo
	one.rootInfo.TableNameIdx = 1
	return one
}

type KORM_Work_D_SelectObj_OrderByObj struct {
	tableName string
	fieldName string
	supper    *KORM_Work_D_SelectObj
}

func (this *KORM_Work_D_SelectObj_OrderByObj) ASC() *KORM_Work_D_SelectObj {
	this.supper.orderBy = append(this.supper.orderBy, this.tableName+".`"+this.fieldName+"` ASC ")
	return this.supper
}
func (this *KORM_Work_D_SelectObj_OrderByObj) DESC() *KORM_Work_D_SelectObj {
	this.supper.orderBy = append(this.supper.orderBy, this.tableName+".`"+this.fieldName+"` DESC ")
	return this.supper
}
func (this *KORM_Work_D_SelectObj) OrderBy_ID() *KORM_Work_D_SelectObj_OrderByObj {
	return &KORM_Work_D_SelectObj_OrderByObj{
		fieldName: "ID",
		tableName: this.joinNode.TableName,
		supper:    this,
	}
}
func (this *KORM_Work_D_SelectObj) OrderBy_Name() *KORM_Work_D_SelectObj_OrderByObj {
	return &KORM_Work_D_SelectObj_OrderByObj{
		fieldName: "Name",
		tableName: this.joinNode.TableName,
		supper:    this,
	}
}
func (this *KORM_Work_D_SelectObj) OrderBy_Url() *KORM_Work_D_SelectObj_OrderByObj {
	return &KORM_Work_D_SelectObj_OrderByObj{
		fieldName: "Url",
		tableName: this.joinNode.TableName,
		supper:    this,
	}
}
func (this *KORM_Work_D_SelectObj) OrderBy_SaveDir() *KORM_Work_D_SelectObj_OrderByObj {
	return &KORM_Work_D_SelectObj_OrderByObj{
		fieldName: "SaveDir",
		tableName: this.joinNode.TableName,
		supper:    this,
	}
}
func (this *KORM_Work_D_SelectObj) OrderBy_State() *KORM_Work_D_SelectObj_OrderByObj {
	return &KORM_Work_D_SelectObj_OrderByObj{
		fieldName: "State",
		tableName: this.joinNode.TableName,
		supper:    this,
	}
}
func (this *KORM_Work_D_SelectObj) OrderBy_Info() *KORM_Work_D_SelectObj_OrderByObj {
	return &KORM_Work_D_SelectObj_OrderByObj{
		fieldName: "Info",
		tableName: this.joinNode.TableName,
		supper:    this,
	}
}
func (this *KORM_Work_D_SelectObj) OrderBy_CreateTime() *KORM_Work_D_SelectObj_OrderByObj {
	return &KORM_Work_D_SelectObj_OrderByObj{
		fieldName: "CreateTime",
		tableName: this.joinNode.TableName,
		supper:    this,
	}
}
func (this *KORM_Work_D_SelectObj) OrderBy_UpdateTime() *KORM_Work_D_SelectObj_OrderByObj {
	return &KORM_Work_D_SelectObj_OrderByObj{
		fieldName: "UpdateTime",
		tableName: this.joinNode.TableName,
		supper:    this,
	}
}

func (this *KORM_Work_D_SelectObj) LimitOffset(limit int, offset int) *KORM_Work_D_SelectObj {
	this.limit = limit
	this.offset = offset
	return this
}

// pageSize: [1, n)
// pageNo:   [1,n)
func (this *KORM_Work_D_SelectObj) SetPageLimit(pageSize int, pageNo int) *KORM_Work_D_SelectObj {
	if pageSize <= 0 || pageNo <= 0 {
		panic("KORM_Work_D_SelectObj SetPageLimit error param")
	}
	this.limit = pageSize
	this.offset = pageSize * (pageNo - 1)
	return this
}
func (this *KORM_Work_D_SelectObj) MustRun_Count() (cnt int64) {
	result, err := this.supper.ExecRawQuery(korm.BuildQueryStringCountExist(korm.BuildQueryStringCountExist_Req{
		MainTableName:       "Work_D",
		MainTableNameAlias:  this.joinNode.TableName,
		RootInfoBufLeftJoin: &this.joinNode.Root.BufLeftJoin,
		BufWhere:            &this.bufWhere,
		IsExist:             false,
	}), this.argsWhere...)
	if err != nil {
		panic(err)
	}
	cnt, err = korm.ScanCount(result)
	if err != nil {
		panic(err)
	}
	return cnt
}

func (this *KORM_Work_D_SelectObj) MustRun_Exist() (exist bool) {
	result, err := this.supper.ExecRawQuery(korm.BuildQueryStringCountExist(korm.BuildQueryStringCountExist_Req{
		MainTableName:       "Work_D",
		MainTableNameAlias:  this.joinNode.TableName,
		RootInfoBufLeftJoin: &this.joinNode.Root.BufLeftJoin,
		BufWhere:            &this.bufWhere,
		IsExist:             true,
	}), this.argsWhere...)
	if err != nil {
		panic(err)
	}
	exist, err = korm.ScanExist(result)
	if err != nil {
		panic(err)
	}
	return exist
}
func (this *KORM_Work_D_SelectObj) MustRun_ResultOne() (info Work_D) {
	info, _ = this.MustRun_ResultOne2()
	return info
}

func (this *KORM_Work_D_SelectObj) MustRun_ResultOne2() (info Work_D, ok bool) {
	this.limit = 1
	korm_fillSelectFieldNameList_Work_D(this.joinNode)

	var buf2 bytes.Buffer
	buf2.WriteString("SELECT ")
	this.joinNode.FillSelect(&buf2, true)

	buf2.WriteString(" FROM `Work_D` " + this.joinNode.TableName + " ")
	buf2.WriteString(this.rootInfo.BufLeftJoin.String())
	buf2.WriteString(this.bufWhere.String())
	if len(this.orderBy) > 0 {
		buf2.WriteString("ORDER BY " + strings.Join(this.orderBy, ",") + " ")
	}
	if this.limit != 0 {
		buf2.WriteString("LIMIT " + strconv.Itoa(this.limit) + " ")
	}
	if this.offset != 0 {
		buf2.WriteString("OFFSET " + strconv.Itoa(this.offset) + " ")
	}

	var result *sql.Rows
	var err error

	result, err = this.supper.ExecRawQuery(buf2.String(), this.argsWhere...)
	if err != nil {
		panic(err)
	}
	defer result.Close()
	if result.Next() == false {
		return info, false
	}
	resp := korm_Work_D_scan(this.joinNode, &info)
	err = result.Scan(resp.argList...)
	if err != nil {
		panic(err)
	}
	for idx, a := range resp.argList {
		v := a.(*sql.NullString).String
		if v == "" {
			continue
		}
		err = resp.argParseFn[idx](v)
		if err != nil {
			panic(err)
		}
	}

	return info, true
}

func (this *KORM_Work_D_SelectObj) MustRun_ResultList() (list []Work_D) {
	korm_fillSelectFieldNameList_Work_D(this.joinNode)

	var buf2 bytes.Buffer
	buf2.WriteString("SELECT ")
	this.joinNode.FillSelect(&buf2, true)

	buf2.WriteString(" FROM `Work_D` " + this.joinNode.TableName + " ")
	buf2.WriteString(this.rootInfo.BufLeftJoin.String())
	buf2.WriteString(this.bufWhere.String())
	if len(this.orderBy) > 0 {
		buf2.WriteString("ORDER BY " + strings.Join(this.orderBy, ",") + " ")
	}
	if this.limit != 0 {
		buf2.WriteString("LIMIT " + strconv.Itoa(this.limit) + " ")
	}
	if this.offset != 0 {
		buf2.WriteString("OFFSET " + strconv.Itoa(this.offset) + " ")
	}

	var result *sql.Rows
	var err error

	result, err = this.supper.ExecRawQuery(buf2.String(), this.argsWhere...)
	if err != nil {
		panic(err)
	}
	defer result.Close()

	for result.Next() {
		var info Work_D
		korm_fillSelectFieldNameList_Work_D(this.joinNode)
		resp := korm_Work_D_scan(this.joinNode, &info)
		err = result.Scan(resp.argList...)
		if err != nil {
			panic(err)
		}
		for idx, a := range resp.argList {
			v := a.(*sql.NullString).String
			if v == "" {
				continue
			}
			err = resp.argParseFn[idx](v)
			if err != nil {
				panic(err)
			}
		}

		list = append(list, info)
	}
	return list
}
func (this *KORM_Work_D_SelectObj) MustRun_ResultMap() (m map[string]Work_D) {
	korm_fillSelectFieldNameList_Work_D(this.joinNode)

	var buf2 bytes.Buffer
	buf2.WriteString("SELECT ")
	this.joinNode.FillSelect(&buf2, true)

	buf2.WriteString(" FROM `Work_D` " + this.joinNode.TableName + " ")
	buf2.WriteString(this.rootInfo.BufLeftJoin.String())
	buf2.WriteString(this.bufWhere.String())
	if len(this.orderBy) > 0 {
		buf2.WriteString("ORDER BY " + strings.Join(this.orderBy, ",") + " ")
	}
	if this.limit != 0 {
		buf2.WriteString("LIMIT " + strconv.Itoa(this.limit) + " ")
	}
	if this.offset != 0 {
		buf2.WriteString("OFFSET " + strconv.Itoa(this.offset) + " ")
	}

	var result *sql.Rows
	var err error

	result, err = this.supper.ExecRawQuery(buf2.String(), this.argsWhere...)
	if err != nil {
		panic(err)
	}
	defer result.Close()

	for result.Next() {
		var info Work_D
		korm_fillSelectFieldNameList_Work_D(this.joinNode)
		resp := korm_Work_D_scan(this.joinNode, &info)
		err = result.Scan(resp.argList...)
		if err != nil {
			panic(err)
		}
		for idx, a := range resp.argList {
			v := a.(*sql.NullString).String
			if v == "" {
				continue
			}
			err = resp.argParseFn[idx](v)
			if err != nil {
				panic(err)
			}
		}

		if m == nil {
			m = map[string]Work_D{}
		}
		m[info.ID] = info

	}
	return m
}
func (this *KORM_Work_D_SelectObj) MustRun_ResultListWithTotalMatch() (list []Work_D, totalMatch int64) {
	var err error
	korm_fillSelectFieldNameList_Work_D(this.joinNode)

	var buf2 bytes.Buffer
	buf2.WriteString("SELECT ")
	if this.supper.mode == korm.InitModeMysql {
		buf2.WriteString("SQL_CALC_FOUND_ROWS  ")
	}
	this.joinNode.FillSelect(&buf2, true)

	buf2.WriteString(" FROM `Work_D` " + this.joinNode.TableName + " ")
	buf2.WriteString(this.rootInfo.BufLeftJoin.String())
	buf2.WriteString(this.bufWhere.String())
	if len(this.orderBy) > 0 {
		buf2.WriteString("ORDER BY " + strings.Join(this.orderBy, ",") + " ")
	}
	if this.limit != 0 {
		buf2.WriteString("LIMIT " + strconv.Itoa(this.limit) + " ")
	}
	if this.offset != 0 {
		buf2.WriteString("OFFSET " + strconv.Itoa(this.offset) + " ")
	}
	var conn *sql.Conn
	var result *sql.Rows
	if this.supper.db != nil {
		var err error
		conn, err = this.supper.db.Conn(context.Background())
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		result, err = conn.QueryContext(context.Background(), buf2.String(), this.argsWhere...)
		if err != nil {
			panic(err)
		}
	} else {
		var err error
		result, err = this.supper.tx.Query(buf2.String(), this.argsWhere...)
		if err != nil {
			panic(err)
		}
	}

	defer result.Close()
	for result.Next() {
		var info Work_D
		korm_fillSelectFieldNameList_Work_D(this.joinNode)
		resp := korm_Work_D_scan(this.joinNode, &info)
		err = result.Scan(resp.argList...)
		if err != nil {
			panic(err)
		}
		for idx, a := range resp.argList {
			v := a.(*sql.NullString).String
			if v == "" {
				continue
			}
			err = resp.argParseFn[idx](v)
			if err != nil {
				panic(err)
			}
		}

		list = append(list, info)
	}
	result.Close()
	nextQuery := ""
	if this.supper.mode == korm.InitModeMysql {
		nextQuery = "select FOUND_ROWS()"
	} else if this.supper.mode == korm.InitModeSqlite {
		buf2.Reset()
		buf2.WriteString("SELECT COUNT(1) ")
		buf2.WriteString("FROM `Work_D` " + this.joinNode.TableName + " ")
		buf2.WriteString(this.rootInfo.BufLeftJoin.String())
		buf2.WriteString(this.bufWhere.String())
		nextQuery = buf2.String()
	} else {
		panic("not support")
	}
	var result2 *sql.Rows
	if conn != nil {
		result2, err = conn.QueryContext(context.Background(), nextQuery)
	} else {
		result2, err = this.supper.tx.Query(nextQuery)
	}
	if err != nil {
		panic(err)
	}
	defer result2.Close()

	if result2.Next() == false {
		panic("MustRun_ResultListWithPageInfo ")
	}
	err = result2.Scan(&totalMatch)
	if err != nil {
		panic(err)
	}

	return list, totalMatch
}

type KORM_Where_KORM_Work_D_SelectObj_ID struct {
	supper      *KORM_Work_D_SelectObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Work_D_SelectObj) Where_ID() *KORM_Where_KORM_Work_D_SelectObj_ID {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_SelectObj_ID{supper: this, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_SelectObj_ID) Equal(ID string) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`ID` ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, ID)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_ID) NotEqual(ID string) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`ID` ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, ID)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_ID) Greater(ID string) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`ID` ")

	this.supper.bufWhere.WriteString(">? ")
	this.supper.argsWhere = append(this.supper.argsWhere, ID)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_ID) GreaterOrEqual(ID string) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`ID` ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, ID)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_ID) Less(ID string) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`ID` ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, ID)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_ID) LessOrEqual(ID string) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`ID` ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, ID)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_ID) In(vList []string) *KORM_Work_D_SelectObj {
	if len(vList) == 0 {
		this.supper.bufWhere.WriteString("= '' AND 0 ") // 什么都不存在, 直接返回
		return this.supper
	}
	this.supper.bufWhere.WriteString("IN (")
	for idx, v := range vList {
		if idx > 0 {
			this.supper.bufWhere.WriteString(", ")
		}
		this.supper.bufWhere.WriteString("?")
		this.supper.argsWhere = append(this.supper.argsWhere, v)
	}
	this.supper.bufWhere.WriteString(") ")
	return this.supper
}

type KORM_Where_KORM_Work_D_SelectObj_ID_Length struct {
	supper      *KORM_Work_D_SelectObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Where_KORM_Work_D_SelectObj_ID) Length() *KORM_Where_KORM_Work_D_SelectObj_ID_Length {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_SelectObj_ID_Length{supper: this.supper, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_SelectObj_ID_Length) Equal(length int) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + this.supper.joinNode.TableName + `.` + "`ID`) ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_ID_Length) NotEqual(length int) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + this.supper.joinNode.TableName + `.` + "`ID`) ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_ID_Length) GreaterOrEqual(length int) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + this.supper.joinNode.TableName + `.` + "`ID`) ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_ID_Length) Less(length int) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + this.supper.joinNode.TableName + `.` + "`ID`) ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_ID_Length) LessOrEqual(length int) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + this.supper.joinNode.TableName + `.` + "`ID`) ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}

type KORM_Where_KORM_Work_D_SelectObj_Name struct {
	supper      *KORM_Work_D_SelectObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Work_D_SelectObj) Where_Name() *KORM_Where_KORM_Work_D_SelectObj_Name {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_SelectObj_Name{supper: this, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Name) Equal(Name string) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`Name` ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Name)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Name) NotEqual(Name string) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`Name` ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Name)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Name) Greater(Name string) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`Name` ")

	this.supper.bufWhere.WriteString(">? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Name)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Name) GreaterOrEqual(Name string) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`Name` ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Name)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Name) Less(Name string) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`Name` ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Name)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Name) LessOrEqual(Name string) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`Name` ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Name)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Name) In(vList []string) *KORM_Work_D_SelectObj {
	if len(vList) == 0 {
		this.supper.bufWhere.WriteString("= '' AND 0 ") // 什么都不存在, 直接返回
		return this.supper
	}
	this.supper.bufWhere.WriteString("IN (")
	for idx, v := range vList {
		if idx > 0 {
			this.supper.bufWhere.WriteString(", ")
		}
		this.supper.bufWhere.WriteString("?")
		this.supper.argsWhere = append(this.supper.argsWhere, v)
	}
	this.supper.bufWhere.WriteString(") ")
	return this.supper
}

type KORM_Where_KORM_Work_D_SelectObj_Name_Length struct {
	supper      *KORM_Work_D_SelectObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Where_KORM_Work_D_SelectObj_Name) Length() *KORM_Where_KORM_Work_D_SelectObj_Name_Length {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_SelectObj_Name_Length{supper: this.supper, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Name_Length) Equal(length int) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + this.supper.joinNode.TableName + `.` + "`Name`) ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Name_Length) NotEqual(length int) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + this.supper.joinNode.TableName + `.` + "`Name`) ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Name_Length) GreaterOrEqual(length int) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + this.supper.joinNode.TableName + `.` + "`Name`) ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Name_Length) Less(length int) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + this.supper.joinNode.TableName + `.` + "`Name`) ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Name_Length) LessOrEqual(length int) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + this.supper.joinNode.TableName + `.` + "`Name`) ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}

type KORM_Where_KORM_Work_D_SelectObj_Url struct {
	supper      *KORM_Work_D_SelectObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Work_D_SelectObj) Where_Url() *KORM_Where_KORM_Work_D_SelectObj_Url {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_SelectObj_Url{supper: this, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Url) Equal(Url string) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`Url` ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Url)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Url) NotEqual(Url string) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`Url` ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Url)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Url) Greater(Url string) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`Url` ")

	this.supper.bufWhere.WriteString(">? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Url)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Url) GreaterOrEqual(Url string) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`Url` ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Url)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Url) Less(Url string) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`Url` ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Url)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Url) LessOrEqual(Url string) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`Url` ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Url)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Url) In(vList []string) *KORM_Work_D_SelectObj {
	if len(vList) == 0 {
		this.supper.bufWhere.WriteString("= '' AND 0 ") // 什么都不存在, 直接返回
		return this.supper
	}
	this.supper.bufWhere.WriteString("IN (")
	for idx, v := range vList {
		if idx > 0 {
			this.supper.bufWhere.WriteString(", ")
		}
		this.supper.bufWhere.WriteString("?")
		this.supper.argsWhere = append(this.supper.argsWhere, v)
	}
	this.supper.bufWhere.WriteString(") ")
	return this.supper
}

type KORM_Where_KORM_Work_D_SelectObj_Url_Length struct {
	supper      *KORM_Work_D_SelectObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Where_KORM_Work_D_SelectObj_Url) Length() *KORM_Where_KORM_Work_D_SelectObj_Url_Length {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_SelectObj_Url_Length{supper: this.supper, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Url_Length) Equal(length int) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + this.supper.joinNode.TableName + `.` + "`Url`) ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Url_Length) NotEqual(length int) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + this.supper.joinNode.TableName + `.` + "`Url`) ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Url_Length) GreaterOrEqual(length int) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + this.supper.joinNode.TableName + `.` + "`Url`) ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Url_Length) Less(length int) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + this.supper.joinNode.TableName + `.` + "`Url`) ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Url_Length) LessOrEqual(length int) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + this.supper.joinNode.TableName + `.` + "`Url`) ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}

type KORM_Where_KORM_Work_D_SelectObj_SaveDir struct {
	supper      *KORM_Work_D_SelectObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Work_D_SelectObj) Where_SaveDir() *KORM_Where_KORM_Work_D_SelectObj_SaveDir {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_SelectObj_SaveDir{supper: this, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_SelectObj_SaveDir) Equal(SaveDir string) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`SaveDir` ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, SaveDir)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_SaveDir) NotEqual(SaveDir string) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`SaveDir` ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, SaveDir)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_SaveDir) Greater(SaveDir string) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`SaveDir` ")

	this.supper.bufWhere.WriteString(">? ")
	this.supper.argsWhere = append(this.supper.argsWhere, SaveDir)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_SaveDir) GreaterOrEqual(SaveDir string) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`SaveDir` ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, SaveDir)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_SaveDir) Less(SaveDir string) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`SaveDir` ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, SaveDir)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_SaveDir) LessOrEqual(SaveDir string) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`SaveDir` ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, SaveDir)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_SaveDir) In(vList []string) *KORM_Work_D_SelectObj {
	if len(vList) == 0 {
		this.supper.bufWhere.WriteString("= '' AND 0 ") // 什么都不存在, 直接返回
		return this.supper
	}
	this.supper.bufWhere.WriteString("IN (")
	for idx, v := range vList {
		if idx > 0 {
			this.supper.bufWhere.WriteString(", ")
		}
		this.supper.bufWhere.WriteString("?")
		this.supper.argsWhere = append(this.supper.argsWhere, v)
	}
	this.supper.bufWhere.WriteString(") ")
	return this.supper
}

type KORM_Where_KORM_Work_D_SelectObj_SaveDir_Length struct {
	supper      *KORM_Work_D_SelectObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Where_KORM_Work_D_SelectObj_SaveDir) Length() *KORM_Where_KORM_Work_D_SelectObj_SaveDir_Length {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_SelectObj_SaveDir_Length{supper: this.supper, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_SelectObj_SaveDir_Length) Equal(length int) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + this.supper.joinNode.TableName + `.` + "`SaveDir`) ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_SaveDir_Length) NotEqual(length int) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + this.supper.joinNode.TableName + `.` + "`SaveDir`) ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_SaveDir_Length) GreaterOrEqual(length int) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + this.supper.joinNode.TableName + `.` + "`SaveDir`) ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_SaveDir_Length) Less(length int) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + this.supper.joinNode.TableName + `.` + "`SaveDir`) ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_SaveDir_Length) LessOrEqual(length int) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + this.supper.joinNode.TableName + `.` + "`SaveDir`) ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}

type KORM_Where_KORM_Work_D_SelectObj_State struct {
	supper      *KORM_Work_D_SelectObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Work_D_SelectObj) Where_State() *KORM_Where_KORM_Work_D_SelectObj_State {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_SelectObj_State{supper: this, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_SelectObj_State) Equal(State int) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`State` ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, State)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_State) NotEqual(State int) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`State` ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, State)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_State) Greater(State int) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`State` ")

	this.supper.bufWhere.WriteString(">? ")
	this.supper.argsWhere = append(this.supper.argsWhere, State)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_State) GreaterOrEqual(State int) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`State` ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, State)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_State) Less(State int) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`State` ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, State)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_State) LessOrEqual(State int) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`State` ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, State)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_State) In(vList []int) *KORM_Work_D_SelectObj {
	if len(vList) == 0 {
		this.supper.bufWhere.WriteString("= '' AND 0 ") // 什么都不存在, 直接返回
		return this.supper
	}
	this.supper.bufWhere.WriteString("IN (")
	for idx, v := range vList {
		if idx > 0 {
			this.supper.bufWhere.WriteString(", ")
		}
		this.supper.bufWhere.WriteString("?")
		this.supper.argsWhere = append(this.supper.argsWhere, v)
	}
	this.supper.bufWhere.WriteString(") ")
	return this.supper
}

type KORM_Where_KORM_Work_D_SelectObj_Info struct {
	supper      *KORM_Work_D_SelectObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Work_D_SelectObj) Where_Info() *KORM_Where_KORM_Work_D_SelectObj_Info {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_SelectObj_Info{supper: this, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Info) Equal(Info string) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`Info` ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Info)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Info) NotEqual(Info string) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`Info` ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Info)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Info) Greater(Info string) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`Info` ")

	this.supper.bufWhere.WriteString(">? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Info)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Info) GreaterOrEqual(Info string) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`Info` ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Info)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Info) Less(Info string) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`Info` ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Info)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Info) LessOrEqual(Info string) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`Info` ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Info)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Info) In(vList []string) *KORM_Work_D_SelectObj {
	if len(vList) == 0 {
		this.supper.bufWhere.WriteString("= '' AND 0 ") // 什么都不存在, 直接返回
		return this.supper
	}
	this.supper.bufWhere.WriteString("IN (")
	for idx, v := range vList {
		if idx > 0 {
			this.supper.bufWhere.WriteString(", ")
		}
		this.supper.bufWhere.WriteString("?")
		this.supper.argsWhere = append(this.supper.argsWhere, v)
	}
	this.supper.bufWhere.WriteString(") ")
	return this.supper
}

type KORM_Where_KORM_Work_D_SelectObj_Info_Length struct {
	supper      *KORM_Work_D_SelectObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Where_KORM_Work_D_SelectObj_Info) Length() *KORM_Where_KORM_Work_D_SelectObj_Info_Length {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_SelectObj_Info_Length{supper: this.supper, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Info_Length) Equal(length int) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + this.supper.joinNode.TableName + `.` + "`Info`) ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Info_Length) NotEqual(length int) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + this.supper.joinNode.TableName + `.` + "`Info`) ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Info_Length) GreaterOrEqual(length int) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + this.supper.joinNode.TableName + `.` + "`Info`) ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Info_Length) Less(length int) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + this.supper.joinNode.TableName + `.` + "`Info`) ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_Info_Length) LessOrEqual(length int) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + this.supper.joinNode.TableName + `.` + "`Info`) ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}

type KORM_Where_KORM_Work_D_SelectObj_CreateTime struct {
	supper      *KORM_Work_D_SelectObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Work_D_SelectObj) Where_CreateTime() *KORM_Where_KORM_Work_D_SelectObj_CreateTime {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_SelectObj_CreateTime{supper: this, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_SelectObj_CreateTime) Equal(CreateTime time.Time) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`CreateTime` ")

	vCreateTime := CreateTime.UTC().Format(time.RFC3339Nano)
	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, vCreateTime)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_CreateTime) NotEqual(CreateTime time.Time) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`CreateTime` ")

	vCreateTime := CreateTime.UTC().Format(time.RFC3339Nano)
	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, vCreateTime)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_CreateTime) GreaterOrEqual(CreateTime time.Time) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`CreateTime` ")

	vCreateTime := CreateTime.UTC().Format(time.RFC3339Nano)
	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, vCreateTime)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_CreateTime) Less(CreateTime time.Time) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`CreateTime` ")

	vCreateTime := CreateTime.UTC().Format(time.RFC3339Nano)
	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, vCreateTime)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_CreateTime) LessOrEqual(CreateTime time.Time) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`CreateTime` ")

	vCreateTime := CreateTime.UTC().Format(time.RFC3339Nano)
	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, vCreateTime)
	return this.supper
}

type KORM_Where_KORM_Work_D_SelectObj_UpdateTime struct {
	supper      *KORM_Work_D_SelectObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Work_D_SelectObj) Where_UpdateTime() *KORM_Where_KORM_Work_D_SelectObj_UpdateTime {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_SelectObj_UpdateTime{supper: this, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_SelectObj_UpdateTime) Equal(UpdateTime time.Time) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`UpdateTime` ")

	vUpdateTime := UpdateTime.UTC().Format(time.RFC3339Nano)
	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, vUpdateTime)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_UpdateTime) NotEqual(UpdateTime time.Time) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`UpdateTime` ")

	vUpdateTime := UpdateTime.UTC().Format(time.RFC3339Nano)
	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, vUpdateTime)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_UpdateTime) GreaterOrEqual(UpdateTime time.Time) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`UpdateTime` ")

	vUpdateTime := UpdateTime.UTC().Format(time.RFC3339Nano)
	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, vUpdateTime)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_UpdateTime) Less(UpdateTime time.Time) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`UpdateTime` ")

	vUpdateTime := UpdateTime.UTC().Format(time.RFC3339Nano)
	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, vUpdateTime)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_SelectObj_UpdateTime) LessOrEqual(UpdateTime time.Time) *KORM_Work_D_SelectObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString(this.supper.joinNode.TableName + `.` + "`UpdateTime` ")

	vUpdateTime := UpdateTime.UTC().Format(time.RFC3339Nano)
	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, vUpdateTime)
	return this.supper
}
func (this *KORM_Work_D_SelectObj) CondMultOpBegin_AND() *KORM_Work_D_SelectObj {
	if this.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.bufWhere.WriteString("WHERE ")
	}
	this.bufWhere.WriteString("( ")
	this.linkOpList = append([]string{"AND"}, this.linkOpList...)
	this.isLinkBegin = true
	return this
}
func (this *KORM_Work_D_SelectObj) CondMultOpBegin_OR() *KORM_Work_D_SelectObj {
	if this.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.bufWhere.WriteString("WHERE ")
	}
	this.bufWhere.WriteString("( ")
	this.linkOpList = append([]string{"OR"}, this.linkOpList...)
	this.isLinkBegin = true
	return this
}
func (this *KORM_Work_D_SelectObj) CondMultOpEnd() *KORM_Work_D_SelectObj {
	if this.isLinkBegin {
		panic("() is not allowed in sql statement") // bad sql: SELECT * FROM u where ()
	}
	this.linkOpList = this.linkOpList[1:]
	this.bufWhere.WriteString(") ")
	return this
}

// Update Work_D
type KORM_Work_D_UpdateObj struct {
	supper      *OrmAll
	bufWhere    bytes.Buffer
	argsWhere   []interface{}
	bufSet      bytes.Buffer
	argsSet     []interface{}
	linkOpList  []string
	isLinkBegin bool
}

func (this *KORM_Work_D) Update() *KORM_Work_D_UpdateObj {
	return &KORM_Work_D_UpdateObj{supper: this.supper, linkOpList: []string{"AND"}, isLinkBegin: true}
}
func (this *KORM_Work_D) MustUpdateBy_ID(info Work_D) (rowsAffected int64) {
	rowsAffected = this.Update().Where_ID().Equal(info.ID).Set_Name(info.Name).Set_Url(info.Url).Set_SaveDir(info.SaveDir).Set_State(info.State).Set_Info(info.Info).Set_CreateTime(info.CreateTime).Set_UpdateTime(info.UpdateTime).MustRun()
	return rowsAffected
}
func (this *KORM_Work_D_UpdateObj) Set_Name(Name string) *KORM_Work_D_UpdateObj {
	if len(this.argsSet) > 0 {
		this.bufSet.WriteString(", ")
	} else {
		this.bufSet.WriteString("SET ")
	}
	this.bufSet.WriteString("`Name` = ? ")
	this.argsSet = append(this.argsSet, Name)
	return this
}
func (this *KORM_Work_D_UpdateObj) Set_Url(Url string) *KORM_Work_D_UpdateObj {
	if len(this.argsSet) > 0 {
		this.bufSet.WriteString(", ")
	} else {
		this.bufSet.WriteString("SET ")
	}
	this.bufSet.WriteString("`Url` = ? ")
	this.argsSet = append(this.argsSet, Url)
	return this
}
func (this *KORM_Work_D_UpdateObj) Set_SaveDir(SaveDir string) *KORM_Work_D_UpdateObj {
	if len(this.argsSet) > 0 {
		this.bufSet.WriteString(", ")
	} else {
		this.bufSet.WriteString("SET ")
	}
	this.bufSet.WriteString("`SaveDir` = ? ")
	this.argsSet = append(this.argsSet, SaveDir)
	return this
}
func (this *KORM_Work_D_UpdateObj) Inc_State(v int) *KORM_Work_D_UpdateObj {
	if len(this.argsSet) > 0 {
		this.bufSet.WriteString(", ")
	} else {
		this.bufSet.WriteString("SET ")
	}
	this.bufSet.WriteString("`State` = `State` + ? ")
	this.argsSet = append(this.argsSet, v)
	return this
}
func (this *KORM_Work_D_UpdateObj) Set_State(State int) *KORM_Work_D_UpdateObj {
	if len(this.argsSet) > 0 {
		this.bufSet.WriteString(", ")
	} else {
		this.bufSet.WriteString("SET ")
	}
	this.bufSet.WriteString("`State` = ? ")
	this.argsSet = append(this.argsSet, State)
	return this
}
func (this *KORM_Work_D_UpdateObj) Set_Info(Info string) *KORM_Work_D_UpdateObj {
	if len(this.argsSet) > 0 {
		this.bufSet.WriteString(", ")
	} else {
		this.bufSet.WriteString("SET ")
	}
	this.bufSet.WriteString("`Info` = ? ")
	this.argsSet = append(this.argsSet, Info)
	return this
}
func (this *KORM_Work_D_UpdateObj) Set_CreateTime(CreateTime time.Time) *KORM_Work_D_UpdateObj {
	if len(this.argsSet) > 0 {
		this.bufSet.WriteString(", ")
	} else {
		this.bufSet.WriteString("SET ")
	}
	this.bufSet.WriteString("`CreateTime` = ? ")
	this.argsSet = append(this.argsSet, CreateTime.UTC().Format(time.RFC3339Nano))
	return this
}
func (this *KORM_Work_D_UpdateObj) Set_UpdateTime(UpdateTime time.Time) *KORM_Work_D_UpdateObj {
	if len(this.argsSet) > 0 {
		this.bufSet.WriteString(", ")
	} else {
		this.bufSet.WriteString("SET ")
	}
	this.bufSet.WriteString("`UpdateTime` = ? ")
	this.argsSet = append(this.argsSet, UpdateTime.UTC().Format(time.RFC3339Nano))
	return this
}
func (this *KORM_Work_D_UpdateObj) MustRun() (RowsAffected int64) {
	if len(this.argsSet) == 0 {
		panic("len(this.argsSet) == 0")
	}
	var buf2 bytes.Buffer
	buf2.WriteString("UPDATE `Work_D` ")
	buf2.WriteString(this.bufSet.String())
	buf2.WriteString(this.bufWhere.String())
	var result sql.Result
	var err error

	result, err = this.supper.ExecRaw(buf2.String(), append(this.argsSet, this.argsWhere...)...)
	if err != nil {
		panic(err)
	}
	RowsAffected, err = result.RowsAffected()
	if err != nil {
		panic(err)
	}
	return RowsAffected
}

type KORM_Where_KORM_Work_D_UpdateObj_ID struct {
	supper      *KORM_Work_D_UpdateObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Work_D_UpdateObj) Where_ID() *KORM_Where_KORM_Work_D_UpdateObj_ID {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_UpdateObj_ID{supper: this, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_ID) Equal(ID string) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`ID` ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, ID)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_ID) NotEqual(ID string) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`ID` ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, ID)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_ID) Greater(ID string) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`ID` ")

	this.supper.bufWhere.WriteString(">? ")
	this.supper.argsWhere = append(this.supper.argsWhere, ID)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_ID) GreaterOrEqual(ID string) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`ID` ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, ID)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_ID) Less(ID string) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`ID` ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, ID)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_ID) LessOrEqual(ID string) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`ID` ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, ID)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_ID) In(vList []string) *KORM_Work_D_UpdateObj {
	if len(vList) == 0 {
		this.supper.bufWhere.WriteString("= '' AND 0 ") // 什么都不存在, 直接返回
		return this.supper
	}
	this.supper.bufWhere.WriteString("IN (")
	for idx, v := range vList {
		if idx > 0 {
			this.supper.bufWhere.WriteString(", ")
		}
		this.supper.bufWhere.WriteString("?")
		this.supper.argsWhere = append(this.supper.argsWhere, v)
	}
	this.supper.bufWhere.WriteString(") ")
	return this.supper
}

type KORM_Where_KORM_Work_D_UpdateObj_ID_Length struct {
	supper      *KORM_Work_D_UpdateObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Where_KORM_Work_D_UpdateObj_ID) Length() *KORM_Where_KORM_Work_D_UpdateObj_ID_Length {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_UpdateObj_ID_Length{supper: this.supper, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_ID_Length) Equal(length int) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`ID`) ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_ID_Length) NotEqual(length int) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`ID`) ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_ID_Length) GreaterOrEqual(length int) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`ID`) ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_ID_Length) Less(length int) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`ID`) ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_ID_Length) LessOrEqual(length int) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`ID`) ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}

type KORM_Where_KORM_Work_D_UpdateObj_Name struct {
	supper      *KORM_Work_D_UpdateObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Work_D_UpdateObj) Where_Name() *KORM_Where_KORM_Work_D_UpdateObj_Name {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_UpdateObj_Name{supper: this, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Name) Equal(Name string) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Name` ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Name)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Name) NotEqual(Name string) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Name` ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Name)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Name) Greater(Name string) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Name` ")

	this.supper.bufWhere.WriteString(">? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Name)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Name) GreaterOrEqual(Name string) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Name` ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Name)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Name) Less(Name string) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Name` ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Name)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Name) LessOrEqual(Name string) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Name` ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Name)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Name) In(vList []string) *KORM_Work_D_UpdateObj {
	if len(vList) == 0 {
		this.supper.bufWhere.WriteString("= '' AND 0 ") // 什么都不存在, 直接返回
		return this.supper
	}
	this.supper.bufWhere.WriteString("IN (")
	for idx, v := range vList {
		if idx > 0 {
			this.supper.bufWhere.WriteString(", ")
		}
		this.supper.bufWhere.WriteString("?")
		this.supper.argsWhere = append(this.supper.argsWhere, v)
	}
	this.supper.bufWhere.WriteString(") ")
	return this.supper
}

type KORM_Where_KORM_Work_D_UpdateObj_Name_Length struct {
	supper      *KORM_Work_D_UpdateObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Where_KORM_Work_D_UpdateObj_Name) Length() *KORM_Where_KORM_Work_D_UpdateObj_Name_Length {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_UpdateObj_Name_Length{supper: this.supper, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Name_Length) Equal(length int) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`Name`) ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Name_Length) NotEqual(length int) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`Name`) ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Name_Length) GreaterOrEqual(length int) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`Name`) ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Name_Length) Less(length int) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`Name`) ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Name_Length) LessOrEqual(length int) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`Name`) ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}

type KORM_Where_KORM_Work_D_UpdateObj_Url struct {
	supper      *KORM_Work_D_UpdateObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Work_D_UpdateObj) Where_Url() *KORM_Where_KORM_Work_D_UpdateObj_Url {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_UpdateObj_Url{supper: this, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Url) Equal(Url string) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Url` ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Url)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Url) NotEqual(Url string) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Url` ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Url)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Url) Greater(Url string) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Url` ")

	this.supper.bufWhere.WriteString(">? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Url)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Url) GreaterOrEqual(Url string) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Url` ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Url)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Url) Less(Url string) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Url` ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Url)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Url) LessOrEqual(Url string) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Url` ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Url)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Url) In(vList []string) *KORM_Work_D_UpdateObj {
	if len(vList) == 0 {
		this.supper.bufWhere.WriteString("= '' AND 0 ") // 什么都不存在, 直接返回
		return this.supper
	}
	this.supper.bufWhere.WriteString("IN (")
	for idx, v := range vList {
		if idx > 0 {
			this.supper.bufWhere.WriteString(", ")
		}
		this.supper.bufWhere.WriteString("?")
		this.supper.argsWhere = append(this.supper.argsWhere, v)
	}
	this.supper.bufWhere.WriteString(") ")
	return this.supper
}

type KORM_Where_KORM_Work_D_UpdateObj_Url_Length struct {
	supper      *KORM_Work_D_UpdateObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Where_KORM_Work_D_UpdateObj_Url) Length() *KORM_Where_KORM_Work_D_UpdateObj_Url_Length {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_UpdateObj_Url_Length{supper: this.supper, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Url_Length) Equal(length int) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`Url`) ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Url_Length) NotEqual(length int) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`Url`) ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Url_Length) GreaterOrEqual(length int) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`Url`) ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Url_Length) Less(length int) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`Url`) ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Url_Length) LessOrEqual(length int) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`Url`) ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}

type KORM_Where_KORM_Work_D_UpdateObj_SaveDir struct {
	supper      *KORM_Work_D_UpdateObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Work_D_UpdateObj) Where_SaveDir() *KORM_Where_KORM_Work_D_UpdateObj_SaveDir {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_UpdateObj_SaveDir{supper: this, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_SaveDir) Equal(SaveDir string) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`SaveDir` ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, SaveDir)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_SaveDir) NotEqual(SaveDir string) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`SaveDir` ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, SaveDir)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_SaveDir) Greater(SaveDir string) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`SaveDir` ")

	this.supper.bufWhere.WriteString(">? ")
	this.supper.argsWhere = append(this.supper.argsWhere, SaveDir)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_SaveDir) GreaterOrEqual(SaveDir string) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`SaveDir` ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, SaveDir)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_SaveDir) Less(SaveDir string) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`SaveDir` ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, SaveDir)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_SaveDir) LessOrEqual(SaveDir string) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`SaveDir` ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, SaveDir)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_SaveDir) In(vList []string) *KORM_Work_D_UpdateObj {
	if len(vList) == 0 {
		this.supper.bufWhere.WriteString("= '' AND 0 ") // 什么都不存在, 直接返回
		return this.supper
	}
	this.supper.bufWhere.WriteString("IN (")
	for idx, v := range vList {
		if idx > 0 {
			this.supper.bufWhere.WriteString(", ")
		}
		this.supper.bufWhere.WriteString("?")
		this.supper.argsWhere = append(this.supper.argsWhere, v)
	}
	this.supper.bufWhere.WriteString(") ")
	return this.supper
}

type KORM_Where_KORM_Work_D_UpdateObj_SaveDir_Length struct {
	supper      *KORM_Work_D_UpdateObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Where_KORM_Work_D_UpdateObj_SaveDir) Length() *KORM_Where_KORM_Work_D_UpdateObj_SaveDir_Length {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_UpdateObj_SaveDir_Length{supper: this.supper, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_SaveDir_Length) Equal(length int) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`SaveDir`) ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_SaveDir_Length) NotEqual(length int) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`SaveDir`) ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_SaveDir_Length) GreaterOrEqual(length int) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`SaveDir`) ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_SaveDir_Length) Less(length int) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`SaveDir`) ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_SaveDir_Length) LessOrEqual(length int) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`SaveDir`) ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}

type KORM_Where_KORM_Work_D_UpdateObj_State struct {
	supper      *KORM_Work_D_UpdateObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Work_D_UpdateObj) Where_State() *KORM_Where_KORM_Work_D_UpdateObj_State {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_UpdateObj_State{supper: this, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_State) Equal(State int) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`State` ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, State)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_State) NotEqual(State int) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`State` ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, State)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_State) Greater(State int) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`State` ")

	this.supper.bufWhere.WriteString(">? ")
	this.supper.argsWhere = append(this.supper.argsWhere, State)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_State) GreaterOrEqual(State int) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`State` ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, State)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_State) Less(State int) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`State` ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, State)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_State) LessOrEqual(State int) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`State` ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, State)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_State) In(vList []int) *KORM_Work_D_UpdateObj {
	if len(vList) == 0 {
		this.supper.bufWhere.WriteString("= '' AND 0 ") // 什么都不存在, 直接返回
		return this.supper
	}
	this.supper.bufWhere.WriteString("IN (")
	for idx, v := range vList {
		if idx > 0 {
			this.supper.bufWhere.WriteString(", ")
		}
		this.supper.bufWhere.WriteString("?")
		this.supper.argsWhere = append(this.supper.argsWhere, v)
	}
	this.supper.bufWhere.WriteString(") ")
	return this.supper
}

type KORM_Where_KORM_Work_D_UpdateObj_Info struct {
	supper      *KORM_Work_D_UpdateObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Work_D_UpdateObj) Where_Info() *KORM_Where_KORM_Work_D_UpdateObj_Info {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_UpdateObj_Info{supper: this, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Info) Equal(Info string) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Info` ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Info)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Info) NotEqual(Info string) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Info` ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Info)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Info) Greater(Info string) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Info` ")

	this.supper.bufWhere.WriteString(">? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Info)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Info) GreaterOrEqual(Info string) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Info` ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Info)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Info) Less(Info string) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Info` ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Info)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Info) LessOrEqual(Info string) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Info` ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Info)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Info) In(vList []string) *KORM_Work_D_UpdateObj {
	if len(vList) == 0 {
		this.supper.bufWhere.WriteString("= '' AND 0 ") // 什么都不存在, 直接返回
		return this.supper
	}
	this.supper.bufWhere.WriteString("IN (")
	for idx, v := range vList {
		if idx > 0 {
			this.supper.bufWhere.WriteString(", ")
		}
		this.supper.bufWhere.WriteString("?")
		this.supper.argsWhere = append(this.supper.argsWhere, v)
	}
	this.supper.bufWhere.WriteString(") ")
	return this.supper
}

type KORM_Where_KORM_Work_D_UpdateObj_Info_Length struct {
	supper      *KORM_Work_D_UpdateObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Where_KORM_Work_D_UpdateObj_Info) Length() *KORM_Where_KORM_Work_D_UpdateObj_Info_Length {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_UpdateObj_Info_Length{supper: this.supper, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Info_Length) Equal(length int) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`Info`) ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Info_Length) NotEqual(length int) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`Info`) ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Info_Length) GreaterOrEqual(length int) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`Info`) ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Info_Length) Less(length int) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`Info`) ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_Info_Length) LessOrEqual(length int) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`Info`) ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}

type KORM_Where_KORM_Work_D_UpdateObj_CreateTime struct {
	supper      *KORM_Work_D_UpdateObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Work_D_UpdateObj) Where_CreateTime() *KORM_Where_KORM_Work_D_UpdateObj_CreateTime {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_UpdateObj_CreateTime{supper: this, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_CreateTime) Equal(CreateTime time.Time) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`CreateTime` ")

	vCreateTime := CreateTime.UTC().Format(time.RFC3339Nano)
	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, vCreateTime)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_CreateTime) NotEqual(CreateTime time.Time) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`CreateTime` ")

	vCreateTime := CreateTime.UTC().Format(time.RFC3339Nano)
	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, vCreateTime)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_CreateTime) GreaterOrEqual(CreateTime time.Time) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`CreateTime` ")

	vCreateTime := CreateTime.UTC().Format(time.RFC3339Nano)
	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, vCreateTime)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_CreateTime) Less(CreateTime time.Time) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`CreateTime` ")

	vCreateTime := CreateTime.UTC().Format(time.RFC3339Nano)
	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, vCreateTime)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_CreateTime) LessOrEqual(CreateTime time.Time) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`CreateTime` ")

	vCreateTime := CreateTime.UTC().Format(time.RFC3339Nano)
	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, vCreateTime)
	return this.supper
}

type KORM_Where_KORM_Work_D_UpdateObj_UpdateTime struct {
	supper      *KORM_Work_D_UpdateObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Work_D_UpdateObj) Where_UpdateTime() *KORM_Where_KORM_Work_D_UpdateObj_UpdateTime {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_UpdateObj_UpdateTime{supper: this, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_UpdateTime) Equal(UpdateTime time.Time) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`UpdateTime` ")

	vUpdateTime := UpdateTime.UTC().Format(time.RFC3339Nano)
	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, vUpdateTime)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_UpdateTime) NotEqual(UpdateTime time.Time) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`UpdateTime` ")

	vUpdateTime := UpdateTime.UTC().Format(time.RFC3339Nano)
	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, vUpdateTime)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_UpdateTime) GreaterOrEqual(UpdateTime time.Time) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`UpdateTime` ")

	vUpdateTime := UpdateTime.UTC().Format(time.RFC3339Nano)
	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, vUpdateTime)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_UpdateTime) Less(UpdateTime time.Time) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`UpdateTime` ")

	vUpdateTime := UpdateTime.UTC().Format(time.RFC3339Nano)
	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, vUpdateTime)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_UpdateObj_UpdateTime) LessOrEqual(UpdateTime time.Time) *KORM_Work_D_UpdateObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`UpdateTime` ")

	vUpdateTime := UpdateTime.UTC().Format(time.RFC3339Nano)
	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, vUpdateTime)
	return this.supper
}
func (this *KORM_Work_D_UpdateObj) CondMultOpBegin_AND() *KORM_Work_D_UpdateObj {
	if this.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.bufWhere.WriteString("WHERE ")
	}
	this.bufWhere.WriteString("( ")
	this.linkOpList = append([]string{"AND"}, this.linkOpList...)
	this.isLinkBegin = true
	return this
}
func (this *KORM_Work_D_UpdateObj) CondMultOpBegin_OR() *KORM_Work_D_UpdateObj {
	if this.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.bufWhere.WriteString("WHERE ")
	}
	this.bufWhere.WriteString("( ")
	this.linkOpList = append([]string{"OR"}, this.linkOpList...)
	this.isLinkBegin = true
	return this
}
func (this *KORM_Work_D_UpdateObj) CondMultOpEnd() *KORM_Work_D_UpdateObj {
	if this.isLinkBegin {
		panic("() is not allowed in sql statement") // bad sql: SELECT * FROM u where ()
	}
	this.linkOpList = this.linkOpList[1:]
	this.bufWhere.WriteString(") ")
	return this
}

type KORM_Work_D_DeleteObj struct {
	supper      *OrmAll
	bufWhere    bytes.Buffer
	argsWhere   []interface{}
	linkOpList  []string
	isLinkBegin bool
}

func (this *KORM_Work_D) Delete() *KORM_Work_D_DeleteObj {
	return &KORM_Work_D_DeleteObj{supper: this.supper, linkOpList: []string{"AND"}, isLinkBegin: true}
}
func (this *KORM_Work_D_DeleteObj) MustRun() (RowsAffected int64) {
	var buf2 bytes.Buffer
	buf2.WriteString("DELETE FROM Work_D ")
	buf2.WriteString(this.bufWhere.String())
	var result sql.Result
	var err error
	result, err = this.supper.ExecRaw(buf2.String(), this.argsWhere...)

	if err != nil {
		panic(err)
	}
	RowsAffected, err = result.RowsAffected()
	if err != nil {
		panic(err)
	}
	return RowsAffected
}

type KORM_Where_KORM_Work_D_DeleteObj_ID struct {
	supper      *KORM_Work_D_DeleteObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Work_D_DeleteObj) Where_ID() *KORM_Where_KORM_Work_D_DeleteObj_ID {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_DeleteObj_ID{supper: this, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_ID) Equal(ID string) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`ID` ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, ID)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_ID) NotEqual(ID string) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`ID` ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, ID)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_ID) Greater(ID string) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`ID` ")

	this.supper.bufWhere.WriteString(">? ")
	this.supper.argsWhere = append(this.supper.argsWhere, ID)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_ID) GreaterOrEqual(ID string) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`ID` ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, ID)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_ID) Less(ID string) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`ID` ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, ID)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_ID) LessOrEqual(ID string) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`ID` ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, ID)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_ID) In(vList []string) *KORM_Work_D_DeleteObj {
	if len(vList) == 0 {
		this.supper.bufWhere.WriteString("= '' AND 0 ") // 什么都不存在, 直接返回
		return this.supper
	}
	this.supper.bufWhere.WriteString("IN (")
	for idx, v := range vList {
		if idx > 0 {
			this.supper.bufWhere.WriteString(", ")
		}
		this.supper.bufWhere.WriteString("?")
		this.supper.argsWhere = append(this.supper.argsWhere, v)
	}
	this.supper.bufWhere.WriteString(") ")
	return this.supper
}

type KORM_Where_KORM_Work_D_DeleteObj_ID_Length struct {
	supper      *KORM_Work_D_DeleteObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Where_KORM_Work_D_DeleteObj_ID) Length() *KORM_Where_KORM_Work_D_DeleteObj_ID_Length {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_DeleteObj_ID_Length{supper: this.supper, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_ID_Length) Equal(length int) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`ID`) ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_ID_Length) NotEqual(length int) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`ID`) ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_ID_Length) GreaterOrEqual(length int) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`ID`) ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_ID_Length) Less(length int) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`ID`) ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_ID_Length) LessOrEqual(length int) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`ID`) ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}

type KORM_Where_KORM_Work_D_DeleteObj_Name struct {
	supper      *KORM_Work_D_DeleteObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Work_D_DeleteObj) Where_Name() *KORM_Where_KORM_Work_D_DeleteObj_Name {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_DeleteObj_Name{supper: this, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Name) Equal(Name string) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Name` ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Name)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Name) NotEqual(Name string) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Name` ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Name)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Name) Greater(Name string) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Name` ")

	this.supper.bufWhere.WriteString(">? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Name)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Name) GreaterOrEqual(Name string) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Name` ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Name)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Name) Less(Name string) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Name` ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Name)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Name) LessOrEqual(Name string) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Name` ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Name)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Name) In(vList []string) *KORM_Work_D_DeleteObj {
	if len(vList) == 0 {
		this.supper.bufWhere.WriteString("= '' AND 0 ") // 什么都不存在, 直接返回
		return this.supper
	}
	this.supper.bufWhere.WriteString("IN (")
	for idx, v := range vList {
		if idx > 0 {
			this.supper.bufWhere.WriteString(", ")
		}
		this.supper.bufWhere.WriteString("?")
		this.supper.argsWhere = append(this.supper.argsWhere, v)
	}
	this.supper.bufWhere.WriteString(") ")
	return this.supper
}

type KORM_Where_KORM_Work_D_DeleteObj_Name_Length struct {
	supper      *KORM_Work_D_DeleteObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Where_KORM_Work_D_DeleteObj_Name) Length() *KORM_Where_KORM_Work_D_DeleteObj_Name_Length {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_DeleteObj_Name_Length{supper: this.supper, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Name_Length) Equal(length int) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`Name`) ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Name_Length) NotEqual(length int) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`Name`) ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Name_Length) GreaterOrEqual(length int) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`Name`) ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Name_Length) Less(length int) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`Name`) ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Name_Length) LessOrEqual(length int) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`Name`) ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}

type KORM_Where_KORM_Work_D_DeleteObj_Url struct {
	supper      *KORM_Work_D_DeleteObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Work_D_DeleteObj) Where_Url() *KORM_Where_KORM_Work_D_DeleteObj_Url {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_DeleteObj_Url{supper: this, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Url) Equal(Url string) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Url` ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Url)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Url) NotEqual(Url string) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Url` ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Url)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Url) Greater(Url string) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Url` ")

	this.supper.bufWhere.WriteString(">? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Url)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Url) GreaterOrEqual(Url string) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Url` ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Url)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Url) Less(Url string) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Url` ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Url)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Url) LessOrEqual(Url string) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Url` ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Url)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Url) In(vList []string) *KORM_Work_D_DeleteObj {
	if len(vList) == 0 {
		this.supper.bufWhere.WriteString("= '' AND 0 ") // 什么都不存在, 直接返回
		return this.supper
	}
	this.supper.bufWhere.WriteString("IN (")
	for idx, v := range vList {
		if idx > 0 {
			this.supper.bufWhere.WriteString(", ")
		}
		this.supper.bufWhere.WriteString("?")
		this.supper.argsWhere = append(this.supper.argsWhere, v)
	}
	this.supper.bufWhere.WriteString(") ")
	return this.supper
}

type KORM_Where_KORM_Work_D_DeleteObj_Url_Length struct {
	supper      *KORM_Work_D_DeleteObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Where_KORM_Work_D_DeleteObj_Url) Length() *KORM_Where_KORM_Work_D_DeleteObj_Url_Length {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_DeleteObj_Url_Length{supper: this.supper, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Url_Length) Equal(length int) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`Url`) ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Url_Length) NotEqual(length int) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`Url`) ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Url_Length) GreaterOrEqual(length int) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`Url`) ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Url_Length) Less(length int) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`Url`) ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Url_Length) LessOrEqual(length int) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`Url`) ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}

type KORM_Where_KORM_Work_D_DeleteObj_SaveDir struct {
	supper      *KORM_Work_D_DeleteObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Work_D_DeleteObj) Where_SaveDir() *KORM_Where_KORM_Work_D_DeleteObj_SaveDir {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_DeleteObj_SaveDir{supper: this, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_SaveDir) Equal(SaveDir string) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`SaveDir` ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, SaveDir)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_SaveDir) NotEqual(SaveDir string) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`SaveDir` ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, SaveDir)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_SaveDir) Greater(SaveDir string) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`SaveDir` ")

	this.supper.bufWhere.WriteString(">? ")
	this.supper.argsWhere = append(this.supper.argsWhere, SaveDir)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_SaveDir) GreaterOrEqual(SaveDir string) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`SaveDir` ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, SaveDir)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_SaveDir) Less(SaveDir string) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`SaveDir` ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, SaveDir)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_SaveDir) LessOrEqual(SaveDir string) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`SaveDir` ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, SaveDir)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_SaveDir) In(vList []string) *KORM_Work_D_DeleteObj {
	if len(vList) == 0 {
		this.supper.bufWhere.WriteString("= '' AND 0 ") // 什么都不存在, 直接返回
		return this.supper
	}
	this.supper.bufWhere.WriteString("IN (")
	for idx, v := range vList {
		if idx > 0 {
			this.supper.bufWhere.WriteString(", ")
		}
		this.supper.bufWhere.WriteString("?")
		this.supper.argsWhere = append(this.supper.argsWhere, v)
	}
	this.supper.bufWhere.WriteString(") ")
	return this.supper
}

type KORM_Where_KORM_Work_D_DeleteObj_SaveDir_Length struct {
	supper      *KORM_Work_D_DeleteObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Where_KORM_Work_D_DeleteObj_SaveDir) Length() *KORM_Where_KORM_Work_D_DeleteObj_SaveDir_Length {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_DeleteObj_SaveDir_Length{supper: this.supper, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_SaveDir_Length) Equal(length int) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`SaveDir`) ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_SaveDir_Length) NotEqual(length int) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`SaveDir`) ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_SaveDir_Length) GreaterOrEqual(length int) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`SaveDir`) ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_SaveDir_Length) Less(length int) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`SaveDir`) ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_SaveDir_Length) LessOrEqual(length int) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`SaveDir`) ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}

type KORM_Where_KORM_Work_D_DeleteObj_State struct {
	supper      *KORM_Work_D_DeleteObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Work_D_DeleteObj) Where_State() *KORM_Where_KORM_Work_D_DeleteObj_State {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_DeleteObj_State{supper: this, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_State) Equal(State int) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`State` ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, State)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_State) NotEqual(State int) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`State` ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, State)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_State) Greater(State int) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`State` ")

	this.supper.bufWhere.WriteString(">? ")
	this.supper.argsWhere = append(this.supper.argsWhere, State)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_State) GreaterOrEqual(State int) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`State` ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, State)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_State) Less(State int) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`State` ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, State)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_State) LessOrEqual(State int) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`State` ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, State)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_State) In(vList []int) *KORM_Work_D_DeleteObj {
	if len(vList) == 0 {
		this.supper.bufWhere.WriteString("= '' AND 0 ") // 什么都不存在, 直接返回
		return this.supper
	}
	this.supper.bufWhere.WriteString("IN (")
	for idx, v := range vList {
		if idx > 0 {
			this.supper.bufWhere.WriteString(", ")
		}
		this.supper.bufWhere.WriteString("?")
		this.supper.argsWhere = append(this.supper.argsWhere, v)
	}
	this.supper.bufWhere.WriteString(") ")
	return this.supper
}

type KORM_Where_KORM_Work_D_DeleteObj_Info struct {
	supper      *KORM_Work_D_DeleteObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Work_D_DeleteObj) Where_Info() *KORM_Where_KORM_Work_D_DeleteObj_Info {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_DeleteObj_Info{supper: this, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Info) Equal(Info string) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Info` ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Info)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Info) NotEqual(Info string) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Info` ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Info)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Info) Greater(Info string) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Info` ")

	this.supper.bufWhere.WriteString(">? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Info)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Info) GreaterOrEqual(Info string) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Info` ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Info)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Info) Less(Info string) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Info` ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Info)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Info) LessOrEqual(Info string) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`Info` ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, Info)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Info) In(vList []string) *KORM_Work_D_DeleteObj {
	if len(vList) == 0 {
		this.supper.bufWhere.WriteString("= '' AND 0 ") // 什么都不存在, 直接返回
		return this.supper
	}
	this.supper.bufWhere.WriteString("IN (")
	for idx, v := range vList {
		if idx > 0 {
			this.supper.bufWhere.WriteString(", ")
		}
		this.supper.bufWhere.WriteString("?")
		this.supper.argsWhere = append(this.supper.argsWhere, v)
	}
	this.supper.bufWhere.WriteString(") ")
	return this.supper
}

type KORM_Where_KORM_Work_D_DeleteObj_Info_Length struct {
	supper      *KORM_Work_D_DeleteObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Where_KORM_Work_D_DeleteObj_Info) Length() *KORM_Where_KORM_Work_D_DeleteObj_Info_Length {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_DeleteObj_Info_Length{supper: this.supper, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Info_Length) Equal(length int) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`Info`) ")

	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Info_Length) NotEqual(length int) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`Info`) ")

	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Info_Length) GreaterOrEqual(length int) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`Info`) ")

	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Info_Length) Less(length int) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`Info`) ")

	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_Info_Length) LessOrEqual(length int) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("length(" + "`Info`) ")

	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, length)
	return this.supper
}

type KORM_Where_KORM_Work_D_DeleteObj_CreateTime struct {
	supper      *KORM_Work_D_DeleteObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Work_D_DeleteObj) Where_CreateTime() *KORM_Where_KORM_Work_D_DeleteObj_CreateTime {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_DeleteObj_CreateTime{supper: this, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_CreateTime) Equal(CreateTime time.Time) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`CreateTime` ")

	vCreateTime := CreateTime.UTC().Format(time.RFC3339Nano)
	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, vCreateTime)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_CreateTime) NotEqual(CreateTime time.Time) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`CreateTime` ")

	vCreateTime := CreateTime.UTC().Format(time.RFC3339Nano)
	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, vCreateTime)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_CreateTime) GreaterOrEqual(CreateTime time.Time) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`CreateTime` ")

	vCreateTime := CreateTime.UTC().Format(time.RFC3339Nano)
	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, vCreateTime)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_CreateTime) Less(CreateTime time.Time) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`CreateTime` ")

	vCreateTime := CreateTime.UTC().Format(time.RFC3339Nano)
	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, vCreateTime)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_CreateTime) LessOrEqual(CreateTime time.Time) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`CreateTime` ")

	vCreateTime := CreateTime.UTC().Format(time.RFC3339Nano)
	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, vCreateTime)
	return this.supper
}

type KORM_Where_KORM_Work_D_DeleteObj_UpdateTime struct {
	supper      *KORM_Work_D_DeleteObj
	isLinkBegin bool
	linkOpList  []string
}

func (this *KORM_Work_D_DeleteObj) Where_UpdateTime() *KORM_Where_KORM_Work_D_DeleteObj_UpdateTime {
	isLinkBeginValue := this.isLinkBegin
	this.isLinkBegin = false
	return &KORM_Where_KORM_Work_D_DeleteObj_UpdateTime{supper: this, isLinkBegin: isLinkBeginValue, linkOpList: this.linkOpList}
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_UpdateTime) Equal(UpdateTime time.Time) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`UpdateTime` ")

	vUpdateTime := UpdateTime.UTC().Format(time.RFC3339Nano)
	this.supper.bufWhere.WriteString("=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, vUpdateTime)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_UpdateTime) NotEqual(UpdateTime time.Time) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`UpdateTime` ")

	vUpdateTime := UpdateTime.UTC().Format(time.RFC3339Nano)
	this.supper.bufWhere.WriteString("!=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, vUpdateTime)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_UpdateTime) GreaterOrEqual(UpdateTime time.Time) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`UpdateTime` ")

	vUpdateTime := UpdateTime.UTC().Format(time.RFC3339Nano)
	this.supper.bufWhere.WriteString(">=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, vUpdateTime)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_UpdateTime) Less(UpdateTime time.Time) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`UpdateTime` ")

	vUpdateTime := UpdateTime.UTC().Format(time.RFC3339Nano)
	this.supper.bufWhere.WriteString("<? ")
	this.supper.argsWhere = append(this.supper.argsWhere, vUpdateTime)
	return this.supper
}
func (this *KORM_Where_KORM_Work_D_DeleteObj_UpdateTime) LessOrEqual(UpdateTime time.Time) *KORM_Work_D_DeleteObj {
	if this.supper.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.supper.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.supper.bufWhere.WriteString("WHERE ")
	}
	this.supper.bufWhere.WriteString("`UpdateTime` ")

	vUpdateTime := UpdateTime.UTC().Format(time.RFC3339Nano)
	this.supper.bufWhere.WriteString("<=? ")
	this.supper.argsWhere = append(this.supper.argsWhere, vUpdateTime)
	return this.supper
}
func (this *KORM_Work_D_DeleteObj) CondMultOpBegin_AND() *KORM_Work_D_DeleteObj {
	if this.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.bufWhere.WriteString("WHERE ")
	}
	this.bufWhere.WriteString("( ")
	this.linkOpList = append([]string{"AND"}, this.linkOpList...)
	this.isLinkBegin = true
	return this
}
func (this *KORM_Work_D_DeleteObj) CondMultOpBegin_OR() *KORM_Work_D_DeleteObj {
	if this.bufWhere.Len() > 0 {
		if this.isLinkBegin == false {
			this.bufWhere.WriteString(this.linkOpList[0] + " ")
		}
	} else {
		this.bufWhere.WriteString("WHERE ")
	}
	this.bufWhere.WriteString("( ")
	this.linkOpList = append([]string{"OR"}, this.linkOpList...)
	this.isLinkBegin = true
	return this
}
func (this *KORM_Work_D_DeleteObj) CondMultOpEnd() *KORM_Work_D_DeleteObj {
	if this.isLinkBegin {
		panic("() is not allowed in sql statement") // bad sql: SELECT * FROM u where ()
	}
	this.linkOpList = this.linkOpList[1:]
	this.bufWhere.WriteString(") ")
	return this
}
