import QtQuick 2.12     // MouseArea

SText {
    property var onClicked: function () {}
    property bool active: false
    
    anchors.centerIn: parent
    font.family: beaufortbold.name
    font.bold: true
    font.pixelSize: 14
    color: active ? "#c7cbd1" : "#3b3b3b"
    // d4d4d4
    MouseArea {
        id: mousearea
        anchors.fill: parent
        cursorShape: Qt.PointingHandCursor
        onClicked: parent.onClicked()
        hoverEnabled: true
    }
} 
