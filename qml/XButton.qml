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

        // Outer border.
        gradient: Gradient {
            GradientStop { 
                id: gradientStart
                position: 0.0;
                color: (hovered ? "#615840" : "#362d14")
            }
            GradientStop {
                id: gradientStop
                position: 1.0;
                color: (hovered ? "#d9b16c" : "#a17b2f")
            }
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
            id: fill
            width: (parent.width-15)
            height: (parent.height-15)
            anchors.centerIn: parent
            color: (hovered ? "#000c14" : "#040405")
        }
    }
    
    PropertyAnimation {
        id: animateIn
        target: fill
        properties: "color";
        to: "#000c14";
        duration: 200
    }

    PropertyAnimation {
        id: animateOut
        target: fill
        properties: "color";
        to: "#040405";
        duration: 100
    }

    // Gradient animations in.
    /*PropertyAnimation {
        id: animateGradientTopIn
        target: gradientStart
        properties: "color"
        to: "#b5ac98"
        duration: 100
    }

    PropertyAnimation {
        id: animateGradientBottomIn
        target: gradientStop
        properties: "color"
        to: "#d9b16c"
        duration: 100
    }

    // Gradient animations out.
    PropertyAnimation {
        id: animateGradientTopOut
        target: gradientStart
        properties: "color"
        to: "#362d14"
        duration: 0
    }

    PropertyAnimation {
        id: animateGradientBottomOut
        target: gradientStop
        properties: "color"
        to: "#a17b2f"
        duration: 0
    }*/

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

    /*onHoveredChanged: {
        if (hovered) {
            animateIn.start();
            animateGradientTopIn.start();
            animateGradientBottomIn.start();
            
        } else {
            animateOut.start();
            animateGradientTopOut.start();
            animateGradientBottomOut.start();
        }
    }*/
}
