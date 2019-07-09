import QtQuick 2.12
import QtQuick.Controls 1.4
import QtQuick.Controls.Styles 1.4

Rectangle {
    id: patcher
    height: 60
    x: 20 

    // Anchors
    anchors.verticalCenter: parent.verticalCenter

    // Transparent background
    color: "#00000000" 

    // Show when we're patching and no error has occurred.
    Item {
        anchors.fill:parent 
        visible: (diablo.patching && !diablo.errored)

        ProgressBar {
            value: diablo.patchProgress
            width: parent.width
            height: 10

            // Anchors
            anchors.verticalCenter: parent.verticalCenter
            
            style: ProgressBarStyle {
                background: Rectangle {
                    radius: 2
                    color: "#381612"
                    border.color: "#141009"
                    border.width: 1
                }
                
                progress: Rectangle {
                    color: "#873d29"
                }
            }
        }

        Text {
            anchors.bottom: parent.bottom;
            color: "#ffffff"
            text: diablo.status
            font.pointSize: 12
            font.family: montserrat.name
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
            source: "assets/error.svg"
        }

        Text {
            color: "#ffffff"
            topPadding: 30
            text: "Couldn't patch game files"
            font.family: montserrat.name
            font.pixelSize: 11
            anchors.horizontalCenter: parent.horizontalCenter
        }

        Text {
            color: "#ffffff"
            topPadding: 30
            text: "Couldn't patch game files"
            font.family: montserrat.name
            font.pixelSize: 11
            anchors.horizontalCenter: parent.horizontalCenter
        }
    }

    // Show when patching is done, no error occurred and the game is playable.
    Item {
        anchors.fill:parent 
        visible: (!diablo.patching && !diablo.errored && diablo.playable)
        
         Text {
            anchors.bottom: parent.bottom;
            color: "#ffffff"
            text: "Game is up to date"
            font.pointSize: 15
            font.family: montserrat.name
        }
    }

    // Show when the Diablo version is invalid, we're not patching and there's no error.
    Item {
        anchors.left: parent.left
        anchors.verticalCenter: parent.verticalCenter
        width: 300
        height: 40
        visible: (!diablo.validVersion && !diablo.patching && !diablo.errored)
        
         Text {
            anchors.left: parent.left
            anchors.verticalCenter: parent.verticalCenter
            color: "#ffffff"
            text: "Game version isn't 1.13c"
            font.pointSize: 15
            font.family: montserrat.name
        }

        Button {
            anchors.top: parent.top
            anchors.right: parent.right
            width: 100
            height: 40
            Text {
                text: "UPDATE"
                color: "#f3e6d0"
                font.family: d2Font.name
                anchors.verticalCenter: parent.verticalCenter
                anchors.horizontalCenter: parent.horizontalCenter
                font.pointSize: 15;
            }

            onClicked: {
                console.log("Apply patches clicked")
                diablo.applyPatches()
            }
        }
    }

    Component.onCompleted: {
        if(settings.NrOfGames > 0) {
            diablo.validateVersion()
        }
    }
}