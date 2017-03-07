package csvdata

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
)

var ErrNoField error = errors.New("csvdata: no filed for encoder")

type Encoder struct {
	*csv.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	enc := &Encoder{csv.NewWriter(w)}
	return enc
}

func (enc *Encoder) Encode(item interface{}) error {
	rv := reflect.ValueOf(item)
	n := rv.NumField()
	ss := make([]string, n)
	pos := 0
	for i := 0; i < n; i++ {
		rfield := rv.Field(i)
		rftk := rfield.Type().Kind()
		switch rftk {
		case reflect.Float64:
			{
				fv := rfield.Float()
				sv := strconv.FormatFloat(fv, 'E', -1, 64)
				ss[pos] = sv
				pos++
			}
		case reflect.Float32:
			{
				fv := rfield.Float()
				sv := strconv.FormatFloat(fv, 'E', -1, 32)
				ss[pos] = sv
				pos++
			}
		case reflect.String:
			{
				sv := rfield.String()
				sv = strings.TrimSpace(sv)
				ss[pos] = sv
				pos++
			}
		case reflect.Int:
			{
				fv := rfield.Int()
				sv := strconv.FormatInt(fv, 10)
				ss[pos] = sv
				pos++
			}
		case reflect.Int8:
			{
				fv := rfield.Int()
				sv := strconv.FormatInt(fv, 10)
				ss[pos] = sv
				pos++
			}
		case reflect.Int16:
			{
				fv := rfield.Int()
				sv := strconv.FormatInt(fv, 10)
				ss[pos] = sv
				pos++
			}
		case reflect.Int32:
			{
				fv := rfield.Int()
				sv := strconv.FormatInt(fv, 10)
				ss[pos] = sv
				pos++
			}
		case reflect.Int64:
			{
				fv := rfield.Int()
				sv := strconv.FormatInt(fv, 10)
				ss[pos] = sv
				pos++
			}
		case reflect.Uint:
			{
				fv := rfield.Uint()
				sv := strconv.FormatUint(fv, 10)
				ss[pos] = sv
				pos++
			}
		case reflect.Uint8:
			{
				fv := rfield.Uint()
				sv := strconv.FormatUint(fv, 10)
				ss[pos] = sv
				pos++
			}
		case reflect.Uint16:
			{
				fv := rfield.Uint()
				sv := strconv.FormatUint(fv, 10)
				ss[pos] = sv
				pos++
			}
		case reflect.Uint32:
			{
				fv := rfield.Uint()
				sv := strconv.FormatUint(fv, 10)
				ss[pos] = sv
				pos++
			}
		case reflect.Uint64:
			{
				fv := rfield.Uint()
				sv := strconv.FormatUint(fv, 10)
				ss[pos] = sv
				pos++
			}
		default:
			{
				// do nothing
				fmt.Println(rftk)
			}
		}

	}

	if pos > 0 {
		err := enc.Write(ss)
		enc.Flush()
		return err
	}

	return ErrNoField
}
