package querybuilder

import (
	"fmt"
	"strings"
)

// Builder represents simple query builder, mainly to prevent SQL injection.
type Builder struct {
	sb      *strings.Builder
	counter int
	args    []interface{}
}

// New creates new query builder.
func New(baseQuery string, data ...interface{}) *Builder {
	var sb strings.Builder

	b := &Builder{
		sb: &sb,
	}
	b.addQuery(baseQuery, data...)
	return b
}

func (b *Builder) AddQuery(format string, data ...interface{}) {
	b.sb.WriteString(" ")
	b.addQuery(format, data...)
}

func (b *Builder) addQuery(format string, data ...interface{}) {
	b.args = append(b.args, data...)

	for i := 0; i < len(data); i++ {
		b.counter++
		format = strings.Replace(format, "?", fmt.Sprintf("$%d", b.counter), 1)
	}

	b.sb.WriteString(format)
}

// AddString add raw string to the query.
//
// A space is prepended before the query.
func (b *Builder) AddString(str string) {
	b.sb.WriteString(" " + str)
}

// Query returns the safe query with placeholder replaced as necessary.
func (b *Builder) Query() string {
	return b.sb.String()
}

// Args returns the data arguments for the query.
func (b *Builder) Args() []interface{} {
	return b.args
}
