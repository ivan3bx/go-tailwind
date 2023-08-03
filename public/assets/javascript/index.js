import { Application } from "@hotwired/stimulus"
import HelloController from "./hello_controller.js"
import DropdownController from "./dropdown_controller.js"

window.Stimulus = Application.start()

//
// register controllers here
//
Stimulus.register("dropdown", DropdownController)