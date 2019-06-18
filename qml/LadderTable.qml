import QtQuick 2.10				//Item
import QtQuick.Controls 1.4		//TableView
import QtQuick.Controls 2.3		//Button
import QtQuick.Layouts 1.3		//ColumnLayout

Rectangle {
	property var errored: false

    id: ladderTableBox
	width: mainWindow.width * 0.30
    height: parent.height - 100
	color: "#00000000"

    anchors.top: parent.top
    anchors.right: parent.right

	ColumnLayout {
		anchors.fill: parent
		
		Header {
			Layout.alignment: Qt.AlignTop
			text: "LADDER TOP 10"
			font.pointSize: 16
			topPadding: 5
			bottomPadding: 5
			visible: true//!ladder.loading
		}

		ListView {
			id: ladderList
			spacing: 3
			visible: true//!ladder.loading

			Layout.fillWidth: true
			Layout.fillHeight: true

			model: ladder.characters
			delegate: LadderTableDelegate{}
		}

		Item {
			Layout.fillWidth: true
			Layout.fillHeight: true
			visible: false//(ladder.loading || errored)

			// Loading bar.			
			CircularProgress {
				anchors.centerIn: parent
				visible: false//ladder.loading
    		}

			// Error item.
			Item {
				anchors.centerIn: parent
				visible: errored
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

	Component.onCompleted: ladder.getLadder("exp")
}