import QtQuick 2.12     // MouseArea

SText {
    property var onClicked: function () {}
    
    anchors.centerIn: parent
    font.family: beaufortbold.name
    font.bold: true
    font.pixelSize: 14
    color: mousearea.containsMouse ? "#ffffe6" : "#c7cbd1"
    
    MouseArea {
        id: mousearea
        anchors.fill: parent
        cursorShape: Qt.PointingHandCursor
        onClicked: parent.onClicked()
        hoverEnabled: true
    }
} 
