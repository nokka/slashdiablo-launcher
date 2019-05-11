import QtQuick 2.0
import QtQuick.Controls 1.4
import QtQuick.Controls.Styles 1.4

Rectangle {
    // Size
    width: parent.width * 0.70
    height: 20
    x: 20 

    // Anchors
    anchors.verticalCenter: parent.verticalCenter

    // Transparent background
    color: "#00000000" 

    ProgressBar {
        value: QmlBridge.patchProgress
        
        width: parent.width
        height: 20
        
        style: ProgressBarStyle {
            background: Rectangle {
                radius: 2
                color: "#00578a"
                border.color: "#002b5c"
                border.width: 1
            }
            
            progress: Rectangle {
                color: "#0983b8"
            }
        }
    }
}