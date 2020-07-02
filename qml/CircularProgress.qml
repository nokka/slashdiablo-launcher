import QtQuick 2.12
import QtQml 2.2

Item {
    id: root

    width: size
    height: size

    property int size: 50               // The size of the circle in pixel
    property real arcBegin: 0           // start arc angle in degree
    property real arcEnd: 120           // end arc angle in degree
    property real arcOffset: 0          // rotation
    property bool isPie: false          // paint a pie instead of an arc
    property bool showBackground: true  // a full circle as a background of the arc
    property real lineWidth: 5          // width of the line
    property string colorCircle: "#ffffff"
    property string colorBackground: "#736c6a"
    property int animationDuration: 800

     NumberAnimation on rotation {
        from: 0; to: 360;
        running: visible
        loops: Animation.Infinite;
        duration: animationDuration;
        easing.type: Easing.InOutCubic
    }

    Canvas {
        id: canvas
        anchors.fill: parent
        rotation: -90 + parent.arcOffset

        onPaint: {
            var ctx = getContext("2d")
            var x = width / 2
            var y = height / 2
            var start = Math.PI * (parent.arcBegin / 180)
            var end = Math.PI * (parent.arcEnd / 180)
            ctx.reset()

            if (root.showBackground) {
                ctx.beginPath();
                ctx.arc(x, y, (width / 2) - parent.lineWidth / 2, 0, Math.PI * 2, false)
                ctx.lineWidth = root.lineWidth
                ctx.strokeStyle = root.colorBackground
                ctx.stroke()
            }
            
            ctx.beginPath();
            ctx.arc(x, y, (width / 2) - parent.lineWidth / 2, start, end, false)
            ctx.lineWidth = root.lineWidth
            ctx.strokeStyle = root.colorCircle
            ctx.stroke()
        }
    }
}
