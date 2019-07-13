import QtQuick 2.12
import QtQuick.Controls 2.5

import "componentCreator.js" as ComponentCreator

Rectangle {
    id: launchView
    color: "#080806"
    width: parent.width; height: parent.height

     // Background image.
    Item {
        id: background
        anchors.fill: parent;
        Image { source: "assets/bg.png"; fillMode: Image.Tile; anchors.fill: parent;  opacity: 1.0 }
    }
    
    // Top bar for the entire app.
    TopBar {
        id: topbar
        anchors.top: launchView.top;
        width: parent.width
        height: 80
    }

    // Content area.
    Item {
        id: launchContent
        width: launchView.width
        height: (launchView.height-topbar.height)
        anchors.top: topbar.bottom

        // Main content area.
        Item {
            width: parent.width * 0.70
            height: parent.height - 80
        }

        // Top ladder table.
        LadderTable{}

        // Bottom bar.
        BottomBar{}
    }
}