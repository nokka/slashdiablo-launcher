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

    Rectangle {
        id: prereqsBox
        anchors.fill: parent
        color: "black"

        Item {
            width: 400
            height: 400
            anchors.centerIn: parent

            Column {
                width: parent.width

                Item {
                    width: parent.width
                    height: 240
                    
                    Item {
                        width: 210.6
                        height: 240.3
                        anchors.top: parent.top
                        anchors.topMargin: 20
                        anchors.horizontalCenter: parent.horizontalCenter
                        Image { source: "assets/logo-bg.png"; anchors.fill: parent; fillMode: Image.PreserveAspectFit; opacity: 1.0 }
                    }

                    Item {
                        width: 216
                        height: 63.9
                        anchors.horizontalCenter: parent.horizontalCenter
                        anchors.top: parent.top
                        anchors.topMargin: 109
                        Image { source: "assets/logo-text.png"; anchors.fill: parent; fillMode: Image.PreserveAspectFit; opacity: 1.0 }
                    }
                }

                Item { 
                    width: parent.width
                    height: 60

                    Title {
                        anchors.horizontalCenter: parent.horizontalCenter
                        anchors.bottom: parent.bottom
                        text: "PREPARING LAUNCHER..."
                    }
                 }

                Item {
                    width: parent.width
                    height: 60

                    CircularProgress {
                        anchors.horizontalCenter: parent.horizontalCenter
                        anchors.verticalCenter: parent.verticalCenter
                        size: 30
                        visible: true
                    }
                }
            }
        }

        OpacityAnimator {
            target: prereqsBox;
            from: 1;
            to: 0;
            duration: 500
            running: settings.prerequisitesLoaded
        }
    }
    
    // Settings popup.
    SettingsPopup{
        id: settingsPopup
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
                settingsLoader.item.open()
            }
            // TODO: Think about if this should be sync,
            // to return an error and display that we couldn't talk to slash API.
            /*var success = settings.getAvailableMods()

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
            }*/
        }
    }
}
