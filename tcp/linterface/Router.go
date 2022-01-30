package linterface

type LRouter interface {
	PreHandle(request LRequest)
	Handle(request LRequest)
	PostHandle(request LRequest)
}
