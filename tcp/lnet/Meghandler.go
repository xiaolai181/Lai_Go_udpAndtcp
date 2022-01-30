package lnet

import (
	"fmt"
	"lai_zinx/tcp/linterface"
	"strconv"
)

type Msghandle struct {
	Apis           map[uint32]linterface.LRouter
	WorkerPoolSize uint32
	TaskQueue      []chan linterface.LRequest
}

func NewMsgHandle() *Msghandle {
	return &Msghandle{
		Apis:           make(map[uint32]linterface.LRouter),
		WorkerPoolSize: uint32(WorkPoolSize),
		TaskQueue:      make([]chan linterface.LRequest, WorkPoolSize),
	}

}

func (mh *Msghandle) SendMsTaskQueue(req linterface.LRequest) {
	workerID := req.GetConnection().GetConnID() & mh.WorkerPoolSize
	mh.TaskQueue[workerID] <- req
}

func (mh *Msghandle) DoMsgHandler(req linterface.LRequest) {
	handler, ok := mh.Apis[req.GetMsgID()]
	if !ok {
		fmt.Println("api msgID = ", req.GetMsgID(), " is not FOUND!")
		return
	}

	//执行对应处理方法
	handler.PreHandle(req)
	handler.Handle(req)
	handler.PostHandle(req)
}

func (mh *Msghandle) AddRouter(msgID uint32, router linterface.LRouter) {
	if _, ok := mh.Apis[msgID]; ok {
		panic("repeated api , msgID = " + strconv.Itoa(int(msgID)))
	}
	mh.Apis[msgID] = router
	fmt.Println("Add api msgID = ", msgID)
}
func (mh *Msghandle) StartOneWorker(workerID int, taskQueue chan linterface.LRequest) {
	fmt.Println("Worker ID = ", workerID, " is started.")
	//不断的等待队列中的消息
	for {
		select {
		//有消息则取出队列的Request，并执行绑定的业务方法
		case request := <-taskQueue:
			mh.DoMsgHandler(request)
		}
	}
}

func (mh *Msghandle) StartWorkerPool() {
	//遍历需要启动worker的数量，依此启动
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		//一个worker被启动
		//给当前worker对应的任务队列开辟空间
		mh.TaskQueue[i] = make(chan linterface.LRequest, WorkPoolSize)
		//启动当前Worker，阻塞的等待对应的任务队列是否有消息传递进来
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}
