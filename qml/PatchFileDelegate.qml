import QtQuick 2.12

Item {
    id: row
    width: parent.width
    height: 35

    Rectangle {
        anchors.fill: parent
        color: (model.index % 2 == 0 ? "#000000" : "#080505")
        opacity: 0.2
    }

    Row {
        TableCell {
            width: row.width * 0.20
            height: row.height
            content: model.name
        }
        TableCell {
            width: row.width * 0.20
            height: row.height
            content: getName()
        }
        TableCell {
            width: row.width * 0.20
            height: row.height
            content: (model.localCRC.length > 0 ? localCRC : "not on disk")
        }
        TableCell {
            width: row.width * 0.20
            height: row.height
            content: model.remoteCRC
        }

         Item { 
            width: row.width * 0.20
            height: row.height

            Text {
                color: (model.fileAction == "download" ? "#64d168" : "#fa5757")
                font.pixelSize: 12
                font.family: beaufortbold.name
                text: model.fileAction
                anchors.verticalCenter: parent.verticalCenter
            }

            Separator{}
        }
    }

    function getName() {
        var path = model.d2Path
        var parts = path.split("/")
        
        var name = parts[parts.length - 1]
        if(name == "") {
            name = "Unknown"
        }

        return name
    }
}
