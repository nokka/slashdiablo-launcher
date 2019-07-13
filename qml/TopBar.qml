import QtQuick 2.12                 // Item
import QtQuick.Layouts 1.3          // RowLayout

Item {
    // Background.
    Rectangle {
        anchors.fill: parent
        color: "#13031a"
        opacity: 0.7
    }

    // Main menu.
    Item {
        width: 300
        height: parent.height

        RowLayout {
            id: mainMenu
            anchors.fill: parent
            Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
            spacing: 6
            
            Item {
                Layout.alignment: Qt.AlignRight | Qt.AlignVCenter
                height: parent.height
                width: 100
                
                MenuItem {
                    text: "HOME"
                }
            }

            Item {
                Layout.alignment: Qt.AlignRight | Qt.AlignVCenter
                height: parent.height
                width: 100
                
                MenuItem {
                    text: "COMMUNITY"
                }
            }

            Item {
                Layout.alignment: Qt.AlignRight | Qt.AlignVCenter
                height: parent.height
                width: 100
                
                MenuItem {
                    text: "LADDER"
                }
            }

            Item {
                Layout.alignment: Qt.AlignRight | Qt.AlignVCenter
                height: parent.height
                width: 100
                
                MenuItem {
                    text: "ARMORY"
                }
            }
        }
    }

    // Status panel.
    Rectangle {
        height: parent.height
        width: parent.width * 0.30
        anchors.right: parent.right
        color: "blue"

        RowLayout {
            id: statusMenu
            anchors.fill: parent
            Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
            spacing: 6

            // Server status.
            Item {
                Layout.alignment: Qt.AlignRight | Qt.AlignVCenter
                height: parent.height
                
                SText {
                    text: "SERVER STATUS"
                    font.bold: true
                    anchors.verticalCenter: parent.verticalCenter
                    anchors.right: parent.right
                    anchors.rightMargin: 20
                }

                // Status circle.
                Rectangle {
                    width: 15
                    height: 15
                    color: "#0B8A9F"
                    radius: width*0.5
                    anchors.verticalCenter: parent.verticalCenter
                    anchors.right: parent.right
                }
            }

            // Users online.
            Item {
                Layout.alignment: Qt.AlignRight | Qt.AlignVCenter
                height: parent.height
                
                SText {
                    text: "129"
                    font.bold: true
                    anchors.verticalCenter: parent.verticalCenter
                    anchors.right: parent.right
                    anchors.rightMargin: 20
                }

                // Status circle.
                Rectangle {
                    width: 15
                    height: 15
                    color: "#0B8A9F"
                    radius: width*0.5
                    anchors.verticalCenter: parent.verticalCenter
                    anchors.right: parent.right
                }
            }
        }
    }
}