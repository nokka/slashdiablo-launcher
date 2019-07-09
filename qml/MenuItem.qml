import QtQuick 2.12     // Text, MouseArea

Text {
    anchors.verticalCenter: parent.verticalCenter
    anchors.horizontalCenter: parent.horizontalCenter
    font.family: robotobold.name
    font.pixelSize: 12
    font.bold: true
    color: "#ffffff"
    
    MouseArea {
        anchors.fill: parent
        cursorShape: Qt.PointingHandCursor

        onClicked: {
            console.log("CLICKED")
        }
    }
} 