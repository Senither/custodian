<div>
    @can('create', \App\Models\Task::class)
        <button x-on:click="$wire.set('showModal', true)" class="btn btn-primary">Add Task</button>

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
                              d="M12 18v-5.25m0 0a6.01 6.01 0 0 0 1.5-.189m-1.5.189a6.01 6.01 0 0 1-1.5-.189m3.75 7.478a12.06 12.06 0 0 1-4.5 0m3.75 2.383a14.406 14.406 0 0 1-3 0M14.25 18v-.192c0-.983.658-1.823 1.508-2.316a7.5 7.5 0 1 0-7.517 0c.85.493 1.509 1.333 1.509 2.316V18">
                        </path>
                    </svg>
                    <h3 class="font-bold text-lg">Creating new Task</h3>
                </div>
                <form wire:submit="createTask" class="flex flex-col gap-4 px-6">
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
                                <option :value="null">Pick one</option>
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
                                <option :value="null">Pick one</option>
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
                        <button wire:click="createTask" class="text-primary-content btn btn-primary">Create Task</button>
                        <button x-on:click="$wire.set('showModal', false)" class="bg-base-200 border border-base-300 btn">Close</button>
                    </div>
                </div>
            </div>
        </dialog>
    @endcan
</div>
