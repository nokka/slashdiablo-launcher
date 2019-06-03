import QtQuick 2.10				//Item
import QtQuick.Controls 1.4		//TableView
import QtQuick.Controls 2.3		//Button
import QtQuick.Layouts 1.3		//ColumnLayout
import CustomQmlTypes 1.0		//CustomTableModel

Item {
    id: ladderTableBox
	width: 300
    height: 400

    anchors.top: parent.top
    anchors.right: parent.right

	ColumnLayout {
		anchors.fill: parent

		TableView {
			id: ladderTable

			Layout.fillWidth: true
			Layout.fillHeight: true

			model: CustomTableModel{}

			TableViewColumn {
                width: (ladderTableBox.width * 0.33)
				role: "Character"
				title: role
			}

			TableViewColumn {
                width: (ladderTableBox.width * 0.33)
				role: "Class"
				title: role
			}

            TableViewColumn {
                width: (ladderTableBox.width * 0.33)
				role: "Level"
				title: role
			}
		}
	}
}