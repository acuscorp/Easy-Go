package marshaling

import (
	"encoding/csv"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)
func MyStuffsOnReflect() {
  var entries []MyData 
  
  func(v interface{}){
    var myData MyData
    vPtr := reflect.ValueOf(v)
    println(vPtr.Type().String()) // prints []marshaling.MyData
    println(vPtr.Type() == reflect.ValueOf(entries).Type()) // prints true
    println(vPtr.Kind().String()) // prints slice
    println(vPtr.Kind() == reflect.Slice) // prints strue
    println(vPtr.Type().Elem().String()) //prints marshaling.MyData
    println(vPtr.Type().Elem() == reflect.TypeOf(myData)) // prints true
    println(vPtr.Type().Elem().Kind().String()) //prints struct
    println(vPtr.Type().Elem().Kind() == reflect.Struct)  // prints true
    
  }(entries)
}

func MyMarshalingCSV() {
  // this is a data example that could come from a csv file
	data := `name,age,has_pet
Jon,"100",true
"Fred ""The Hammer"" Smith",42,false
Martha,37,"true"
`
  // read the data and save it to a byte array allData
	r := csv.NewReader(strings.NewReader(data))
	allData, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
  // data read form the file and save it in  MyData structure array using Unmarshal
	var entries []MyData
	println("Unmarshaling")
	Unmarshal(allData, &entries)
	fmt.Println(entries)

	//now to turn entries into output
	println("Marshaling")
	out, err := Marshal(entries)
	if err != nil {
		panic(err)
	}
	sb := &strings.Builder{}
	w := csv.NewWriter(sb)
	w.WriteAll(out)
	fmt.Println(sb)
}

type MyData struct {
	Name   string `csv:"name"`
	Age    int    `csv:"age"`
	HasPet bool   `csv:"has_pet"`
}

// Marshal maps all of structs in a slice of structs to a slice of slice of strings.
// The first row written is the header with the column names.
func Marshal(v interface{}) ([][]string, error) {
	sliceVal := reflect.ValueOf(v)
	if sliceVal.Kind() != reflect.Slice {
		return nil, errors.New("must be a slice of structs")
	}
	structType := sliceVal.Type().Elem()
	if structType.Kind() != reflect.Struct {
		return nil, errors.New("must be a slice of structs")
	}
	var out [][]string
	header := marshalHeader(structType)
	out = append(out, header)
	for i := 0; i < sliceVal.Len(); i++ {
		row, err := marshalOne(sliceVal.Index(i))
		if err != nil {
			return nil, err
		}
		out = append(out, row)
	}
	return out, nil
}

func marshalHeader(vt reflect.Type) []string {
	var row []string
	for i := 0; i < vt.NumField(); i++ {
		field := vt.Field(i)
		if curTag, ok := field.Tag.Lookup("csv"); ok {
			row = append(row, curTag)
		}
	}
	return row
}

func marshalOne(vv reflect.Value) ([]string, error) {
	var row []string
	vt := vv.Type()
	for i := 0; i < vv.NumField(); i++ {
		fieldVal := vv.Field(i)
		if _, ok := vt.Field(i).Tag.Lookup("csv"); !ok {
			continue
		}
		switch fieldVal.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			row = append(row, strconv.FormatInt(fieldVal.Int(), 10))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			row = append(row, strconv.FormatUint(fieldVal.Uint(), 10))
		case reflect.String:
			row = append(row, fieldVal.String())
		case reflect.Bool:
			row = append(row, strconv.FormatBool(fieldVal.Bool()))
		default:
			return nil, fmt.Errorf("cannot handle field of kind %v", fieldVal.Kind())
		}
	}
	return row, nil
}

// why do we receive an interface?
// because it means you can pass any data type
func Unmarshal(data [][]string, v interface{}) error {
	sliceValPtr := reflect.ValueOf(v)         // why do I get an sliceValPtr?  from interfaces you get a reflect.Ptr  /
	if sliceValPtr.Kind() != reflect.Ptr {    // so in is necessary to have the ponter in order to save the data to the real data type
		return errors.New("must be a pointer to a slice of a struct")
	}

	sliceVal := sliceValPtr.Elem()            // from a pointer you can get the 'values or value' using Elem() so you can latter check /
	if sliceVal.Kind() != reflect.Slice {     //with Kind() return the type of the structure, could be a map, slice, array, etc,.
		return errors.New("must be a pointer to a slice of a struct")
	}

	structType := sliceVal.Type().Elem()      // from sliceVal.Type() return the type of the slice and it is a MyData and then I get from Elem() the values so it returns/
                                            // the values and of the MyData
	if structType.Kind() != reflect.Struct {  // check if this is and structure
		return errors.New("must be a pointer to a slice of a struct")
	}

	header := data[0]

	namePos := make(map[string]int, len(header))

	for k, v := range header {
		namePos[v] = k
	}

	for _, row := range data[1:] {
		newVal := reflect.New(structType).Elem()
		err := unmarshalOne(row, namePos, newVal)
		if err != nil {
			return err
		}
		sliceVal.Set(reflect.Append(sliceVal, newVal))
	}
	return nil

}

func unmarshalOne(row []string, namePos map[string]int, vv reflect.Value) error {
	vt := vv.Type()
	for i := 0; i < vv.NumField(); i++ {
		typeField := vt.Field(i)
		pos, ok := namePos[typeField.Tag.Get("csv")]
		if !ok {
			continue
		}
		val := row[pos]
		field := vv.Field(i)
		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			i, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return err
			}
			field.SetInt(i)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			i, err := strconv.ParseUint(val, 10, 64)
			if err != nil {
				return err
			}
			field.SetUint(i)
		case reflect.String:
			field.SetString(val)
		case reflect.Bool:
			b, err := strconv.ParseBool(val)
			if err != nil {
				return err
			}
			field.SetBool(b)
		default:
			return fmt.Errorf("cannot handle field of kind %v", field.Kind())
		}
	}
	return nil
}