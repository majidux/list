package usecases

import "github.com/majidux/list/entities"

type TodosRepository interface {
	GetAllTodos() ([]entities.Todo, error)
}