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
            height: 4
            value: diablo.patchProgress
            width: parent.width
            anchors.verticalCenter: parent.verticalCenter
            
            style: ProgressBarStyle {
                background: Rectangle {
                    radius: 3
                    color: "#00686b"
                }
                
                progress: Rectangle {
                    radius: 3
                    color: "#069499"
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

        Title {
            anchors.left: parent.left
            anchors.verticalCenter: parent.verticalCenter
            anchors.leftMargin: 30
            text: "Games are up to date"
            font.pixelSize: 15
        }


        Dropdown{
            id: gameInstances
            anchors.right: parent.right
            anchors.verticalCenter: parent.verticalCenter
            anchors.rightMargin: 30
            currentIndex: 0
            model: [ "Slashdiablo", "Battle.net"]
            height: 30
            width: 120

            onActivated: {
                diablo.setGateway(this.currentText)
            }
        }
    }

    // Show when the Diablo version is invalid, we're not patching and there's no error.
    Item {
        anchors.left: parent.left
        anchors.verticalCenter: parent.verticalCenter
        width: 350
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

        Title {
            anchors.left: parent.left
            anchors.verticalCenter: parent.verticalCenter
            anchors.leftMargin: 30
            text: "Games need to be updated"
            font.pixelSize: 15
        }

        XButton {
            width: 120
            height: 40
            label: "UPDATE NOW"
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