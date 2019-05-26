import QtQuick 2.12
import QtQuick.Controls 2.1
import QtQuick.Layouts 1.3
import QtQuick.Dialogs 1.3

Popup {
    id: settingsDialog
    modal: true
    focus: true
    
    background: Rectangle {
        anchors.fill: parent
        color: "#100b17"
    }

    ColumnLayout {
        anchors.fill: parent

        Column {
            id: fileDialogBox
            width: (mainWindow.width/2)
            Layout.alignment: Qt.AlignHCenter
            spacing: 5

            Column {
                Label {
                    text: "Add gamepath for Diablo II"
                    font.pointSize: 16
                    color: "#ffffff"
                    font.bold: true
                }
            }

            Row {
                TextField {
                    id: d2pathInput
                    width: fileDialogBox.width * 0.80
                    readOnly: true
                    text: ""

                    background: Rectangle {
                        radius: 3
                        color: "#1d1924"
                    }
                }

                Button {
                    id: chooseD2Path
                    width: fileDialogBox.width * 0.20
                    text: "CHOOSE"

                    contentItem: Text {
                        text: chooseD2Path.text
                        font: chooseD2Path.font
                        color: "#ffffff"
                        horizontalAlignment: Text.AlignHCenter
                        verticalAlignment: Text.AlignVCenter
                        elide: Text.ElideRight
                    }

                    background: Rectangle {
                        color: chooseD2Path.down ? "#49c5f2" : "#0b86ba"
                    }

                    onClicked: d2PathDialog.open()
                }
            }

            Column {
                Button {
                    id: saveGamePath
                    width: fileDialogBox.width
                    text: "SAVE"
                    
                    contentItem: Text {
                        text: saveGamePath.text
                        font: saveGamePath.font
                        color: "#ffffff"
                        horizontalAlignment: Text.AlignHCenter
                        verticalAlignment: Text.AlignVCenter
                        elide: Text.ElideRight
                    }

                    background: Rectangle {
                        color: saveGamePath.down ? "#49c5f2" : "#0b86ba"
                        radius: 2
                    }
                    
                    onClicked: {
                        console.log("save" + d2pathInput.text)
                        var success = settings.setGamePaths(d2pathInput.text, "")
                        if (success) {
                            settingsDialog.close()
                        }
                    }
                }
            }
        }

        FileDialog {
            id: d2PathDialog
            selectFolder: true
            title: "Please choose a file"
            folder: shortcuts.home
            onAccepted: {
                var path = d2PathDialog.fileUrl.toString()
                path = path.replace(/^(file:\/{2})/,"")
                d2pathInput.text = path
            }
            onRejected: {
                console.log("Canceled")
            }
        }
    }
}