import QtQuick 2.12
import QtQuick.Controls 1.4
import QtQuick.Controls.Styles 1.4

Item {
    id: patcher
    height: 80
    anchors.left: parent.left
    anchors.leftMargin: 20
    anchors.verticalCenter: parent.verticalCenter

    // Show when we're patching and no error has occurred.
    Item {
        anchors.fill:parent 
        visible: (diablo.patching && !diablo.errored)

        ProgressBar {
            height: 7
            value: diablo.patchProgress
            width: parent.width
            anchors.verticalCenter: parent.verticalCenter
            
            style: ProgressBarStyle {
                background: Rectangle {
                    radius: 3
                    color: "#262626"
                    border.color: "#191919"
                    border.width: 2
                }
                
                progress: Rectangle {
                    radius: 3
                    color: "#600303"
                    border.color: "#191919"
                    border.width: 2
                }
            }
        }

        SText {
            anchors.bottom: parent.bottom;
            text: diablo.status
            font.pixelSize: 12
        }
    }

    // Show when patcher errors.
    Item {
        anchors.fill:parent 
        visible: diablo.errored
        
        Image {
            id: patcherError
            fillMode: Image.PreserveAspectFit
            anchors.horizontalCenter: parent.horizontalCenter
            anchors.top: parent.top
            width: 20
            height: 20
            source: "assets/svg/error.svg"
        }

        SText {
            anchors.horizontalCenter: parent.horizontalCenter
            topPadding: 30
            text: "Couldn't patch game files"
            font.pixelSize: 11
        }
    }

    // Show when patching is done, no error occurred and the game is playable.
    Item {
        anchors.fill:parent 
        visible: (!diablo.patching && !diablo.errored && diablo.playable)

        SText {
            anchors.bottom: parent.bottom;
            text: "Game is up to date"
            font.pixelSize: 15
        }
    }

    // Show when the Diablo version is invalid, we're not patching and there's no error.
    Item {
        anchors.left: parent.left
        anchors.verticalCenter: parent.verticalCenter
        width: 320
        height: 40
        visible: (!diablo.validVersion && !diablo.patching && !diablo.errored)

        Image {
            id: versionError
            fillMode: Image.PreserveAspectFit
            anchors.verticalCenter: parent.verticalCenter
            anchors.left: parent.left
            width: 20
            height: 20
            source: "assets/svg/error.svg"
        }

        SText {
            anchors.left: parent.left
            anchors.verticalCenter: parent.verticalCenter
            anchors.leftMargin: 30
            text: "Game version isn't 1.13c"
            font.pixelSize: 15
        }

        SButton {
            width: 100
            height: 40
            label: "UPDATE"
            fontSize: 10
            anchors.top: parent.top
            anchors.right: parent.right
            cursorShape: Qt.PointingHandCursor

            onClicked: {
                diablo.applyPatches()
            }
        }
    }

    Component.onCompleted: {
        if(settings.games.rowCount() > 0) {
            diablo.validateVersion()
        }
    }
}