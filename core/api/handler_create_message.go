package api

import (
	"app/configs"
	"app/constants"
	"app/models"
	"app/permissions"
	"app/router"
	"app/store"
	"app/validations"
	"app/ws"
	"mime/multipart"
	"slices"

	"github.com/gofiber/fiber/v2"
)

type CreateMessageRequest struct {
	Text     string                `form:"text" json:"text"`
	RoomName string                `form:"room_name" json:"room_name"`
	Room     *models.Room          `form:"-" json:"-"`
	AuthInfo *models.ReqInfo       `form:"-" json:"-"`
	File     *multipart.FileHeader `form:"-" json:"-"`
	Type     models.MessageType    `form:"-" json:"-"`
}

func (r *CreateMessageRequest) ValidateAndNormalize(s *store.Store, c *fiber.Ctx) error {
	if len(r.Text) != 0 {
		if !validations.IsValidMessageText(r.Text) {
			return router.BadRequestError(constants.InvalidData)
		}
	}
	room, err := s.Room.GetByName(r.RoomName)
	if err != nil {
		return err
	}
	if room == nil {
		return router.NotFoundError()
	}
	r.Type = models.MessageTypeText
	r.Room = room
	r.AuthInfo = router.GetReqInfo(c)
	if !permissions.CanWriteIntoRoom(r.AuthInfo.User, room) {
		return router.PermissionError()
	}
	file, err := c.FormFile("file")
	if err == nil {
		if !validations.IsValidFilename(file.Filename) {
			return router.BadRequestError(constants.InvalidData)
		}
		contentType := file.Header.Get("Content-Type")
		allowedTypes := []string{"image/jpeg", "image/png"}
		if !slices.Contains(allowedTypes, contentType) {
			return router.BadRequestError(constants.InvalidFile)
		}
		r.Type = models.MessageTypeImage
		r.File = file
	}
	return nil
}

type CreateMessageHandler struct {
	s  *store.Store
	ws *ws.WS
}

func (h *CreateMessageHandler) V1(c *fiber.Ctx) error {
	var req CreateMessageRequest
	if err := router.GetRequest(&req, h.s, c); err != nil {
		return err
	}
	var vID int64 = 0
	if req.File != nil {
		path := configs.UploadFilePath(req.File)
		if err := c.SaveFile(req.File, path); err != nil {
			return err
		}
		file := &models.File{
			OwnerUsername: req.AuthInfo.Username,
			Title:         req.File.Filename,
			Path:          path,
			Size:          req.File.Size,
			MimeType:      req.File.Header.Get("Content-Type"),
			IsActive:      true,
		}
		if err := h.s.File.Create(file); err != nil {
			return err
		}
		vID = file.ID
	}
	msg := &models.Message{
		RoomName: req.Room.Name,
		Username: req.AuthInfo.Username,
		Text:     req.Text,
		ValueID:  vID,
		Type:     req.Type,
	}
	if err := h.s.Message.Create(msg); err != nil {
		return err
	}
	var err error
	msg.Room, err = h.s.Room.GetByName(msg.RoomName)
	if err != nil {
		return err
	}
	h.ws.Broadcast(ws.NewCreateMessageResponse(msg), msg.Room.Members.GetUsernames())
	return c.JSON(router.NewSuccessGlobalResponse(
		constants.ChatSubmitSuccessfuly,
		ChatResponse{}.Clone(*msg),
	))
}
