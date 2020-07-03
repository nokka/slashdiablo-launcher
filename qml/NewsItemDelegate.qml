import QtQuick 2.12

Item {
    id: newsItem
    width: 440
    height: (description.contentHeight + 70) // Content + Title height.

    Row {
        width: parent.width
        height: parent.height
        spacing: 10
        anchors.topMargin: 20

        Item {
            height: parent.height
            width: newsItem.width * 0.70;
            anchors.verticalCenter: parent.verticalCenter

            Title {
                id: title
                textFormat: Text.RichText
                text: (model.link != "" ? "<a style='color:#c7cbd1; text-decoration:none;' href='"+model.link+"'>"+model.title+"</a>" : model.title)
                font.pixelSize: 16
                rightPadding: 20

                onLinkActivated: Qt.openUrlExternally(link)

                MouseArea {
                    anchors.fill: parent
                    acceptedButtons: Qt.NoButton // we don't want to eat clicks on the Text
                    cursorShape: parent.hoveredLink ? Qt.PointingHandCursor : Qt.ArrowCursor
                }

                Image {
                    visible: (model.link != "" ? true : false)
                    id: linkoutIcon
                    fillMode: Image.Pad
                    anchors.top: parent.top
                    anchors.right: parent.right
                    anchors.topMargin: 1
                    width: 16
                    height: 16
                    source: "assets/icons/out.png"
                    opacity: 0.3
                }
            }

            // Timestamp.
            Title {
                id: timestamp
                color: "#a19a97"
                text: model.date + " " + model.year
                anchors.top: title.bottom
            }

            SText {
                id: description
                text: model.text
                color: "#736c6a"
                width: parent.width * 0.95
                wrapMode: Text.WordWrap
                anchors.top: timestamp.bottom
                anchors.topMargin: 10
                elide: Text.ElideRight
            }
        }
    }
}
