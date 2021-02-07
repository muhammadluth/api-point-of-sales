package handler

import "sync"

type ISetupRouter interface {
	Router(wg *sync.WaitGroup)
}
