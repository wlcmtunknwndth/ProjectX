<script setup>
import {computed, onMounted, ref} from "vue";

import {useEventsStore} from "../store/events.js";
const eventsStore = useEventsStore();

const isOpen = ref(false)

const additionalClasses = computed(() => ({
  'Tag_focused__QoTGa': isOpen.value
}))

onMounted(() => {
  addEventListener('click', () => isOpen.value = false)
})
</script>

<template>
  <div style="position: relative;" @click="(e) => e.stopPropagation()">
    <div
        class="Select_rawInputContainer__6AfZX"
        @click="isOpen = !isOpen"
    >
      <div
          class="Tag_container__aLkPQ SelectRawInput_container__DRZ5_"
          :class="additionalClasses"
      >
        <span class="Tag_children__r1Gec">
          Доступно
        </span>
        <div class="Tag_iconContainer__NZM9K">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" class="Tag_icon__yzj_d Tag_triangleIcon__iuwqO">
            <path fill="currentColor"
                  d="M8.78 10.024a1 1 0 0 1-1.56 0l-1.92-2.4A1 1 0 0 1 6.08 6h3.84a1 1 0 0 1 .78 1.625l-1.92 2.399Z"></path>
          </svg>
        </div>
      </div>
    </div>

    <div
        v-if="isOpen"
        class="Select_select__e_I5E-dropdown Select_multiple__XFlmk Select_select__e_I5E-dropdown-placement-bottomLeft"
        style="top: calc(100% + 12px);"
    >
      <div>
        <div class="Select_desktopDropdownContainer__XlOBU">
          <div class="rc-virtual-list">
            <div class="rc-virtual-list-holder">
              <div>
                <div class="rc-virtual-list-holder-inner"
                     style="max-height: 300px; overflow: auto;"
                >

                  <template
                      v-for="item in eventsStore.selector"
                      :key="item.type"
                  >
                    <div
                        :aria-selected="item.isSelected"
                        class="Select_select__e_I5E-item Select_select__e_I5E-item-option"
                        :class="{ 'Select_select__e_I5E-item-option-selected': item.isSelected }"
                        :title="item.title"
                        @click="item.isSelected = !item.isSelected"
                    >
                      <div class="Select_select__e_I5E-item-option-content">
                        {{ item.title }}
                      </div>
                      <span
                          class="Select_select__e_I5E-item-option-state"
                          unselectable="on"
                          aria-hidden="true"
                          style="user-select: none;"
                      >
                        <label class="Checkbox_label__uNY3B" for=":r2p:">
                          <svg v-if="!item.isSelected" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"
                               class="Checkbox_icon__JExvi Checkbox_uncheckIcon__khYnj"><defs><clipPath
                              id="Uncheck_svg__b"><use href="#Uncheck_svg__a"></use></clipPath><path id="Uncheck_svg__a"
                                                                                                     d="M4.686 4.686c-1.313 1.313-1.41 2.62-1.605 5.237C3.031 10.6 3 11.3 3 12s.03 1.399.081 2.077c.195 2.616.292 3.924 1.605 5.237 1.313 1.313 2.62 1.41 5.237 1.605C10.6 20.969 11.3 21 12 21s1.399-.03 2.077-.081c2.616-.195 3.924-.292 5.237-1.605 1.313-1.313 1.41-2.62 1.605-5.237C20.969 13.4 21 12.7 21 12s-.03-1.399-.081-2.077c-.195-2.616-.292-3.924-1.605-5.237-1.313-1.313-2.62-1.41-5.237-1.605A28.108 28.108 0 0 0 12 3c-.7 0-1.399.03-2.077.081-2.616.195-3.924.292-5.237 1.605Z"></path></defs><use
                              fill="transparent" stroke="currentColor" stroke-width="4" clip-path="url(#Uncheck_svg__b)"
                              href="#Uncheck_svg__a"></use></svg>
                          <svg v-if="item.isSelected" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"
                               class="Checkbox_icon__JExvi Checkbox_checkIcon__fdfAb"><path fill="currentColor"
                                                                                            fill-rule="evenodd"
                                                                                            d="M4.686 4.686c-1.313 1.313-1.41 2.62-1.605 5.237C3.031 10.6 3 11.3 3 12s.03 1.399.081 2.077c.195 2.616.292 3.924 1.605 5.237 1.313 1.313 2.62 1.41 5.237 1.605C10.6 20.969 11.3 21 12 21s1.399-.03 2.077-.081c2.616-.195 3.924-.292 5.237-1.605 1.313-1.313 1.41-2.62 1.605-5.237C20.969 13.4 21 12.7 21 12s-.03-1.399-.081-2.077c-.195-2.616-.292-3.924-1.605-5.237-1.313-1.313-2.62-1.41-5.237-1.605A28.108 28.108 0 0 0 12 3c-.7 0-1.399.03-2.077.081-2.616.195-3.924.292-5.237 1.605Zm11.521 5.521a1 1 0 0 0-1.414-1.414L11 12.586l-1.293-1.293a1 1 0 0 0-1.414 1.414l2 2a1 1 0 0 0 1.414 0l4.5-4.5Z"
                                                                                            clip-rule="evenodd"></path></svg>
                        </label>
                      </span>
                    </div>
                  </template>

                </div>
              </div>
            </div>
          </div>
          <div class="SelectDropdownFooter_container__Otbwz">
            <button
                type="button"
                class="Button_button__znY7u Button_secondaryTheme___s8Ei Button_medium__UJyor"
                @click="eventsStore.resetFilteredEvents()"
            >
              Сбросить
            </button>
            <button
                type="button"
                class="Button_button__znY7u Button_primaryTheme__xtq5p Button_medium__UJyor"
                @click="eventsStore.getFilteredEvents()"
            >
              Применить
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
