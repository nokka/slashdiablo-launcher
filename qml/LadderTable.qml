import QtQuick 2.12				//Item
import QtQuick.Controls 2.3		//TableView, Button
import QtQuick.Layouts 1.3		//ColumnLayout

Item {
    id: ladderTableBox
	width: parent.width
    height: parent.height
    anchors.top: parent.top
    anchors.right: parent.right

	property string activeMode: "exp"

	ColumnLayout {
		anchors.fill: parent
		anchors.leftMargin: 10
		anchors.rightMargin: 10

		// Shown when there's characters to show.
		Item {
			visible: (!ladder.loading && !ladder.error)
			Layout.alignment: Qt.AlignLeft
			height: 100
			width: 285
			Column {
				Row {
					height: 50
					spacing: 10

					Layout.alignment: Qt.AlignCenter

					ModeItem {
						width: 90
						height: 40
						label: "STANDARD"

						onClicked: {
							active = true
						}
					}

					ModeItem {
						width: 90
						height: 40
						label: "HARDCORE"
						
						onClicked: {}
					}
				}

				// Header to the list.
				Row {
					id: header
					height: 50
					Layout.alignment: Qt.AlignBottom

					LadderCell {
						width: ladderList.width * 0.10
						height: 50
						content: "Rank"
					}

					LadderCell {
						width: ladderList.width * 0.10
						height: 50
						content: "Level"
					}

					LadderCell {
						width: ladderList.width * 0.10
						height: 50
						content: "Class"
					}

					LadderCell {
						width: ladderList.width * 0.40
						height: 50
						content: "Level"
					}

					LadderCell {
						width: ladderList.width * 0.10
						height: 50
						content: "Title"
					}

					Item {
						width: ladderList.width * 0.20
						height: 50

						Text {
							color: "#b5b5b5"
							font.pixelSize: 12
							font.bold: true
							font.family: beaufortbold.name
							text: "Status"
							anchors.verticalCenter: parent.verticalCenter
							anchors.right: parent.right
						}

						Separator{}
					}
				}
			}
		}
		
		ListView {
			id: ladderList
			spacing: 0
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
