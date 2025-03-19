package repository

import (
	"context"
	"strconv"
	"testing"

	"github.com/senither/custodian/database"
	"github.com/senither/custodian/database/model"
	"github.com/stretchr/testify/assert"
)

var defaultTestingTask = model.Task{
	Message: "Test Task",
	Status:  false,
}

func setupTaskTesting() (context.Context, model.User, func()) {
	database.InitiateDatabaseConnection(database.MemorySQLiteDSN)
	ctx := context.Background()

	user := model.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password",
	}

	CreateUser(ctx, user)
	createdUser, _ := FindUserByEmail(ctx, user.Email)

	return ctx, createdUser, func() {
		database.Disconnect()
	}
}

func TestCreateTaskForUser(t *testing.T) {
	ctx, user, close := setupTaskTesting()
	defer close()

	t.Run("it can create a new task for a user", func(t *testing.T) {
		err := CreateTaskForUser(ctx, &user, defaultTestingTask)
		assert.NoError(t, err)

		tasks, err := GetTasksWithSearchForUserWithRelations(ctx, &user, map[string]interface{}{})
		assert.NoError(t, err)
		assert.Len(t, tasks, 1)
		assert.Equal(t, defaultTestingTask.Message, tasks[0].Message)
	})
}

func TestGetTasksWithSearchForUserWithRelations(t *testing.T) {
	ctx, user, close := setupTaskTesting()
	defer close()

	CreateTaskForUser(ctx, &user, defaultTestingTask)

	for i := 0; i < 5; i++ {
		CreateTaskForUser(ctx, &user, model.Task{
			Message: "Example Task #" + strconv.Itoa(i),
			Status:  true,
		})
	}

	t.Run("it can get all tasks for a user", func(t *testing.T) {
		tasks, err := GetTasksWithSearchForUserWithRelations(ctx, &user, nil)
		assert.NoError(t, err)
		assert.Len(t, tasks, 6)
	})

	t.Run("it can get tasks by status for a user", func(t *testing.T) {
		tasks, err := GetTasksWithSearchForUserWithRelations(ctx, &user, map[string]interface{}{
			"status = ?": 0,
		})

		assert.NoError(t, err)
		assert.Len(t, tasks, 1)
		assert.Equal(t, defaultTestingTask.Message, tasks[0].Message)
	})

	t.Run("it can get tasks by message for a user", func(t *testing.T) {
		tasks, err := GetTasksWithSearchForUserWithRelations(ctx, &user, map[string]interface{}{
			"message LIKE ? ": "%task #3%",
		})

		assert.NoError(t, err)
		assert.Len(t, tasks, 1)
		assert.Equal(t, "Example Task #3", tasks[0].Message)
	})

	t.Run("it can get tasks by status and message for a user", func(t *testing.T) {
		tasks, err := GetTasksWithSearchForUserWithRelations(ctx, &user, map[string]interface{}{
			"status = ?":     1,
			"message LIKE ?": "%#3%",
		})

		assert.NoError(t, err)
		assert.Len(t, tasks, 1)
		assert.Equal(t, "Example Task #3", tasks[0].Message)
	})
}

func TestFindTaskForUser(t *testing.T) {
	ctx, user, close := setupTaskTesting()
	defer close()

	CreateTaskForUser(ctx, &user, defaultTestingTask)
	tasks, _ := GetTasksWithSearchForUserWithRelations(ctx, &user, nil)

	t.Run("it can find a task for a user by ID", func(t *testing.T) {
		task, err := FindTaskForUser(ctx, user, tasks[0].ID)
		assert.NoError(t, err)
		assert.Equal(t, defaultTestingTask.Message, task.Message)
	})

	t.Run("it returns an error when the task does not exist", func(t *testing.T) {
		_, err := FindTaskForUser(ctx, user, 999)
		assert.Error(t, err)
	})
}

func TestUpdateTask(t *testing.T) {
	ctx, user, close := setupTaskTesting()
	defer close()

	CreateTaskForUser(ctx, &user, defaultTestingTask)
	tasks, _ := GetTasksWithSearchForUserWithRelations(ctx, &user, nil)

	t.Run("it can update a task", func(t *testing.T) {
		task := tasks[0]
		err := UpdateTask(ctx, task, map[string]interface{}{
			"message": "Updated Task",
		})
		assert.NoError(t, err)

		updatedTask, _ := FindTaskForUser(ctx, user, task.ID)
		assert.Equal(t, "Updated Task", updatedTask.Message)
	})
}

func TestDeleteTask(t *testing.T) {
	ctx, user, close := setupTaskTesting()
	defer close()

	CreateTaskForUser(ctx, &user, defaultTestingTask)
	tasks, _ := GetTasksWithSearchForUserWithRelations(ctx, &user, nil)

	t.Run("it can delete a task", func(t *testing.T) {
		task := tasks[0]
		err := DeleteTask(ctx, task)
		assert.NoError(t, err)

		_, err = FindTaskForUser(ctx, user, task.ID)
		assert.Error(t, err)
	})
}
