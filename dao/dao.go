package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Dao struct {
	*gorm.DB
}

func NewDao(user, password, address, database string) *Dao {
	db, err := gorm.Open("mysql", user+":"+password+"@tcp("+address+")/"+database+"?charset=utf8")
	if err != nil {
		panic("open db error : " + err.Error())
	}
	db.LogMode(true)
	db.SingularTable(true)
	return &Dao{db}
}

func (d *Dao) createTableIfNotExist(table interface{}) *Dao {
	if !d.HasTable(table) {
		d.DB = d.DB.CreateTable(table)
	}
	return d
}

func (d *Dao) CreateTable(table interface{}) {
	d.createTableIfNotExist(table)
}

func (d *Dao) Get(query, result interface{}) {
	d.Where(query).First(result)
}

func (d *Dao) GetAll(query, result interface{}) {
	d.Where(query).Find(result)
}

func (d *Dao) GetByOrder(query, result interface{}, orderField string) {
	d.Where(query).Order(orderField).Find(result)
}

func (d *Dao) GetByField(table interface{}, field string, value interface{}) {
	d.First(table, field+"=?", value)
}

func (d *Dao) Put(value interface{}) {
	d.Create(value)
}

func (d *Dao) Update(table, value interface{}) {
	d.Model(table).Update(value)
}

func (d *Dao) Set(table interface{}, name string, value interface{}) {
	d.Where(table).Set(name, value)
}

func (d *Dao) LimitFind(query, value interface{}, limit int, order interface{}) {
	d.Limit(limit).Where(query).Order(order, true).Find(value)
}

func (d *Dao) LimitOffsetFind(query, value, limit, offset, order interface{}) {
	d.Limit(limit).Offset(offset).Where(query).Order(order, true).Find(value)
}
