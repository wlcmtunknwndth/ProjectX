<script setup>
import {computed, onMounted, ref} from "vue";

import {useRoute} from 'vue-router'
const route = useRoute()

const isOpen = ref(false)

const additionalClasses = computed(() => ({
  'CoreLayoutMenuItem_moreDropdownOpen__l4gDA': isOpen.value
}))

onMounted(() => {
  addEventListener('click', () => isOpen.value = false)
})
</script>

<template>
  <div
      style="position: relative;"
      @click="(e) => e.stopPropagation()"
  >
    <div
        class="CoreLayoutMenuItem_container__luLtl CoreLayoutMenuItem_moreContainer__7mrXT CoreLayoutMenu_moreMenuItem__AhIlC"
        style="padding-right: 0;"
        @click="isOpen = !isOpen"
    >
      <div
          class="CoreLayoutMenuItem_moreInnerContainer__XBqEn"
          :class="additionalClasses"
      >
        Ещё
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" class="CoreLayoutMenuItem_triangleIcon__qJYXN">
          <path fill="currentColor"
                d="M8.78 10.024a1 1 0 0 1-1.56 0l-1.92-2.4A1 1 0 0 1 6.08 6h3.84a1 1 0 0 1 .78 1.625l-1.92 2.399Z"></path>
        </svg>
      </div>
    </div>

    <div
        v-if="isOpen"
        style="max-width: 240px; position: absolute; top: calc(100% + 12px); right: 0;"
    >
      <ul class="CoreLayoutMenu_submenu__PzcHT">
        <li class="CoreLayoutMenuItem_container__luLtl CoreLayoutMenu_menuItem__JDxAI">
          <a class="CoreLayoutMenuItem_link__Bo_mY" target="_self" href="/giftcard">
            Подарочная карта
          </a>
        </li>
        <li class="CoreLayoutMenuItem_container__luLtl CoreLayoutMenu_menuItem__JDxAI">
          <a class="CoreLayoutMenuItem_link__Bo_mY" target="_self" href="/lifestyle">
            LifeStyle
          </a>
        </li>
        <li class="CoreLayoutMenuItem_container__luLtl CoreLayoutMenu_menuItem__JDxAI">
          <a class="CoreLayoutMenuItem_link__Bo_mY" target="_self" href="/sankt-peterburg/collections/show">
            Шоу
          </a>
        </li>
        <li class="CoreLayoutMenuItem_container__luLtl CoreLayoutMenu_menuItem__JDxAI">
          <a class="CoreLayoutMenuItem_link__Bo_mY" target="_self" href="/sankt-peterburg/collections/festivals">
            Фестивали
          </a>
        </li>
        <li class="CoreLayoutMenuItem_container__luLtl CoreLayoutMenu_menuItem__JDxAI">
          <a class="CoreLayoutMenuItem_link__Bo_mY" target="_self" href="/sankt-peterburg/collections/exhibitions">
            Выставки
          </a>
        </li>
        <li class="CoreLayoutMenuItem_container__luLtl CoreLayoutMenu_menuItem__JDxAI">
          <a class="CoreLayoutMenuItem_link__Bo_mY" target="_self" href="/sankt-peterburg/collections/excursions">
            Экскурсии
          </a>
        </li>
        <li class="CoreLayoutMenuItem_container__luLtl CoreLayoutMenu_menuItem__JDxAI">
          <a class="CoreLayoutMenuItem_link__Bo_mY" target="_self" href="/sankt-peterburg/collections/sport">
            Спорт
          </a>
        </li>
        <li
            class="CoreLayoutMenuItem_container__luLtl CoreLayoutMenu_menuItem__JDxAI"
            :class="{'CoreLayoutMenuItem_active__hUtMn': route.name === 'main' }"
            style="white-space: normal;height: auto;line-height: normal;"
        >
          <router-link :to="{ name: 'main' }">
            <div class="CoreLayoutMenuItem_link__Bo_mY">
              Доступно для людей с ограниченными возможностями
            </div>
          </router-link>
        </li>
      </ul>
    </div>
  </div>
</template>