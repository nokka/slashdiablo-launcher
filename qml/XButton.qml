import QtQuick 2.12
import QtQuick.Controls 2.5     // Button
import QtGraphicalEffects 1.0   // Gradient

Button {
    id: xbutton

    property int fontSize: 15
    property string label: ""
    
    Text {
        text: label
        color: "#f3e6d0"
        font.family: beaufortbold.name
        anchors.verticalCenter: parent.verticalCenter
        anchors.horizontalCenter: parent.horizontalCenter
        font.pixelSize: fontSize;
    }

    background: Rectangle {
        anchors.fill: parent

        // Outer border.
        gradient: Gradient {
            GradientStop { position: 0.0; color: "#362d14" }
            GradientStop { position: 1.0; color: "#a17b2f" }
        }

        // Inner fill.
        Rectangle {
            width: (parent.width-2)
            height: (parent.height-2)
            anchors.centerIn: parent
            color: "#010912"
        }

        // Inner border.
        Rectangle {
            width: (parent.width-12)
            height: (parent.height-12)
            anchors.centerIn: parent
            color: "#042029"
        }

        // Most inner fill.
        Rectangle {
            width: (parent.width-15)
            height: (parent.height-15)
            anchors.centerIn: parent
            color: "#040405"
        }
    }

    MouseArea {
        id: mouseArea
        hoverEnabled: true
        anchors.fill: parent

        cursorShape: containsMouse ? Qt.PointingHandCursor : Qt.ArrowCursor

        // Disable click on mouse area, making the event propagate
        // to the parent button. We need the mouse area to override
        // the button mouse cursor property.
        onPressed: mouse.accepted = false
    }
}