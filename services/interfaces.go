package services

type Logic interface {
	SayHello(userID string) (string, error)
}
