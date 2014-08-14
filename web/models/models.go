package models

import (
	"github.com/aaudis/GoRedisSession"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

const (
	_DB_NAME      = "default"
	_MYSQL_DRIVER = "mysql"
)

var (
	Redis_session *rsess.SessionConnect
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"null"`
	Views           int64
	TopicTime       time.Time `orm:"null"`
	ToppicCout      int64
	TopiclastUserId int64
}
type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCout       int64
	ReplyLastUserId int64
	Recommend       int64
	Zhaosi          string
}

func RegisterDB() {

	orm.RegisterModel(new(Category), new(Topic), new(Customers))

	orm.RegisterDriver(_MYSQL_DRIVER, orm.DR_MySQL)
	orm.RegisterDataBase("default", _MYSQL_DRIVER, "root:root@(127.0.0.1:3306)/default?charset=utf8")
	/*SESSION初始化*/
	rsess.Prefix = "sess:" // session prefix (in Redis)
	rsess.Expire = 1800    // 30 minute session expiration

	// Connecting to Redis and creating storage instance
	temp_sess, err := rsess.New("sid", 0, "127.0.0.1", 6379)
	if err != nil {
		log.Printf("%s", err)
	}

	Redis_session = temp_sess // assing to global variable
}
