package models

import (
	iotgrpcapi "github.com/SKF/proto/v2/iot"
)

type Tasks []Task

func (tasks Tasks) ToGRPC() (td *iotgrpcapi.TaskDescriptions) {
	td = &iotgrpcapi.TaskDescriptions{}
	td.TaskDescriptionArr = []*iotgrpcapi.TaskDescription{}

	for _, task := range tasks {
		taskElement := task.ToGRPC()
		td.TaskDescriptionArr = append(td.TaskDescriptionArr, &taskElement)
	}

	return
}

func (tasks Tasks) FilterOnUncompleted() (result Tasks) {
	for _, task := range tasks {
		if !task.IsCompleted && task.TaskStatus != "COMPLETED" {
			result = append(result, task)
		}
	}
	return
}
