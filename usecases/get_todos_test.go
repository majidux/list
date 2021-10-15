package usecases_test

import (
	"fmt"
	"testing"

	"github.com/majidux/list/entities"
	"github.com/majidux/list/usecases"
)

type MockTodosRepo struct {}

func (MockTodosRepo) GetAllTodos() ([]entities.Todo, error) {
	return nil, fmt.Errorf("Somthing went wrong")
}

func TestGetTodos(t *testing.T){
	repo := new(MockTodosRepo)
	todos, err := usecases.GetTodos(repo)
	if err != usecases.ErrInternal{
		t.Fatalf("Expected ErrInternal, got %v",  err)
	}
	if todos != nil {
		t.Fatalf("Expected nil, got %v",  todos)
	}
}