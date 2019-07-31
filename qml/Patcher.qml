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
            height: 2
            value: diablo.patchProgress
            width: parent.width
            anchors.verticalCenter: parent.verticalCenter
            
            style: ProgressBarStyle {
                background: Rectangle {
                    radius: 3
                    color: "#0d0d0d"
                }
                
                progress: Rectangle {
                    radius: 3
                    color: "#600303"
                }
            }
        }

        SText {
            anchors.bottom: parent.bottom;
            anchors.bottomMargin: 10
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
            anchors.left: parent.left
            anchors.verticalCenter: parent.verticalCenter
            width: 25
            height: 25
            source: "assets/svg/error.svg"
        }

        SText {
            anchors.left: parent.left
            anchors.verticalCenter: parent.verticalCenter
            text: "Couldn't patch game files"
            font.pixelSize: 15
            anchors.leftMargin: 30
            topPadding: 5
        }
    }

    // Show when patching is done, no error occurred and the game is playable.
    Item {
        anchors.fill:parent 
        visible: (!diablo.patching && !diablo.errored && diablo.playable)

        SText {
            anchors.left: parent.left
            anchors.verticalCenter: parent.verticalCenter
            anchors.leftMargin: 30
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
            text: "Games aren't up to date"
            font.pixelSize: 15
        }

        XButton {
            width: 100
            height: 40
            label: "UPDATE"
            fontSize: 10
            anchors.top: parent.top
            anchors.right: parent.right

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