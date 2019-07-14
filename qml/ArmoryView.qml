import QtQuick 2.12
import QtQuick.Controls 2.5

import "componentCreator.js" as ComponentCreator

Rectangle {
    id: armoryView
    color: "#09030a"

    Item {
        height: 100
        anchors.verticalCenter: parent.verticalCenter
        anchors.horizontalCenter: parent.horizontalCenter
        
        SText {
            text: "ARMORY VIEW"
            anchors.top: parent.top
            anchors.horizontalCenter: parent.horizontalCenter
        }

        SButton {
            label: "Close"
            width: 100; height: 50
            anchors.bottom: parent.bottom
            anchors.horizontalCenter: parent.horizontalCenter
            cursorShape: Qt.PointingHandCursor

            onClicked: stack.pop()
        }
    }
}