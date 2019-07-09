import QtQuick 2.12
import QtQuick.Controls 2.5

Rectangle {
    id: settingsView
    color: "#080806"

    Button {
        text: "POP"
        onClicked: stack.pop()
    }

    Component.onCompleted: {
        console.log("Settings loaded")
    }

    Component.onDestruction: {
        console.log("Settings destroyed")
    }
}