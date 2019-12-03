import QtQuick 2.12

Item {
    property int fontSize: 12
    property string fontFamily: beaufortbold.name
    property string content: ""

    Text {
        color: "#b5b5b5"
        font.pixelSize: fontSize
        font.family: fontFamily
        text: content
        anchors.verticalCenter: parent.verticalCenter
    }

    Separator{}
}
