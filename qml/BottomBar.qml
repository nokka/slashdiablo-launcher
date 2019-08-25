import QtQuick 2.12

Item {
        id: bottombar

        // Background image.
        Item {
            id: background
            anchors.fill: parent;
            Image { source: "assets/bottom_bg.jpg"; fillMode: Image.Stretch; anchors.fill: parent;  opacity: 0.4 }
        }

        Separator {
            color: "#030202"
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
