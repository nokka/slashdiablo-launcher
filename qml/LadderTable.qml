import QtQuick 2.10				//Item
import QtQuick.Controls 1.4		//TableView
import QtQuick.Controls 2.3		//Button
import QtQuick.Layouts 1.3		//ColumnLayout

Rectangle {
    id: ladderTableBox
	width: mainWindow.width * 0.30
    height: (parent.height - 100)
	color: "#00000000"

    anchors.top: parent.top
    anchors.right: parent.right

	ColumnLayout {
		anchors.fill: parent

		Header {
			text: "LADDER TOP 10"
			font.pointSize: 16
			topPadding: 15
			bottomPadding: 15
		}

		ListView {
			id: ladderList
			spacing: 3

			Layout.fillWidth: true
			Layout.fillHeight: true

			model: QmlBridge.ladderCharacters
			delegate: LadderTableDelegate{}
		}
	}
}