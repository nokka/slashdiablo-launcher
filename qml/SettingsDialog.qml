import QtQuick 2.12
import QtQuick.Controls 2.1
import QtQuick.Layouts 1.3
import QtQuick.Dialogs 1.3

Rectangle {
    color: "#080806"
    
    ColumnLayout {
        anchors.fill: parent
        Column {
            Header {
                text: "SETTINGS"
                font.pointSize: 20
            }

            id: fileDialogBox
            width: (mainWindow.width/2)
            Layout.alignment: Qt.AlignTop | Qt.AlignHCenter
            Layout.topMargin: 50    
            spacing: 5

            Column {
                topPadding: 15

                Label {
                    text: "Set Diablo II directory"
                    font.pointSize: 13
                    font.family: montserrat.name
                    color: "#ffffff"
                    font.bold: true
                }
            }

            Row {
                TextField {
                    id: d2pathInput
                    width: fileDialogBox.width * 0.80
                    readOnly: true
                    text: settings.D2Location

                    background: Rectangle {
                        radius: 3
                        color: "#1d1924"
                    }
                }

                DefaultButton {
                    id: chooseD2Path
                    text: "Open"
                    width: fileDialogBox.width * 0.20
                    onClicked: d2PathDialog.open()
                }
            }

            Column {
                topPadding: 15
                Header {
                    text: "Number of instances"
                }
                
                Dropdown{
                    id: d2Instances
                    height: 30
                    width: 60
                }

                Check{
                    checked: true
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
                        id: hdEnabled
                        anchors.right: parent.right
                        anchors.verticalCenter: parent.verticalCenter
                        checked: false
                    }
                }

                Item {
                    width: parent.width * 0.80
                    height: 35
                     Layout.alignment: Qt.AlignLeft

                     Header {
                         anchors.verticalCenter: parent.verticalCenter
                         text: "Do you have a HD Diablo II?"
                     }
                }
            }

            Column {
                topPadding: 15
                Label {
                    text: "Set HD directory"
                    font.pointSize: 13
                    font.family: montserrat.name
                    color: "#ffffff"
                    font.bold: true
                }

                visible: hdEnabled.checked
            }

            Row {
                TextField {
                    id: hdPathInput
                    width: fileDialogBox.width * 0.80
                    readOnly: true
                    text: settings.HDLocation
                    background: Rectangle {
                        radius: 3; color: "#1d1924"
                    }
                }

                DefaultButton {
                    id: chooseHDPath
                    text: "Open"
                    width: fileDialogBox.width * 0.20
                    onClicked: hdPathDialog.open()
                }

                visible: hdEnabled.checked
            }

            Column {
                visible: hdEnabled.checked
                topPadding: 15

                Header {
                    text: "Number of HD instances"
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

                Check{
                    checked: true
                }
            }

            // Save button.
            Column {
                topPadding: 25

                DefaultButton {
                    id: saveGamePath
                    text: "SAVE SETTINGS"

                    onClicked: {
                        var hdPath = hdPathInput.text
                        var hdi = hdInstances.currentText
                        
                        // HD isn't enable, reset the HD fields.
                        if(!hdEnabled.checked) {
                            hdPath = ""
                            hdi = 0
                        }

                        // Update settings in the backend.
                        var success = settings.update(
                            d2pathInput.text,
                            d2Instances.currentText,
                            hdPath,
                            hdi
                        )

                        if (success) {
                            settingsDialog.visible = false
                            diablo.validateVersion()
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