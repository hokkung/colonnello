package entity

type User struct {
	Name string
}

func (e User) GetID() interface{} {
	return ""
}
