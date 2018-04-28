package sence

import (
	"github.com/laonsx/statemachines/westworld"
	"sync"
)

var EntityManager sync.Map

func RegistEntity(entity westworld.BaseEntity)  {

	EntityManager.Store(entity.GetId(),entity)
}

func GetEntityById(id int) westworld.BaseEntity{

	if entity,ok:=EntityManager.Load(id);ok{
		return entity.(westworld.BaseEntity)
	}

	return nil
}

