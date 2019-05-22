import QtQuick 2.12
import QtQuick.Controls 2.1
import QtQuick.Dialogs 1.3

Rectangle {
    id: gamepath_view
    width: parent.width
    height: parent.height
    anchors.fill: parent;
    color: "#100b17"

    Column {
        anchors.centerIn: parent

        TextField {
            id: input
    
            anchors.horizontalCenter: parent.horizontalCenter
            placeholderText: "Write something ..."
        }

        Button {
            anchors.horizontalCenter: parent.horizontalCenter
            text: "and click me!"
            onClicked: fileDialog.open()
        }
    }

    FileDialog {
        id: fileDialog
        selectFolder: true
        title: "Please choose a file"
        folder: shortcuts.home
        onAccepted: {
            console.log("You chose: " + fileDialog.fileUrls)
        }
        onRejected: {
            console.log("Canceled")
        }
    }
}