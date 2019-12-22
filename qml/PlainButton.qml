import QtQuick 2.12
import QtQuick.Controls 2.5

Button {
    id: plainbutton

    property int fontSize: 15
    property string label: ""
    property string backgroundColor: "#1e1b26"
    property string colorHovered: "#ab4432"
    property string borderColor: "#000000"
    property int radius: 0
    property bool active: false
    property bool activatable: false
    property bool clickable: false
    
    Text {
        text: label
        color: (activatable ? (active ? "#fff" : "#737373") : "#fff")
        font.family: beaufortbold.name
        anchors.verticalCenter: parent.verticalCenter
        anchors.horizontalCenter: parent.horizontalCenter
        font.pixelSize: fontSize;
    }

    background: Rectangle {
        anchors.fill: parent
        color: borderColor
        radius: radius

        // Inner fill.
        Rectangle {
            radius: radius
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
