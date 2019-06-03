package main

import "github.com/therecipe/qt/core"

// Model Roles
const (
	Character = int(core.Qt__UserRole) + 1<<iota
	Class
	Level
)

// LadderItem ...
type LadderItem struct {
	Character string
	Class     string
	Level     int
}

// LadderModel ...
type LadderModel struct {
	core.QAbstractTableModel

	_         func() `constructor:"init"`
	modelData []LadderItem
}

func (m *LadderModel) init() {
	m.modelData = []LadderItem{
		{
			Character: "meanski",
			Class:     "Paladin",
			Level:     99,
		},
	}

	m.ConnectRoleNames(m.roleNames)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectData(m.data)
}

func (m *LadderModel) roleNames() map[int]*core.QByteArray {
	return map[int]*core.QByteArray{
		Character: core.NewQByteArray2("Character", -1),
		Class:     core.NewQByteArray2("Class", -1),
		Level:     core.NewQByteArray2("Level", -1),
	}
}

func (m *LadderModel) rowCount(*core.QModelIndex) int {
	return len(m.modelData)
}

func (m *LadderModel) columnCount(*core.QModelIndex) int {
	return 3
}

func (m *LadderModel) data(index *core.QModelIndex, role int) *core.QVariant {
	item := m.modelData[index.Row()]
	switch role {
	case Character:
		return core.NewQVariant14(item.Character)
	case Class:
		return core.NewQVariant14(item.Class)
	case Level:
		return core.NewQVariant14("99")
	}
	return core.NewQVariant()
}

func init() {
	LadderModel_QRegisterMetaType()
	LadderItem_QRegisterMetaType()
}
