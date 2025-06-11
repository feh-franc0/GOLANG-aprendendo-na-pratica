package pkg

// Example is a simple structure that can be reused in other projects.
type Example struct {
	Name  string
	Value int
}

// NewExample creates a new instance of Example.
func NewExample(name string, value int) Example {
	return Example{Name: name, Value: value}
}

// GetName returns the name of the Example.
func (e Example) GetName() string {
	return e.Name
}

// GetValue returns the value of the Example.
func (e Example) GetValue() int {
	return e.Value
}