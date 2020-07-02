import QtQuick 2.12
import QtQuick.Controls 2.5

Button {
    id: toggleButton

    property bool active: false
    property int fontSize: 11
    property string label: ""
    property alias cursorShape: mouseArea.cursorShape
    
    Text {
        text: label
        color: active ? "#fff" : "#57555e"
        font.family: beaufortbold.name
        anchors.verticalCenter: parent.verticalCenter
        anchors.horizontalCenter: parent.horizontalCenter
        font.pixelSize: fontSize;
    }

    background: Rectangle {
        color: "#0C0C0C"
        border.color: "#000000"
        radius: 0
        border.width: 1
    }

    MouseArea {
        id: mouseArea
        hoverEnabled: true
        anchors.fill: parent

        cursorShape: containsMouse? Qt.PointingHandCursor : Qt.ArrowCursor

        // Disable click on mouse area, making the event propagate
        // to the parent button. We need the mouse area to override
        // the button mouse cursor property.
        onPressed: {
            active = !active
            mouse.accepted = false
        }
    }
}
