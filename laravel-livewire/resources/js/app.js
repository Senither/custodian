import './bootstrap'

Livewire.hook('component.init', function () {
    if (typeof themeChange === 'function') {
        themeChange(false)
    }
})
