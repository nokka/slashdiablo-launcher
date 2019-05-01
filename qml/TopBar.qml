import QtQuick 2.0
import QtQuick.Controls 2.5

Rectangle {
        id: topbar
        anchors.top: root.top;
        width: parent.width
        height: 80
        color: "#100b17"

        Text {
            id: title
            text: "Slashdiablo"
            color: "#b0adb3"
            x: 25; anchors.verticalCenter: parent.verticalCenter
            font.pointSize: 24; font.bold: true
        }

        Button {
            id: closeButton
            width: 25; height: 25
            
            anchors.top: parent.top;
            anchors.right: parent.right;
            
            background: Rectangle { 
                color: "#00000000"
                Image {
                    source: "../assets/close.png";
                    fillMode: Image.PreserveAspectFit
                    anchors.fill: parent; 
                }
            }

            onClicked: QmlBridge.closeLauncher()
        }

        Button {
            id: minimizeButton
            width: 25; height: 25
            
            x: 25; anchors.verticalCenter: parent.verticalCenter
            
            text: "Minimize"

            onClicked: QmlBridge.minimizeLauncher()
        }

}