package store

import (
	"app/models"
	"errors"

	"gorm.io/gorm"
)

type RoomStore struct {
	db *gorm.DB
}

func NewRoomStore(db *gorm.DB) *RoomStore {
	return &RoomStore{
		db: db,
	}
}

//TODO: impl cache system and map for accociations

func (s *RoomStore) GetRoomsOfUser(user *models.User) ([]models.Room, error) {
	if user == nil {
		return nil, nil
	}
	var roomNames []string
	if err := s.db.Model(&models.RoomMember{}).
		Select("room_name").Where("username=?", user.Username).
		Order("room_name ASC").Find(&roomNames).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rooms []models.Room
	if err := s.db.Find(&rooms, "name IN (?)", roomNames).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rooms, nil
}

func (s *RoomStore) CreateGroup(name, title string, ownerUser *models.User, users []models.User) error {
	members := []models.RoomMember{
		{
			RoomName: name,
			Username: ownerUser.Username,
			Role:     models.MemberRoleOwner,
		},
	}
	for _, user := range users {
		members = append(members, models.RoomMember{
			RoomName: name,
			Username: user.Username,
			Role:     models.MemberRoleNormal,
		})
	}
	return s.db.Create(&models.Room{
		Name:     name,
		Title:    title,
		Type:     models.RoomTypeGroup,
		IsActive: true,
		Members:  members,
	}).Error
}

func (s *RoomStore) CreateDirect(name string, reqUser *models.User, user *models.User) error {
	members := []models.RoomMember{
		{
			RoomName: name,
			Username: reqUser.Username,
			Role:     models.MemberRoleOwner,
		},
		{
			RoomName: name,
			Username: user.Username,
			Role:     models.MemberRoleOwner,
		},
	}
	return s.db.Create(&models.Room{
		Name:     name,
		Title:    "",
		Type:     models.RoomTypeDirect,
		IsActive: true,
		Members:  members,
	}).Error
}

func (s *RoomStore) GetByName(name string) (*models.Room, error) {
	var ret models.Room
	if err := s.db.Preload("Members").Where("name=?", name).Take(&ret).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &ret, nil
}
