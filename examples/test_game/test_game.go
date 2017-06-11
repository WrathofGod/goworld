package main

import (
	"time"

	"github.com/xiaonanln/goTimer"
	"github.com/xiaonanln/goworld"
	"github.com/xiaonanln/goworld/components/game"
	"github.com/xiaonanln/goworld/gwlog"
)

func init() {

}

type gameDelegate struct {
	game.GameDelegate
}

func main() {
	goworld.SetSpaceDelegate(&SpaceDelegate{})
	goworld.RegisterEntity("OnlineService", &OnlineService{})
	goworld.RegisterEntity("Monster", &Monster{})
	goworld.RegisterEntity("Avatar", &Avatar{})

	goworld.Run(&gameDelegate{})
}

func (game gameDelegate) OnReady() {
	game.GameDelegate.OnReady()

	eids := goworld.ListEntityIDs("OnlineService")
	gwlog.Info("Found saved OnlineService ids: %v", eids)

	if len(eids) == 0 {
		goworld.CreateEntityAnywhere("OnlineService")
	} else {
		// already exists
		onlineServiceID := eids[0]
		goworld.LoadEntityAnywhere("OnlineService", onlineServiceID)
	}

	timer.AddCallback(time.Millisecond*1000, game.checkGameStarted)
}

func (game gameDelegate) checkGameStarted() {
	ok := game.isGameStarted()
	gwlog.Info("checkGameStarted: %v", ok)
	if ok {
		game.onGameStarted()
	} else {
		timer.AddCallback(time.Millisecond*1000, game.checkGameStarted)
	}
}

func (game gameDelegate) isGameStarted() bool {
	if len(goworld.GetServiceProviders("OnlineService")) == 0 {
		return false
	}
	return true
}

func (game gameDelegate) onGameStarted() {
	goworld.CreateSpaceAnywhere()
}