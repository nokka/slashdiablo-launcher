import QtQuick 2.12
import QtQuick.Controls 2.5     // Button
import QtGraphicalEffects 1.0   // Gradient

Button {
    id: plainbutton

    property int fontSize: 15
    property string label: ""
    property bool clickable: true
    property bool enabled: clickable
    property string backgroundColor: "#5c0202"
    property string colorHovered: "#3b0000"
    property string borderColor: "#000000"
    
    Text {
        text: label
        color: clickable ? "#fff" : "#737373"
        font.family: beaufortbold.name
        anchors.verticalCenter: parent.verticalCenter
        anchors.horizontalCenter: parent.horizontalCenter
        font.pixelSize: fontSize;
    }

    background: Rectangle {
        anchors.fill: parent
        color: borderColor
        radius: 2

        // Inner fill.
        Rectangle {
            radius: 2
            width: (parent.width-2)
            height: (parent.height-2)
            anchors.centerIn: parent
            color: (hovered ? colorHovered : backgroundColor)
        }
    }

    MouseArea {
        id: mouseArea
        hoverEnabled: true
        anchors.fill: parent

        cursorShape: containsMouse ? ((clickable) ? Qt.PointingHandCursor : Qt.ForbiddenCursor) : Qt.ArrowCursor

        // Disable click on mouse area, making the event propagate
        // to the parent button. We need the mouse area to override
        // the button mouse cursor property.
        onPressed: {
            if(!clickable) {
                return false
            }
            mouse.accepted = false
        }
    }
}
