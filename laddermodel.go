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
	Name  string
	Class string
	Level string
}

// LadderModel ...
type LadderModel struct {
	core.QAbstractListModel

	_ map[int]*core.QByteArray `property:"roles"`
	_ func()                   `constructor:"init"`

	characters []Character
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
}

func (m *LadderModel) rowCount(*core.QModelIndex) int {
	fmt.Println("ROW COUNT", len(m.characters))
	return len(m.characters)
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

	if index.Row() >= len(m.characters) {
		return core.NewQVariant()
	}

	c := m.characters[0]

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
			return core.NewQVariant1(c.Level)
		}

	default:
		{
			return core.NewQVariant()
		}
	}
}

// AddCharacter ...
func (m *LadderModel) AddCharacter(c *Character) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.characters), len(m.characters))
	m.characters = append(m.characters, *c)
	m.EndInsertRows()
}

func init() {
	LadderModel_QRegisterMetaType()
}
