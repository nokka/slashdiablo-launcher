import QtQuick 2.12
import QtQuick.Controls 1.4
import QtQuick.Controls.Styles 1.4

Rectangle {
    id: patcher
    height: 20
    x: 20 

    // Anchors
    anchors.verticalCenter: parent.verticalCenter

    // Transparent background
    color: "#00000000" 

    ProgressBar {
        value: QmlBridge.patchProgress
        
        width: parent.width
        height: 10
        
        style: ProgressBarStyle {
            background: Rectangle {
                radius: 2
                color: "#381612"
                border.color: "#141009"
                border.width: 1
            }
            
            progress: Rectangle {
                color: "#873d29"
            }
        }
    }

    Component.onCompleted: {
        if(settings.D2Location.length > 0) {
            //QmlBridge.patchGame()
        }
    }
}