import QtQuick 2.12             // Gradient
import QtQuick.Controls 2.5     // Button

Button {
    id: sbutton

    property int fontSize: 12
    property string label: ""
    property int borderRadius: (parent.width * 0.5)
    property string backgroundColor: "#00000000"
    property string borderColor: "#800507"
    property alias cursorShape: mouseArea.cursorShape
    
    Text {
        text: label
        color: "#fff"
        font.family: roboto.name
        anchors.verticalCenter: parent.verticalCenter
        anchors.horizontalCenter: parent.horizontalCenter
        font.pixelSize: fontSize;
    }

    background: Rectangle {
        color: backgroundColor
        radius: borderRadius
        border.width: 2
        border.color: borderColor
    }

    PropertyAnimation {
        id: animateIn
        target: sbutton
        properties: "background.border.color";
        to: "#646466";
        duration: 100
    }

    PropertyAnimation {
        id: animateOut
        target: sbutton
        properties: "background.border.color";
        to: borderColor;
        duration: 200
    }

    MouseArea {
        id: mouseArea
        anchors.fill: parent

        // Disable click on mouse area, making the event propagate
        // to the parent button. We need the mouse area to override
        // the button mouse cursor property.
        onPressed:  mouse.accepted = false
    }

    onHoveredChanged: hovered ? animateIn.start() : animateOut.start();
}
