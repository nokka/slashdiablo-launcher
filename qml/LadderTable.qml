import QtQuick 2.12
import QtQuick.Controls 2.3
import QtQuick.Layouts 1.3

Item {
    id: ladderTableBox
	width: parent.width
    height: parent.height
    anchors.top: parent.top
    anchors.right: parent.right

	property string activeMode: "exp"

	function getLadderByMode(mode) {
        ladder.characters.clear()
		ladder.getLadder(mode)
    }

	ColumnLayout {
		anchors.fill: parent
		anchors.leftMargin: 50
		anchors.rightMargin: 50

		Item {
			Layout.alignment: Qt.AlignCenter
			height: 55
			width: 306

			Row {
				height: 55
				spacing: 10

				Layout.alignment: Qt.AlignCenter

				PlainButton {
					width: 90
					height: 40
					label: "STANDARD"
					activatable: true
					fontSize: 12
					active: (activeMode == "standard")
					backgroundColor: "#000000"
					colorHovered: "#000000"
					anchors.verticalCenter: parent.verticalCenter

					onClicked: {
						activeMode = "standard"
						getLadderByMode("standard")
					}
				}

				PlainButton {
					width: 90
					height: 40
					label: "HARDCORE"
					fontSize: 12
					activatable: true
					active: (activeMode == "hardcore")
					backgroundColor: "#000000"
					colorHovered: "#000000"
					anchors.verticalCenter: parent.verticalCenter
					
					onClicked: {
						activeMode = "hardcore"
						getLadderByMode("hardcore")
					}
				}

				PlainButton {
					width: 36
					height: 40
					label: "EXP"
					fontSize: 12
					activatable: true
					active: (activeMode == "exp")
					backgroundColor: "#000000"
					colorHovered: "#000000"
					anchors.verticalCenter: parent.verticalCenter
					
					onClicked: {
						activeMode = "exp"
						getLadderByMode("exp")
					}
				}

				PlainButton {
					width: 60
					height: 40
					label: "EXP HC"
					fontSize: 12
					activatable: true
					active: (activeMode == "exphardcore")
					backgroundColor: "#000000"
					colorHovered: "#000000"
					anchors.verticalCenter: parent.verticalCenter
					
					onClicked: {
						activeMode = "exphardcore"
						getLadderByMode("exphardcore")
					}
				}
			}

			Separator{}
		}

		// Shown when there's characters to show.
		Item {
			visible: (!ladder.loading && !ladder.error)
			Layout.alignment: Qt.AlignLeft
			height: 40

			// Header to the list.
			Row {
				id: header
				height: 40
				Layout.alignment: Qt.AlignBottom

				TableCell {
					width: ladderList.width * 0.10
					height: parent.height
					content: "Rank"
				}

				TableCell {
					width: ladderList.width * 0.10
					height: parent.height
					content: "Level"
				}

				TableCell {
					width: ladderList.width * 0.10
					height: parent.height
					content: "Class"
				}

				TableCell {
					width: ladderList.width * 0.40
					height: parent.height
					content: "Level"
				}

				TableCell {
					width: ladderList.width * 0.10
					height: parent.height
					content: "Title"
				}

				Item {
					width: ladderList.width * 0.20
					height: parent.height

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
		
		ListView {
			id: ladderList
			spacing: 0
			visible: (!ladder.loading && !ladder.error)
			height: 320

			Layout.fillWidth: true
			Layout.fillHeight: true

			// Disable scroll.
			interactive: false

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
				showBackground: false
    		}

			// Error item.
			Item {
				anchors.centerIn: parent
				visible: ladder.error
				height: 100

				Title {
					color: "#ffffff"
					topPadding: 30
					text: "Unable to get ladder characters"
					font.pixelSize: 13
					anchors.horizontalCenter: parent.horizontalCenter
				}
			}
		}
	}

	Component.onCompleted: {
		getLadderByMode("exp")
	}
}
