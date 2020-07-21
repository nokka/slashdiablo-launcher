import QtQuick 2.12

Item {
    id: root
    width: 1024; height: 600

    // Load fonts.
    FontLoader { id: roboto; source: "assets/fonts/Roboto-Regular.ttf" }
    FontLoader { id: robotobold; source: "assets/fonts/Roboto-Bold.ttf" }
    FontLoader { id: beaufort; source: "assets/fonts/Beaufort-Regular.ttf" }
    FontLoader { id: beaufortbold; source: "assets/fonts/Beaufort-Bold.ttf" }

    // Background image.
    Rectangle {
        id: background
        anchors.fill: parent;
        color: "#0a0a0d"
        Image { source: "assets/bg.png"; fillMode: Image.PreserveAspectCrop; anchors.fill: parent;}
    }

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
       
        // Loads pages dynamically, launcher view is default.
        Loader {
            id: contentLoader
            anchors.fill: parent
            source: "LauncherView.qml"
        }
    }

    Loader { 
        id: settingsLoader
        anchors.centerIn: root
        width: 850
        height: 500
    }

    // Loading screen that shows up to load prerequisites.
    LoadingScreen{}
    
    // Settings popup.
    SettingsPopup{
        id: settingsPopup
    }

    // Patch popup.
    PatchPopup{
        id: patchPopup
    }

    // This is a bit of a hack to get a popup to display right after
    // the parent loads, if we remove the timer we get an error saying
    // there's no parent to create the popup from.
    Timer {
        interval: 0; running: true; repeat: false
        onTriggered: {
            // Load all prerequisite data async.
            settings.getPrerequisites()

            if(settings.games.rowCount() == 0) {
                settingsPopup.open()
            }
        }
    }
}
