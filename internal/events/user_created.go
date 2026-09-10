package events

const UserCreatedEventName = "user.created"

type UserCreatedEvent struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"name"`
	Login    string `json:"login"`
}

func NewUserCreatedEvent(userID int, name, login string) *UserCreatedEvent {
	return &UserCreatedEvent{
		UserID:   userID,
		UserName: name,
		Login:    login,
	}
}

func (e *UserCreatedEvent) Name() string {
	return "user.created"
}
