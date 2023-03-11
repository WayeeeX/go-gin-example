package models

import (
	"errors"
	"fmt"
	"github.com/WayeeeX/go-gin-example/models/request"
	"github.com/WayeeeX/go-gin-example/pkg/util"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/WayeeeX/go-gin-example/pkg/setting"
	"time"
)

var DB *gorm.DB

// Setup initializes the database instance
func Setup() {
	var err error
	DB, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(DB *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	DB.LogMode(true)
	DB.SingularTable(true)
	DB.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	DB.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	DB.Callback().Delete().Replace("gorm:delete", deleteCallback)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer DB.Close()
}

// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `ModifiedOn` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

// deleteCallback will set `DeletedOn` where deleting
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

// addExtraSpaceIfExist adds a separator
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}

// 通用 CRUD
// 创建数据(可以创建[单条]数据, 也可[批量]创建)
func Create[T any](data *T) {
	err := DB.Create(&data).Error
	if err != nil {
		panic(err)
	}
}

// [单条]数据查询
func GetOne[T any](data T, query string, args ...any) T {
	err := DB.Where(query, args...).First(&data).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) { // 记录找不到 err 不 panic
		panic(err)
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		panic("无此记录")
	}

	return data
}

// [单行]更新: 传入对应结构体[传递主键用] 和 带有对应更新字段值的[结构体]，结构体不更新零值
func Update[T any](data *T, slt ...string) {
	if len(slt) > 0 {
		DB.Model(&data).Select(slt).Updates(&data)
		return
	}
	err := DB.Model(&data).Updates(&data).Error
	if err != nil {
		panic(err)
	}
}

// [批量]更新: map 的字段就是要更新的字段 (map 可以更新零值), 通过条件可以实现[单行]更新
func UpdatesMap[T any](data *T, maps map[string]any, query string, args ...any) {
	err := DB.Model(&data).Where(query, args...).Updates(maps).Error
	if err != nil {
		panic(err)
	}
}

// [批量]更新: 结构体的属性就是要更新的字段 (结构体不更新零值), 通过条件可以实现[单行]更新
func Updates[T any](data *T, query string, args ...any) {
	err := DB.Model(&data).Where(query, args...).Updates(&data).Error
	if err != nil {
		panic(err)
	}
}

// 数据列表
func List[T any](data T, req request.PageQuery, query string, args ...any) (T, int64) {
	var total int64
	DB.Model(&data).Count(&total).Where(query, args).Limit(req.PageSize).Offset(util.GetOffset(req)).
		Find(&data)
	return data, total
}

// [批量]删除数据, 通过条件控制可以删除单条数据
func Delete[T any](data T, query string, args ...any) {
	err := DB.Where(query, args...).Delete(&data).Error
	if err != nil {
		panic(err)
	}
}

// 统计数量
func Count[T any](data T, query string, args ...any) int64 {
	var total int64
	DB := DB.Model(data)
	if query != "" {
		DB = DB.Where(query, args...)
	}
	if err := DB.Count(&total).Error; err != nil {
		panic(err)
	}
	return total
}
