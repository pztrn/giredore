package structs

// Reply defined reply data structure that giredored and giredorectl
// will use.
type Reply struct {
	Status Status
	Errors []Error
	Data   interface{}
}
