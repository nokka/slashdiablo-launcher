import QtQuick 2.12
import QtQuick.Controls 2.5     // Button
import QtGraphicalEffects 1.0   // Gradient

Button {
    id: plainbutton

    property int fontSize: 15
    property string label: ""
    property string activeColor: "#fff"
    property bool active: false
    
    Text {
        text: label
        color: (active ? activeColor : "grey")
        font.family: beaufortbold.name
        anchors.verticalCenter: parent.verticalCenter
        anchors.horizontalCenter: parent.horizontalCenter
        font.pixelSize: fontSize;
    }

    background: Rectangle {
        anchors.fill: parent
        color: "#00000000"
    }
}
