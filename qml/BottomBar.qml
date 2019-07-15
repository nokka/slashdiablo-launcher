import QtQuick 2.12

Item {
        id: bottombar
        anchors.bottom: parent.bottom;
        width: parent.width; height: 80

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
            SButton {
                label: "PLAY"
                fontSize: 15
                enabled: diablo.playable
                width: parent.width; height: 50
                anchors.verticalCenter: parent.verticalCenter
                cursorShape: Qt.PointingHandCursor

                onClicked: diablo.launchGame()
            }
        }
}