import QtQuick 2.12 // Rectangle
import QtQuick.Controls 2.5 // ItemDelegate

ItemDelegate {
    width: parent.width
    height: 30

    contentItem: Title {
        id:textItem
        text: modelData
        color: hovered ? "#ffffff" : "#f3e6d0"
        verticalAlignment: Text.AlignVCenter
        horizontalAlignment: Text.AlignLeft
    }
    background: Rectangle {
        color: hovered ? "#111111" : "#050505"
        width: parent.width
        border.width: 0
        opacity: 0.2
    }
}
