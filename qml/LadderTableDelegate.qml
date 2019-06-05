import QtQuick 2.4

Rectangle{
    property var character: model
    
    width: 200
    height: 20
    color: ListView.isCurrentItem ? "#003366" : "#585858"
    border.color: "gray"
    border.width: 1

    Text{
        anchors.centerIn: parent
        color: "black"
        text: model.class
    }
}