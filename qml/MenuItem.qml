import QtQuick 2.12     // Text, MouseArea

SText {
    anchors.verticalCenter: parent.verticalCenter
    anchors.horizontalCenter: parent.horizontalCenter
    font.bold: true
    
    MouseArea {
        anchors.fill: parent
        cursorShape: Qt.PointingHandCursor

        onClicked: {
            console.log("CLICKED")
        }
    }
} 