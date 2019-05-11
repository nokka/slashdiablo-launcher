import QtQuick 2.0
import QtQuick.Controls 2.5

Rectangle {
        id: bottombar
        anchors.bottom: parent.bottom;
        width: parent.width
        height: 100
        color: "#100b17"


        // Commented out for now.
        //Progress{}

        ProgressBar {
            value: QmlBridge.patchProgress
        }

        // Launch button.
        Button {
            width: parent.width * 0.20; height: 50
            anchors.verticalCenter: parent.verticalCenter
            anchors.right: parent.right;
            anchors.margins: 20 

            Text {
                id: launchtext
                text: "LAUNCH"
                color: "#ffffff"
                anchors.verticalCenter: parent.verticalCenter
                anchors.horizontalCenter: parent.horizontalCenter
                font.pointSize: 16;
            }

            background: Rectangle {
                color: "#0b86ba"
                radius: 2
            }

            onClicked: QmlBridge.launchGame()
        }
}