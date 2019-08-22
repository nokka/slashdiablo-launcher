import QtQuick 2.12

Item {
        id: bottombar
        anchors.bottom: parent.bottom;
        width: parent.width; height: 80

        // Background.
        Rectangle {
            anchors.fill: parent
            color: "#000000"
            opacity: 0.6
        }

        Separator {
            color: "#3d3b36"
            anchors.top: parent.top
            anchors.bottom: undefined
        }

        // Patcher including progress bar.
        Patcher{
            width: parent.width * 0.65
        }

        Item {
            width: parent.width * 0.30; height: parent.height
            anchors.verticalCenter: parent.verticalCenter
            anchors.right: parent.right;
            anchors.rightMargin: 20

             // Launch button.
            XButton {
                label: "PLAY"
                fontSize: 15
                clickable: diablo.validVersion
                width: parent.width; height: 50
                anchors.verticalCenter: parent.verticalCenter

                onClicked: diablo.launchGame()
            }
        }
}
