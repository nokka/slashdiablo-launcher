import QtQuick 2.4

Rectangle {
    property var character: model
 
    width: parent.width - 15
    height: 40
    color: "#00000000"
    
    Rectangle {
		anchors.fill: parent
        height: parent.height
        radius: 5
        opacity: 0.7
		color: "#1f1b16"
	}

    Text {
        id: classItem 
        font.family: montserrat.name
        color: "#f3e6d0"
        text: model.class
        anchors.verticalCenter: parent.verticalCenter
        anchors.left: parent.left
        anchors.leftMargin: 20
    }

    Text {
        color: "#f3e6d0"
        font.family: montserrat.name
        text: model.name
        anchors.verticalCenter: parent.verticalCenter
        anchors.left: parent.left
        anchors.leftMargin: classItem.width + 30

    }

    Text {
        color: "#517d8a"
        font.family: montserrat.name
        text: model.level
        anchors.verticalCenter: parent.verticalCenter
        anchors.rightMargin: 20
        anchors.right: parent.right

    }
}