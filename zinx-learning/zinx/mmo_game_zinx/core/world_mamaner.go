package core

import "sync"

/**
当前游戏的世界总管理模块
*/
type WorldManager struct {
	// AOIManager 当前世界地图AOI的管理模块
	AoiMgr *AOIManager
	// 当前全部在线的Player集合
	Players map[int32]*Player
	// 保护Players集合的锁
	pLock sync.RWMutex
}

// 提供一个对外的世界管理模块的句柄（全局）
var WorldMgr *WorldManager

// 初始化方法
func init() {
	WorldMgr = &WorldManager{
		// AOIManager 当前世界地图AOI的管理模块
		AoiMgr: NewAOIManager(AOI_MIN_X, AOI_MAX_X, AOI_CNTS_X, AOI_MIN_Y, AOI_MAX_Y, AOI_CNTS_Y),
		// 初始化在线的Player集合
		Players: make(map[int32]*Player),
	}
}

//添加一个玩家
func (vm *WorldManager) AddPlayer(player *Player) {
	vm.pLock.Lock()
	vm.Players[player.Pid] = player
	defer vm.pLock.Unlock()

	// 将player添加到AOIManager中
	vm.AoiMgr.AddToGridByPos(int(player.Pid), player.X, player.Z)
}

//删除一个玩家
func (vm *WorldManager) RemovePlayerByPid(pid int32) {

	// 获取当前玩家
	player, ok := vm.Players[pid]
	if ok {
		vm.pLock.Lock()
		delete(vm.Players, pid)
		vm.pLock.Unlock()
		// 并且移除玩家的坐标信息
		vm.AoiMgr.RemoveFromGridByPos(int(pid), player.X, player.Z)
	}

}

//通过玩家ID查询Player对象
func (vm *WorldManager) GetPlayerByPid(pid int32) *Player {
	vm.pLock.RLock()
	defer vm.pLock.RUnlock()

	return vm.Players[pid]
}

//获取全部的在线玩家
func (vm *WorldManager) GetAllPlayers() []*Player {
	vm.pLock.RLock()
	defer vm.pLock.RUnlock()

	players := make([]*Player, 0)
	// 添加到切片中
	for _, p := range vm.Players {
		players = append(players, p)
	}
	return players
}
