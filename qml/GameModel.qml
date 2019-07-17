import QtQuick 2.0

ListModel {
    ListElement {
        gameId: 1
        location: "D:/Games/Diablo II"
        instances: 2
        maphack: true
        hd: true
    }
    ListElement {
        gameId: 1
        location: "D:/Games/Diablo II-HD"
        instances: 4
        maphack: false
        hd: false
    }
}