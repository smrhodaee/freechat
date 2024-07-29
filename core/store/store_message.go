package store

import (
	"app/models"

	"gorm.io/gorm"
)

type MessageStore struct {
	db *gorm.DB
}

func NewMessageStore(db *gorm.DB) *MessageStore {
	return &MessageStore{
		db: db,
	}
}

func (s *MessageStore) Create(msg *models.Message) error {
	return s.db.Create(msg).Error
}

func (s *MessageStore) DeleteByID(id int64) error {
	return s.db.Delete(&models.Message{}, "id=?", id).Error
}

func (s *MessageStore) GetByID(id int64) (*models.Message, error) {
	var ret models.Message
	if err := s.db.Model(&models.Message{}).Preload("Room.Members").Take(&ret, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &ret, nil
}

func (s *MessageStore) GetMessagesOfRoom(room *models.Room) ([]models.Message, error) {
	var ret []models.Message
	if err := s.db.Where("room_name=?", room.Name).Order("created_at").Find(&ret).Error; err != nil {
		return nil, err
	}
	return ret, nil
}
