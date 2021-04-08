package main

import (
	"errors"
	"time"

	"github.com/spidernest-go/logger"
)

var TASKS []*Task

var ErrNoInsert = errors.New("Task was not inserted.")

type Task struct {
	ID          uint      `db:"id"`
	Event       string    `db:"event"`
	RequestedBy string    `db:"requested_by"`
	Affects     string    `db:"affects"`
	Deadline    time.Time `db:"deadline"`
}

func retrieveAllTasks() {
	DB.Collection("tasks").Find().All(&TASKS)
}

func writeTask(task *Task) (*Task, error) {
	DB.InsertInto("tasks").Values(task).Returning("id").Iterator().Next(&task)

	// HACK: Does r actually mean failure on false?
	/*if r == false {
		logger.Error().Msg("Task could not be inserted into the table.")
	}*/

	return task, nil
}

func delegateTask(task *Task) {
	switch task.Event {
	case "unsilence":
		go taskUnsilence(task)
	}
}

func removeTask(task *Task) {
	DB.Collection("tasks").Find().Where("id = ", task.ID).Delete()
}

func taskUnsilence(task *Task) {
	time.Sleep(time.Until(task.Deadline))

	err := discord.GuildMemberRoleRemove(GUILDIDENT, task.Affects, ROLESILENCED)
	if err != nil {
		logger.Warn().Err(err).Msg("User was unable to be unsilenced, as they were already unaffected, left the guild, or it had been removed prior. Either way, this task will have to be manually removed from the task pool.")

		return
	}

	removeTask(task)
}
