package main
import (
  "fmt"
)
/*
func universal(v interface{}) (string ,error) {
	switch t := v.(type) {
	case int, int32, int64:
		return fmt.Sprint(t, " is int"), nil
	case float32, float64:
		return fmt.Sprint(t, " is float"), nil
	case string:
		return fmt.Sprint(t, " is string"), nil
	case bool:
		return fmt.Sprint(t, " is bool"), nil
	default:
		return fmt.Sprint(t, " is unknown"), errors.New("OMG you used universal function function with invalid argument")
	}
}

func main() {
	res, _ := universal(5)
	fmt.Println(res)
	res, _ = universal("hello")
	fmt.Println("hello:", res)
	res, _ = universal(3.14)
	fmt.Println(res)
	res, _ = universal(true)
	fmt.Println(res)
	var struc struct{}
	res, _ = universal(struc)
	fmt.Println("{}   :", res)
}
*/
func getData() []interface{} {
  return []interface{}{
    1,
    3.14,
    "hello",
    true,
    struct{}{},
  }
}

func getTypeName(v interface{}) string {
  switch v.(type) {
  case int, int32, int64:
    return "int"
  case float64, float32:
    return "float"
  case bool:
    return "bool"
  case string:
    return "string"
  default:
    return "unknown"
  }
}

func main() {
  data := getData()
  for i := 0; i < len(data); i++ {
    fmt.Printf("%v is %v\n", data[i], getTypeName(data[i]))
  }
}
