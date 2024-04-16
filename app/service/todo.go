package service

import (
	"flag"
	"fmt"
	"net/http"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/swarajbhagat11/todo-app/app/repository"
	"github.com/swarajbhagat11/todo-app/err"
	"github.com/swarajbhagat11/todo-app/helpers"
	"github.com/swarajbhagat11/todo-app/models"
)

var httpGet = http.Get

type TodoService struct {
	repo repository.Todo
}

func NewTodoService(repo repository.Todo) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) addAllTodo(todoListToAdd []models.Todo) (bool, *err.AppError) {
	return s.repo.AddAllTodo(todoListToAdd)
}

func (s *TodoService) GetByIndex(index int) (models.Todo, *err.AppError) {
	return s.repo.GetByIndex(index)
}

func (s *TodoService) fetchTodo(url string, todoChan chan<- models.Todo, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := httpGet(url)
	if err != nil {
		log.Errorln("[TodoService => fetchTodo] Error fetching TODO:", err)
		return
	}
	defer resp.Body.Close()

	todo, err := helpers.DecodeTodoJSON(resp.Body)
	if err != nil {
		log.Errorln("[TodoService => fetchTodo] Error decoding JSON:", err)
		return
	}

	todoChan <- todo
}

func (s *TodoService) ReadData() {
	flag.Parse()

	todoChan := make(chan models.Todo)
	var wg sync.WaitGroup

	// Fetch the first 20 even numbered TODOs
	for i := 2; i <= 40; i += 2 {
		url := fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%d", i)
		wg.Add(1)
		go s.fetchTodo(url, todoChan, &wg)
	}

	// Start a goroutine to close the channel when all TODOs are fetched
	go func() {
		wg.Wait()
		close(todoChan)
	}()

	var todoList = []models.Todo{}
	// Print the title and completion status of each TODO
	for todo := range todoChan {
		fmt.Printf("Title: %s, Completed: %t\n", todo.Title, todo.Completed)
		todoList = append(todoList, todo)
	}

	success, err := s.addAllTodo(todoList)
	if success {
		log.Println("[TodoService => ReadData] All to do list added successfully!")
	}
	if err != nil {
		log.Errorln("[TodoService => ReadData] Error while adding Todo list", err)
	}
}
