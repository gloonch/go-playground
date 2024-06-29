package task

import (
	"fmt"
	"todocli/entity"
)

type ServiceRepository interface {
	//DoesThisUserHaveCategoryID(userID, categoryID int) bool
	CreateNewTask(t entity.Task) (entity.Task, error)
	ListUserTasks(userID int) ([]entity.Task, error) // it's not domain driven
}

type Service struct {
	repository ServiceRepository
}

func NewService(repository ServiceRepository) Service {

	return Service{repository: repository}

}

type CreateRequest struct {
	Title               string
	DueDate             string
	CategoryID          int
	AuthenticatedUserID int
}

type CreateResponse struct {
	Task entity.Task
}

func (t Service) Create(req CreateRequest) (CreateResponse, error) {

	//if t.repository.DoesThisUserHaveCategoryID(req.AuthenticatedUserID, req.CategoryID) {
	//	return CreateResponse{}, fmt.Errorf("User does not have category %d", req.CategoryID)
	//}

	createdTask, cErr := t.repository.CreateNewTask(entity.Task{
		Title:      req.Title,
		DueDate:    req.DueDate,
		CategoryID: req.CategoryID,
		IsDone:     false,
		UserID:     req.AuthenticatedUserID,
	})
	if cErr != nil {
		return CreateResponse{}, fmt.Errorf("Failed to create task: %v", cErr)
	}

	return CreateResponse{Task: createdTask}, nil
}

type ListRequest struct {
	UserID int
}

type ListResponse struct {
	Tasks []entity.Task
}

func (t Service) List(req ListRequest) (ListResponse, error) {
	tasks, err := t.repository.ListUserTasks(req.UserID)
	if err != nil {
		return ListResponse{}, fmt.Errorf("cant't list user tasks: %v", err)
	}

	return ListResponse{Tasks: tasks}, nil

}
