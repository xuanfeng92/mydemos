package core

import "fmt"


// 定义一些AOI的边界值
const (
	AOI_MIN_X int = 85
	AOI_MAX_X int = 410
	AOI_CNTS_X int = 10
	AOI_MIN_Y int = 10
	AOI_MAX_Y int = 750
	AOI_CNTS_Y int = 20
)

/**
AOI区域管理模块
*/
type AOIManager struct {
	MinX int //整个区域左边界坐标

	MaxX int //整个区域右边界坐标

	CntsX int //X方向格子数量

	MinY int //整个区域上边界坐标

	MaxY int //整个区域下边界坐标

	CntsY int //Y方向格子数量

	grids map[int]*Grid //当前区域有哪些格子- map-key=格子ID value=格子指针对象
}

//初始化一个AOI管理区域模块
func NewAOIManager(minX, maxX, cntsx, minY, maxY, cntsy int) *AOIManager {
	aoiMgr := &AOIManager{
		MinX:  minX,
		MaxX:  maxX,
		CntsX: cntsx,
		MinY:  minY,
		MaxY:  maxY,
		CntsY: cntsy,
		grids: make(map[int]*Grid),
	}
	// 初始化ＡＯＩ区域的格子，并对这些格子进行编号和初始化
	for y := 0; y < cntsy; y++ {
		for x := 0; x < cntsx; x++ {
			// 计算格子ID，根据x,y编号
			//格子编号： id= idy*cntsX + idx
			gid := y*cntsx + x

			// 初始化格子
			aoiMgr.grids[gid] = NewGrid(gid,
				aoiMgr.MinX+x*aoiMgr.gridWidth(),
				aoiMgr.MinX+(x+1)*aoiMgr.gridWidth(),
				aoiMgr.MinY+y*aoiMgr.gridLength(),
				aoiMgr.MinY+(y+1)*aoiMgr.gridLength())
		}
	}
	return aoiMgr
}

//调试使用打印当前AOI模块
func (m *AOIManager) String() string {
	s := fmt.Sprintf("AOIManagr:\n[minX:%d], [maxX:%d], [cntsX:%d], [minY:%d], [maxY:%d], [cntsY:%d]\nGrids in AOI Manager:\n",
		m.MinX, m.MaxX, m.CntsX, m.MinY, m.MaxY, m.CntsY)
	for _, grid := range m.grids {
		s += fmt.Sprintln(grid)
	}

	return s
}

//根据格子的gID得到当前周边的九宫格信息
func (m *AOIManager) GetSurroundGridsByGid(gID int) (grids []*Grid) {
	//判断gID是否存在
	if _, ok := m.grids[gID]; !ok {
		return
	}

	//将当前gid添加到九宫格中
	grids = append(grids, m.grids[gID])

	// 根据gID, 得到格子所在的坐标
	x, y := gID%m.CntsX, gID/m.CntsX

	// 新建一个临时存储周围格子的数组
	surroundGid := make([]int, 0)

	// 新建8个方向向量: 左上: (-1, -1), 左中: (-1, 0), 左下: (-1,1), 中上: (0,-1), 中下: (0,1), 右上:(1, -1)
	// 右中: (1, 0), 右下: (1, 1), 分别将这8个方向的方向向量按顺序写入x, y的分量数组
	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	// 根据8个方向向量, 得到周围点的相对坐标, 挑选出没有越界的坐标, 将坐标转换为gid
	for i := 0; i < 8; i++ {
		newX := x + dx[i]
		newY := y + dy[i]

		if newX >= 0 && newX < m.CntsX && newY >= 0 && newY < m.CntsY {
			surroundGid = append(surroundGid, newY*m.CntsX+newX)
		}
	}

	// 根据没有越界的gid, 得到格子信息
	for _, gid := range surroundGid {
		grids = append(grids, m.grids[gid])
	}

	return
}

//通过横纵坐标获取对应的格子ID
func (m *AOIManager) GetGidByPos(x, y float32) int {
	gx := (int(x) - m.MinX) / m.gridWidth()
	gy := (int(y) - m.MinY) / m.gridLength()

	return gy*m.CntsX + gx
}

//通过横纵坐标得到周边九宫格内的全部PlayerIDs
//通过Player坐标得到当前Player周边九宫格内全部PlayerIDs
func (m *AOIManager) GetPidsByPos(x, y float32) (playerIDs []int) {
	//根据横纵坐标得到当前坐标属于哪个格子ID
	gID := m.GetGidByPos(x, y)

	//根据格子ID得到周边九宫格的信息
	grids := m.GetSurroundGridsByGid(gID)
	for _, v := range grids {
		playerIDs = append(playerIDs, v.GetPlayerIDs()...)
		//fmt.Printf("===> grid ID : %d, pids : %v  ====\n", v.GID, v.GetPlayerIDs())
	}

	return
}

//通过GID获取全部的PlayerID
func (m *AOIManager) GetPidsByGid(gID int) (playerIDs []int) {
	playerIDs = m.grids[gID].GetPlayerIDs()
	return
}

//添加一个PlayerID 到一个格子中
func (m *AOIManager) AddPidToGrid(pID, gID int) {
	m.grids[gID].Add(pID)
}

//移除一个格子中的playerID
func (m *AOIManager) RemovePidFromGrid(pID, gID int) {
	m.grids[gID].Remove(pID)
}

//通过坐标将Player添加到一个格子中
//通过横纵坐标添加一个Player到一个格子中
func (m *AOIManager) AddToGridByPos(pID int, x, y float32) {
	gID := m.GetGidByPos(x, y)
	grid := m.grids[gID]
	grid.Add(pID)
}

//通过坐标把一个Player从格子中删除
//通过横纵坐标把一个Player从对应的格子中删除
func (m *AOIManager) RemoveFromGridByPos(pID int, x, y float32) {
	gID := m.GetGidByPos(x, y)
	grid := m.grids[gID]
	grid.Remove(pID)
}

//得到每个格子在x轴方向的宽度
func (m *AOIManager) gridWidth() int {
	return (m.MaxX - m.MinX) / m.CntsX
}

//得到每个格子在x轴方向的长度
func (m *AOIManager) gridLength() int {
	return (m.MaxY - m.MinY) / m.CntsY
}
