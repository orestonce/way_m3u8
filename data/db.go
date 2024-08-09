package data

import (
	"database/sql"
	_ "github.com/glebarez/go-sqlite"
	"github.com/orestonce/korm"
	"gom3u8/model"
	"sync"
)

var gDbInfo *model.OrmAll
var gDbInfoInit sync.Once

func GetDbInstance() *model.OrmAll {
	gDbInfoInit.Do(func() {
		const f = "./workinfo.sqlite3"
		db, err := sql.Open("sqlite", f)
		if err != nil {
			panic(err)
		}
		model.KORM_MustInitTableAll(db, korm.InitModeSqlite)
		gDbInfo = model.OrmAllNew(db, korm.InitModeSqlite)
	})
	return gDbInfo
}
