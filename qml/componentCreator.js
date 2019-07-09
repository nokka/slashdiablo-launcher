function createSettingsView(parent, model) {
    var component = Qt.createComponent("SettingsView.qml")
    var view = component.createObject(null)

    if (view == null) {
        console.log("failed to create settings view")
    }
    
    return view
}