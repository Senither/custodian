<div>
    @can('update', $form->task)
        <button x-on:click="$wire.set('showModal', true)"
                class="bg-base-100 btn btn-square hover:btn-primary">
            <svg class="w-5 h-5 text-content" data-slot="icon" fill="none" stroke-width="1.5"
                 stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"
                 aria-hidden="true">
                <path stroke-linecap="round" stroke-linejoin="round"
                      d="m16.862 4.487 1.687-1.688a1.875 1.875 0 1 1 2.652 2.652L10.582 16.07a4.5 4.5 0 0 1-1.897 1.13L6 18l.8-2.685a4.5 4.5 0 0 1 1.13-1.897l8.932-8.931Zm0 0L19.5 7.125M18 14v4.75A2.25 2.25 0 0 1 15.75 21H5.25A2.25 2.25 0 0 1 3 18.75V8.25A2.25 2.25 0 0 1 5.25 6H10">
                </path>
            </svg>
        </button>

        @teleport('body')
            <dialog class="backdrop-blur-sm backdrop-grayscale modal"
                    x-bind:open="$wire.showModal"
                    x-on:close.stop="$wire.set('showModal', false)"
                    x-on:keydown.escape.window="$wire.set('showModal', false)"
                    x-on:keydown.tab.prevent="$event.shiftKey || nextFocusable().focus()"
                    x-on:keydown.shift.tab.prevent="prevFocusable().focus()">
                <div class="p-0 modal-box">
                    <div class="flex items-center gap-4 p-6">
                        <svg class="w-8 h-8 text-primary" data-slot="icon" fill="none" stroke-width="1.5" stroke="currentColor"
                             viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
                            <path stroke-linecap="round" stroke-linejoin="round"
                                  d="m16.862 4.487 1.687-1.688a1.875 1.875 0 1 1 2.652 2.652L10.582 16.07a4.5 4.5 0 0 1-1.897 1.13L6 18l.8-2.685a4.5 4.5 0 0 1 1.13-1.897l8.932-8.931Zm0 0L19.5 7.125M18 14v4.75A2.25 2.25 0 0 1 15.75 21H5.25A2.25 2.25 0 0 1 3 18.75V8.25A2.25 2.25 0 0 1 5.25 6H10">
                            </path>
                        </svg>
                        <h3 class="font-bold text-lg">Editing Task</h3>
                    </div>
                    <form wire:submit="saveTask" class="flex flex-col gap-4 px-6">
                        <label class="form-control w-full">
                            <div class="label">
                                <span class="label-text">What is the task?</span>
                            </div>
                            <input wire:model="form.message" type="text" placeholder="Eg. Do the thing..."
                                   class="input-bordered w-full input" />
                            <x-input-error :messages="$errors->get('form.message')" class="mt-2" />
                        </label>

                        <div class="gap-4 grid grid-cols-1 md:grid-cols-2">
                            <label class="form-control">
                                <div class="label">
                                    <span class="label-text">Category</span>
                                </div>
                                <select wire:model="form.category" class="select-bordered select">
                                    <option disabled>Pick one</option>
                                    @foreach ($categories as $category)
                                        <option value="{{ $category->id }}">{{ $category->name }}</option>
                                    @endforeach
                                </select>
                                <x-input-error :messages="$errors->get('form.category')" class="mt-2" />
                            </label>

                            <label class="form-control">
                                <div class="label">
                                    <span class="label-text">Priority</span>
                                </div>
                                <select wire:model="form.priority" class="select-bordered select">
                                    <option disabled>Pick one</option>
                                    @foreach ($priorities as $priority)
                                        <option value="{{ $priority->id }}">{{ $priority->name }}</option>
                                    @endforeach
                                </select>
                                <x-input-error :messages="$errors->get('form.priority')" class="mt-2" />
                            </label>
                        </div>
                    </form>

                    <div method="dialog" class="modal-backdrop">
                        <div class="gap-4 grid grid-cols-2 mt-4 px-6 py-4 border-t border-base-300">
                            <button wire:click="saveTask" class="text-primary-content btn btn-primary">Save Task</button>
                            <button x-on:click="$wire.set('showModal', false)" class="bg-base-200 border border-base-300 btn">Close</button>
                        </div>
                    </div>
                </div>
            </dialog>
        @endteleport
    @endcan
</div>
