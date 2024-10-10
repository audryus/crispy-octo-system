class Toast {
    constructor(message, level, duration) {
        this.id = Date.now()
        this.message = message
        this.level = `alert-${level}`

        if (!level) {
            this.level = `alert-info`
        }

        if (!duration) {
            duration = 3000
        }

        setTimeout(() => {
            Alpine.store("toaster").remove(this.id)
        }, duration)
    }

    get html() {
        return `
        <div id="${this.id}" role="alert" class="alert ${this.level} shadow-lg bg-base flex justify-between duration-300 transition-opacity opacity-100 ease-in-out">
            <div class="w-56">
                <div class="flex justify-between">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                            <path stroke-linecap="round" stroke-linejoin="round" d="M11.25 11.25l.041-.02a.75.75 0 011.063.852l-.708 2.836a.75.75 0 001.063.853l.041-.021M21 12a9 9 0 11-18 0 9 9 0 0118 0zm-9-3.75h.008v.008H12V8.25z"></path>
                        </svg>
                    </svg>
                    <button class="btn btn-circle btn-sm" onclick="Alpine.store('toaster').remove(${this.id})">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                        </svg>
                    </button>
                </div>
                <div>
                    <h3 class="font-bold" data-i18n="RLY_ERR"></h3>
                    <div class="text-lg">This is a info</div>
                </div>
            </div>
        </div>
        `
    }
}

document.addEventListener('alpine:init', () => {
    Alpine.store('toaster', {
        init() {
            this.items = []
        },
        remove(id) {
            const ele = document.getElementById(id)
            ele.classList.remove("opacity-100")
            ele.classList.add("opacity-0")
            console.log("remove", id)

            const idx = this.items.findIndex((e) => e.id === id)
            setTimeout(() => {
                this.items.splice(0, 1)
            }, 1000)
        },
        delete(id) {
            const idx = this.items.findIndex((e) => e.id === id)
            delete myArray[idx]
        },
        items:[],
        get toasts() {
            return this.items
        },
        push(toast) {
            this.items.push(toast)
        }
    })
})

document.addEventListener("response", function(evt) {
    Alpine.store("toaster").push(new Toast(evt.detail.value))
})
document.addEventListener('htmx:afterSettle', function(evt) {
    localize('[data-i18n]')
});