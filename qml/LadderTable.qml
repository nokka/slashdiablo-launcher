import QtQuick 2.12				//Item
import QtQuick.Controls 2.3		//TableView, Button
import QtQuick.Layouts 1.3		//ColumnLayout

Item {
    id: ladderTableBox
	width: parent.width * 0.30
    height: parent.height - 80

	Rectangle {
		anchors.fill: parent
		color: "#000000"
		opacity: 0.5
		radius: 5
	}

    anchors.top: parent.top
    anchors.right: parent.right

	ColumnLayout {
		anchors.fill: parent
		anchors.leftMargin: 10
		anchors.rightMargin: 10

		// Shown when there's characters to show.
		Item {
			visible: (!ladder.loading && !ladder.error)
			Layout.alignment: Qt.AlignHCenter
			height: 40
			width: 285
			
			Title {
				text: "LADDER TOP 10"
				font.pixelSize: 13
				font.bold: true
				topPadding: 20
			}
		}
		
		ListView {
			id: ladderList
			spacing: 2
			visible: (!ladder.loading && !ladder.error)
			height: 320

			Layout.fillWidth: true
			Layout.fillHeight: true

			model: ladder.characters
			delegate: LadderTableDelegate{}
		}

		// Show if we're loading on if there's been an error.
		Item {
			Layout.fillWidth: true
			Layout.fillHeight: true
			visible: (ladder.loading || ladder.error)

			// Loading circle.			
			CircularProgress {
				anchors.centerIn: parent
				visible: ladder.loading
    		}

			// Error item.
			Item {
				anchors.centerIn: parent
				visible: ladder.error
				height: 100

				Image {
					id: ladderError
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
					text: "Couldn't get ladder characters"
					font.family: roboto.name
					font.pixelSize: 11
					anchors.horizontalCenter: parent.horizontalCenter
				}
			}
		}
	}

	Component.onCompleted: {
		ladder.characters.clear()
		ladder.getLadder("exp")
	}
}