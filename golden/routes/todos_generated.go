package routes

import (
	"github.com/bketelsen/factor/markup"
)

type Todos struct {
}

var TodosTemplate = `<div class="todos">
    {{ range .Todos }}
    <!--
        TodoComponent is an autogenerated template function that works
        somehow. Maybe it creates a new components.Todo and calls Todo.Render?
    -->
    {{ .TodoComponent . }}
    {{ end }}
</div>`
var TodosStyles = ``

func (t *Todos) Render() string {
	return TodosTemplate
}

func init() {
	markup.Register(&Todos{})
}
