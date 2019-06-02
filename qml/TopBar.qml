import QtQuick 2.4
import QtQuick.Controls 2.5

Rectangle {

    /*Text {
        id: title
        font.family: d2Font.name
        text: "Slashdiablo"
        color: "#fcfcfc"
        x: 25; anchors.verticalCenter: parent.verticalCenter
        font.pointSize: 24; font.bold: true
    }*/

    // Repeated background image for the top bar.
    Row {
        Repeater {
            model: 12
            Image {
                source: "assets/top_border_repeat_darker.png";
            }
        }
    }

    // Crackled top bar to the left.
    Rectangle {
        width: 100
        height: 39
        anchors.top: parent.top
        anchors.left: parent.left

        Image {
            source: "assets/top_left_darker.png"
        }
    }
    
    // The skull in the middle.
    Rectangle {
        width: 442
        height: 71
        anchors.top: parent.top
        anchors.horizontalCenter: parent.horizontalCenter
        
        Image {
            source: "assets/top_skull_eyes.png"
        }
    }

    // Crackled top bar to the right.
    Rectangle {
        width: 100
        height: 39
        anchors.top: parent.top
        anchors.right: parent.right

        Image {
            source: "assets/top_right_darker.png"
        }

        // Close button.
        Button {
            id: closeButton
            width: 16; height: 16
            anchors.margins: 15
            
            anchors.top: parent.top;
            anchors.right: parent.right;
            
            background: Rectangle { 
                color: "#751a15"
                radius: 8
            }

            MouseArea {
                anchors.fill: parent
                cursorShape: Qt.PointingHandCursor
                onClicked: QmlBridge.closeLauncher()
            }
        }

        // Minimize button.
        Button {
            id: minimizeButton
            width: 16; height: 16
            anchors.margins: 45;
            anchors.topMargin: 15
            
            anchors.top: parent.top;
            anchors.right: parent.right;

            background: Rectangle { 
                color: "#ffc130"
                radius: 8
            }

            MouseArea {
                anchors.fill: parent
                cursorShape: Qt.PointingHandCursor
                onClicked: mainWindow.showMinimized()
            }
        }
    }

    // Draggable area to move the window.
    Item {
        id: draggable
        width: parent.width * 0.80
        height: 40
        anchors.top: parent.top;
        anchors.horizontalCenter: parent.horizontalCenter

        MouseArea {
            id: draggable_mousearea
            anchors.fill: parent;
            property variant clickPos: "1,1"

            onPressed: {
                clickPos  = Qt.point(mouse.x,mouse.y)
            }

            onPositionChanged: {
                var delta = Qt.point(mouse.x-clickPos.x, mouse.y-clickPos.y)
                mainWindow.x += delta.x;
                mainWindow.y += delta.y;
            }
        }
    }
}