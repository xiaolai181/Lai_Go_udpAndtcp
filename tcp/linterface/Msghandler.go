package linterface

type LMsgHandle interface {
	DoMsgHandler(request LRequest)
	AddRouter(msgId uint32, router LRouter)
	StartWorkerPool()
	SendMsTaskQueue(request LRequest)
}
