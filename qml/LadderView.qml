import QtQuick 2.12
import QtQuick.Controls 2.5
import QtGraphicalEffects 1.13

Rectangle {
    color: "#030202"

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
