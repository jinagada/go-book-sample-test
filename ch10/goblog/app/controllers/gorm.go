package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
	"goblog/app/models"
	"golang.org/x/crypto/bcrypt"
)

var db *gorm.DB

type GormController struct {
	*revel.Controller
	Txn *gorm.DB
}

func InitDB() {
	var (
		driver, spec string
		found        bool
	)
	// revel.ERROR.Fatal 형태의 로그는 현재는 지원하지 않음
	// 대신 controller.Log or revel.AppLog 를 사용하는 것이 권장됨
	// When logging in Revel you should use the controller.Log If you have services running in the background you should use the revel.AppLog
	// URL : https://revel.github.io/manual/logging.html
	if driver, found = revel.Config.String("db.driver"); !found {
		revel.AppLog.Fatal("No db.driver found.")
	}
	if spec, found = revel.Config.String("db.spec"); !found {
		revel.AppLog.Fatal("No db.spec found.")
	}
	var err error
	db, err = gorm.Open(driver, spec)
	if err != nil {
		revel.AppLog.Fatal(err.Error())
	}
	db.LogMode(true)
	migrate()
}

const (
	DefaultName     = "Admin"
	DefaultRole     = "admin"
	DefaultUsername = "admin"
	DefaultPassword = "admin"
)

func migrate() {
	db.AutoMigrate(&models.Post{}, &models.Comment{}, &models.User{})
	bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(DefaultPassword), bcrypt.DefaultCost)
	db.Where(models.User{Name: DefaultName, Role: DefaultRole, Username: DefaultUsername}).
		Attrs(models.User{Password: string(bcryptPassword)}).
		FirstOrCreate(&models.User{})
}

func (c *GormController) Begin() revel.Result {
	c.Txn = db.Begin()
	return nil
}

func (c *GormController) Rollback() revel.Result {
	if c.Txn != nil {
		c.Txn.Rollback()
		c.Txn = nil
	}
	return nil
}

func (c *GormController) Commit() revel.Result {
	if c.Txn != nil {
		c.Txn.Commit()
		c.Txn = nil
	}
	return nil
}
