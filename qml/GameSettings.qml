import QtQuick 2.12
import QtQuick.Layouts 1.3		// ColumnLayout
import QtQuick.Controls 2.1     // TextField
import QtQuick.Dialogs 1.3      // FileDialog

Item {
    property var game: {}
    property bool depApplied: false
    property bool depError: false
    property int activeHDIndex: 0

    function setGame(current) {
        // Set current game instance to the view.
        game = current
        
        // Textfield needs to be set explicitly since it's read only.
        if(game.location != undefined) {
            d2pathInput.text = game.location
        }

        // Update the switches initial state without triggering an animation.
        maphackSwitch.update()
        overrideMaphackCfgSwitch.update()
        updateToggleBoxes(current)
        updateModVersions(current)

    }

    function updateToggleBoxes(current) {
        if(current.flags != null) {
            windowModeFlag.active = current.flags.includes("-w")
            gfxFlag.active = current.flags.includes("-3dfx")
            skipFlag.active = current.flags.includes("-skiptobnet")
        } else {
            windowModeFlag.active = false
            gfxFlag.active = false
            skipFlag.active = false
        }
    }

    // updateModVersions will set the correct index of the hd mod dropdown.
    function updateModVersions(current) {
        if(settings.availableHDMods.length > 0) {
            if(current.hd_version == "") {
                activeHDIndex = 0
                return
            }

            // Find the correct index.
            for(var i = 0; i < settings.availableHDMods.length; i++) {
                if(settings.availableHDMods[i] == current.hd_version) {
                    activeHDIndex = i
                    return
                }
            }
        }

        // Default to 0.
        activeHDIndex = 0 
    }

    function makeFlagList() {
        var flags = []
        if(windowModeFlag.active) {
            flags.push("-w")
        }
        
        if(gfxFlag.active) {
            flags.push("-3dfx")
        }

        if(skipFlag.active) {
            flags.push("-skiptobnet")
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
                flags: makeFlagList(),
                hd_version: hdVersion.currentText
            }
            
            settings.upsertGame(JSON.stringify(body))
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
                        width: fileDialogBox.width * 0.55; height: 35
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
                        width: fileDialogBox.width * 0.15; height: 35
                        cursorShape: Qt.PointingHandCursor

                        onClicked: d2PathDialog.open()
                    }

                    Item {
                        width: (fileDialogBox.width - (d2pathInput.width + chooseD2Path.width)); height: 35

                        Row {
                            spacing: 2
                            leftPadding: 2

                            ToggleButton {
                                id: windowModeFlag
                                label: "-w"
                                width: 47
                                height: 35
                                onClicked: updateGameModel()
                            }

                            ToggleButton {
                                id: gfxFlag
                                label: "-3dfx"
                                width: 47
                                height: 35
                                onClicked: updateGameModel()
                            }

                            ToggleButton {
                                id: skipFlag
                                label: "-skip"
                                width: 47
                                height: 35
                                onClicked: updateGameModel()
                            }
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
                            text: "HD MOD VERSION"
                            font.pixelSize: 13
                        }

                        SText {
                            text: "Select if you want any HD mod installed"
                            font.pixelSize: 11
                            topPadding: 5
                            color: "#454545"
                        }
                    }
                    Column {
                        id: includeHD
                        width: 90

                        Dropdown{
                            id: hdVersion
                            currentIndex: activeHDIndex
                            model: settings.availableHDMods
                            height: 30
                            width: 90

                            onActivated: {
                                updateGameModel()
                            }
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
                                var success = diablo.applyDEP(d2pathInput.text)

                                if(success) {
                                    depApplied = true
                                    // Remove message after a timeout.
                                    depAppliedTimer.restart()
                                } else {
                                    depError = true
                                    // Remove message after a timeout.
                                    depErrorTimer.restart()
                                }
                            }
                        }
                    } 
                }

                // DEP success message.
                Rectangle {
                    visible: depApplied
                    width: parent.width
                    height: parent.height
                    color: "#00632e"
                    border.width: 1
                    border.color: "#000000"

                    SText {
                        text: "DEP fix successfully applied - don't forget to reboot!"
                        font.pixelSize: 11
                        anchors.centerIn: parent
                        color: "#ffffff"
                    }
                }

                // DEP error message.
                Rectangle {
                    visible: depError
                    width: parent.width
                    height: parent.height
                    color: "#8f3131"
                    border.width: 1
                    border.color: "#000000"

                    SText {
                        text: "There was an error while applying DEP, please try again!"
                        font.pixelSize: 11
                        anchors.centerIn: parent
                        color: "#ffffff"
                    }
                }
            }
        }
    }

    Timer {
        id: depAppliedTimer
        interval: 3000; running: false; repeat: false
        onTriggered: depApplied = false
    }

    Timer {
        id: depErrorTimer
        interval: 3000; running: false; repeat: false
        onTriggered: depError = false
    }
}
