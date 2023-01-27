package models

func (dt TodoDb) Map() *Todo {
	return &Todo{
		ID:        dt.ID.Hex(),
		Name:      dt.Name,
		Completed: dt.Completed,
	}
}

func (t *Todo) Map(userId string) *TodoDb {
	return &TodoDb{
		UserId:    userId,
		Completed: t.Completed,
		Name:      t.Name,
	}
}
