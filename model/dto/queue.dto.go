package dto

type QueueEvent struct {
	UserID            string  `json:"user_id"`
	QueueNumber       int     `json:"queue_number"`
	EstimatedTime     float64 `json:"estimated_time"`
	PercetageProgress float64 `json:"percentage_progress"`
	IsFinished        bool    `json:"is_finished"`
	Message           string  `json:"message,omitempty"`
}

type RegisterQueueResponse struct {
	WaitingRoomToken string `json:"waiting_room_token"`
	MainRoomToken    string `json:"main_room_token"`
}
