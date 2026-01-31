package repositories

import (
	"rest-api-go-gin/internal/database"
	"rest-api-go-gin/internal/models"
)

type EventRepository struct{}

func (r *EventRepository) Create(event *models.Event) error {
	return database.DB.Create(event).Error
}

func (r *EventRepository) FindAll() ([]models.Event, error) {
	var events []models.Event
	err := database.DB.Find(&events).Error
	return events, err
}

func (r *EventRepository) FindByID(id uint) (*models.Event, error) {
	var event models.Event
	err := database.DB.First(&event, id).Error
	return &event, err
}

func (r *EventRepository) FindByOwnerID(ownerID uint) ([]models.Event, error) {
	var events []models.Event
	err := database.DB.Where("owner_id = ?", ownerID).Find(&events).Error
	return events, err
}

func (r *EventRepository) Update(event *models.Event) error {
	return database.DB.Save(event).Error
}

func (r *EventRepository) Delete(id uint) error {
	return database.DB.Delete(&models.Event{}, id).Error
}
