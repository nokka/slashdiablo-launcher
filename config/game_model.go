package config

import (
	"github.com/therecipe/qt/core"
)

// Model Roles.
const (
	ID = int(core.Qt__UserRole) + 1<<iota
	Location
	Instances
	OverrideBHCfg
	Flags
	HDVersion
	MaphackVersion
)

// GameModel represents a Diablo game.
type GameModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*Game                  `property:"games"`

	_ func(*Game) `slot:"addGame"`
}

func (m *GameModel) init() {
	m.SetRoles(map[int]*core.QByteArray{
		ID:             core.NewQByteArray2("id", -1),
		Location:       core.NewQByteArray2("location", -1),
		Instances:      core.NewQByteArray2("instances", -1),
		OverrideBHCfg:  core.NewQByteArray2("override_bh_config", -1),
		Flags:          core.NewQByteArray2("flags", -1),
		HDVersion:      core.NewQByteArray2("hd_version", -1),
		MaphackVersion: core.NewQByteArray2("maphack_version", -1),
	})

	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectAddGame(m.addGame)
}

func (m *GameModel) rowCount(*core.QModelIndex) int {
	return len(m.Games())
}

func (m *GameModel) columnCount(*core.QModelIndex) int {
	return 1
}

func (m *GameModel) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *GameModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.Games()) {
		return core.NewQVariant()
	}

	item := m.Games()[index.Row()]

	switch role {
	case ID:
		return core.NewQVariant1(item.ID)
	case Location:
		return core.NewQVariant1(item.Location)
	case Instances:
		return core.NewQVariant1(item.Instances)
	case OverrideBHCfg:
		return core.NewQVariant1(item.OverrideBHCfg)
	case Flags:
		return core.NewQVariant1(item.Flags)
	case HDVersion:
		return core.NewQVariant1(item.HDVersion)
	case MaphackVersion:
		return core.NewQVariant1(item.MaphackVersion)
	default:
		return core.NewQVariant()
	}
}

// addGame adds a game to the model.
func (m *GameModel) addGame(g *Game) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.Games()), len(m.Games()))
	m.SetGames(append(m.Games(), g))
	m.EndInsertRows()
}

// updateGame will notify the UI of the updated model item.
func (m *GameModel) updateGame(index int) {
	var fIndex = m.Index(0, 0, core.NewQModelIndex())
	var lIndex = m.Index(index, 0, core.NewQModelIndex())
	m.DataChanged(fIndex, lIndex, []int{Location, Instances, OverrideBHCfg, Flags, HDVersion, MaphackVersion})
}

func (m *GameModel) removeGame(index int) {
	m.BeginRemoveRows(core.NewQModelIndex(), index, index)
	m.SetGames(append(m.Games()[:index], m.Games()[index+1:]...))
	m.EndRemoveRows()
}

func init() {
	GameModel_QRegisterMetaType()
	Game_QRegisterMetaType()
}
