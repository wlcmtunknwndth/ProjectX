import { defineStore } from 'pinia'
import {eventsMock} from "../mock/events.js";
// import {http} from "../network/http.js";

export const useEventsStore = defineStore('events', {
  state: () => ({
    selector: [
      {title: "Маломобильным гражданам", isSelected: true, type: "disability"},
      {title: "Глухим и слабослышащим", isSelected: true, type: "deaf"},
      {title: "Незрячим и слабовидящим", isSelected: true, type: "blind"},
      {title: "С нейроотличиями", isSelected: true, type: "neuro"},
    ],

    events: [],
    filteredEvents: [],
    detailEvent: null,
  }),

  actions: {
    getEvents() {
      this.events = eventsMock;
      this.resetFilteredEvents();
    },

    getDetailEvent(eventId) {
      this.detailEvent = eventsMock.find((event) => event.id === eventId);
    },

    getFilteredEvents() {
      const resultArr = [];

      for (const e of this.events) {
        for (const sf of this.selectedFeatures) {
          if (e.feature.includes(sf)) {
            resultArr.push(e)
          }
        }
      }

      this.filteredEvents = Array.from(new Set(resultArr));
    },

    resetFilteredEvents() {
      this.filteredEvents = this.events;

      for (const s of this.selector) {
        s.isSelected = true;
      }
    }
  },

  getters: {
    selectedFeatures: (state) => {
      return state.selector.filter((s) => s.isSelected === true).map((s) => s.type)
    },
  }
})

// http.get('/event?id=1').then((response) => {
//   console.log(response)
// })
//
// http.get('/events').then((response) => {
//   console.log(response)
// })

// http.post('/create_event', {
//   price: 2000,
//   restrictions: 10,
//   date: "2024-05-10T17:00:00+00:00",
//   feature: "deaf",
//   city: "moscow",
//   address: "Двор Гостинки",
//   name: "Концерт группы «Три дня дождя». Summer Sound x билайн",
//   description: "chainsaw gutsfuck"
// }).then((response) => {
//   console.log(response)
// })