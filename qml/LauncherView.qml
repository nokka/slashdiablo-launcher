import QtQuick 2.12
import QtQuick.Controls 2.5

Item {
    id: launchView
    width: parent.width; height: parent.height
    anchors.leftMargin: 20

    Item {
        width: 405
        height: parent.height
        anchors.right: parent.right

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
    //LadderTable{}
    
}
