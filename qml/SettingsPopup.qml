import QtQuick 2.4
import QtQuick.Controls 2.1
import QtQuick.Controls.Material 2.1
import QtQuick.Layouts 1.3

Popup {
    id: settingsPopup

    modal: true
    focus: true
    width: 1024
    height: 600
    dim: true
    anchors.centerIn: root
    closePolicy: Popup.CloseOnEscape | Popup.CloseOnPressOutsideParent

    background: Rectangle {
        color: "#000000"
        opacity: 0.5
        anchors.fill: parent
    }
}