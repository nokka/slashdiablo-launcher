import QtQuick 2.4

Rectangle {
    property int fontSize: 12

    width: parent.width - 15
    height: 32
    color: "#00000000"
    
    Rectangle {
		anchors.fill: parent
        height: parent.height
        radius: 5
        opacity: 0.7
		color: "#1f1b16"
	}

    Text {
        id: rankItem 
        font.family: montserrat.name
        font.pixelSize: fontSize
        color: "#ffffff"
        text: model.rank
        anchors.verticalCenter: parent.verticalCenter
        anchors.left: parent.left
        anchors.leftMargin: 10
    }

    Text {
        id: classItem 
        font.family: montserrat.name
        font.pixelSize: fontSize
        color: "#f3e6d0"
        text: model.class
        anchors.verticalCenter: parent.verticalCenter
        anchors.left: parent.left
        anchors.leftMargin: rankItem.width + 20
    }

    Text {
        color: "#f3e6d0"
        font.family: montserrat.name
        font.pixelSize: fontSize
        text: model.name
        anchors.verticalCenter: parent.verticalCenter
        anchors.left: parent.left
        anchors.leftMargin: classItem.width + 40

    }

    Text {
        color: "#517d8a"
        font.family: montserrat.name
        font.pixelSize: fontSize
        text: model.level
        anchors.verticalCenter: parent.verticalCenter
        anchors.rightMargin: 20
        anchors.right: parent.right
    }
}