import { Application } from "@hotwired/stimulus"
import HelloController from "./hello_controller.js"

window.Stimulus = Application.start()
Stimulus.register("hello", HelloController)