package main

import (
	"fmt"

	"github.com/therecipe/qt/core"
)

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
	core.QAbstractListModel

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*Character             `property:"characters"`

	_ func()                     `constructor:"init"`
	_ func(character *Character) `slot:"addCharacter"`
}

func (m *LadderModel) init() {
	m.SetRoles(map[int]*core.QByteArray{
		Name:  core.NewQByteArray2("name", -1),
		Class: core.NewQByteArray2("class", -1),
		Level: core.NewQByteArray2("level", -1),
	})

	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectAddCharacter(m.addCharacter)
}

func (m *LadderModel) rowCount(*core.QModelIndex) int {
	fmt.Println("ROW COUNT", len(m.Characters()))
	return len(m.Characters())
}

func (m *LadderModel) columnCount(*core.QModelIndex) int {
	return 1
}

func (m *LadderModel) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *LadderModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.Characters()) {
		return core.NewQVariant()
	}

	var c = m.Characters()[len(m.Characters())-1-index.Row()]
	if c == nil {
		return core.NewQVariant()
	}

	fmt.Println("GOT PASSED NIL CHECK")
	fmt.Println(c)

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
	m.BeginInsertRows(core.NewQModelIndex(), 0, 0)
	m.SetCharacters(append(m.Characters(), c))
	m.EndInsertRows()
}

func init() {
	LadderModel_QRegisterMetaType()
	Character_QRegisterMetaType()
}
