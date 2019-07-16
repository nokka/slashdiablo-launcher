import QtQuick 2.12             // Gradient
import QtQuick.Controls 2.5     // Button

Button {
    id: sbutton

    property int fontSize: 12
    property string label: ""
    property int borderRadius: (parent.width * 0.5)
    property alias cursorShape: mouseArea.cursorShape
    
    Text {
        text: label
        color: "#f3e6d0"
        font.family: roboto.name
        anchors.verticalCenter: parent.verticalCenter
        anchors.horizontalCenter: parent.horizontalCenter
        font.pixelSize: fontSize;
    }

    background: Rectangle {
        color: "#00000000"
        radius: borderRadius
        border.width: 2
        border.color: "#6E3E87"
    }

    PropertyAnimation {
        id: animateIn
        target: sbutton
        properties: "background.border.color";
        to: "#CE82F5";
        duration: 200
    }

    PropertyAnimation {
        id: animateOut
        target: sbutton
        properties: "background.border.color";
        to: "#6E3E87";
        duration: 300
    }

    MouseArea {
        id: mouseArea
        anchors.fill: parent

        // Disable click on mouse area, making the event propagate
        // to the parent button.
        onPressed:  mouse.accepted = false
    }

    onHoveredChanged: hovered ? animateIn.start() : animateOut.start();
}