import QtQuick 2.0

Rectangle {
    id: root
    width: 1024; height: 600
    color: "#1a1324"

    // Background image.
    Rectangle {
        id: background
        anchors.fill: parent;
        Image { source: "../assets/bg.jpg"; fillMode: Image.Tile; anchors.fill: parent;  opacity: 1.0 }
    }

    // Top bar for the entire app.
    TopBar {}

    // Progress bar.
    Progress{}

    // Launch button.
    Button{}
}