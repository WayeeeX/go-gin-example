package models

import (
	"database/sql/driver"
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

var db *gorm.DB

type Model struct {
	ID         uint64     `gorm:"primary_key" json:"id"`
	CreateTime *LocalTime `gorm:"autoCreateTime" json:"create_time"`
	UpdateTime *LocalTime `gorm:"autoUpdateTime" json:"update_time"`
}
type LocalTime time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
	zone        = "Asia/Shanghai"
)

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = LocalTime(now)
	return
}

// MarshalJSON implements json marshal interface.
func (t LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}
func (t LocalTime) String() string {
	return time.Time(t).Format(timeFormart)
}

func (t LocalTime) local() time.Time {
	loc, _ := time.LoadLocation(zone)
	return time.Time(t).In(loc)
}

// Value ...
func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	var ti = time.Time(t)
	if ti.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return ti, nil
}

// Scan valueof time.Time 注意是指针类型 method
func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// Setup initializes the database instance
func Setup() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer db.Close()
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
	err := db.Create(&data).Error
	if err != nil {
		panic(err)
	}
}

// [单条]数据查询
func GetOne[T any](data T, query string, args ...any) T {
	err := db.Where(query, args...).First(&data).Error
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
		db.Model(&data).Select(slt).Updates(&data)
		return
	}
	err := db.Model(&data).Updates(&data).Error
	if err != nil {
		panic(err)
	}
}

// [批量]更新: map 的字段就是要更新的字段 (map 可以更新零值), 通过条件可以实现[单行]更新
func UpdatesMap[T any](data *T, maps map[string]any, query string, args ...any) {
	err := db.Model(&data).Where(query, args...).Updates(maps).Error
	if err != nil {
		panic(err)
	}
}

// [批量]更新: 结构体的属性就是要更新的字段 (结构体不更新零值), 通过条件可以实现[单行]更新
func Updates[T any](data *T, query string, args ...any) {
	err := db.Model(&data).Where(query, args...).Updates(&data).Error
	if err != nil {
		panic(err)
	}
}

// 数据列表
func List[T any](data T, req request.PageQuery, query string, args ...any) (T, int64) {
	var total int64
	db.Model(&data).Count(&total).Where(query, args).Limit(req.PageSize).Offset(util.GetOffset(req)).
		Find(&data)
	return data, total
}

// [批量]删除数据, 通过条件控制可以删除单条数据
func Delete[T any](data T, query string, args ...any) {
	err := db.Where(query, args...).Delete(&data).Error
	if err != nil {
		panic(err)
	}
}

// 统计数量
func Count[T any](data T, query string, args ...any) int64 {
	var total int64
	db := db.Model(data)
	if query != "" {
		db = db.Where(query, args...)
	}
	if err := db.Count(&total).Error; err != nil {
		panic(err)
	}
	return total
}
