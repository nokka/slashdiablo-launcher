import QtQuick 2.0
import QtQuick.Controls 2.5

Rectangle {
        id: bottombar
        anchors.bottom: parent.bottom;
        width: parent.width
        height: 100
        color: "#00000000"

        // Patcher including progress bar.
        Patcher{
            width: parent.width * 0.65
        }

        // Launch button.
        Button {
            width: parent.width * 0.20; height: 50
            anchors.verticalCenter: parent.verticalCenter
            anchors.right: parent.right;
            anchors.margins: 20
            anchors.rightMargin: 72

            Text {
                text: "LAUNCH"
                color: "#f3e6d0"
                font.family: d2Font.name
                anchors.verticalCenter: parent.verticalCenter
                anchors.horizontalCenter: parent.horizontalCenter
                font.pointSize: 16;
            }

            background: Rectangle {
                color: "#790905"
                radius: 2
            }

            onClicked: QmlBridge.launchGame()
        }

        // Settings button.
        Button {
            width: 52; height: 52
            anchors.verticalCenter: parent.verticalCenter
            anchors.right: parent.right;
            anchors.margins: 20
            anchors.rightMargin: 20

            background: Rectangle {
                color: "#000000"
                opacity: 0.5
                radius: 2
            }

            Image {
                id: settingsIcon
                fillMode: Image.PreserveAspectFit
                anchors.centerIn: parent
                width: 20
                height: 20
                source: "assets/settings.svg"
            }

            MouseArea {
                anchors.fill: parent
                cursorShape: Qt.PointingHandCursor
                onClicked: settingsDialog.open()
            }
        }
}