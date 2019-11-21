import QtQuick 2.12                 // Item
import QtQuick.Layouts 1.3          // RowLayout

Item {
    id: derp
    property string activeMenuItem: "launch"
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
        opacity: 0.3
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
                
                MenuItem {
                    text: "LAUNCH"
                    active: (activeMenuItem == "launch")

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
                
                MenuItem {
                    text: "COMMUNITY"
                    active: (activeMenuItem == "community")

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
                
                MenuItem {
                    text: "RULES"
                    active: (activeMenuItem == "rules")

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
            
                MenuItem {
                    text: "ARMORY"
                    active: (activeMenuItem == "armory")

                    onClicked: function() {
                        activeMenuItem = "armory"
                        contentLoader.source = menuSources.armory
                    }
                }
            }
        }
    }

    // Settings.
    Item {
        width: 120; height: parent.height
        anchors.right: parent.right
        anchors.rightMargin: 20

        Item {
            width: 120
            height: 20
            anchors.verticalCenter: parent.verticalCenter
            anchors.horizontalCenter: parent.horizontalCenter

            Image {
                id: optionsIcon
                fillMode: Image.PreserveAspectFit
                anchors.verticalCenter: parent.verticalCenter
                anchors.right: parent.right
                anchors.rightMargin: 15
                width: 16
                height: 16
                source: "assets/icons/cog.png"

                MouseArea {
                    anchors.fill: parent
                    cursorShape: Qt.PointingHandCursor
                    onClicked: settingsPopup.open()
                }
            }
        }
    }

    Separator{}
}
