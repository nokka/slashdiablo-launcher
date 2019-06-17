package ladder

import (
	"fmt"

	"github.com/therecipe/qt/core"
)

// Model Roles
const (
	Rank = int(core.Qt__UserRole) + 1<<iota
	Name
	Class
	Level
)

// TopLadderModel ...
type TopLadderModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*Character             `property:"characters"`

	_ func(*Character) `slot:"addCharacter"`
}

func (m *TopLadderModel) init() {
	fmt.Println("INIT")
	m.SetRoles(map[int]*core.QByteArray{
		Rank:  core.NewQByteArray2("rank", -1),
		Name:  core.NewQByteArray2("name", -1),
		Class: core.NewQByteArray2("class", -1),
		Level: core.NewQByteArray2("level", -1),
	})

	fmt.Println("CONNECTING")
	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectAddCharacter(m.addCharacter)
}

func (m *TopLadderModel) rowCount(*core.QModelIndex) int {
	fmt.Println("ROW COUNT")
	return len(m.Characters())
}

func (m *TopLadderModel) columnCount(*core.QModelIndex) int {
	fmt.Println("COLUMN COUNT")
	return 1
}

func (m *TopLadderModel) roleNames() map[int]*core.QByteArray {
	fmt.Println("ROLE NAMES")
	return m.Roles()
}

func (m *TopLadderModel) data(index *core.QModelIndex, role int) *core.QVariant {
	fmt.Println("DATA CALLED")
	if !index.IsValid() {
		return core.NewQVariant()
	}

	fmt.Println("ADDING ROW", index.Row())
	chars := m.Characters()

	fmt.Println(len(chars))
	fmt.Println(chars[0])

	item := Character{
		Name:  "test",
		Class: "pal",
		Level: 99,
	}

	switch role {
	case Rank:
		{
			return core.NewQVariant1(item.Rank)
		}
	case Name:
		{
			return core.NewQVariant1(item.Name)
		}

	case Class:
		{
			return core.NewQVariant1(item.Class[:3])
		}
	case Level:
		{
			return core.NewQVariant1(item.Level)
		}

	default:
		{
			return core.NewQVariant()
		}
	}
}

// AddCharacter adds a character to the model.
func (m *TopLadderModel) addCharacter(c *Character) {
	fmt.Println("ADD CHARACTER CALLED")
	m.BeginInsertRows(core.NewQModelIndex(), 0, 0)
	m.SetCharacters(append(m.Characters(), c))
	m.EndInsertRows()
}

func init() {
	TopLadderModel_QRegisterMetaType()
	Character_QRegisterMetaType()
}
