<div>
    <button x-on:click="$wire.set('showModal', true)"
            class="bg-base-100 btn btn-square hover:btn-error">
        <svg class="w-5 h-5 text-content" data-slot="icon" fill="none" stroke-width="1.5"
             stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"
             aria-hidden="true">
            <path stroke-linecap="round" stroke-linejoin="round"
                  d="m14.74 9-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 0 1-2.244 2.077H8.084a2.25 2.25 0 0 1-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 0 0-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 0 1 3.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 0 0-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 0 0-7.5 0">
            </path>
        </svg>
    </button>

    <dialog class="backdrop-blur-sm backdrop-grayscale modal"
            x-bind:open="$wire.showModal"
            x-on:close.stop="$wire.set('showModal', false)"
            x-on:keydown.escape.window="$wire.set('showModal', false)"
            x-on:keydown.tab.prevent="$event.shiftKey || nextFocusable().focus()"
            x-on:keydown.shift.tab.prevent="prevFocusable().focus()">
        <div class="p-0 modal-box">
            <div class="flex items-center gap-4 p-6">
                <svg class="w-8 h-8 text-warning" data-slot="icon" fill="none" stroke-width="1.5" stroke="currentColor"
                     viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
                    <path stroke-linecap="round" stroke-linejoin="round"
                          d="M12 9v3.75m-9.303 3.376c-.866 1.5.217 3.374 1.948 3.374h14.71c1.73 0 2.813-1.874 1.948-3.374L13.949 3.378c-.866-1.5-3.032-1.5-3.898 0L2.697 16.126ZM12 15.75h.007v.008H12v-.008Z">
                    </path>
                </svg>
                <h3 class="font-bold text-lg">Are you sure you want to continue?</h3>
            </div>
            <div class="px-6 prose">
                <p>
                    You're about the delete the "{{ $task->message }}" task, this is a permanent action and cannot be reversed.
                </p>
                <p>Are you sure you want to continue?</p>
            </div>
            <form method="dialog" class="modal-backdrop">
                <div class="gap-4 grid grid-cols-2 mt-4 px-6 py-4 border-t border-base-300">
                    <button wire:click="delete" class="btn btn-warning">Delete Task</button>
                    <button x-on:click="$wire.set('showModal', false)" class="bg-base-200 border border-base-300 btn">Close</button>
                </div>
            </form>
        </div>
    </dialog>
</div>
