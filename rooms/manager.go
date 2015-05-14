package rooms

type Manager struct {
	Rooms map[*conn.Connection]bool
}

func New() *Manager {
	return &Manager{
		Rooms: make(map[*Room]bool),
	}
}

func (m *Manager) AddRoom() error {
	return nil
}

func (m *Manager) RemoveRoom() error {
	return nil
}

func (m *Manager) GetAll() ([]*Room, error) {
	return nil, nil
}
