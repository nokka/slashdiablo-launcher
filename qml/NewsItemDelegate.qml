import QtQuick 2.12

Item {
    id: newsItem
    width: 405
    height: 120

    Row {
        width: parent.width
        height: 120
        spacing: 10
        anchors.topMargin: 20

        // Timestamp.
        /*Item {
            width: 80
            height: 80
            anchors.verticalCenter: parent.verticalCenter

            Column {
                width: 80
                height: 80
                Title {
                    color: "#ab4432"
                    text: model.date
                    anchors.horizontalCenter: parent.horizontalCenter
                }

                Title {
                    color: "#453a2c"
                    text: model.year
                    anchors.horizontalCenter: parent.horizontalCenter
                }
            }
        }*/

        Item {
            width: newsItem.width * 0.70; height: 90
            anchors.verticalCenter: parent.verticalCenter

            Title {
                id: title
                text: model.title
                font.pixelSize: 16
            }

            SText {
                id: description
                text: model.text
                color: "#a6987c"
                width: parent.width * 0.95
                wrapMode: Text.WordWrap
                anchors.top: title.bottom
                anchors.topMargin: 10
                elide: Text.ElideRight
            }
        }

        Title {
            text: "READ MORE"
            font.pixelSize: 11
            font.underline: true
            anchors.rightMargin: 20
            anchors.verticalCenter: parent.verticalCenter
        }
    }

    // Border bottom.
    Image {
        width: newsItem.width; height: 9
        anchors.left: parent.left
        anchors.bottom: parent.bottom
        fillMode: Image.PreserveAspectFit
        source: "assets/item_bg.png"
    }
}
