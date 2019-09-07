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
                text: model.title
                font.pixelSize: 16
            }

            // Timestamp.
            Title {
                id: timestamp
                color: "#915a23"
                text: model.date + " " + model.year
                anchors.top: title.bottom
            }

            SText {
                id: description
                text: model.text
                color: "#a6987c"
                width: parent.width * 0.95
                wrapMode: Text.WordWrap
                anchors.top: timestamp.bottom
                anchors.topMargin: 10
                elide: Text.ElideRight
            }
        }
    }

    // Border bottom.
    /*Image {
        width: newsItem.width; height: 9
        anchors.left: parent.left
        anchors.bottom: parent.bottom
        fillMode: Image.PreserveAspectFit
        source: "assets/item_bg.png"
    }*/
}
