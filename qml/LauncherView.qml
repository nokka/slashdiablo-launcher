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
            color: "#3F2A2A"
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

        Item {
            anchors.bottom: parent.bottom
            width: parent.width; height: 70
            anchors.left: parent.left
            anchors.bottomMargin: 15

            Separator{
                color: "#3F2A2A"
                anchors.top: parent.top
                anchors.bottom: undefined
            }
    
            Title {
                text: "DIABLO GATEWAY"
                anchors.bottom: gameInstances.top
                anchors.left: parent.left
                anchors.leftMargin: 20
                anchors.bottomMargin: 5
            }

            Dropdown{
                id: gameInstances
                anchors.left: parent.left
                anchors.bottom: parent.bottom
                anchors.leftMargin: 20
                currentIndex: 0
                model: ["Slashdiablo", "Battle.net"]
                height: 30
                width: 300

                onActivated: {
                    diablo.setGateway(this.currentText)
                }
            }
        }        
    }

        
    // Bottom bar.
    BottomBar{
        id: bottombar
        width: (parent.width-sidebar.width); height: 80
        anchors.bottom: parent.bottom;
    }

    // Top ladder table.
    //LadderTable{}
    
}
