import QtQuick 2.12

Item {
    id: root
    width: 1024; height: 600

    // Load fonts.
    FontLoader { id: roboto; source: "assets/fonts/Roboto-Regular.ttf" }
    FontLoader { id: robotobold; source: "assets/fonts/Roboto-Bold.ttf" }
    FontLoader { id: beaufort; source: "assets/fonts/Beaufort-Regular.ttf" }
    FontLoader { id: beaufortbold; source: "assets/fonts/Beaufort-Bold.ttf" }

    property bool prerequisitesLoaded: false

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
        }
    }

    Loader { 
        id: settingsLoader
        anchors.centerIn: root
        width: 850
        height: 500
    }

    Rectangle {
        anchors.fill: parent
        color: "black"
        visible: !prerequisitesLoaded

        Item {
            anchors.centerIn: parent

            // Loading circle.			
            CircularProgress {
                anchors.horizontalCenter: parent.horizontalCenter
                size: 30
                visible: true
            }

            Title {
                text: "TALKING TO SLASH API"
            }
        }
        
        
    }
    
    // Settings popup.
    /*SettingsPopup{
        id: settingsPopup
    }*/

    // This is a bit of a hack to get a popup to display right after
    // the parent loads, if we remove the timer we get an error saying
    // there's no parent to create the popup from.
    Timer {
        interval: 2000; running: true; repeat: false
        onTriggered: {
            // TODO: Think about if this should be sync,
            // to return an error and display that we couldn't talk to slash API.
            var success = settings.getAvailableMods()

            if(success) {
                // Allow content to load since all prerequisites are loaded.
                settingsLoader.source = "SettingsPopup.qml"
               
                // Allow content to load since all prerequisites are loaded.
                contentLoader.source = "LauncherView.qml"
                
                if(settings.games.rowCount() == 0) {
                    settingsLoader.item.open()
                }

                // Tell the UI we're done with prerequisites.
                prerequisitesLoaded = true
            } else {
                // TODO: Show error button
            }  
        }
    }
}
