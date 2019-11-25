import QtQuick 2.12
import QtQuick.Layouts 1.3		// ColumnLayout
import QtQuick.Controls 2.1     // TextField
import QtQuick.Dialogs 1.3      // FileDialog

Item {
    property var game: {}

    function setGame(current) {
        // Set current game instance to the view.
        this.game = current

        // Textfield needs to be set explicitly since it's read only.
        if(this.game.location != undefined) {
            d2pathInput.text = this.game.location
        }

        // Update the switches initial state without triggering an animation.
        maphackSwitch.update()
        overrideMaphackCfgSwitch.update()
        hdSwitch.update()
        updateToggleBoxes()
    }

    function updateToggleBoxes() {
        if(this.game.flags != null) {
            windowModeFlag.active = this.game.flags.includes("-w")
            gfxFlag.active = this.game.flags.includes("-gfx")
        }
    }

    function makeFlagList() {
        var flags = []
        if(windowModeFlag.active) {
            flags.push("-w")
        }
        
        if(gfxFlag.active) {
            flags.push("-gfx")
        }

        return flags
    }

    function updateGameModel() {
        if(game != undefined) {
            var body = {
                id: game.id,
                location: d2pathInput.text,
                instances: (gameInstances.currentIndex+1),
                maphack: maphackSwitch.checked,
                override_bh_cfg: overrideMaphackCfgSwitch.checked,
                hd: hdSwitch.checked,
                flags: makeFlagList()
            }
            
            var success = settings.upsertGame(JSON.stringify(body))
            
            // TODO: Implement error handling.
        }
    }

    Item {
        id: currentGame
        width: parent.width * 0.95
        height: 400

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
                    topPadding: 0
                    spacing: 5

                    Title {
                        text: "SET DIABLO II DIRECTORY"
                        font.pixelSize: 13
                    }

                    SText {
                        text: "Specify your Diablo II game directory in order for the launcher to use it."
                        font.pixelSize: 11
                        color: "#454545"
                    }
                }

                Row {
                    anchors.bottom: parent.bottom
                    anchors.bottomMargin: 15

                    TextField {
                        id: d2pathInput
                        width: fileDialogBox.width * 0.60; height: 35
                        font.pixelSize: 11
                        color: "#454545"
                        readOnly: true
                        text: (game != undefined ? game.location : "")

                        background: Rectangle {
                            color: "#1a1a17"
                        }
                    }

                    SButton {
                        id: chooseD2Path
                        label: "Open"
                        borderRadius: 0
                        borderColor: "#373737"
                        width: fileDialogBox.width * 0.20; height: 35
                        cursorShape: Qt.PointingHandCursor

                        onClicked: d2PathDialog.open()
                    }

                    Item {
                        width: (fileDialogBox.width - (d2pathInput.width + chooseD2Path.width)); height: 35

                        ToggleButton {
                            id: windowModeFlag
                            label: "-w"
                            width: 45
                            height: 35
                            anchors.left: parent.left
                            anchors.leftMargin: 5
                            onClicked: updateGameModel()
                        }

                        ToggleButton {
                            id: gfxFlag
                            label: "-gfx"
                            width: 45
                            height: 35
                            anchors.right: parent.right
                            onClicked: updateGameModel()
                        }
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
                            
                            // Update the game model.
                            updateGameModel()
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
                        
                        Title {
                            text: "INSTANCES TO LAUNCH"
                            font.pixelSize: 13
                        }

                        SText {
                            text: "Number of this specific install that will launch when playing the game."
                            font.pixelSize: 11
                            topPadding: 5
                            color: "#454545"
                        }
                    }
                    Column {
                        id: instancesDropdown
                        width: 60
                        Dropdown{
                            id: gameInstances
                            currentIndex: (game != undefined ? (game.instances-1) : 0)
                            model: [ 1, 2, 3, 4 ]
                            height: 30
                            width: 60

                            onActivated: updateGameModel()
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
                        Title {
                            text: "INCLUDE MAPHACK"
                            font.pixelSize: 13
                        }

                        SText {
                            text: "Maphack will be downloaded automatically for this specific install."
                            font.pixelSize: 11
                            topPadding: 5
                            color: "#454545"
                        }
                    }
                    Column {
                        id: includeMaphack
                        width: 60
                        SSwitch{
                            id: maphackSwitch
                            checked: ((game != undefined && game.maphack != undefined) ? game.maphack : false)
                            onToggled: updateGameModel()
                        }
                    } 
                }
                
                Separator{}
            }

            // Use default maphack config.
            Item {
                Layout.preferredWidth: settingsLayout.width
                Layout.preferredHeight: 60

                Row {
                    topPadding: 10

                    Column {
                        width: (settingsLayout.width - overrideMaphackCfg.width)
                        Title {
                            text: "OVERRIDE MAPHACK CONFIG"
                            font.pixelSize: 13
                        }

                        SText {
                            text: "If you want to provide your own custom BH.cfg."
                            font.pixelSize: 11
                            topPadding: 5
                            color: "#454545"
                        }
                    }
                    Column {
                        id: overrideMaphackCfg
                        width: 60
                        SSwitch{
                            id: overrideMaphackCfgSwitch
                            checked: ((game != undefined && game.override_bh_cfg != undefined) ? game.override_bh_cfg : false)
                            onToggled: updateGameModel()
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
                        Title {
                            text: "INCLUDE HD MOD"
                            font.pixelSize: 13
                        }

                        SText {
                            text: "HD mod will be installed automatically for this specific install."
                            font.pixelSize: 11
                            topPadding: 5
                            color: "#454545"
                        }
                    }
                    Column {
                        id: includeHD
                        width: 60
                        SSwitch{
                            id: hdSwitch
                            checked: ((game != undefined && game.hd != undefined) ? game.hd : false)
                            onToggled: updateGameModel()
                        }
                    }
                }
                
                Separator{}
            }

             // Dep fix.
            Item {
                Layout.preferredWidth: settingsLayout.width
                Layout.preferredHeight: 60

                Row {
                    topPadding: 10

                    Column {
                        width: (settingsLayout.width - depFixButton.width)
                        Title {
                            text: "DISABLE DEP (REQUIRES ADMIN)"
                            font.pixelSize: 13
                        }

                        SText {
                            text: "Run if this install has troubles with crashing - requires reboot after."
                            font.pixelSize: 11
                            topPadding: 5
                            color: "#454545"
                        }
                    }
                    Column {
                        id: depFixButton
                        width: 100
                        
                        PlainButton {
                            width: 100
                            height: 40
                            label: "Run"

                            onClicked: {
                                diablo.applyDEP(d2pathInput.text)
                            }
                        }
                    } 
                }
            }
        }
    }
}
