import QtQuick 2.12
import QtQuick.Controls 2.5

import "componentCreator.js" as ComponentCreator

Item {
        id: bottombar
        anchors.bottom: parent.bottom;
        width: parent.width; height: 80

        // Patcher including progress bar.
        Patcher{
            width: parent.width * 0.65
        }

        Item {
            width: parent.width * 0.28; height: parent.height
            anchors.verticalCenter: parent.verticalCenter
            anchors.right: parent.right;
            anchors.rightMargin: 18

             // Launch button.
            Button {
                enabled: diablo.playable
                width: (parent.width - settingsButton.width); height: 50
                anchors.verticalCenter: parent.verticalCenter
                anchors.right: parent.right;
                anchors.margins: 20
                anchors.rightMargin: settingsButton.width

                Text {
                    text: "PLAY"
                    color: "#f3e6d0"
                    font.family: roboto.name
                    anchors.verticalCenter: parent.verticalCenter
                    anchors.horizontalCenter: parent.horizontalCenter
                    font.pointSize: 16;
                }

                background: Rectangle {
                    radius: 3
                    gradient: Gradient {
                        GradientStop { position: 0.0; color: "#4398d1" }
                        GradientStop { position: 1.0; color: "#347bad" }
                    }
                }

                onClicked: diablo.launchGame()
            }

            // Settings button.
            Button {
                id: settingsButton
                width: 52; height: 52
                anchors.verticalCenter: parent.verticalCenter
                anchors.right: parent.right;
                anchors.margins: 20
                anchors.rightMargin: 0

                background: Rectangle {
                    color: "#000000"
                    opacity: 0.5
                    radius: 2
                }

                MouseArea {
                    anchors.fill: parent
                    cursorShape: Qt.PointingHandCursor
                    onClicked: stack.push(ComponentCreator.createSettingsView(this, null))
                }
            }

            Text {
                text: "Run DEP fix"
                color: "#ffffff"
                anchors.left: parent.left
                anchors.bottom: parent.bottom

                MouseArea {
                    anchors.fill: parent
                    cursorShape: Qt.PointingHandCursor
                    onClicked: diablo.runDEPFix()
                }
            }
        }
}