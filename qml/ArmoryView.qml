import QtQuick 2.12
import QtQuick.Controls 2.5

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
    }
}