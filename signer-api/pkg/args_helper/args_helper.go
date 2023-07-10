package args_helper

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func GetArgs(queryArgs url.Values, args interface{}) (interface{}, error) {
	structType := reflect.TypeOf(args)
	result := reflect.New(structType).Elem()

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		jsonKey := field.Tag.Get("json")
		value := queryArgs.Get(jsonKey)
		if value != "" {
			t := GetPropertyType(args, field.Name)
			switch t.Kind() {
			case reflect.Int:
				num, err := strconv.Atoi(value)
				if err != nil {
					return nil, fmt.Errorf("\nError converting string to int: %v\n", err.Error())
				}
				result.FieldByName(field.Name).SetInt(int64(num))
			case reflect.String:
				result.FieldByName(field.Name).SetString(value)
			case reflect.Slice:
				sl := GetIntSlice(value)
				result.FieldByName(field.Name).Set(reflect.ValueOf(sl))
			}
		}
	}

	return result.Interface(), nil
}

func GetPropertyType(data interface{}, key string) reflect.Type {
	value := reflect.ValueOf(data)
	field, _ := value.Type().FieldByName(key)

	return field.Type
}

func GetIntSlice(value string) []int {
	strSlice := strings.Split(value, ",")
	intSlice := make([]int, len(strSlice))

	for i, s := range strSlice {
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Errorf("\nError converting '%s' to int: %v\n", s, err)
		}
		intSlice[i] = num
	}

	return intSlice
}

func AddArgToQuery(builder *strings.Builder, args ...string) {
	const separator = " "

	for _, arg := range args {
		builder.WriteString(separator)
		builder.WriteString(arg)
	}
}

func GetArgsIn(str string) string {
	const start = "("
	const end = ")"
	return start + str + end
}

func GetNumsString(nums []int) string {
	strNums := make([]string, len(nums))

	for i, num := range nums {
		strNums[i] = strconv.Itoa(num)
	}

	return strings.Join(strNums, ",")
}
