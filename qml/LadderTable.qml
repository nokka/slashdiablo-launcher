import QtQuick 2.10				//Item
import QtQuick.Controls 1.4		//TableView
import QtQuick.Controls 2.3		//Button
import QtQuick.Layouts 1.3		//ColumnLayout

Item {
    id: ladderTableBox
	width: 300
    height: 400

    anchors.top: parent.top
    anchors.right: parent.right

	ColumnLayout {
		anchors.fill: parent

		ListView {
			id: ladderList

			Layout.fillWidth: true
			Layout.fillHeight: true

			model: QmlBridge.ladderCharacters
			delegate: LadderTableDelegate{}
		}
	}
}