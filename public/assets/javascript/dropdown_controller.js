import { Controller } from "@hotwired/stimulus"
import { toggle } from "el-transition"

export default class extends Controller {
    static targets = ["button", "content"]

    toggle() {
        toggle(this.contentTarget)
    }
}