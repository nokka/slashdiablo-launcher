import QtQuick 2.12

Rectangle {
    property int fontSize: 12

    width: parent.width
    height: 33
    radius: 5
    color: "#00000000"

    Title {
        id: rankItem
        width: 15
        font.pixelSize: fontSize
        color: "#9e998b"
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
        color: "#3d3b36"
        text: model.class
        anchors.verticalCenter: parent.verticalCenter
        anchors.left: parent.left
        anchors.leftMargin: rankItem.width + 20
    }

    Title {
        color: mousearea.containsMouse ? "#ffffe6" : "#c4b58b"
        font.pixelSize: 13
        text: model.name
        anchors.verticalCenter: parent.verticalCenter
        anchors.left: parent.left
        anchors.leftMargin: classItem.width + 40

    }

    Text {
        color: "#069499"
        font.family: roboto.name
        font.pixelSize: fontSize
        text: model.level
        anchors.verticalCenter: parent.verticalCenter
        anchors.rightMargin: 20
        anchors.right: parent.right
    }

    MouseArea {
        id: mousearea
        anchors.fill: parent
        cursorShape: Qt.PointingHandCursor
        onClicked: {
            console.log("clicked ladder item")
            Qt.openUrlExternally("https://armory.slashdiablo.net/character/"+model.name.toLowerCase())
        }
        hoverEnabled: true
    }

    Separator{
        color: "#3d3b36"
    }
}