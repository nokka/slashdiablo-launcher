function createSettingsView(parent) {
    var component = Qt.createComponent("SettingsView.qml")
    var view = component.createObject(parent)

    if (view == null) {
        console.log("failed to create settings view")
    }
    
    return view
}

function createArmoryView(parent) {
    var component = Qt.createComponent("ArmoryView.qml")
    var view = component.createObject(parent)

    if (view == null) {
        console.log("failed to create armory view")
    }
    
    return view
}

function createRulesView(parent) {
    var component = Qt.createComponent("RulesView.qml")
    var view = component.createObject(parent)

    if (view == null) {
        console.log("failed to create rules view")
    }
    
    return view
}

function createCommunityView(parent) {
    var component = Qt.createComponent("CommunityView.qml")
    var view = component.createObject(parent)

    if (view == null) {
        console.log("failed to create community view")
    }
    
    return view
}

function createSettingsPopup(parent) {
    var component = Qt.createComponent("SettingsPopup.qml")
    var popup = component.createObject(parent)

    if (popup == null) {
        console.log("Error creating settings popup")
    }
    return popup
}