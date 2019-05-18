import QtQuick 2.2
import QtQuick.Dialogs 1.2

Rectangle {
    id: root
    width: 1024; height: 600
    color: "#1a1324"
    

    // View that will ask to set gamepath.
    //GamePath {}

    // Background image.
    Item {
        id: background
        anchors.fill: parent;
        Image { source: "assets/bg.jpg"; fillMode: Image.Tile; anchors.fill: parent;  opacity: 1.0 }
    }

    // Top bar for the entire app.
    TopBar {}

    // Bottom bar.
    BottomBar{}
}