import QtQuick 2.12				//Item

Item {
    id: settingsDelegate
    width: parent.width
    height: 50

    // Left active indicator border.
    Rectangle {
        color: settingsDelegate.ListView.isCurrentItem ? "#ab4432" : "#57555e"
        width: 3
        height: parent.height
        anchors.left: parent.left
    }

    // Name and location.
    Item {
        height: parent.height
        width: 200

        anchors.top: parent.top
        anchors.left: parent.left
        anchors.leftMargin: 20
        anchors.topMargin: 10

        Title {
            color: settingsDelegate.ListView.isCurrentItem ? "#fff" : "#bababa"
            anchors.top: parent.top
            text: getName()
            font.pixelSize: 14
            anchors.topMargin: 5
        }
    }

    // Mods.
    Item {
        height: 25
        width: 60
        anchors.right: parent.right
        anchors.verticalCenter: parent.verticalCenter
        anchors.rightMargin: 28

        Row {
            spacing: 2
            layoutDirection: Qt.RightToLeft
            width: parent.width

            // HD circle.
            Rectangle { 
                visible: model.hd
                color: "#009fb8"
                width: 25
                height: 25
                radius: (width * 0.5)
                
                SText {
                    anchors.centerIn: parent
                    text: "HD"
                    font.pixelSize: 10
                }
            }

            // Maphack circle.
            Rectangle {
                visible: model.maphack
                color: "#038a66"
                width: 25
                height: 25
                radius: (width * 0.5)
                
                SText {
                    anchors.centerIn: parent
                    text: "MH"
                    font.pixelSize: 10
                }
            }
        }
    }

    // Delete button.
    Image {
        id: deleteIcon
        height: 18
        anchors.right: parent.right
        anchors.verticalCenter: parent.verticalCenter
        anchors.rightMargin: 5
        fillMode: Image.PreserveAspectFit
        source: "assets/svg/delete.svg"

        MouseArea {
            anchors.fill: parent
            cursorShape: Qt.PointingHandCursor
            onClicked: {
                settings.deleteGame(model.id)
            }
        }
    }

    Separator{
        color: "#21211f"
    }

    MouseArea {
        id: mousearea
        anchors.top: parent.top
        anchors.left: parent.left
        width: (parent.width * 0.90)
        height: parent.height
        hoverEnabled: true
        cursorShape: Qt.PointingHandCursor
        onClicked: {
            settingsDelegate.ListView.view.currentIndex = index;
        }
    }

    function getName() {
        var path = model.location
        var parts = path.split("/")
        
        var name = parts[parts.length - 1]
        if(name == "") {
            name = "New game"
        }

        return name
    }
}
