package ws

import (
	"app/store"
	"sync"

	"github.com/sirupsen/logrus"
)

type WS struct {
	store      *store.Store
	log        *logrus.Logger
	sessions   map[string][]*WsClient // <user-id>
	register   chan *WsClient
	unregister chan *WsClient
	message    chan *WsMessage
	lock       *sync.Mutex
}

func NewWS(log *logrus.Logger, store *store.Store) *WS {
	log.SetReportCaller(false)
	return &WS{
		log:        log,
		store:      store,
		sessions:   make(map[string][]*WsClient, 0),
		register:   make(chan *WsClient),
		unregister: make(chan *WsClient),
		message:    make(chan *WsMessage),
		lock: &sync.Mutex{},
	}
}
