import QtQuick 2.12				//Item
import QtQuick.Controls 1.4		//TableView
import QtQuick.Controls 2.3		//Button
import QtQuick.Layouts 1.3		//ColumnLayout

Rectangle {
    id: ladderTableBox
	width: mainWindow.width * 0.30
    height: parent.height - 100
	color: "#00000000"

    anchors.top: parent.top
    anchors.right: parent.right

	ColumnLayout {
		anchors.fill: parent

		// Shown when there's characters to show.
		Header {
			Layout.alignment: Qt.AlignTop
			text: "LADDER TOP 10"
			font.pointSize: 16
			topPadding: 20
			bottomPadding: 15
			visible: (!ladder.loading && !ladder.error)
		}

		ListView {
			id: ladderList
			spacing: 3
			visible: (!ladder.loading && !ladder.error)

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
					source: "assets/error.svg"
				}

				Text {
					color: "#ffffff"
					topPadding: 30
					text: "Couldn't get ladder characters"
					font.family: montserrat.name
					font.pixelSize: 11
					anchors.horizontalCenter: parent.horizontalCenter
				}
			}
		}
	}

	Component.onCompleted: {
		//ladder.getLadder("exp")
	}
}