import QtQuick 2.12

Rectangle {
    property int fontSize: 12

    width: parent.width
    height: 32
    radius: 5
    color: "#00000000"

    Text {
        id: rankItem
        width: 15
        font.family: roboto.name
        font.pixelSize: fontSize
        color: "#ffffff"
        text: model.rank
        anchors.verticalCenter: parent.verticalCenter
        anchors.left: parent.left
        anchors.leftMargin: 10
    }

    Text {
        id: classItem 
        width: 30
        font.family: roboto.name
        font.pixelSize: fontSize
        color: "#f3e6d0"
        text: model.class
        anchors.verticalCenter: parent.verticalCenter
        anchors.left: parent.left
        anchors.leftMargin: rankItem.width + 20
    }

    Text {
        color: "#f3e6d0"
        font.family: roboto.name
        font.pixelSize: fontSize
        text: model.name
        anchors.verticalCenter: parent.verticalCenter
        anchors.left: parent.left
        anchors.leftMargin: classItem.width + 40

    }

    Text {
        color: "#ffe6a1"
        font.family: roboto.name
        font.pixelSize: fontSize
        text: "lvl " + model.level
        anchors.verticalCenter: parent.verticalCenter
        anchors.rightMargin: 20
        anchors.right: parent.right
    }
}