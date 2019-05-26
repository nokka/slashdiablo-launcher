import QtQuick 2.4
import QtQuick.Controls 2.5

ApplicationWindow {
    id: mainWindow
    objectName: "mainWindow"

    visible: true
    flags: (Qt.WindowMinimizeButtonHint | Qt.FramelessWindowHint | Qt.Window)

    width: 1024; height: 600
    color: "#1a1324"

     // Top bar for the entire app.
    TopBar {
        id: topbar
        anchors.top: mainWindow.top;
        width: parent.width
        height: 80
        color: "#100b17"
    }

    // Content area.
    Item {
        id: contentArea
        anchors.top: topbar.bottom
        width: mainWindow.width
        height: (mainWindow.height-topbar.height)

        // Background image.
        Item {
            id: background
            anchors.fill: parent;
            Image { source: "assets/bg.jpg"; fillMode: Image.Tile; anchors.fill: parent;  opacity: 1.0 }
        }

        // Bottom bar.
        BottomBar{}
    }

    // Game path dialog, used when the Diablo game path hasn't been set.
    Item {
        SettingsDialog {
            id: settingsDialog
            x: 0; y: 0
            width: mainWindow.width
            height: mainWindow.height
        }
    }

    Component.onCompleted: {
        if(settings.D2Location.length == 0) {
            settingsDialog.open()
        }
    }
}