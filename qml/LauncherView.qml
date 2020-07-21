import QtQuick 2.12

Item {
    id: launchView
    property bool configButtonHovered: false
    width: parent.width; height: parent.height
    
    Item {
        id: logobg
        width: 210.6
        height: 240.3
        anchors.top: parent.top
        anchors.topMargin: 26
        anchors.horizontalCenter: parent.horizontalCenter
        Image { source: "assets/logo-bg.png"; anchors.fill: parent; fillMode: Image.PreserveAspectFit; opacity: 1.0 }
    }

    Item {
        id: logotext
        width: 216
        height: 63.9
        anchors.horizontalCenter: parent.horizontalCenter
        anchors.top: parent.top
        anchors.topMargin: 116
        Image { source: "assets/logo-text.png"; anchors.fill: parent; fillMode: Image.PreserveAspectFit; opacity: 1.0 }
    }

    // Sidebar to the right.
    Item {
        id: sidebar
        width: 350
        height: parent.height
        anchors.right: parent.right
        
        NewsTable{}

        Item {
            width: 80
            height: 40
            anchors.bottom: parent.bottom
            anchors.right: parent.right

            Title {
                text: settings.buildVersion
                font.pixelSize: 10
                anchors.left: parent.left
                anchors.verticalCenter: parent.verticalCenter
            }

            Image {
                id: debugIcon
                fillMode: Image.Pad
                anchors.verticalCenter: parent.verticalCenter
                anchors.right: parent.right
                anchors.rightMargin: 20
                width: 16
                height: 16
                source: "assets/icons/bug.png"
                opacity: configButtonHovered ? 1.0 : 0.5

                MouseArea {
                    anchors.fill: parent
                    hoverEnabled: true
                    cursorShape: Qt.PointingHandCursor
                    onClicked: {
                        settings.openConfigPath()
                        configButtonHovered = false
                    }
                    
                    onEntered: {
                        configButtonHovered = true
                    }
                    onExited: {
                        configButtonHovered = false
                    }
                }
            }
        }
    }
        
    // Bottom bar.
    Item {
        id: bottombar
        width: (parent.width-sidebar.width); height: 80
        anchors.bottom: parent.bottom;

        // Patcher including progress bar.
        Patcher{
            width: 650
        }
    }
}
