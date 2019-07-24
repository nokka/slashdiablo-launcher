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
            text: getName()
            font.pixelSize: 12
            font.bold: true
        }

        SText {
            color: "#8d8d8d"
            anchors.topMargin: 15
            anchors.top: parent.top
            text: (model.location != "" ? model.location : "No directory set")
            font.pixelSize: 12
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

            Rectangle { 
                visible: model.hd
                color: "#D47F31"
                width: 25
                height: 25
                radius: (width * 0.5)
                
                SText {
                    anchors.centerIn: parent
                    text: "HD"
                    font.pixelSize: 10
                    font.bold: true      
                }
            }
            
            Rectangle {
                visible: model.maphack
                color: "#1A8EBF"
                width: 25
                height: 25
                radius: (width * 0.5)
                
                SText {
                    anchors.centerIn: parent
                    text: "MH"
                    font.pixelSize: 10
                    font.bold: true      
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
        color: "#6E3E87"
    }

    MouseArea {
        anchors.top: parent.top
        anchors.left: parent.left
        width: (parent.width * 0.90)
        height: parent.height
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