import QtQuick 2.12
import QtQuick.Controls 2.5
import QtQuick.Layouts 1.3		//ColumnLayout

Item {
    id: launchView
    width: parent.width; height: parent.height
    anchors.leftMargin: 20
    
    Item {
        id: logobg
        width: 234
        height: 267
        anchors.top: parent.top
        anchors.topMargin: 20
        anchors.horizontalCenter: parent.horizontalCenter
        Image { source: "assets/logo-bg.png"; anchors.fill: parent; opacity: 1.0 }
    }

    Item {
        id: logotext
        width: 240
        height: 71
        anchors.horizontalCenter: parent.horizontalCenter
        anchors.top: parent.top
        anchors.topMargin: 117
        Image { source: "assets/logo-text.png"; anchors.fill: parent; opacity: 1.0 }
    }

    // Sidebar to the right.
    Item {
        id: sidebar
        width: 350
        height: parent.height
        anchors.right: parent.right

        Separator{
            width: 1
            color: "#4a402c"
            anchors.right: undefined
            anchors.top: parent.top
        }

        NewsTable{}
    }
        
    // Bottom bar.
    BottomBar{
        id: bottombar
        width: (parent.width-sidebar.width); height: 80
        anchors.bottom: parent.bottom;
    }

}
