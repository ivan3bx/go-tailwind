import { Controller } from "@hotwired/stimulus"

export default class extends Controller {
    static targets = ["name"]
  
    connect() {
      console.log("Hello, Stimulus i am here.")
    }
}