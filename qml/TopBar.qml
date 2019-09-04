import QtQuick 2.12                 // Item
import QtQuick.Layouts 1.3          // RowLayout

Item {
    id: derp
    property string activeMenuItem: "launch"
    property string menuGradientStart: "#00000000"
    property string menuGradientStop: "#363636"
    property var menuSources: { 
        "launch": "LauncherView.qml",
        "rules": "RulesView.qml",
        "community": "CommunityView.qml",
        "armory": "ArmoryView.qml"
    }

    // Background.
    Rectangle {
        anchors.fill: parent
        color: "#000000"
        opacity: 0.0
    }

    // Main menu.
    Item {
        width: 300
        height: parent.height
        anchors.left: parent.left
        anchors.leftMargin: 20
        
        RowLayout {
            id: mainMenu
            anchors.fill: parent
            Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
            spacing: 6

            Item {
                Layout.alignment: Qt.AlignRight | Qt.AlignVCenter
                height: parent.height
                width: 100

                Rectangle {
                    visible: activeMenuItem == "launch"
                    opacity: 0.1
                    anchors.fill: parent
                    gradient: Gradient {
                           GradientStop { position: 0.2; color: menuGradientStart }
                            GradientStop { position: 1.0; color: menuGradientStop }
                    }
                }
                
                MenuItem {
                    text: "LAUNCH"
                    onClicked: function() {
                        activeMenuItem = "launch"
                        contentLoader.source = menuSources.launch
                    }
                }
            }

            Item {
                Layout.alignment: Qt.AlignRight | Qt.AlignVCenter
                height: parent.height
                width: 100

                Rectangle {
                    visible: activeMenuItem == "community"
                    opacity: 0.1
                    anchors.fill: parent
                    gradient: Gradient {
                           GradientStop { position: 0.2; color: menuGradientStart }
                            GradientStop { position: 1.0; color: menuGradientStop }
                    }
                }
                
                MenuItem {
                    text: "COMMUNITY"

                    onClicked: function() {
                        activeMenuItem = "community"
                        contentLoader.source = menuSources.community
                    }
                }
            }

            Item {
                Layout.alignment: Qt.AlignRight | Qt.AlignVCenter
                height: parent.height
                width: 100

                Rectangle {
                    visible: activeMenuItem == "rules"
                    opacity: 0.1
                    anchors.fill: parent
                    gradient: Gradient {
                           GradientStop { position: 0.2; color: menuGradientStart }
                            GradientStop { position: 1.0; color: menuGradientStop }
                    }
                }
                
                MenuItem {
                    text: "RULES"

                    onClicked: function() {
                        activeMenuItem = "rules"
                        contentLoader.source = menuSources.rules
                    }
                }
            }

            Item {
                Layout.alignment: Qt.AlignRight | Qt.AlignVCenter
                height: parent.height
                width: 100

                Rectangle {
                    visible: activeMenuItem == "armory"
                    opacity: 0.1
                    anchors.fill: parent
                    gradient: Gradient {
                           GradientStop { position: 0.2; color: menuGradientStart }
                            GradientStop { position: 1.0; color: menuGradientStop }
                    }
                }
            
                MenuItem {
                    text: "ARMORY"
                    onClicked: function() {
                        activeMenuItem = "armory"
                        contentLoader.source = menuSources.armory
                    }
                }
            }
        }
    }

    // Status panel.
    Item {
        height: parent.height
        width: parent.width * 0.3215
        anchors.right: parent.right

        RowLayout {
            id: statusMenu
            anchors.fill: parent
            Layout.alignment: Qt.AlignRight | Qt.AlignVCenter
            spacing: 0

            // Server status.
            Item {
                width: 110
                Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
                height: parent.height
                
                Item {
                    width: 110
                    height: 20
                    anchors.verticalCenter: parent.verticalCenter
                    anchors.horizontalCenter: parent.horizontalCenter

                    Title {
                        text: "SERVER STATUS"    
                        anchors.verticalCenter: parent.verticalCenter                  
                    }

                    // Status circle.
                    Rectangle {
                        width: 12
                        height: 12
                        color: "#0aa646"
                        radius: (width * 0.5)
                        border.color: "#61e897"
                        border.width: 1
                        anchors.verticalCenter: parent.verticalCenter
                        anchors.right: parent.right 
                    }
                }
            }

            // Users online.
            /*Item {
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
            }*/

             // Settings.
            Item {
                width: 120
                Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
                height: parent.height

                Item {
                    width: 120
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

                    Title {
                        text: "GAME SETTINGS"
                        font.bold: true
                        anchors.verticalCenter: parent.verticalCenter
                        anchors.right: parent.right
                        anchors.rightMargin: 5

                        MouseArea {
                            anchors.fill: parent
                            cursorShape: Qt.PointingHandCursor
                            onClicked: settingsPopup.open()
                        }
                    }
                }
            }
        }
    }

    Separator{
        color: "#4a402c"
    }
}
