package SessionManager

import (
	"net"
	"sync"
)

type ClientManage struct {
	clients map[string]net.Conn
	mutex   sync.RWMutex
}

type ClientManager interface {
	AddClient(conn net.Conn, clientId string) error
	DeleteClient(clientId string) error
}

func NewClientManager() ClientManager {
	return &ClientManage{
		clients: make(map[string]net.Conn),
	}
}

func (cm *ClientManage) AddClient(conn net.Conn, clientId string) error {
	cm.mutex.Lock()
	cm.clients[clientId] = conn
	defer cm.mutex.Unlock()

	return nil
}

func (cm *ClientManage) DeleteClient(clientId string) error {
	cm.mutex.Lock()
	delete(cm.clients, clientId)
	defer cm.mutex.Unlock()

	return nil
}
