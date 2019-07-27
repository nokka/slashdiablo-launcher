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
    height: 510
    margins: 0
    padding: 0
    
    anchors.centerIn: root
    closePolicy: Popup.NoAutoClose

    Rectangle {
        color: "#050000"
        anchors.fill: parent

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

                SText {
                    text: "MY GAMES"
                    anchors.top: parent.top
                    anchors.left: parent.left
                    anchors.topMargin: 20
                    font.pixelSize: 15
                    font.bold: true
                    leftPadding: 15
                }

                ListView {
                    id: gamesList
                    width: parent.width - 15;
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
                SText {
                    text: "+ Add new game"
                    anchors.top: gamesList.bottom
                    anchors.left: parent.left
                    anchors.topMargin: 20
                    font.bold: true
                    leftPadding: 25

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

                // Close button.
                Item {
                    width: 35
                    height: 35
                    anchors.top: parent.top
                    anchors.right: parent.right
                    anchors.rightMargin: 10
                    anchors.topMargin: 5

                    Image {
                        fillMode: Image.PreserveAspectFit
                        anchors.centerIn: parent
                        width: 35
                        height: 35
                        source: "assets/svg/close.svg"
				    }
                    
                   
                    MouseArea {
                        anchors.fill: parent
                        cursorShape: containsMouse ? Qt.ArrowCursor : Qt.PointingHandCursor
                        onClicked: {
                            settingsPopup.close()

                            // Validate the game versions after we've made updates.
                            diablo.validateVersion()
                        }
                    }
                }
            }
        }
    }
}