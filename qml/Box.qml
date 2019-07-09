import QtQuick 2.12

Rectangle {
    property string title: ""
    border.width: 2
    border.color: "#000000"
    color: "#00000000"

    Rectangle {
        id: background
        width: parent.width-4
        height: parent.height-4
        anchors.verticalCenter: parent.verticalCenter
        anchors.horizontalCenter: parent.horizontalCenter
        border.width: 1
        border.color: "#85817d"
        
        Item {
            width: parent.width-2
            height: parent.height-2
            anchors.verticalCenter: parent.verticalCenter
            anchors.horizontalCenter: parent.horizontalCenter
            Image { source: "assets/tmp/tyriel.jpg"; fillMode: Image.Stretch; anchors.fill: parent;  opacity: 1.0 }
        }
    }

    Text {
        color: "#ffffff"
        font.family: roboto.name
        font.pixelSize: 20
        text: title
        anchors.left: parent.left
        anchors.bottom: parent.bottom
        anchors.margins: 20
    }
}