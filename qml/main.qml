import QtQuick 2.12
import QtQuick.Controls 2.5

import "componentCreator.js" as ComponentCreator

Item {
    id: root
    width: 1024; height: 600

    // Load fonts.
    FontLoader { id: roboto; source: "assets/fonts/Roboto-Regular.ttf" }
    FontLoader { id: robotobold; source: "assets/fonts/Roboto-Bold.ttf" }
    
    // Top bar for the entire app.
    TopBar {
        id: topbar
        anchors.top: parent.top;
        width: parent.width
        height: 80
    }

    // Content area.
    Item {
        width: parent.width
        height: (parent.height - topbar.height)
        anchors.top: topbar.bottom
       
        // Loads pages dynamically.
        Loader {
            id: contentLoader
            anchors.fill: parent
            source: "LauncherView.qml"
        }
    }

    // This is a bit of a hack to get a popup to display right after
    // the parent loads, if we remove the timer we get an error saying
    // there's no parent to create the popup from.
    Timer {
        interval: 0; running: true; repeat: false
        onTriggered: {
            console.log("on triggred running")
            ComponentCreator.createSettingsPopup(root).open();
        }
    }
}