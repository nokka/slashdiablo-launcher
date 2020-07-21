import QtQuick 2.12

Item {
    id: row
    width: parent.width
    height: 35

    Rectangle {
        anchors.fill: parent
        color: (model.index % 2 == 0 ? "#000000" : "#080505")
        opacity: 0.2
    }

    Row {

        TableCell {
            width: row.width * 0.10
            height: row.height
            content: "#" + model.rank
        }

        TableCell {
            width: row.width * 0.10
            height: row.height
            content: model.level
        }

        TableCell {
            width: row.width * 0.10
            height: row.height
            content: model.class
        }

        Item { 
            width: row.width * 0.40
            height: row.height

            Text {
                color: mousearea.containsMouse ? "#ffffe6" : "#c4b58b"
                font.pixelSize: 12
                font.family: beaufortbold.name
                text: model.name
                anchors.verticalCenter: parent.verticalCenter
            }

            Separator{}
        }

        TableCell {
            width: row.width * 0.10
            height: row.height
            content: model.title
        }

        Item { 
            width: row.width * 0.20
            height: row.height

            Text {
                color: (model.status == "alive" ? "#64d168" : "#fa5757")
                font.pixelSize: 12
                font.family: beaufortbold.name
                text: model.status
                anchors.verticalCenter: parent.verticalCenter
                anchors.right: parent.right
            }

            Separator{}
        }
    }

    MouseArea {
        id: mousearea
        anchors.fill: parent
        cursorShape: Qt.PointingHandCursor
        onClicked: {
            Qt.openUrlExternally("https://armory.slashdiablo.net/character/"+model.name.toLowerCase())
        }
        hoverEnabled: true
    }
}
