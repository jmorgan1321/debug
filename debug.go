// debug is a library that makes it simple to add tracing and logging into a project.
package debug

type IndentLevel int

func (l *IndentLevel) Increment() {
	*l++
}
func (l *IndentLevel) Decrement() {
	*l--
}
func (l *IndentLevel) String() string {
	s := ""
	for i := *l; i > 0; i-- {
		s += "\t"
	}
	return s
}

var IndentationLevel *IndentLevel

func init() {
	IndentationLevel = new(IndentLevel)
}
