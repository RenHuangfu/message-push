package entity

type CreateNameParam struct {
	Name string
}

type CreateNameResult struct {
	Message string
}

type SearchNameParam struct {
	Id int32
}

type SearchNameResult struct {
	Exist bool
	Name  string
}
