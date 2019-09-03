import QtQuick 2.12
import QtQuick.Controls 2.5     // Button
import QtGraphicalEffects 1.0   // Gradient

Button {
    id: xbutton

    property int fontSize: 15
    property string label: ""
    property bool clickable: true
    property bool enabled: clickable
    
    Text {
        text: label
        color: clickable ? "#f3e6d0" : "#737373"
        font.family: beaufortbold.name
        anchors.verticalCenter: parent.verticalCenter
        anchors.horizontalCenter: parent.horizontalCenter
        font.pixelSize: fontSize;
    }

    background: Rectangle {
        anchors.fill: parent
        color: "#21060D"
        radius: 5

        // Inner fill.
        Rectangle {
            width: (parent.width-2)
            height: (parent.height-2)
            anchors.centerIn: parent
            color: (hovered ? "#0391C4" : "#027EB4")
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
