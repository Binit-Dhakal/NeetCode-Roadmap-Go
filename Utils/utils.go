package utils

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func HelperTest(t *testing.T, name string, output interface{}, fn interface{}, args ...interface{}) {
	t.Helper()

	t.Run(name, func(t *testing.T) {
		result, err := callFunction(fn, args...)
		if err != nil {
			t.Errorf("Error %v", err)
			return
		}

		deepEqual, err := CheckUnorderedSliceEquality(result[0], output)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if !deepEqual {
			t.Errorf("got %v, expected %v", result[0], output)
		}
	})
}

func NormalizeAndStringify(slice interface{}) (string, error) {
	v := reflect.ValueOf(slice)

	if v.Kind() == reflect.Slice {
		// Recursively normalize each element in the slice
		elements := make([]string, v.Len())
		for i := 0; i < v.Len(); i++ {
			elemStr, err := NormalizeAndStringify(v.Index(i).Interface())
			if err != nil {
				return "", nil
			}
			elements[i] = elemStr
		}
		sort.Strings(elements)
		return fmt.Sprintf("[%s]", join(elements, ", ")), nil
	}

	return fmt.Sprint(v.Interface()), nil
}

func join(elements []string, sep string) string {
	result := ""
	for i, e := range elements {
		if i > 0 {
			result += sep
		}
		result += e
	}
	return result
}

func CheckUnorderedSliceEquality(got, want interface{}) (bool, error) {
	normalizedStr1, err1 := NormalizeAndStringify(got)
	normalizedStr2, err2 := NormalizeAndStringify(want)

	if err1 != nil || err2 != nil {
		return false, fmt.Errorf("error normalizing slices :%v, %v", err1, err2)
	}

	return normalizedStr1 == normalizedStr2, nil
}

func callFunction(fn interface{}, args ...interface{}) ([]interface{}, error) {
	fnValue := reflect.ValueOf(fn)

	if fnValue.Kind() != reflect.Func {
		return nil, fmt.Errorf("provided argument is not a function")
	}

	reflectArgs := make([]reflect.Value, len(args))
	for i, arg := range args {
		reflectArgs[i] = reflect.ValueOf(arg)
	}

	results := fnValue.Call(reflectArgs)

	interfaceResults := make([]interface{}, len(results))
	for i, result := range results {
		interfaceResults[i] = result.Interface()
	}

	return interfaceResults, nil
}
