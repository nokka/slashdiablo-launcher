import QtQuick 2.12
import QtQuick.Controls 2.5

Rectangle {
    id: mainWindow
    color: "#080806"
    width: 1024; height: 600

    // Load fonts.
    FontLoader { id: d2Font; source: "assets/fonts/EXL.ttf" }
    FontLoader { id: montserrat; source: "assets/fonts/Montserrat-Light.ttf" }

    // Background image.
    Item {
        id: background
        anchors.fill: parent;
        Image { source: "assets/tmp/bg.png"; fillMode: Image.Tile; anchors.fill: parent;  opacity: 1.0 }
    }

    // Top bar for the entire app.
    TopBar {
        id: topbar
        anchors.top: mainWindow.top;
        width: parent.width
        height: 50
    }

    // Content area.
    Rectangle {
        id: contentArea
        width: mainWindow.width
        height: (mainWindow.height-topbar.height)
        anchors.top: topbar.bottom
        //color: "#080806"
        color: "#00000000"

        // Main content area.
        ContentArea{}

        // Top ladder table.
        LadderTable{}

        // Bottom bar.
        BottomBar{}
    }

    // Game path dialog, used when the Diablo game path hasn't been set.
    SettingsDialog {
        id: settingsDialog
        visible: false
        x: 0; y: 0
        width: mainWindow.width
        height: mainWindow.height
    }

    Component.onCompleted: {
        if(settings.NrOfGames == 0) {
            settingsDialog.visible = true
        }
    }
}