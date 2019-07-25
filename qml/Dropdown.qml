import QtQuick 2.12
import QtQuick.Controls 2.1

ComboBox {
    id: dropdown

    contentItem: Text {
        text: dropdown.displayText
        font: dropdown.font
        color: dropdown.pressed ? "#969696" : "#ffffff"
        topPadding: 7
        leftPadding: 12
    }

    background: Rectangle {
        color: "#1d1924"
        border.color: "#800507"
        radius: height/2
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

    delegate: DropdownDelegate{}
}