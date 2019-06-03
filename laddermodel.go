package main

import "github.com/therecipe/qt/core"

// Model Roles
const (
	Name = int(core.Qt__UserRole) + 1<<iota
	Class
	Level
)

// Character ...
type Character struct {
	core.QObject

	Name  string
	Class string
	Level int
}

// LadderModel ...
type LadderModel struct {
	core.QAbstractTableModel

	_ []*Character `property:"characters"`

	_ func()                     `constructor:"init"`
	_ func(character *Character) `slot:"addCharacter"`
}

func (m *LadderModel) init() {
	m.ConnectRoleNames(m.roleNames)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectData(m.data)
}

func (m *LadderModel) roleNames() map[int]*core.QByteArray {
	return map[int]*core.QByteArray{
		Name:  core.NewQByteArray2("Character", -1),
		Class: core.NewQByteArray2("Class", -1),
		Level: core.NewQByteArray2("Level", -1),
	}
}

func (m *LadderModel) rowCount(*core.QModelIndex) int {
	return len(m.Characters())
}

func (m *LadderModel) columnCount(*core.QModelIndex) int {
	return 3
}

func (m *LadderModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.Characters()) {
		return core.NewQVariant()
	}

	var c = m.Characters()[index.Row()]

	switch role {
	case Name:
		{
			return core.NewQVariant1(c.Name)
		}

	case Class:
		{
			return core.NewQVariant1(c.Class)
		}
	case Level:
		{
			return core.NewQVariant1("99")
		}

	default:
		{
			return core.NewQVariant()
		}
	}
}

func (m *LadderModel) addCharacter(c *Character) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.Characters()), len(m.Characters()))
	m.SetCharacters(append(m.Characters(), c))
	m.EndInsertRows()
}

func init() {
	LadderModel_QRegisterMetaType()
	Character_QRegisterMetaType()
}
