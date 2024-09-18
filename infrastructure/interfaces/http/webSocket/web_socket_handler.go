package webSocket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/oscargh945/go-Chat/infrastructure/webSocket/models"
	"net/http"
)

type Handler struct {
	hub *models.Hub
}

func NewWebSocketHandler(hub *models.Hub) *Handler {
	return &Handler{
		hub: hub,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var request *models.RoomRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.hub.Rooms[request.ID] = &models.Room{
		ID:      request.ID,
		Name:    request.Name,
		Clients: make(map[string]*models.Client),
	}
	c.JSON(http.StatusOK, request)
}

func (h *Handler) JoinRoom(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	roomID := c.Param("roomID")
	clientID := c.Query("clientID")
	username := c.Query("username")

	client := &models.Client{
		Socket:   conn,
		Message:  make(chan *models.Message, 10),
		ID:       clientID,
		RoomID:   roomID,
		UserName: username,
	}

	message := &models.Message{
		Content:  "A new user has joined the room",
		RoomID:   roomID,
		UserName: username,
	}

	h.hub.Register <- client
	h.hub.Broadcast <- message

	go client.Write()
	client.Read(h.hub)
}

func (h *Handler) GetRooms(c *gin.Context) {
	rooms := make([]models.RoomResponse, 0)

	for _, room := range h.hub.Rooms {
		rooms = append(rooms, models.RoomResponse{
			ID:   room.ID,
			Name: room.Name,
		})
	}
	c.JSON(http.StatusOK, rooms)
}

func (h *Handler) GetClients(c *gin.Context) {
	var clients []models.ClientResponse
	roomID := c.Param("roomID")

	if _, ok := h.hub.Rooms[roomID]; !ok {
		clients = make([]models.ClientResponse, 0)
		c.JSON(http.StatusOK, clients)
	}

	for _, c := range h.hub.Rooms[roomID].Clients {
		clients = append(clients, models.ClientResponse{
			ID:       c.ID,
			UserName: c.UserName,
		})
	}
	c.JSON(http.StatusOK, clients)
}
