package csvdata

import (
	"encoding/csv"
	_ "errors"
	"fmt"
	"io"
	"reflect"
	"strconv"
)

type Decoder struct {
	*csv.Reader
}

func NewDecoder(r io.Reader) *Decoder {
	dec := &Decoder{csv.NewReader(r)}
	dec.TrimLeadingSpace = true
	return dec
}

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
	fmt.Println(fields)
	re := rv.Elem()
	fieldnum := re.NumField()
	fmt.Println(fieldnum)
	for _, fv := range fields {

		for {
			if pos == fieldnum {
				break
			}
			done := true
			rf := re.Field(pos)
			//fmt.Println("name:", rf.Type().Name())
			fmt.Println("name:", re.Field(pos).T)
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
					fmt.Println(pos, trf)
					rf.SetFloat(trf)
				}
			case reflect.Float64:
				{
					trf, err = strconv.ParseFloat(fv, 64)
					if err != nil {
						return err
					}
					fmt.Println(pos, trf)
					rf.SetFloat(trf)
				}

			case reflect.String:
				{
					fmt.Println(pos, fv)
					rf.SetString(fv)
				}
			case reflect.Int:
				{
					tri, err = strconv.ParseInt(fv, 10, 64)
					if err != nil {
						return err
					}
					fmt.Println(pos, tri)
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
					fmt.Println(pos, trui)
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
