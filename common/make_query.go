package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/url"
	"reflect"
	"strings"
)

func ToQueryByJsonTag(ptr interface{}) (query string, err error) {
	return ToQueryByTag(ptr, "json")
}

func ToQueryByTag(ptr interface{}, tag string) (query string, err error) {
	// Check if ptr is a map
	ptrVal := reflect.ValueOf(ptr)
	var pointed interface{}
	if ptrVal.Kind() == reflect.Ptr {
		ptrVal = ptrVal.Elem()
		pointed = ptrVal.Interface()
	}
	if ptrVal.Kind() == reflect.Map &&
		ptrVal.Type().Key().Kind() == reflect.String {
		if pointed != nil {
			ptr = pointed
		}
		return mapToQuery(ptr)
	}

	return toQueryByPtr(ptr, tag)
}

func mapToQuery(ptr interface{}) (query string, err error) {
	el := reflect.TypeOf(ptr).Elem()

	if el.Kind() == reflect.Slice {
		ptrMap, ok := ptr.(map[string][]string)
		if !ok {
			return "", errors.New("cannot convert to map slices of strings")
		}
		vals := make(url.Values)
		for k, v := range ptrMap {
			vals.Set(k, strings.Join(v, ","))
		}

		return vals.Encode(), nil
	}

	ptrMap, ok := ptr.(map[string]string)
	if !ok {
		return "", errors.New("cannot convert to map of strings")
	}
	vals := make(url.Values)
	for k, v := range ptrMap {
		vals.Set(k, v)
	}

	return vals.Encode(), nil
}

func toQueryByPtr(ptr interface{}, tag string) (query string, err error) {
	vals := make(url.Values)
	err = mapping(reflect.ValueOf(ptr), vals, tag)
	if err != nil {
		return "", err
	}
	query = vals.Encode()
	return
}

func mapping(value reflect.Value, setter url.Values, tag string) error {
	var vKind = value.Kind()

	if vKind == reflect.Ptr {
		vPtr := value
		if value.IsNil() {
			return nil
		}
		err := mapping(vPtr.Elem(), setter, tag)
		return err
	}

	if vKind != reflect.Struct {
		return nil
	}

	if vKind == reflect.Struct {
		tValue := value.Type()

		for i := 0; i < value.NumField(); i++ {
			sf := tValue.Field(i)
			if sf.PkgPath != "" && !sf.Anonymous { // unexported
				continue
			}

			val := value.Field(i)
			fld := tValue.Field(i)
			vkd := val.Kind()
			switch vkd {
			case reflect.Ptr, reflect.Slice, reflect.Map:
				if val.IsNil() {
					continue
				}
			}

			tagStr := fld.Tag.Get(tag)
			var tags []string
			if len(tagStr) > 0 {
				tags = strings.Split(tagStr, ",")
			}
			keyName := ""
			if len(tags) > 0 {
				if tags[0] == "-" {
					continue
				}
				keyName = tags[0]
				tags = tags[1:]
				tagMap := make(map[string]struct{})
				for _, item := range tags {
					tagMap[item] = struct{}{}
				}
				if _, ok := tagMap["omitempty"]; ok && val.IsZero() {
					continue
				}
			} else {
				keyName = fld.Name
			}

			var valData []byte
			var err error
			switch vkd {
			case reflect.Slice, reflect.Array:
				var buf bytes.Buffer
				for i := 0; i < val.Len(); i++ {
					v := val.Index(i)
					vBytes, err := json.Marshal(v.Interface())
					if err != nil {
						return err
					}
					if i > 0 {
						buf.WriteByte(',')
					}
					buf.Write(vBytes[1 : len(vBytes)-1])
				}
				valData = buf.Bytes()
			case reflect.Map, reflect.Struct, reflect.String, reflect.Ptr:
				valData, err = json.Marshal(val.Interface())
				if err != nil {
					return err
				}
				if len(valData) > 0 && valData[0] == '"' {
					valData = valData[1 : len(valData)-1]
				}
			default:
				valData, err = json.Marshal(val.Interface())
				if err != nil {
					return err
				}
			}

			setter.Set(keyName, string(valData))
		}
		return nil
	}
	return nil
}
