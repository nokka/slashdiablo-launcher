import QtQuick 2.12
import QtQuick.Layouts 1.3		// ColumnLayout
import QtQuick.Controls 2.1     // TextField
import QtQuick.Dialogs 1.3      // FileDialog

Item {
    property var game: {}

    function setGame(current) {
        // Set current game instance to the view.
        this.game = current

        // Update the switches initial state without triggering an animation.
        maphackSwitch.update()
        hdSwitch.update()
    }

    Rectangle {
        id: currentGame
        width: parent.width * 0.95
        height: 400
        color: "#0d0a08"
        radius: 5
        border.color: "#343434"
        border.width: 1

        anchors.horizontalCenter: parent.horizontalCenter

        ColumnLayout {
            id: settingsLayout
            width: (currentGame.width * 0.95)
            spacing: 2
            
            anchors.horizontalCenter: parent.horizontalCenter

            // D2 Directory box.
            Item {
                id: fileDialogBox
                Layout.preferredWidth: settingsLayout.width
                Layout.preferredHeight: 100

                Column {
                    anchors.top: parent.top
                    topPadding: 10
                    spacing: 5

                    SText {
                        text: "Set Diablo II directory"
                        font.pixelSize: 13
                        font.bold: true
                    }

                    SText {
                        text: "Specify your Diablo II game directory in order for the launcher to use it."
                        font.pixelSize: 12
                        color: "#505050"
                    }
                }

                Row {
                    anchors.bottom: parent.bottom
                    anchors.bottomMargin: 10

                    TextField {
                        id: d2pathInput
                        width: fileDialogBox.width * 0.80; height: 40
                        readOnly: true
                        text: game.location

                        background: Rectangle {
                            color: "#1d1924"
                        }
                    }

                    SButton {
                        id: chooseD2Path
                        label: "Open"
                        borderRadius: 0
                        width: fileDialogBox.width * 0.20; height: 40
                        cursorShape: Qt.PointingHandCursor

                        onClicked: d2PathDialog.open()
                    }

                    // File dialog.
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
                }
                
                Separator{}
            }

            // Game instances box.
            Item {
                Layout.preferredWidth: settingsLayout.width
                Layout.preferredHeight: 60

                Row {
                    topPadding: 10

                    Column {
                        width: (settingsLayout.width - instancesDropdown.width)
                        SText {
                            text: "Instances to launch"
                            font.pixelSize: 13
                            font.bold: true
                        }

                        SText {
                            text: "Number of this specific install that will launch when playing the game."
                            font.pixelSize: 12
                            color: "#505050"
                        }
                    }
                    Column {
                        id: instancesDropdown
                        width: 60
                        Dropdown{
                            id: gameInstances
                            currentIndex: (game.instances-1)
                            model: [ 1, 2, 3, 4 ]
                            height: 30
                            width: 60
                        }
                    }
                    
                }
                
                Separator{}
            }

            // Include maphack box.
            Item {
                Layout.preferredWidth: settingsLayout.width
                Layout.preferredHeight: 60

                Row {
                    topPadding: 10

                    Column {
                        width: (settingsLayout.width - includeMaphack.width)
                        SText {
                            text: "Include maphack"
                            font.pixelSize: 13
                            font.bold: true
                        }

                        SText {
                            text: "Maphack will be downloaded automatically for this specific install."
                            font.pixelSize: 12
                            color: "#505050"
                        }
                    }
                    Column {
                        id: includeMaphack
                        width: 60
                        SSwitch{
                            id: maphackSwitch
                            checked: game.maphack
                        }
                    } 
                }
                
                Separator{}
            }

            // Include HD box.
            Item {
                Layout.preferredWidth: settingsLayout.width
                Layout.preferredHeight: 60

                Row {
                    topPadding: 10

                    Column {
                        width: (settingsLayout.width - includeHD.width)
                        SText {
                            text: "Include HD mod"
                            font.pixelSize: 13
                            font.bold: true
                        }

                        SText {
                            text: "HD mod will be installed automatically for this specific install."
                            font.pixelSize: 12
                            color: "#505050"
                        }
                    }
                    Column {
                        id: includeHD
                        width: 60
                        SSwitch{
                            id: hdSwitch
                            checked: game.hd
                        }
                    }
                }
                
                Separator{}
            }

            // Save button.
            Item {
                Layout.preferredWidth: settingsLayout.width
                Layout.preferredHeight: 60

                Row {
                    topPadding: 15

                    SButton {
                        id: saveSettings
                        label: "SAVE"
                        width: 100; height: 40
                        cursorShape: Qt.PointingHandCursor

                        onClicked: {
                            console.log("Saving settings")

                            var body = {
                                id: game.id,
                                location: d2pathInput.text,
                                instances: parseInt(gameInstances.currentText, 10),
                                maphack: maphackSwitch.checked,
                                hd: hdSwitch.checked
                            }

                            var success = settings.upsertGame(JSON.stringify(body))
                            
                            console.log(success)
                        }
                    }
                }
            }
        }
    }
}