package d2

import (
	"github.com/therecipe/qt/core"
)

// File ...
type File struct {
	core.QObject
	Name       string
	D2Path     string
	RemoteCRC  string
	LocalCRC   string
	FileAction string
}

// Model Roles.
const (
	Name = int(core.Qt__UserRole) + 1<<iota
	D2Path
	RemoteCRC
	LocalCRC
	FileAction
)

// FileModel represents a patch file.
type FileModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*File                  `property:"files"`

	_ func(*File) `slot:"addFile"`
	_ func()      `slot:"clear"`
}

func (m *FileModel) init() {
	m.SetRoles(map[int]*core.QByteArray{
		Name:       core.NewQByteArray2("name", -1),
		D2Path:     core.NewQByteArray2("d2Path", -1),
		RemoteCRC:  core.NewQByteArray2("remoteCRC", -1),
		LocalCRC:   core.NewQByteArray2("localCRC", -1),
		FileAction: core.NewQByteArray2("fileAction", -1),
	})

	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectAddFile(m.addFile)
	m.ConnectClear(m.clear)
}

func (m *FileModel) rowCount(*core.QModelIndex) int {
	return len(m.Files())
}

func (m *FileModel) columnCount(*core.QModelIndex) int {
	return 1
}

func (m *FileModel) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *FileModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.Files()) {
		return core.NewQVariant()
	}

	item := m.Files()[index.Row()]

	switch role {
	case Name:
		return core.NewQVariant1(item.Name)
	case D2Path:
		return core.NewQVariant1(item.D2Path)
	case RemoteCRC:
		return core.NewQVariant1(item.RemoteCRC)
	case LocalCRC:
		return core.NewQVariant1(item.LocalCRC)
	case FileAction:
		return core.NewQVariant1(item.FileAction)
	default:
		return core.NewQVariant()
	}
}

// addFile adds a file to the model.
func (m *FileModel) addFile(g *File) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.Files()), len(m.Files()))
	m.SetFiles(append(m.Files(), g))
	m.EndInsertRows()
}

// updateFile will notify the UI of the updated model item.
func (m *FileModel) updateFile(index int) {
	var fIndex = m.Index(0, 0, core.NewQModelIndex())
	var lIndex = m.Index(index, 0, core.NewQModelIndex())
	m.DataChanged(fIndex, lIndex, []int{Name, D2Path, LocalCRC, RemoteCRC, FileAction})
}

// removeFile will remove a file from the model.
func (m *FileModel) removeFile(index int) {
	m.BeginRemoveRows(core.NewQModelIndex(), index, index)
	m.SetFiles(append(m.Files()[:index], m.Files()[index+1:]...))
	m.EndRemoveRows()
}

func (m *FileModel) clear() {
	m.BeginResetModel()
	m.SetFiles([]*File{})
	m.EndResetModel()
}

func init() {
	FileModel_QRegisterMetaType()
	File_QRegisterMetaType()
}
