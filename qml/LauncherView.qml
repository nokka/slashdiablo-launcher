import QtQuick 2.12
import QtQuick.Controls 2.5

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

        Item {
            height: parent.height
            anchors.top: parent.top
            anchors.left: parent.left
            anchors.topMargin: 15
            anchors.leftMargin: 20

            // News list.
            ListView {
                id: newsList
                spacing: 4
                visible: (!news.loading && !news.error)
                
                model: news.items
                delegate: NewsItemDelegate{}
		    }            
        }

        // News loading circle.			
        CircularProgress {
            size: 25
            anchors.centerIn: parent
            visible: news.loading
        }

        // Error item.
        Item {
            anchors.verticalCenter: parent.verticalCenter
            anchors.horizontalCenter: parent.horizontalCenter
            width: parent.width
            height: 100

            visible: news.error
            
            Image {
                id: newsError
                fillMode: Image.PreserveAspectFit
                anchors.horizontalCenter: parent.horizontalCenter
                anchors.top: parent.top
                width: 20
                height: 20
                source: "assets/svg/error.svg"
            }

            Text {
                color: "#ffffff"
                topPadding: 30
                text: "Couldn't get news items"
                font.family: roboto.name
                font.pixelSize: 11
                anchors.horizontalCenter: parent.horizontalCenter
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

    Component.onCompleted: {
		news.items.clear()
		news.getNews()
	}
    
}
