package item

import "time"

type gateway interface {
	Create(Item) (Item, error)
	Update(Item) (Item, error)
	GetById(uint) (Item, error)
	GetAllFinished() ([]Item, error)
	GetAllUnfinished() ([]Item, error)
}

type Manager struct {
	gateway gateway
}

func (manager *Manager) AddItem(description string) (Item, error) {
	item := Item{
		Description: description,
		IsFinished:  false,
		CreatedAt:   time.Now(),
	}

	return manager.gateway.Create(item)
}

func (manager *Manager) FinishItem(id uint) (Item, error) {
	item, err := manager.GetItem(id)

	if err != nil {
		return item, err
	}

	item.Finish()

	return manager.gateway.Update(item)
}

func (manager *Manager) GetItem(id uint) (Item, error) {
	item, err := manager.gateway.GetById(id)

	return item, err
}

func (manager *Manager) GetFinishedItems() ([]Item, error) {
	return manager.gateway.GetAllFinished()
}

func (manager *Manager) GetUnfinishedItems() ([]Item, error) {
	return manager.gateway.GetAllUnfinished()
}
