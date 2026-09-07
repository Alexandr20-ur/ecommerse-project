package templates

import (
	"html/template"
)

type Store struct {
	Tmpl *template.Template
}

func New() (*Store, error) {
	tmpl, err := template.ParseFiles(
		"templates/base.gohtml",
		"templates/header.gohtml",
		"templates/content.gohtml",
		"templates/footer.gohtml",
		"templates/nav.gohtml",
	)
	if err != nil {
		return nil, err
	}

	return &Store{tmpl}, nil
}

// NewWithContent Задаем контент для страницы
func (s *Store) NewWithContent(contentFiles ...string) (*template.Template, error) {
	//Создаем копию базового шаблона
	tmpl := template.Must(s.Tmpl.Clone())

	//Добавляем контент страницы
	if _, err := tmpl.ParseFiles(contentFiles...); err != nil {
		return nil, err
	}

	return tmpl, nil
}
