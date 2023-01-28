package models

func (dt *TodoDB) Map() *Todo {
	return &Todo{
		ID:        dt.ID.Hex(),
		Name:      dt.Name,
		Completed: dt.Completed,
	}
}

func (t *Todo) Map(userID string) *TodoDB {
	return &TodoDB{
		UserID:    userID,
		Completed: t.Completed,
		Name:      t.Name,
	}
}
