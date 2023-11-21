import { Application } from "@hotwired/stimulus"
import DropdownController from "./dropdown_controller.js"

window.Stimulus = Application.start()

//
// register controllers here
//
Stimulus.register("dropdown", DropdownController)
