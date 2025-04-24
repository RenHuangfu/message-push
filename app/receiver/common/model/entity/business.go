package entity

type CreateAPPParam struct {
	AppName string
	AppKey  string
}

type CreateAPPResult struct {
	AppID int32
}

type GetAPPParam struct {
	AppID int32
}

type GetAPPResult struct {
	AppID   int32
	AppName string
}

type CreateClientParam struct {
	ClientName string
	ClientKey  string
	AppID      int32
}

type CreateClientResult struct {
	ClientID int32
}

type GetClientParam struct {
	ClientID int32
}

type GetClientResult struct {
	ClientID   int32
	ClientName string
}

type GetClientListParam struct {
	AppID int32
}

type GetClientListResult struct {
	Clients []*GetClientResult
}
