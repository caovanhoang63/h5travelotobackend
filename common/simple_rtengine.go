package common

type SimpleRealtimeEngine interface {
	EmitToRoom(room string, key string, data interface{}) error
	EmitToUser(userId int, key string, data interface{}) error
}
