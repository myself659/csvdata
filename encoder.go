package csvdata

import (
	"encoding/csv"
	"errors"
	"io"
	"reflect"
	"strconv"
	"strings"
)

//ErrNoField no field error
var ErrNoField error = errors.New("csvdata: no filed for encoder")

//Encoder  Encoder description
type Encoder struct {
	*csv.Writer
}

//NewEncoder create encoder for csv writer
func NewEncoder(w io.Writer) *Encoder {
	enc := &Encoder{csv.NewWriter(w)}
	return enc
}

//Encode encode csv item to  csv writer
func (enc *Encoder) Encode(item interface{}) error {
	rv := reflect.ValueOf(item)
	n := rv.NumField()
	ss := make([]string, n)
	pos := 0
	for i := 0; i < n; i++ {
		rfield := rv.Field(i)
		rftk := rfield.Type().Kind()
		switch rftk {
		case reflect.Float64, reflect.Float32:
			{
				fv := rfield.Float()
				sv := strconv.FormatFloat(fv, 'E', -1, 64)
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
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			{
				fv := rfield.Int()
				sv := strconv.FormatInt(fv, 10)
				ss[pos] = sv
				pos++
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			{
				fv := rfield.Uint()
				sv := strconv.FormatUint(fv, 10)
				ss[pos] = sv
				pos++
			}
		default:
			{
				// do nothing
				//fmt.Println(rftk)
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
