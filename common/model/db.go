package model

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"go-micro-doc/common/constant"
	"go-micro-doc/config"
)

var db *gorm.DB

type TransFunc func(tx *gorm.DB) error

/**
 * 获取数据库连接
 */
func GetDB() *gorm.DB {
	if db == nil {
		RegisterDB()
	}
	return db
}

/**
 * 注册数据库连接
 */
func RegisterDB() {
	var err error
	dbConfig := config.AppConfig.Database
	dbConfig.URL = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database, dbConfig.Charset)
	fmt.Println("register database")
	db, err = gorm.Open(dbConfig.Dialect, dbConfig.URL)
	if err != nil {
		panic(fmt.Errorf("fatal error: connect database: %s\n", err))
	}

	if config.GetEnv() == constant.EnvProd {
		db.SetLogger(logrus.StandardLogger())
	} else {
		db.LogMode(true)
	}

	db.DB().SetMaxIdleConns(dbConfig.MaxIdleConnNum)
	db.DB().SetMaxOpenConns(dbConfig.MaxOpenConnNum)
	db.BlockGlobalUpdate(true)
	db.InstantSet("gorm:save_associations", false)
	db.InstantSet("gorm:association_save_reference", false)
}

/**
 * 关闭数据库连接
 */
func CloseDB() {
	if db != nil {
		fmt.Println("close database")
		_ = db.Close()
	}
}

/**
 * 执行事物操作
 */
func Transaction(closures ...TransFunc) (err error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			switch r.(type) {
			case error:
				err = r.(error)
			case string:
				err = errors.New(r.(string))
			default:
				err = errors.New("system internal error")
			}
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	for _, closure := range closures {
		if err := closure(tx); err != nil {
			tx.Rollback()
			return err
		}
		if tx.Error != nil {
			tx.Rollback()
			return tx.Error
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
