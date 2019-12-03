import QtQuick 2.12

Rectangle {
    color: "#000000"

    Item {
        id: ladderView
        width: parent.width; height: parent.height
        anchors.left: parent.left
        anchors.right: parent.right
        anchors.leftMargin: 20
        anchors.rightMargin: 20
        
        LadderTable{}
    }
}
