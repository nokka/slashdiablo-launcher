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
            Text {
                text: "SETTINGS"
                font.family: d2Font.name
                font.pointSize: 22
                color: "#ffffff"
                horizontalAlignment: Text.AlignHCenter
                verticalAlignment: Text.AlignVCenter
                elide: Text.ElideRight
            }

            id: fileDialogBox
            width: (mainWindow.width/2)
            Layout.alignment: Qt.AlignTop | Qt.AlignHCenter
            Layout.topMargin: 100
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
                topPadding: 10
                
                Heading {
                    text: "Do you have a HD version?"
                }
            }

            Row {
                spacing: 5

                Switch {
                    id: control
                    text: qsTr("Do you have a HD version?")

                    indicator: Rectangle {
                        implicitWidth: 48
                        implicitHeight: 26
                        x: control.leftPadding
                        y: parent.height / 2 - height / 2
                        radius: 13
                        color: control.checked ? "#17a81a" : "#ffffff"
                        border.color: control.checked ? "#17a81a" : "#cccccc"

                        Rectangle {
                            x: control.checked ? parent.width - width : 0
                            width: 26
                            height: 26
                            radius: 13
                            color: control.down ? "#cccccc" : "#ffffff"
                            border.color: control.checked ? (control.down ? "#17a81a" : "#21be2b") : "#999999"
                        }
                    }

                    contentItem: Text {
                        text: control.text
                        font: control.font
                        opacity: enabled ? 1.0 : 0.3
                        color: control.down ? "#17a81a" : "#21be2b"
                        verticalAlignment: Text.AlignVCenter
                        leftPadding: control.indicator.width + control.spacing
                    }
                }
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

            // Save button.
            Column {
                topPadding: 10

                DefaultButton {
                    id: saveGamePath 
                    text: "SAVE SETTINGS"

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