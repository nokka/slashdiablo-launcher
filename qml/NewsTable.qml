import QtQuick 2.12				//Item
import QtQuick.Layouts 1.3		//ColumnLayout

Item {
    id: newsTableBox
	width: parent.width
    height: parent.height

    anchors.top: parent.top
    anchors.right: parent.right

	ColumnLayout {
		anchors.fill: parent
		anchors.leftMargin: 15
		anchors.rightMargin: 10
		anchors.topMargin: 15
		
		ListView {
			id: newsList
			spacing: 0
			visible: (!news.loading && !news.error)
			height: 320

			// Disable scroll.
			interactive: false

			Layout.fillWidth: true
			Layout.fillHeight: true

			model: news.items
			delegate: NewsItemDelegate{}
		}

		// Show if we're loading on if there's been an error.
		Item {
			Layout.fillWidth: true
			Layout.fillHeight: true
			visible: (news.loading || news.error)

			// Loading circle.			
			CircularProgress {
				size: 25
				anchors.centerIn: parent
				visible: news.loading
			}

			// Error item.
			Item {
				anchors.centerIn: parent
				visible: news.error
				height: 100

				Title {
					color: "#ffffff"
					topPadding: 30
					text: "Unable to get news"
					font.pixelSize: 13
					anchors.horizontalCenter: parent.horizontalCenter
				}
			}
		}
	}

	Component.onCompleted: {
		news.items.clear()
		news.getNews()
	}
}
