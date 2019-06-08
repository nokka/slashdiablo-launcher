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
    }
}