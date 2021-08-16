/*
@Author :  
@Group : Rosefintech
@Description :
@Data: 2020/11/27 15:47
*/

package tools

import (
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"strings"
)

//take the decimal place
func CutOutFloat64(f float64, length int) float64 {
	s := strconv.FormatFloat(f, 'f', -1, 64)
	res := CancelTargetAfterIndex(s, ".", length)
	return res
}

func CancelTargetAfterIndex(source, tarStr string, length int) float64 {
	result, err := strconv.ParseFloat(source, 64)
	if err != nil {
		return 0
	}
	index := strings.Index(source, tarStr)
	if index < 0 {
		return result
	}
	if length != 0 {
		length += 1
	}
	if len(source[index:]) < length {
		length = len(source[index:])
	}
	str := source[:index+length]
	//fmt.Println(source[:index+1])
	result, err = strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return result
}

func CancelAfterIndex(source string, length int) float64 {
	result, err := strconv.ParseFloat(source, 64)
	if err != nil {
		return 0
	}

	if len(source)-length <= 0 {
		return result
	}
	str := source[:len(source)-length]
	result, err = strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return result
}

func Int64StrToFloat64(source string, decimalsLen int) (float64, error) {
	if len(source)-decimalsLen == 0 {
		source = strings.Join([]string{"0.", source}, "")
		result, err := strconv.ParseFloat(source, 64)
		return result, err
	} else if len(source)-decimalsLen > 0 {
		str1 := source[:len(source)-decimalsLen]
		result, err := strconv.ParseFloat(str1, 64)
		str2 := source[len(source)-decimalsLen:]
		source = strings.Join([]string{"0.", str2}, "")
		decimals, err := strconv.ParseFloat(source, 64)
		if err != nil {
			return result, err
		}
		result += decimals
		return result, err
	}
	base := "0."
	for i := len(source) - decimalsLen; i < 0; i++ {
		base = strings.Join([]string{base, "0"}, "")
	}
	source = strings.Join([]string{base, source}, "")
	result, err := strconv.ParseFloat(source, 64)
	return result, err
}

//get decimal length
func GetFloat64DecimalsLength(f float64) int {
	index := strings.Index(strconv.FormatFloat(f, 'f', -1, 64), ".")
	if index < 0 {
		return 0
	}
	return len(strings.Split(fmt.Sprintf("%v", f), ".")[1])
}

// interface{}=> []interface{}
func CreateAnyTypeSlice(slice interface{}) ([]interface{}, bool) {
	val, ok := isSlice(slice)

	if !ok {
		return nil, false
	}

	sliceLen := val.Len()

	out := make([]interface{}, sliceLen)

	for i := 0; i < sliceLen; i++ {
		out[i] = val.Index(i).Interface()
	}

	return out, true
}

func isSlice(arg interface{}) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)

	if val.Kind() == reflect.Slice {
		ok = true
	}

	return
}

func Reverse(arg interface{}) (result interface{}, err error) {
	slice, ok := CreateAnyTypeSlice(arg)
	if !ok {
		return nil, errors.New("not slice")
	}
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice, nil
}

func ReserveDecimals(bigIntStr string, base int64) (string, error) {
	n := new(big.Int)
	a, ok := new(big.Int).SetString(bigIntStr, 10)
	if !ok {
		return "0", errors.New("can not new bigInt")
	}
	if 0 == a.Cmp(n) {
		return bigIntStr, nil
	}
	b := new(big.Int).SetInt64(base)
	a = a.Div(a, b)
	return a.String(), nil
}
