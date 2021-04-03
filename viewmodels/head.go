package viewmodels

type Head struct {
	Title string
}

func NewHead(title string) Head {
	return Head{Title: "Creating Web Application with Go | " + title}
}
