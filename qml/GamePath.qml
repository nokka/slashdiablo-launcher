
import QtQuick 2.2

Rectangle {
    id: gamepath_view
    width: parent.width
    height: parent.height
    anchors.fill: parent;

    color: "#100b17"

    Text {
        text: "Set your game path"
        color: "#b0adb3"
        x: 25; anchors.verticalCenter: parent.verticalCenter
        font.pointSize: 24; font.bold: true
    }

    //Component.onCompleted: QmlBridge.patchGame()
}