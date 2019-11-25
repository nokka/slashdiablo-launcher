import QtQuick 2.12
import QtQuick.Controls 2.12

Switch {
    id: control

    indicator: Rectangle {
        id: indicator
        implicitWidth: 48
        implicitHeight: 26
        x: control.leftPadding
        y: parent.height / 2 - height / 2
        radius: 13
        color: control.checked ? "#ab4432" : "#1f1f1f"
        border.color: control.checked ? "#ab4432" : "#1f1f1f"

        Rectangle {
            id: circle
            width: 26
            height: 26
            radius: 13
            color: control.down ? "#cccccc" : "#ffffff"
        }
    }

    // Animation that runs when the switch is turned "on".
    PropertyAnimation {
        id: animateOn
        target: circle
        properties: "x";
        to: (indicator.width-circle.width);
        duration: 100
    }

    // Animation that runs when the switch is turned "off".
    PropertyAnimation {
        id: animateOff
        target: circle
        properties: "x";
        to: 0
        duration: 100
    }

    onClicked: control.checked ? animateOn.start() : animateOff.start();

    // Update updates the state of the switch without triggering an animation.
    function update() {
        if(control.checked) {
            circle.x = (indicator.width-circle.width)
        } else {
            circle.x = 0
        }
    }
}
