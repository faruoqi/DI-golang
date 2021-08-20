package ds

type DataStore interface {
	UserNameForID(userID string) (string, bool)
}
