import QtQuick 2.4
import QtQuick.Controls 2.1
import QtQuick.Controls.Material 2.1
import QtQuick.Layouts 1.3

Popup {
    id: settingsPopup

    property int itemHeight: 50
    property var gameRoles: { 
        "id": 257,
        "location": 258,
        "instances": 260,
        "maphack": 264,
        "hd": 272
    }

    modal: true
    focus: true
    width: 850
    height: 500
    margins: 0
    padding: 0
    
    anchors.centerIn: root
    closePolicy: Popup.NoAutoClose

    Rectangle {
        color: "#0d0d0a"
        border.color: "#785A29"
        border.width: 1
        anchors.fill: parent

        // Bottom background.
        Image { 
            source: "assets/settings_bg.jpg";
            fillMode: Image.PreserveAspectFit;
            anchors.bottom: parent.bottom
            anchors.horizontalCenter: parent.horizontalCenter
            anchors.bottomMargin: 1
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
                    text: "MY GAMES"
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
                                "hd": model.data(model.index(this.currentIndex, 0), gameRoles.hd)
                            })
                        }
                    }
                }

                // Add new game button.
                Title {
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
                        anchors.topMargin: 49
                        anchors.horizontalCenter: parent.horizontalCenter
                    }
                }
                
                XButton {
                    label: "DONE"
                    width: 100
                    height: 50
                    anchors.bottom: parent.bottom
                    anchors.left: parent.left
                    anchors.bottomMargin: -25
                    anchors.leftMargin: 65

                    onClicked: {
                        var success = settings.persistGameModel()
                        if(success) {
                            settingsPopup.close()
                            return
                        }

                        // TODO: Add error handling.
                    }
                }
            }
        }
    }
}