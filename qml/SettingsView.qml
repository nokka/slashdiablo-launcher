import QtQuick 2.12
import QtQuick.Layouts 1.3          // RowLayout

Rectangle {
    property int itemHeight: 50
    property int gameListHeight: settings.games.rowCount() * itemHeight
    property bool gameLocationSet: settings.games.rowCount() > 0
    property var gameRoles: { 
        "id": 257,
        "location": 258,
        "instances": 260,
        "maphack": 264,
        "hd": 272
    }

    color: "#09030a"

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
                height: gameListHeight
                anchors.top: parent.top
                anchors.right: parent.right
                anchors.topMargin: 50

                model: settings.games
                delegate: SettingsDelegate{}

                onCurrentItemChanged: {
                    gameSettings.setGame({
                        "id": model.data(model.index(this.currentIndex, 0), gameRoles.id),
                        "location": model.data(model.index(this.currentIndex, 0), gameRoles.location),
                        "instances": model.data(model.index(this.currentIndex, 0), gameRoles.instances),
                        "maphack": model.data(model.index(this.currentIndex, 0), gameRoles.maphack),
                        "hd": model.data(model.index(this.currentIndex, 0), gameRoles.hd)
                    })
                }
            }

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

                        var rows = settings.games.rowCount()
                        gameListHeight = rows * itemHeight
                        gamesList.currentIndex = (rows-1)

                        // Update if any games has been set yet.
                        gameLocationSet = rows > 0
                    }
                }
            }
        }
        
        // Right column.
        Item {
            Layout.fillWidth: true
            Layout.fillHeight: true

            Item {
                visible: !gameLocationSet
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
                visible: gameLocationSet
                anchors.fill: parent

                SText {
                    text: "SETTINGS"
                    anchors.top: parent.top
                    anchors.left: parent.left
                    anchors.topMargin: 20
                    font.pixelSize: 15
                    font.bold: true
                    leftPadding: 20
                }

                GameSettings {
                    id: gameSettings
                    anchors.left: parent.left
                    anchors.top: parent.top
                    anchors.topMargin: 49
                    anchors.horizontalCenter: parent.horizontalCenter
                }
            }
        }
    }

    // Close button.
    Item {
        width: 100
        height: 30
        anchors.top: parent.top
        anchors.right: parent.right
        anchors.rightMargin: 20
        
        SText {
            anchors.verticalCenter: parent.verticalCenter
            anchors.right: parent.right
            text: "CLOSE SETTINGS"
            font.pixelSize: 12
            font.bold: true
        }

        MouseArea {
            anchors.fill: parent
            cursorShape: Qt.PointingHandCursor
            onClicked: stack.pop()
        }
    }
}