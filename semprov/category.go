package semprov

type Category struct {
	Name   string
	Childs []*Category
}

var (
	Root = &Category{Name: "root", Childs: []*Category{
		Health, Justice,
	}}

	Health = &Category{Name: "health", Childs: []*Category{
		Psychologist, Medic,
	}}
	Psychologist = &Category{Name: "psychologist"}
	Medic        = &Category{Name: "medic"}

	Justice = &Category{Name: "justice", Childs: []*Category{
		Lawyer, Attorney,
	}}
	Lawyer   = &Category{Name: "lawyer"}
	Attorney = &Category{Name: "attorney"}
)
