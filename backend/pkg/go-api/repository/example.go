package repository

type IExampleRepository interface {
	GetName(name string) string
}

type ExampleRepository struct {
}

func NewExampleRepository() IExampleRepository {
	return &ExampleRepository{}
}

func (r *ExampleRepository) GetName(name string) string {
	return name
}
