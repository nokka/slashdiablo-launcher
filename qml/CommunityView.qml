import QtQuick 2.12
import QtQuick.Controls 2.5

Rectangle {
    id: communityView
    color: "#09030a"

    Item {
        height: 100
        anchors.verticalCenter: parent.verticalCenter
        anchors.horizontalCenter: parent.horizontalCenter
        
        SText {
            text: "COMMUNITY VIEW"
            anchors.top: parent.top
            anchors.horizontalCenter: parent.horizontalCenter
        }
    }
}