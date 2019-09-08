import QtQuick 2.4
import QtQuick.Controls 2.5
import QtQuick.Controls.Material 2.1
import QtQuick.Layouts 1.3

Popup {
    id: settingsPopup

    property int itemHeight: 50
    property bool errored: false

    property var gameRoles: { 
        "id": 257,
        "location": 258,
        "instances": 260,
        "maphack": 264,
        "override_bh_cfg": 272,
        "hd": 288
    }

    modal: true
    focus: true
    width: 850
    height: 500
    margins: 0
    padding: 0
    
    anchors.centerIn: root
    closePolicy: Popup.NoAutoClose

    Overlay.modal: Item {
        Rectangle {
            anchors.fill: parent
            color: "#000000"
            opacity: 0.8
        }
    }

    Rectangle {
        color: "#0f0f0f"
        border.color: "#785A29"
        border.width: 1
        anchors.fill: parent

        // Bottom background.
        Image {
            width: 848
            source: "assets/stone_bg.png";
            fillMode: Image.Stretch;
            anchors.bottom: parent.bottom
            anchors.horizontalCenter: parent.horizontalCenter
            anchors.bottomMargin: 1
            opacity: 1.0
        }

        RowLayout {
            id: settingsLayout
            anchors.fill: parent
            spacing: 8

             // Left column.
            Item {
                Layout.fillWidth: true
                Layout.minimumWidth: 300
                Layout.preferredWidth: 300
                Layout.maximumWidth: 300
                Layout.fillHeight: true

                Title {
                    text: "GAME SETTINGS"
                    color: "#c4b58b"
                    anchors.top: parent.top
                    anchors.left: parent.left
                    anchors.topMargin: 20
                    font.pixelSize: 15
                    font.bold: true
                    leftPadding: 30
                }

                ListView {
                    id: gamesList
                    width: parent.width - 30;
                    height: gamesList.count * itemHeight
                    anchors.top: parent.top
                    anchors.right: parent.right
                    anchors.topMargin: 50

                    model: settings.games
                    delegate: SettingsDelegate{}

                    onCurrentItemChanged: {   
                        if(gamesList.count > 0) {
                            gameSettings.setGame({
                                "id": model.data(model.index(this.currentIndex, 0), gameRoles.id),
                                "location": model.data(model.index(this.currentIndex, 0), gameRoles.location),
                                "instances": model.data(model.index(this.currentIndex, 0), gameRoles.instances),
                                "maphack": model.data(model.index(this.currentIndex, 0), gameRoles.maphack),
                                "override_bh_cfg": model.data(model.index(this.currentIndex, 0), gameRoles.override_bh_cfg),
                                "hd": model.data(model.index(this.currentIndex, 0), gameRoles.hd)
                            })
                        }
                    }
                }

                // Add new game button.
                Title {
                    visible: (gamesList.count <= 3)
                    text: "+ Add new game"
                    anchors.top: gamesList.bottom
                    anchors.left: parent.left
                    anchors.topMargin: 20
                    font.bold: true
                    leftPadding: 30

                    MouseArea {
                        anchors.fill: parent
                        cursorShape: Qt.PointingHandCursor
                        onClicked: {
                            // Add the game to the model, without persisting it to the store.
                            settings.addGame()

                            // Set last index as current.
                            gamesList.currentIndex = (gamesList.count-1)
                        }
                    }
                }
            }

             // Right column.
            Item {
                Layout.fillWidth: true
                Layout.fillHeight: true

                // Visible if there are no games set up.
                Item {
                    visible: (gamesList.count == 0)
                    anchors.centerIn: parent
                    width: (parent.width * 0.80)

                    SText {
                        text: "Before you can play, you need to setup your game location in the menu to the left."
                        width: parent.width
                        anchors.verticalCenter: parent.verticalCenter
                        anchors.left: parent.left
                        font.pixelSize: 15
                        wrapMode: Text.WordWrap 
                    }
                }

                 // Settings shown if there are games already setup    
                Item {
                    visible: (gamesList.count > 0)
                    anchors.fill: parent

                    GameSettings {
                        id: gameSettings
                        anchors.left: parent.left
                        anchors.top: parent.top
                        anchors.topMargin: 45
                        anchors.horizontalCenter: parent.horizontalCenter
                    }
                }

                // Error popup.
                Item {
                    id: errorPopup
                    visible: errored
                    width: 300
                    height: 50
                    anchors.horizontalCenter: doneButton.horizontalCenter
                    anchors.bottom: doneButton.top
                    anchors.bottomMargin: 15

                    Rectangle {
                        anchors.fill: parent
                        color: "#8f3131"
                        border.width: 1
                        border.color: "#000000"
                    }


                    SText {
                        text: "New Game doesn't have Diablo II directory set."
                        font.pixelSize: 11
                        anchors.centerIn: parent
                        color: "#ffffff"
                    }
                }
                
                XButton {
                    id: doneButton
                    visible: (gamesList.count > 0)
                    label: "DONE"
                    width: 100
                    height: 50
                    anchors.bottom: parent.bottom
                    anchors.left: parent.left
                    anchors.bottomMargin: -25
                    anchors.leftMargin: 65

                    onClicked: {
                        // Reset error.
                        errored = false
                        
                        if(validateGames()) {
                            var success = settings.persistGameModel()
                            if(success) {
                                // Validate the game versions after changes has been made to the settings.
                                diablo.validateVersion()
                                settingsPopup.close()
                            }
                        } else {
                            // Show error.
                            errored = true

                            // Remove error after a timeout.
                            errorTimer.restart()
                        }
                    }
                }
            }
        }
    }

    // validateGames will validate that the input is correctly set.
    function validateGames() {
       for(var i = 0; i < gamesList.count; i++) {
            var location = gamesList.model.data(gamesList.model.index(i, 0), gameRoles.location)
            if(location.length == 0) {
                return false
            }
        }
        
        return true
    }

    Timer {
        id: errorTimer
        interval: 5000; running: false; repeat: false
        onTriggered: errored = false
    }
}
