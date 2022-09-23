# Reflection

## Why using reflection

It allows to work with variables at runtime, using information that wansn't.
Allows to get adeeper comparation rather  if you use == 

## Types, kinds and values

On reflections you can get the type, kind and values of a variable 

### A type
Dedfines the properties of a variable
 - use reflect.TypeOf(\<variable a\>)
 - types can get the name like \<type>.Name()


### Checking if a varliable is nil

``` go

  func hasNoValue(i interface{}) bool {
	iv := reflect.ValueOf(i)
	if !iv.IsValid() {
		return true
	}
	switch iv.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Func, reflect.Interface:
		return iv.IsNil()
	default:
		return false
	}
}

func main() {
	var a interface{}
	fmt.Println(a == nil, hasNoValue(a)) // prints true true

	var b *int
	fmt.Println(b == nil, hasNoValue(b)) // prints true true

	var c interface{} = b
	fmt.Println(c == nil, hasNoValue(c)) // prints false true

	var d int
	fmt.Println(hasNoValue(d)) // prints false

	var e interface{} = d
	fmt.Println(e == nil, hasNoValue(e)) // prints false false
}
```

## Use Reflection to Write a Data Marshaler

Reflection is what the standard library uses to implement marshaling and unmarshaling. 
