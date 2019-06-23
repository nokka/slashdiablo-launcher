import QtQuick 2.12

Rectangle {
    id: mainContent
    width: mainWindow.width * 0.70
    height: parent.height - 100
    color: "#00000000"

    
    Text {
        color: "#ffffff"
        text: "Slashdiablo"
        font.pointSize: 30
        font.family: montserrat.name
        anchors.top: parent.top
        anchors.left: parent.left
        elide: Text.ElideRight
        anchors.topMargin: 20
        anchors.leftMargin: 30
    }

    // News box.
    Box {
        title: "Ladder reset 8PM EST, 12 July"
        width: 350
        height: 200
        anchors.verticalCenter: parent.verticalCenter
        anchors.left: parent.left
        anchors.leftMargin: 30
    }
}