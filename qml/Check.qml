
import QtQuick 2.12
import QtQuick.Controls 1.4
import QtQuick.Controls.Styles 1.4

CheckBox {
    text: "Include maphack"
    
    style: CheckBoxStyle {
        indicator: Rectangle {
                color: "#000000"
                implicitWidth: 16
                implicitHeight: 16
                radius: 3
                border.color: control.activeFocus ? "#ffffff" : "#1c1c1c"
                border.width: 1
                Rectangle {
                    visible: control.checked
                    color: "#6E3E87"
                    border.color: "#6E3E87"
                    radius: 1
                    anchors.margins: 4
                    anchors.fill: parent
                }
        }
    }
}