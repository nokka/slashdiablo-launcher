import QtQuick 2.12

Text {
    color: "#ffffff"
    text: parent.text
    font.family: d2Font.name
    horizontalAlignment: Text.AlignHCenter
    verticalAlignment: Text.AlignVCenter
    elide: Text.ElideRight
}