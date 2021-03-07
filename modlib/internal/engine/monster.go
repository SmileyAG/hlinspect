package engine

import (
	"unsafe"
)

const (
	RouteMFToTargetent  = 1 << 0
	RouteMFToEnemy      = 1 << 1
	RouteMFToCover      = 1 << 2
	RouteMFToDetour     = 1 << 3
	RouteMFToPathcorner = 1 << 4
	RouteMFToNode       = 1 << 5
	RouteMFToLocation   = 1 << 6
	RouteMFIsGoal       = 1 << 7
	RouteMFDontSimplify = 1 << 8
)

// MonsterOffsets store offsets to class members
var MonsterOffsets monsterOffsets = monsterOffsets{
	Schedule:      0x178,
	ScheduleIndex: 0x17c,
	Cine:          0x290,
	Route:         0x180,
	RouteIndex:    0x204,
	WaypointSize:  0x10,
}

type monsterOffsets struct {
	// Look inside CBaseMonster::ChangeSchedule.
	Schedule      uintptr
	ScheduleIndex uintptr
	// Look inside CBaseMonster::GetScheduleOfType. Search for "Script failed for %s"
	Cine uintptr
	// Found from CBaseMonster::RouteNew
	Route        uintptr
	RouteIndex   uintptr
	WaypointSize uintptr
}

// Waypoint represents Waypoint_t
type Waypoint struct {
	ptr unsafe.Pointer
}

// MakeWaypoint creates a new instance of Waypoint
func MakeWaypoint(pointer unsafe.Pointer) Waypoint {
	return Waypoint{ptr: pointer}
}

// Location returns Waypoint_t::vecLocation
func (wp Waypoint) Location() [3]float32 {
	return *(*[3]float32)(unsafe.Pointer(uintptr(wp.ptr) + 0x0))
}

// Type returns Waypoint_t::iType
func (wp Waypoint) Type() int {
	return *(*int)(unsafe.Pointer(uintptr(wp.ptr) + 0xc))
}

// Monster represents CBaseMonster
type Monster struct {
	ptr unsafe.Pointer
}

// MakeMonster creates a new instance of Monster
func MakeMonster(pointer unsafe.Pointer) Monster {
	return Monster{ptr: pointer}
}

// Schedule returns CBaseMonster::m_pSchedule
func (monster Monster) Schedule() *Schedule {
	ptr := *(*unsafe.Pointer)(unsafe.Pointer(uintptr(monster.ptr) + MonsterOffsets.Schedule))
	if ptr == nil {
		return nil
	}
	schedule := MakeSchedule(uintptr(ptr))
	return &schedule
}

// ScheduleIndex returns CBaseMonster::m_iScheduleIndex
func (monster Monster) ScheduleIndex() int {
	return int(*(*int32)(unsafe.Pointer(uintptr(monster.ptr) + MonsterOffsets.ScheduleIndex)))
}

// Cine returns CBaseMonster::m_pCine
func (monster Monster) Cine() Cine {
	return MakeCine(*(*unsafe.Pointer)(unsafe.Pointer(uintptr(monster.ptr) + MonsterOffsets.Cine)))
}

// Routes returns the array CBaseMontser::m_Route with ROUTE_SIZE == 8
func (monster Monster) Routes() [8]Waypoint {
	base := unsafe.Pointer(uintptr(monster.ptr) + MonsterOffsets.Route)
	waypoints := [8]Waypoint{}
	for i := range waypoints {
		waypoints[i] = MakeWaypoint(unsafe.Pointer(uintptr(base) + uintptr(i)*MonsterOffsets.WaypointSize))
	}
	return waypoints
}

// RouteIndex returns CBaseMontser::m_iRouteIndex
func (monster Monster) RouteIndex() int {
	return *(*int)(unsafe.Pointer(uintptr(monster.ptr) + MonsterOffsets.RouteIndex))
}
