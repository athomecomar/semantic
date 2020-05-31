package semprov

type Category struct {
	Name   string
	Childs []*Category
}

func Loc(s string) *Category {
	return LocScoped(Root, s)
}

func (c *Category) String() string {
	return c.Name
}

func LocScoped(c *Category, s string) *Category {
	if c.Name == s {
		return c
	}
	for _, child := range c.Childs {
		if loc := LocScoped(child, s); loc != nil {
			return loc
		}
	}
	return nil
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
