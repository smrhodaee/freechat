package configs

import (
	"app/models"
	"fmt"
	"strings"
)

func GetDirectRoomName(reqUser *models.User, user *models.User) string {
	first, second := reqUser.Username, user.Username
	if strings.Compare(user.Username, reqUser.Username) == -1 {
		first, second = user.Username, reqUser.Username
	}
	return fmt.Sprintf("direct.%s.%s", first, second)
}

func GetMemberRole(r models.MemberRule) string {
	switch r {
	case models.MemberRoleAdmin:
		return "ADMIN"
	case models.MemberRoleOwner:
		return "OWNER"
	}
	return "NORMAL"
}

func GetMessageType(t models.MessageType) string {
	switch t {
	case models.MessageTypeImage:
		return "IMAGE"
	}
	return "TEXT"
}
