package csvdata

import (
	"encoding/csv"
	"io"
	"reflect"
	"strconv"
)

//Decoder Decoder description
type Decoder struct {
	*csv.Reader
}

//NewDecoder create decoder for  csv reader
func NewDecoder(r io.Reader) *Decoder {
	dec := &Decoder{csv.NewReader(r)}
	dec.TrimLeadingSpace = true
	return dec
}

//Decode Decode
func (dec *Decoder) Decode(item interface{}) error {
	var (
		fields []string
		err    error
		tri    int64
		trui   uint64
		trf    float64
	)
	rv := reflect.ValueOf(item)
	fields, err = dec.Read()
	if err != nil {
		return err
	}
	pos := 0

	re := rv.Elem()
	fieldnum := re.NumField()
	for _, fv := range fields {

		for {
			if pos == fieldnum {
				break
			}
			done := true
			rf := re.Field(pos)
			pos++
			if rf.CanSet() == false {
				break
			}

			rfkind := rf.Kind()
			switch rfkind {
			case reflect.Float32:
				{
					trf, err = strconv.ParseFloat(fv, 32)
					if err != nil {
						return err
					}
					rf.SetFloat(trf)
				}
			case reflect.Float64:
				{
					trf, err = strconv.ParseFloat(fv, 64)
					if err != nil {
						return err
					}
					rf.SetFloat(trf)
				}

			case reflect.String:
				{
					rf.SetString(fv)
				}
			case reflect.Int:
				{
					tri, err = strconv.ParseInt(fv, 10, 64)
					if err != nil {
						return err
					}
					rf.SetInt(tri)
				}
			case reflect.Int8:
				{
					tri, err = strconv.ParseInt(fv, 10, 64)
					if err != nil {
						return err
					}
					rf.SetInt(tri)
				}
			case reflect.Int16:
				{
					tri, err = strconv.ParseInt(fv, 10, 64)
					if err != nil {
						return err
					}
					rf.SetInt(tri)
				}
			case reflect.Int32:
				{
					tri, err = strconv.ParseInt(fv, 10, 64)
					if err != nil {
						return err
					}
					rf.SetInt(tri)
				}
			case reflect.Int64:
				{
					tri, err = strconv.ParseInt(fv, 10, 64)
					if err != nil {
						return err
					}
					rf.SetInt(tri)
				}
			case reflect.Uint:
				{
					trui, err = strconv.ParseUint(fv, 10, 64)
					if err != nil {
						return err
					}
					rf.SetUint(trui)
				}
			case reflect.Uint8:
				{
					trui, err = strconv.ParseUint(fv, 10, 64)
					if err != nil {
						return err
					}
					rf.SetUint(trui)
				}
			case reflect.Uint16:
				{
					trui, err = strconv.ParseUint(fv, 10, 64)
					if err != nil {
						return err
					}
					rf.SetUint(trui)
				}
			case reflect.Uint32:
				{
					trui, err = strconv.ParseUint(fv, 10, 64)
					if err != nil {
						return err
					}
					rf.SetUint(trui)
				}
			case reflect.Uint64:
				{
					trui, err = strconv.ParseUint(fv, 10, 64)
					if err != nil {
						return err
					}
					rf.SetUint(trui)
				}
			default:
				{
					// do nothing
					done = false
				}

			}
			if done == true {
				break
			}
		}
	}

	return nil
}
