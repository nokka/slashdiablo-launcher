package ladder

import (
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
	m.SetRoles(map[int]*core.QByteArray{
		Rank:  core.NewQByteArray2("rank", -1),
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

func (m *TopLadderModel) rowCount(*core.QModelIndex) int {
	return len(m.Characters())
}

func (m *TopLadderModel) columnCount(*core.QModelIndex) int {
	return 1
}

func (m *TopLadderModel) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *TopLadderModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.Characters()) {
		return core.NewQVariant()
	}

	item := m.Characters()[index.Row()]

	switch role {
	case Rank:
		return core.NewQVariant1(item.Rank)
	case Name:
		return core.NewQVariant1(item.Name)
	case Class:
		return core.NewQVariant1(item.Class[:3])
	case Level:
		return core.NewQVariant1(item.Level)
	default:
		return core.NewQVariant()
	}
}

// AddCharacter adds a character to the model.
func (m *TopLadderModel) addCharacter(c *Character) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.Characters()), len(m.Characters()))
	m.SetCharacters(append(m.Characters(), c))
	m.EndInsertRows()
}

func (m *TopLadderModel) clear() {
	m.BeginResetModel()
	m.SetCharacters([]*Character{})
	m.EndResetModel()
}

func init() {
	TopLadderModel_QRegisterMetaType()
	Character_QRegisterMetaType()
}
