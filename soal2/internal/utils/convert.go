package utils

import (
	"database/sql"
	"errors"
	"github.com/mitchellh/mapstructure"
	"reflect"
	"strconv"
)

func S2ID(idS string) (int64, error) {
	if idS == "" {
		return 0, errors.New("id不能为空")
	}
	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		return 0, err
	}
	if id <= 0 {
		return 0, errors.New("id不能小于0")
	}
	return id, nil
}

func S2N(idS string) (int64, error) {
	if idS == "" {
		return 0, nil
	}
	num, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func S2F64(idS string) (float64, error) {
	if idS == "" {
		return 0, nil
	}
	num, err := strconv.ParseFloat(idS, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func S2I32(idS string) (int32, error) {
	if idS == "" {
		return 0, nil
	}
	num, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		return 0, err
	}
	return int32(num), nil
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func Map2Struct(input, result interface{}) {
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &result,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		panic(err)
	}

	err = decoder.Decode(input)
	if err != nil {
		panic(err)
	}
}

func ID(result sql.Result, err error) (int64, error) {
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

type vFn func() error

func V(fns ...vFn) error {
	for _, fn := range fns {
		err := fn()
		if err != nil {
			return err
		}
	}
	return nil
}