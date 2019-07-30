import QtQuick 2.12     // MouseArea

SText {
    property var onClicked: function () {}
    
    anchors.centerIn: parent
    font.family: beaufortbold.name
    font.bold: true
    font.pixelSize: 15
    color: "#c4b58b"
    
    MouseArea {
        anchors.fill: parent
        cursorShape: Qt.PointingHandCursor
        onClicked: parent.onClicked()
    }
} 