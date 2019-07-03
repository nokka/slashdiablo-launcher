import QtQuick 2.12
import QtQuick.Controls 2.5

Button {
    id: defaultButton
    property var colorUp: "#0b86ba"
    property var colorDown: "#1c99c7"

    contentItem: Text {
        color: "#ffffff"
        text: parent.text
        font.family: montserrat.name
        horizontalAlignment: Text.AlignHCenter
        verticalAlignment: Text.AlignVCenter
        elide: Text.ElideRight
    }

    background: Rectangle {
        color: defaultButton.down ? colorDown : colorUp
    }
}
