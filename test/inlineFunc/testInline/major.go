package testInline

type Person struct {
}

func (this *Person) GetName(name string) string {
	return name
}
