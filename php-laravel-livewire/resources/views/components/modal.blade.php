<dialog x-data="{ show: @entangle($attributes->wire('model')) }"
        class="backdrop-blur-sm backdrop-grayscale modal"
        x-on:close.stop="show = false"
        x-on:keydown.escape.window="show = false"
        x-trap="show"
        x-bind:open="show">
    <div class="p-0 modal-box">
        {{ $slot }}
    </div>
</dialog>
