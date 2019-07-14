import QtQuick 2.12                 // Item
import QtQuick.Layouts 1.3          // RowLayout

import "componentCreator.js" as ComponentCreator


Item {
    // Background.
    Rectangle {
        anchors.fill: parent
        color: "#13031a"
        opacity: 0.4
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
    Item {
        height: parent.height
        width: parent.width * 0.30
        anchors.right: parent.right

        RowLayout {
            id: statusMenu
            anchors.fill: parent
            Layout.alignment: Qt.AlignRight | Qt.AlignVCenter
            spacing: 0

            // Server status.
            Item {
                width: 120
                Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
                height: parent.height
                
                Item {
                    width: 110
                    height: 20
                    anchors.verticalCenter: parent.verticalCenter
                    anchors.horizontalCenter: parent.horizontalCenter

                    SText {
                        text: "SERVER STATUS"
                        font.bold: true      
                        anchors.verticalCenter: parent.verticalCenter                  
                    }

                    // Status circle.
                    Rectangle {
                        width: 12
                        height: 12
                        color: "#0B8A0F"
                        radius: (width * 0.5)
                        anchors.verticalCenter: parent.verticalCenter
                        anchors.right: parent.right
                        
                    }
                }
            }

            // Users online.
            Item {
                width: 80
                Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
                height: parent.height

                Item {
                    width: 50
                    height: 20
                    anchors.verticalCenter: parent.verticalCenter
                    anchors.horizontalCenter: parent.horizontalCenter

                    SText {
                        text: "520"
                        font.bold: true
                        anchors.verticalCenter: parent.verticalCenter
                    }

                    Image {
                        id: usersIcon
                        fillMode: Image.PreserveAspectFit
                        anchors.verticalCenter: parent.verticalCenter
                        anchors.right: parent.right
                        width: 20
                        height: 20
                        source: "assets/svg/users.svg"
                    }
                }
            }

             // Options.
            Item {
                width: 90
                Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
                height: parent.height

                Item {
                    width: 90
                    height: 20
                    anchors.verticalCenter: parent.verticalCenter
                    anchors.horizontalCenter: parent.horizontalCenter

                    Image {
                        id: optionsIcon
                        fillMode: Image.PreserveAspectFit
                        anchors.verticalCenter: parent.verticalCenter
                        anchors.left: parent.left
                        width: 13
                        height: 13
                        source: "assets/svg/options.svg"
                    }

                    SText {
                        text: "OPTIONS"
                        font.bold: true
                        anchors.verticalCenter: parent.verticalCenter
                        anchors.right: parent.right
                        anchors.rightMargin: 20

                        MouseArea {
                            anchors.fill: parent
                            cursorShape: Qt.PointingHandCursor
                            onClicked: stack.push(ComponentCreator.createSettingsView(this, null))
                        }
                    }
                }
            }
        }
    }
}