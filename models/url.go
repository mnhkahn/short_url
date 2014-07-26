package models

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lunny/xorm"
	"os"
	"short_url/service"
)

var (
	DEFAULT_SQL_CONN *xorm.Engine
)

func init() {
	var err error
	DEFAULT_SQL_CONN, err = xorm.NewEngine("mysql", beego.AppConfig.String("db.url"))
	// defer engine.Close()
	DEFAULT_SQL_CONN.ShowSQL = true

	f, err := os.Create("sql.log")
	if err != nil {
		panic(err.Error())
		return
	}
	DEFAULT_SQL_CONN.Logger = f
}

type Url struct {
	LongUrl  string `json:"longurl" form:"longurl"`
	ShortUrl string `json:"shorturl" form:"shorturl"`
}

type ShortUrl struct {
	Id  int64
	Url string
}

func (u *Url) Shorten() {
	url_db := ShortUrl{}
	has, _ := DEFAULT_SQL_CONN.Where("url=?", u.LongUrl).Get(&url_db)

	if !has {
		url_db.Url = u.LongUrl
		url_db.Id, _ = DEFAULT_SQL_CONN.Insert(url_db)
	}
	u.ShortUrl = service.Encode(uint64(url_db.Id))

}

func (u *Url) Lengthen() {
	url_db := ShortUrl{}
	url_db.Id = int64(service.Decode(u.ShortUrl))

	has, _ := DEFAULT_SQL_CONN.Where("id=?", url_db.Id).Get(&url_db)
	if has {
		u.LongUrl = url_db.Url
	} else {
		u.LongUrl = beego.AppConfig.String("default.no")
	}
}
