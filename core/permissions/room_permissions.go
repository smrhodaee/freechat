package permissions

import (
	"app/models"
	"app/tools"
)

func GetMembership(reqUser *models.User, room *models.Room) *models.RoomMember {
	return tools.FindElement(room.Members, func(value *models.RoomMember) bool {
		return reqUser.Username == value.Username
	})
}

func CanAddMemberToRoom(reqUser *models.User, room *models.Room) bool {
	if reqUser == nil || room == nil {
		return false
	}
	if reqUser.Role == models.UserRoleSuperAdmin {
		return true
	}
	switch room.Type {
	case models.RoomTypeStorage:
		fallthrough
	case models.RoomTypeDirect:
		return false
	case models.RoomTypeGroup:
		fallthrough
	case models.RoomTypeChannel:
		membership := GetMembership(reqUser, room)
		return membership != nil && membership.Role == models.MemberRoleAdmin || membership.Role == models.MemberRoleOwner
	}
	return false
}

func CanReadFromRoom(reqUser *models.User, room *models.Room) bool {
	if reqUser == nil || room == nil {
		return false
	}
	if reqUser.Role == models.UserRoleSuperAdmin {
		return true
	}
	return GetMembership(reqUser, room) != nil
}

func CanWriteIntoRoom(reqUser *models.User, room *models.Room) bool {
	if reqUser == nil || room == nil {
		return false
	}
	if reqUser.Role == models.UserRoleSuperAdmin {
		return true
	}
	membership := GetMembership(reqUser, room)
	if membership == nil {
		return false
	}
	switch room.Type {
	case models.RoomTypeStorage:
		fallthrough
	case models.RoomTypeDirect:
		fallthrough
	case models.RoomTypeGroup:
		return true
	case models.RoomTypeChannel:
		return membership.Role == models.MemberRoleAdmin || membership.Role == models.MemberRoleOwner
	}
	return false
}

func CanDeleteChat(reqUser *models.User, msg *models.Message) bool {
	if reqUser == nil || msg == nil || len(msg.Room.Members) == 0 {
		return false
	}
	if reqUser.Role == models.UserRoleSuperAdmin {
		return true
	}
	if GetMembership(reqUser, msg.Room) == nil {
		return false
	}
	return reqUser.Username == msg.Username
}

func CanBroadCast(reqUser *models.User, msg *models.Message) bool {
	return reqUser.Username == msg.Username
}
