import QtQuick 2.12
import QtQuick.Controls 2.5

import "componentCreator.js" as ComponentCreator

Item {
    id: root
    width: 1024; height: 600

    // Load fonts.
    FontLoader { id: roboto; source: "assets/fonts/Roboto-Regular.ttf" }
    FontLoader { id: robotobold; source: "assets/fonts/Roboto-Bold.ttf" }

    StackView {
        id: stack
        //initialItem: LauncherView{}
        initialItem: SettingsView{}
        anchors.fill: parent

        pushEnter: Transition {
            PropertyAnimation {
                duration: 0
            }
        }
        
        popEnter: Transition {
            PropertyAnimation {
                duration: 0
            }
        }

        popExit: Transition {
            PropertyAnimation {
                duration: 0
            }
        }

        pushExit: Transition {
            PropertyAnimation {
                duration: 0
            }
        }
    }

    Component.onCompleted: {
        //stack.push(ComponentCreator.createSettingsView(stack))
    }
}