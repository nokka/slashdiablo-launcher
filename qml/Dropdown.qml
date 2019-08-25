import QtQuick 2.12
import QtQuick.Controls 2.2

ComboBox {
    id: dropdown

    contentItem: Title {
        text: dropdown.displayText
        color: dropdown.pressed ? "#969696" : "#f3e6d0"
        topPadding: 7
        leftPadding: 12
    }

    background: Rectangle {
        color: "#080808"
        border.color: "#171717"
        radius: 5
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