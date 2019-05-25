import QtQuick 2.12
import QtQuick.Controls 2.1
import QtQuick.Layouts 1.3
import QtQuick.Dialogs 1.3

Popup {
    id: gamePathDialog
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
                    id: pathDialogInput
                    width: fileDialogBox.width * 0.90
                    readOnly: true
                    text: ""

                    background: Rectangle {
                        radius: 3
                        color: "#1d1924"
                    }
                }

                Button {
                    text: "Open"
                    width: fileDialogBox.width * 0.10
                    onClicked: fileDialog.open()

                    background: Rectangle {
                        color: "#1d1924"
                    }
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
                        console.log("save" + pathDialogInput.text)
                        var success = settings.setGamePaths(pathDialogInput.text, "")
                        if (success) {
                            gamePathDialog.close()
                        }
                    }
                }
            }
        }

        FileDialog {
            id: fileDialog
            selectFolder: true
            title: "Please choose a file"
            folder: shortcuts.home
            defaultSuffix: "derp"
            onAccepted: {
                var path = fileDialog.fileUrl.toString()
                path = path.replace(/^(file:\/{2})/,"")
                pathDialogInput.text = path
            }
            onRejected: {
                console.log("Canceled")
            }
        }
    }
}