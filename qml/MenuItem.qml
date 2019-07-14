import QtQuick 2.12     // MouseArea

SText {
    property var onClicked: function () {}

    anchors.verticalCenter: parent.verticalCenter
    anchors.horizontalCenter: parent.horizontalCenter
    font.bold: true
    
    MouseArea {
        anchors.fill: parent
        cursorShape: Qt.PointingHandCursor
        onClicked: parent.onClicked()
    }
} 