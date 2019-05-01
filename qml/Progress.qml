import QtQuick 2.0
import QtQuick.Controls 1.4
import QtQuick.Controls.Styles 1.4

Rectangle {
    // Size
    width: 500
    height: 80
    x: 20

    // Anchors
    anchors.bottom: parent.bottom;

    // Transparent background
    color: "#00000000" 

    ProgressBar {
        value: 0.5
        
        width: parent.width
        height: 50
        
        anchors.horizontalCenter: parent.horizontalCenter
        anchors.verticalCenter: parent.verticalCenter
        
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