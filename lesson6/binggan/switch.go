package main

func main() {
	var x interface{}
	switch x.(type) {
	case nil: // ...
	case int: // ...
	case bool: // ...
	case string: // ...
	default: // ...
	}
}