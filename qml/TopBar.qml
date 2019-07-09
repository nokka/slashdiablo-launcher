import QtQuick 2.12                 // Item
import QtQuick.Layouts 1.3          // RowLayout

Item {
    // Background.
    Rectangle {
        anchors.fill: parent
        color: "#1F1109"
        opacity: 0.5
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

}