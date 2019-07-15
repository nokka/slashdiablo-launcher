import QtQuick 2.12

Rectangle { // size controlled by height
    id: root
    
// public
    property bool checked: false; // oncheckedChanged:  print('oncheckedChanged', checked)       

// private
    width: 250;  height: 50 // default size    
    border.width: 0.05 * root.height
    radius:       0.5  * root.height
    color:        checked? 'white': 'black' // background
    opacity:      enabled? 1: 0.3 // disabled state
    
    Text {
        text:  checked?    'On': 'Off'
        color: checked? 'black': 'white'
        x:    (checked? 0: pill.width) + (parent.width - pill.width - width) / 2
        font.pixelSize: 0.5 * root.height
        anchors.verticalCenter: parent.verticalCenter        
    }
    
    MouseArea { // must be beneath pill MouseArea
        anchors.fill: parent
        onPressed:    parent.opacity = 0.5 // down state
        onReleased:   parent.opacity = 1
        onCanceled:   parent.opacity = 1
        onClicked:    checked = !checked
    }
    
    Rectangle { // pill
        id: pill
        
        x: checked? root.width - pill.width: 0 // binding must not be broken with imperative x = ...
        width: root.height;  height: width // square
        border.width: parent.border.width
        radius:       parent.radius
        
        MouseArea {
            anchors.fill: parent
            
            drag {
                target:   pill
                axis:     Drag.XAxis
                minimumX: 0
                maximumX: root.width - pill.width
            }
            
            onPressed:    parent.opacity = 0.5 // down state
            onReleased: { // releasing at the end of drag
                parent.opacity = 1
                if( checked  &amp;&amp;  pill.x &lt; root.width - pill.width)  checked = false // right to left
                if(!checked  &amp;&amp;  pill.x)                            checked = true  // left  to right
            }
            onCanceled:   parent.opacity = 1
            onClicked:  checked = !checked // clicking on pill             
        }
    }
}