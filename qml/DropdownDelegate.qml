import QtQuick 2.12 // Rectangle
import QtQuick.Controls 2.5 // ItemDelegate

ItemDelegate {
    width: parent.width

    contentItem: Text {
        id:textItem
        text: modelData
        color: hovered ? "#ffffff" : "#898989"
        elide: Text.ElideRight
        verticalAlignment: Text.AlignVCenter
        horizontalAlignment: Text.AlignLeft
    }
    background: Rectangle {
        color: hovered ? "#111111" : "#050505"
        width: parent.width
        border.width: 0
    }
}