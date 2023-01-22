package fiber

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"time"

	"example/m/model"
)

type TodoRequest struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

var todos = []model.Todo{
	{
		Id:     0,
		Title:  "hello world",
		Detail: "hello",
		Date:   time.Now(),
	},
	{
		Id:     1,
		Title:  "hello mars",
		Detail: "hello",
		Date:   time.Now(),
	},
	{
		Id:     2,
		Title:  "hello neptune",
		Detail: "hello",
		Date:   time.Now(),
	},
}

func (f *FiberServer) addTodoRoute(router fiber.Router) {
	r := router.Group("todos")

	r.Get("/", f.getAllTodo)
	r.Post("/", f.createTodo)
	r.Get("/:id", f.getAllById)
	r.Put("/:id", f.updateTodo)
	r.Delete("/:id", f.deleteTodo)
}

func (f *FiberServer) getAllById(c *fiber.Ctx) error {
	s := c.Params("id")

	i, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	todo := todos[i]

	return c.JSON(todo)
}

func (f *FiberServer) getAllTodo(c *fiber.Ctx) error {
	return c.JSON(todos)
}

func (f *FiberServer) createTodo(c *fiber.Ctx) error {
	var t TodoRequest
	if err := c.BodyParser(&t); err != nil {
		panic(err)
	}

	newTodo := model.Todo{
		Id:     todos[len(todos)-1].Id + 1,
		Title:  t.Title,
		Detail: t.Detail,
		Date:   time.Now(),
	}
	todos = append(todos, newTodo)

	return c.Send([]byte("OK"))
}

func (f *FiberServer) updateTodo(c *fiber.Ctx) error {
	s := c.Params("id")
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	var t TodoRequest
	if err := c.BodyParser(&t); err != nil {
		panic(err)
	}

	todos[i].Title = t.Title
	todos[i].Detail = t.Detail
	todos[i].Date = time.Now()

	return c.Send([]byte("OK"))
}

func (f *FiberServer) deleteTodo(c *fiber.Ctx) error {
	s := c.Params("id")
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	fmt.Println("Befor: ", todos)
	todos = removeFromSlice(todos, i)
	fmt.Println("After: ", todos)

	return c.Send([]byte("OK"))
}

func removeFromSlice(sli []model.Todo, i int) []model.Todo {
	return append(sli[:i], sli[i+1:]...)
}
