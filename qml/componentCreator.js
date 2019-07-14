function createSettingsView(parent) {
    var component = Qt.createComponent("SettingsView.qml")
    var view = component.createObject(null)

    if (view == null) {
        console.log("failed to create settings view")
    }
    
    return view
}

function createArmoryView() {
    var component = Qt.createComponent("ArmoryView.qml")
    var view = component.createObject(null)

    if (view == null) {
        console.log("failed to create armory view")
    }
    
    return view
}

function createRulesView() {
    var component = Qt.createComponent("RulesView.qml")
    var view = component.createObject(null)

    if (view == null) {
        console.log("failed to create rules view")
    }
    
    return view
}

function createCommunityView() {
    var component = Qt.createComponent("CommunityView.qml")
    var view = component.createObject(null)

    if (view == null) {
        console.log("failed to create community view")
    }
    
    return view
}