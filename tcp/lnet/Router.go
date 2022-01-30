package lnet

import "lai_zinx/tcp/linterface"

type BaseRouter struct{}

//这里之所以BaseRouter的方法都为空，
// 是因为有的Router不希望有PreHandle或PostHandle
// 所以Router全部继承BaseRouter的好处是，不需要实现PreHandle和PostHandle也可以实例化

//PreHandle -
func (br *BaseRouter) PreHandle(req linterface.LRequest) {}

//Handle -
func (br *BaseRouter) Handle(req linterface.LRequest) {}

//PostHandle -
func (br *BaseRouter) PostHandle(req linterface.LRequest) {}
