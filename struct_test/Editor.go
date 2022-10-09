package struct_test

type Editor struct {
	firstName string
}

func newEditor(firstName string) *Editor {
	return &Editor{firstName: firstName}
}
