package util

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/SaiHLu/rest-template/common/constant"
	"github.com/SaiHLu/rest-template/common/logger"
	"github.com/SaiHLu/rest-template/internal/core/entity"
	"github.com/SaiHLu/rest-template/internal/infrastructure/cache"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func AuditLogger(db *gorm.DB, cacheStorage cache.Cache) {
	db.Callback().Create().After("gorm:create").Register("create_audit_log", createAuditLog("CREATE", cacheStorage))
	db.Callback().Update().After("gorm:update").Register("update_audit_log", createAuditLog("UPDATE", cacheStorage))
	db.Callback().Delete().After("gorm:delete").Register("delete_audit_log", createAuditLog("DELETE", cacheStorage))
}

func createAuditLog(action string, cacheStorage cache.Cache) func(db *gorm.DB) {
	var (
		userId uuid.UUID
		err    error
	)
	return func(db *gorm.DB) {
		if db.Statement.Schema != nil && db.Statement.Schema.Table == "audit_log" || db.Error != nil {
			return
		}

		idBytes, _ := cacheStorage.Get(string(constant.UserIdCtx))
		userId, err = uuid.Parse(string(idBytes))
		if err != nil {
			userId = uuid.Nil
		}

		recordMap, err := getDataBeforeOperation(db)
		if err != nil {
			return
		}

		objId := getKeyFromData("id", recordMap)
		parsedUUID, err := uuid.Parse(objId)
		if err != nil {
			logger.Error(fmt.Sprintf("error parsing UUID: %s", err.Error()))
			return
		}

		auditLog := &entity.AuditLog{
			EntityType: db.Statement.Schema.Table,
			Action:     action,
			EntityId:   parsedUUID,
			Data:       prepareData(recordMap),
			UserId:     &userId,
		}

		if err := db.Session(&gorm.Session{SkipHooks: true, NewDB: true}).Table("audit_log").Create(auditLog).Error; err != nil {
			logger.Error(fmt.Sprintf("error in audit log creation: %s", err.Error()))
			return
		}
	}
}

func getKeyFromData(key string, data map[string]interface{}) string {
	objId, ok := data[key]
	if !ok {
		return ""
	}
	return objId.(string)
}

func prepareData(data map[string]interface{}) []byte {
	dataByte, _ := json.Marshal(&data)
	return dataByte
}

func getDataBeforeOperation(db *gorm.DB) (map[string]interface{}, error) {
	objMap := map[string]interface{}{}
	if db.Error == nil && !db.DryRun {
		objectType := reflect.TypeOf(db.Statement.ReflectValue.Interface())

		// Create a new instance of the object type
		targetObj := reflect.New(objectType).Interface()

		var primaryKeyValue string
		value := db.Statement.ReflectValue

		// Check if the value is a struct
		if value.Kind() == reflect.Struct {
			idField := value.FieldByName("ID")
			if idField.IsValid() && idField.CanInterface() {
				primaryKeyValue = idField.Interface().(uuid.UUID).String()
			}
		}

		// Fetch the target object separately
		if err := db.Session(&gorm.Session{SkipHooks: true, NewDB: true}).Where("id = ?", primaryKeyValue).Where("deleted_at is null or deleted_at is not null").Unscoped().First(&targetObj).Error; err != nil {
			logger.Error(fmt.Sprintf("gorm callback: error while finding target object: %s", err.Error()))
			return nil, err
		}

		logger.Debug(fmt.Sprintf("targetObj: %+v", targetObj))

		jsonBytes, err := json.Marshal(targetObj)
		if err != nil {
			logger.Error(fmt.Sprintf("gorm callback: error while marshalling json data: %s", err.Error()))
			return nil, err
		}
		json.Unmarshal(jsonBytes, &objMap)
	}
	return objMap, nil
}
