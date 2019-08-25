import QtQuick 2.12
import QtQuick.Controls 2.5

Item {
    id: launchView
    width: parent.width; height: parent.height
    anchors.leftMargin: 20

    Item {
        // Background image.
        Item {
            id: background
            anchors.fill: parent;
            Image { source: "assets/diablo.png"; fillMode: Image.Tile; anchors.fill: parent;  opacity: 0.2 }
        }
        
        width: parent.width * 0.68
        height: parent.height
        anchors.left: parent.left

        // News list.
        ListView {
			id: newsList
			spacing: 4
			height: 300

			anchors.top: parent.top
            anchors.left: parent.left
            anchors.topMargin: 20

			model: NewsModel{}
			delegate: NewsItemDelegate{}
		}

        /*Item {
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
        }*/
    }

    // Top ladder table.
    LadderTable{}
    
}
