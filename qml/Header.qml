import QtQuick 2.12     // Text

Text {
    color: "#ffffff"
    text: parent.text
    font.family: roboto.name
    horizontalAlignment: Text.AlignHCenter
    verticalAlignment: Text.AlignVCenter
    elide: Text.ElideRight
}