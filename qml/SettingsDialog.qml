import QtQuick 2.12
import QtQuick.Controls 2.1
import QtQuick.Layouts 1.3
import QtQuick.Dialogs 1.3

Popup {
    id: settingsDialog
    property var usingHD: false
    
    modal: true
    focus: true
    
    background: Rectangle {
        anchors.fill: parent
        color: "#100b17"
    }

    ColumnLayout {
        anchors.fill: parent

        Column {
            Heading {
                text: "SETTINGS"
                font.pointSize: 20
            }

            id: fileDialogBox
            width: (mainWindow.width/2)
            Layout.alignment: Qt.AlignTop | Qt.AlignHCenter
            Layout.topMargin: 100
            spacing: 5

            Column {
                topPadding: 15

                Label {
                    text: "SET DIABLO II DIRECTORY"
                    font.pointSize: 13
                    font.family: d2Font.name
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

                DefaultButton {
                    id: chooseD2Path
                    text: "CHOOSE"
                    width: fileDialogBox.width * 0.20
                    onClicked: d2PathDialog.open()
                }
            }

            Column {
                topPadding: 15
                Heading {
                    text: "NUMBER OF D2 INSTANCES TO LAUNCH"
                }
                
                ComboBox {
                    id: d2Instances
                    model: [ 1, 2, 3, 4 ]
                    height: 30
                    width: 60

                    background: Rectangle {
                        color: "#1d1924"
                        border.color: "#f0681f"
                        radius: height/2
                    }
                }
            }

            Row {
                topPadding: 15
                width: parent.width
                
                layoutDirection: Qt.RightToLeft

                Item {
                    width: parent.width * 0.20
                    height: 35

                    Switch {
                        anchors.right: parent.right
                        anchors.verticalCenter: parent.verticalCenter
                        checked: false

                        onClicked: {
                            usingHD = !usingHD
                        }
                    }
                }

                Item {
                    width: parent.width * 0.80
                    height: 35
                     Layout.alignment: Qt.AlignLeft

                     Heading {
                         anchors.verticalCenter: parent.verticalCenter
                         text: "Do you have a HD Diablo II?"
                     }
                }
            }

            Column {
                topPadding: 15
                Label {
                    text: "SET HD DIRECTORY"
                    font.pointSize: 13
                    font.family: d2Font.name
                    color: "#ffffff"
                    font.bold: true
                }

                visible: usingHD
            }

            Row {
                TextField {
                    id: hdPathInput
                    width: fileDialogBox.width * 0.80
                    readOnly: true
                    background: Rectangle {
                        radius: 3; color: "#1d1924"
                    }
                }

                DefaultButton {
                    id: chooseHDPath
                    text: "CHOOSE"
                    width: fileDialogBox.width * 0.20
                    onClicked: hdPathDialog.open()
                }

                visible: usingHD
            }

            Column {
                topPadding: 15

                Heading {
                    text: "NUMBER OF HD INSTANCES TO LAUNCH"
                }
                
                ComboBox {
                    id: hdInstances
                    model: [ 1, 2, 3, 4 ]
                    height: 30
                    width: 60

                    background: Rectangle {
                        color: "#1d1924"
                        border.color: "#f0681f"
                        radius: height/2
                    }
                }

                visible: usingHD
            }

            // Save button.
            Column {
                topPadding: 25

                DefaultButton {
                    id: saveGamePath
                    text: "SAVE SETTINGS"

                    onClicked: {
                        console.log("save" + d2pathInput.text)
                        var success = settings.setGamePaths(d2pathInput.text, "")
                        if (success) {
                            settingsDialog.close()
                            QmlBridge.patchGame()
                        }
                    }
                }
            }
        }   
        
        // File dialogs.
        FileDialog {
            id: d2PathDialog
            selectFolder: true
            folder: shortcuts.home
            
            onAccepted: {
                var path = d2PathDialog.fileUrl.toString()
                path = path.replace(/^(file:\/{2})/,"")
                d2pathInput.text = path
            }
        }

        FileDialog {
            id: hdPathDialog
            selectFolder: true
            folder: shortcuts.home

            onAccepted: {
                var path = hdPathDialog.fileUrl.toString()
                path = path.replace(/^(file:\/{2})/,"")
                hdPathInput.text = path
            }
        }
    }
}