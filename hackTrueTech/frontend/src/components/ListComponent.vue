<script setup>

import CardComponent from "./CardComponent.vue";

import {useEventsStore} from "../store/events.js";
const eventsStore = useEventsStore();
</script>

<template>
  <div class="PreviewGrid_gridContainer__3vrH_">
    <template
        v-for="event in eventsStore.filteredEvents"
        :key="event.id"
    >
      <CardComponent :event="event" />
    </template>
  </div>
</template>