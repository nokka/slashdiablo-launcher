import QtQuick 2.12
import QtQuick.Controls 2.5

ItemDelegate {
    width: parent.width
    height: 30

    contentItem: Title {
        id:textItem
        text: modelData
        color: hovered ? "#ffffff" : "#57555e"
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
