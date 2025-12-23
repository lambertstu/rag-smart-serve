package database

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
)

func UniGetObjId(id interface{}) (*primitive.ObjectID, error) {
	if id == nil {
		return nil, errors.New("id is required")
	}
	var objId primitive.ObjectID
	if _, ok := id.(string); ok {
		_objId, err := primitive.ObjectIDFromHex(id.(string))
		if err != nil {
			return nil, err
		}
		objId = _objId
	} else if _, ok := id.(primitive.ObjectID); ok {
		objId = id.(primitive.ObjectID)
	} else {
		return nil, errors.New("id type is not supported")
	}
	if objId.IsZero() {
		return nil, errors.New("id is required")
	}
	return &objId, nil
}

func UniGetObjIdList(idList any) ([]primitive.ObjectID, error) {
	if idList == nil {
		return nil, errors.New("idList is required")
	}
	// 通过反射判断是不是数组
	t := reflect.TypeOf(idList).Kind()
	if t != reflect.Slice {
		return nil, errors.New("idList type is not supported")
	}

	var objIdList []primitive.ObjectID
	// 遍历数组
	vv := reflect.ValueOf(idList)
	for i := 0; i < vv.Len(); i++ {
		objId, err := UniGetObjId(vv.Index(i).Interface())
		if err != nil {
			return nil, err
		}
		objIdList = append(objIdList, *objId)
	}
	if len(objIdList) == 0 {
		return nil, errors.New("idList is required")
	}
	return objIdList, nil
}
