import QtQuick 2.12
//import QtQuick.Controls 2.5
import QtQuick.Layouts 1.3

Item {
    // Background
    Rectangle {
        id: topbar_bg
        anchors.fill: parent
        color: "#000000"
        opacity: 0.2
    }

    // Main menu.
    Item {
        width: 300
        height: 50

        RowLayout {
            id: mainMenu
            anchors.fill: parent
            Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
            spacing: 6
            
            Item {
                Layout.alignment: Qt.AlignRight | Qt.AlignVCenter
                height: parent.height
                width: 100
                
                Text {
                    anchors.verticalCenter: parent.verticalCenter
                    anchors.horizontalCenter: parent.horizontalCenter
                    id: redditItem
                    font.family: montserrat.name
                    font.pixelSize: 14
                    color: "#ffffff"
                    text: "REDDIT"
                } 
            }

            Item {
                Layout.alignment: Qt.AlignRight | Qt.AlignVCenter
                height: parent.height
                width: 100
                
                Text {
                    anchors.verticalCenter: parent.verticalCenter
                    anchors.horizontalCenter: parent.horizontalCenter
                    id: armoryItem
                    font.family: montserrat.name
                    font.pixelSize: 14
                    color: "#ffffff"
                    text: "ARMORY"
                } 
            }

            Item {
                Layout.alignment: Qt.AlignRight | Qt.AlignVCenter
                height: parent.height
                width: 100
                
                Text {
                    anchors.verticalCenter: parent.verticalCenter
                    anchors.horizontalCenter: parent.horizontalCenter
                    id: ladderItem
                    font.family: montserrat.name
                    font.pixelSize: 14
                    color: "#ffffff"
                    text: "LADDER"
                } 
            }
        }
    }
    
    // Repeated background image for the top bar.
    /*Row {
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
    }*/
}