import QtQuick 2.12				//Item

Rectangle {
    id: settingsDelegate
    width: parent.width
    height: 50
    color: ListView.isCurrentItem ? "#121212" : "#00000000"

    Image {
        id: chevronRight
        width: 13
        height: 13
        fillMode: Image.PreserveAspectFit
        anchors.verticalCenter: parent.verticalCenter
        anchors.left: parent.left
        source: "assets/svg/chevron-right.svg"
    }

    // Name and location.
    Item {
        height: parent.height
        width: 200

        anchors.top: parent.top
        anchors.left: parent.left
        anchors.leftMargin: 20
        anchors.topMargin: 10

        SText {
            anchors.top: parent.top
            text: "Diablo II"
            font.pixelSize: 12
            font.bold: true
        }

        SText {
            color: "#8d8d8d"
            anchors.topMargin: 15
            anchors.top: parent.top
            text: model.location
            font.pixelSize: 12
        }
    }

    // Mods.
    Item {
        height: parent.height
        width: 80
        anchors.right: parent.right
    }

    Rectangle {
        height: 1

        color: 'purple'
        anchors {
            left: settingsDelegate.left
            right: settingsDelegate.right
            bottom: settingsDelegate.bottom
        }
    }

    MouseArea {
        anchors.fill: parent
        cursorShape: Qt.PointingHandCursor
        onClicked: {
            settingsDelegate.ListView.view.currentIndex = index;
        }
    }
}