import QtQuick 2.12
import QtQuick.Controls 2.5

Item {
    id: launchView
    width: parent.width; height: parent.height
    anchors.leftMargin: 20

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

        // News list.
        ListView {
			id: newsList
			spacing: 4
			height: 300

			anchors.top: parent.top
            anchors.left: parent.left
            anchors.topMargin: 10
            anchors.leftMargin: 20

			model: NewsModel{}
			delegate: NewsItemDelegate{}
		}
    }
        
    // Bottom bar.
    BottomBar{
        id: bottombar
        width: (parent.width-sidebar.width); height: 60
        anchors.bottom: parent.bottom;
    }

    // Top ladder table.
    //LadderTable{}
    
}
