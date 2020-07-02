import QtQuick 2.12
import QtQuick.Controls 2.2

ComboBox {
    id: dropdown

    contentItem: Title {
        text: dropdown.displayText
        color: dropdown.pressed ? "#57555e" : "#969696" 
        topPadding: 7
        leftPadding: 12
    }

    background: Rectangle {
        color: "#0C0C0C"
        border.color: "#000000"
    }

    popup: Popup {
        y: (dropdown.height + 5)
        width: dropdown.width
        implicitHeight: contentItem.implicitHeight
        padding: 1

        contentItem: ListView {
            clip: true
            implicitHeight: contentHeight
            model: dropdown.popup.visible ? dropdown.delegateModel : null
            currentIndex: dropdown.highlightedIndex

            ScrollIndicator.vertical: ScrollIndicator { }
        }

        background: Rectangle {
            border.width: 0
            color: "#050505"
        }
    }

    indicator: Canvas {
        id: canvas
        x: dropdown.width - width - dropdown.rightPadding
        y: dropdown.topPadding + (dropdown.availableHeight - height) / 2
        width: 7
        height: 4
        contextType: "2d"

        Connections {
            target: dropdown
            onPressedChanged: canvas.requestPaint()
        }

        onPaint: {
            var ctx = canvas.getContext('2d');

            ctx.reset();
            ctx.moveTo(0, 0);
            ctx.lineTo(width, 0);
            ctx.lineTo(width / 2, height);
            ctx.closePath();
            ctx.fillStyle = dropdown.pressed ? "#57555e" : "#969696";
            ctx.fill();
        }
    }


    delegate: DropdownDelegate{}
}
