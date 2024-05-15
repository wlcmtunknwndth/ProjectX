<script setup>
import {onMounted} from "vue";

import SelectorComponent from "../components/SelectorComponent.vue";
import ListComponent from "../components/ListComponent.vue";

import {useEventsStore} from "../store/events.js";
const eventsStore = useEventsStore();

onMounted(() => {
  eventsStore.getEvents();
});
</script>

<template>
  <main id="page"
        class="flexboxgrid_container__YXz7C CoreLayout_page__LDIK4 CoreLayout_pagePaddingTop__WezD8 CoreLayout_pagePaddingBottom__vI8wg">
    <h1 class="_collectionAlias__title__dFmHC">Афиша для людей с ограниченными возможностями и особенностями здоровья</h1>
    <div class="jsx-1021369771 DateRangePicker_container__EctZ_">
      <div class="DateRangePickerCalendar_container__tVkUb">
        <button type="button"
                class="Button_button__znY7u Button_secondaryTheme___s8Ei Button_medium__UJyor DateRangePickerCalendar_button__oj7i0">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
            <path fill="currentColor" d="M15 13.75a1.25 1.25 0 1 1 0 2.5 1.25 1.25 0 0 1 0-2.5Z"></path>
            <path fill="currentColor" fill-rule="evenodd"
                  d="M7 3.42c-.905.206-1.609.561-2.314 1.266-1.313 1.313-1.41 2.62-1.605 5.237C3.031 10.6 3 11.3 3 12s.03 1.399.081 2.077c.195 2.616.292 3.924 1.605 5.237 1.313 1.313 2.62 1.41 5.237 1.605C10.6 20.969 11.3 21 12 21s1.399-.03 2.077-.081c2.616-.195 3.924-.292 5.237-1.605 1.313-1.313 1.41-2.62 1.605-5.237C20.969 13.4 21 12.7 21 12s-.03-1.399-.081-2.077c-.195-2.616-.292-3.924-1.605-5.237-.705-.705-1.409-1.06-2.314-1.267V2a1 1 0 1 0-2 0v1.152a97.922 97.922 0 0 0-.923-.071A28.108 28.108 0 0 0 12 3c-.7 0-1.399.03-2.077.081-.327.024-.634.047-.923.071V2a1 1 0 0 0-2 0v1.42ZM16 7a1 1 0 0 1-1-1v-.838a72.028 72.028 0 0 0-1.071-.086A26.108 26.108 0 0 0 12 5c-.642 0-1.292.028-1.929.076A63.4 63.4 0 0 0 9 5.162V6a1 1 0 0 1-2 0v-.503c-.37.137-.62.323-.9.603-.603.603-.774 1.066-.938 2.9h13.676c-.164-1.834-.335-2.297-.938-2.9-.28-.28-.53-.466-.9-.603V6a1 1 0 0 1-1 1ZM5.021 11a24.166 24.166 0 0 0 .055 2.929c.204 2.75.318 3.265 1.024 3.97.706.707 1.22.82 3.971 1.025.637.048 1.287.076 1.929.076.642 0 1.292-.028 1.929-.076 2.75-.204 3.265-.318 3.97-1.024.707-.706.82-1.22 1.025-3.971A24.044 24.044 0 0 0 18.978 11H5.022Z"
                  clip-rule="evenodd"></path>
          </svg>
        </button>
      </div>
      <div class="jsx-1021369771 DateRangePicker_sliderContainer__sMFFs">
        <div class="DateRangePickerControl_container__alln0 DateRangePickerControl_indentAfterCalendar__C9yQw">
          <button type="button"
                  class="Button_button__znY7u Button_lightTheme__mOFjb Button_small__9c32H DateRangePickerControl_button__ZQHIi">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
              <path
                  d="M15 5.012c0-.9-1.077-1.35-1.706-.714l-3.096 3.13C8.066 9.582 7 10.66 7 12c0 1.34 1.066 2.417 3.198 4.573l3.096 3.129c.63.637 1.706.186 1.706-.715 0-.267-.105-.524-.293-.714l-3.095-3.13c-1.106-1.118-1.79-1.815-2.223-2.389-.39-.517-.39-.697-.39-.751v-.006c0-.053 0-.234.39-.751.432-.574 1.117-1.272 2.223-2.39l3.095-3.129c.188-.19.293-.447.293-.715Z"></path>
            </svg>
          </button>
        </div>
        <div class="jsx-1021369771 DateRangePicker_slider__LMkgi CalendarFilter_calendarSlider__iZ5SK">
          <div class="DateRangePickerMonth_container__6QV6K">
            <button type="button"
                    class="Button_button__znY7u Button_lightTheme__mOFjb Button_small__9c32H DateRangePickerMonth_month__4hpUD">
              май
            </button>
            <div class="DateRangePickerMonth_days__TVJ3i">
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34 DateRangePickerDay_afterCalendarDay__U6TGO"
                      disabled="" data-date="09.05.2024"><span>9</span><span
                  class="DateRangePickerDay_weekday__5_oP9">чт</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="10.05.2024"><span>10</span><span
                  class="DateRangePickerDay_weekday__5_oP9">пт</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34 DateRangePickerDay_weekend__cpxdC"
                      data-date="11.05.2024"><span>11</span><span
                  class="DateRangePickerDay_weekday__5_oP9">сб</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34 DateRangePickerDay_weekend__cpxdC"
                      data-date="12.05.2024"><span>12</span><span
                  class="DateRangePickerDay_weekday__5_oP9">вс</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="13.05.2024"><span>13</span><span
                  class="DateRangePickerDay_weekday__5_oP9">пн</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="14.05.2024"><span>14</span><span
                  class="DateRangePickerDay_weekday__5_oP9">вт</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="15.05.2024"><span>15</span><span
                  class="DateRangePickerDay_weekday__5_oP9">ср</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="16.05.2024"><span>16</span><span
                  class="DateRangePickerDay_weekday__5_oP9">чт</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="17.05.2024"><span>17</span><span
                  class="DateRangePickerDay_weekday__5_oP9">пт</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34 DateRangePickerDay_weekend__cpxdC"
                      data-date="18.05.2024"><span>18</span><span
                  class="DateRangePickerDay_weekday__5_oP9">сб</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34 DateRangePickerDay_weekend__cpxdC"
                      data-date="19.05.2024"><span>19</span><span
                  class="DateRangePickerDay_weekday__5_oP9">вс</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="20.05.2024"><span>20</span><span
                  class="DateRangePickerDay_weekday__5_oP9">пн</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="21.05.2024"><span>21</span><span
                  class="DateRangePickerDay_weekday__5_oP9">вт</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="22.05.2024"><span>22</span><span
                  class="DateRangePickerDay_weekday__5_oP9">ср</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="23.05.2024"><span>23</span><span
                  class="DateRangePickerDay_weekday__5_oP9">чт</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="24.05.2024"><span>24</span><span
                  class="DateRangePickerDay_weekday__5_oP9">пт</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34 DateRangePickerDay_weekend__cpxdC"
                      data-date="25.05.2024"><span>25</span><span
                  class="DateRangePickerDay_weekday__5_oP9">сб</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34 DateRangePickerDay_weekend__cpxdC"
                      data-date="26.05.2024"><span>26</span><span
                  class="DateRangePickerDay_weekday__5_oP9">вс</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="27.05.2024"><span>27</span><span
                  class="DateRangePickerDay_weekday__5_oP9">пн</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="28.05.2024"><span>28</span><span
                  class="DateRangePickerDay_weekday__5_oP9">вт</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="29.05.2024"><span>29</span><span
                  class="DateRangePickerDay_weekday__5_oP9">ср</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="30.05.2024"><span>30</span><span
                  class="DateRangePickerDay_weekday__5_oP9">чт</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="31.05.2024"><span>31</span><span
                  class="DateRangePickerDay_weekday__5_oP9">пт</span></button>
            </div>
          </div>
          <div class="DateRangePickerMonth_container__6QV6K">
            <button type="button"
                    class="Button_button__znY7u Button_lightTheme__mOFjb Button_small__9c32H DateRangePickerMonth_month__4hpUD">
              июнь
            </button>
            <div class="DateRangePickerMonth_days__TVJ3i">
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34 DateRangePickerDay_weekend__cpxdC"
                      data-date="01.06.2024"><span>1</span><span class="DateRangePickerDay_weekday__5_oP9">сб</span>
              </button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34 DateRangePickerDay_weekend__cpxdC"
                      data-date="02.06.2024"><span>2</span><span class="DateRangePickerDay_weekday__5_oP9">вс</span>
              </button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="03.06.2024"><span>3</span><span class="DateRangePickerDay_weekday__5_oP9">пн</span>
              </button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="04.06.2024"><span>4</span><span class="DateRangePickerDay_weekday__5_oP9">вт</span>
              </button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="05.06.2024"><span>5</span><span class="DateRangePickerDay_weekday__5_oP9">ср</span>
              </button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="06.06.2024"><span>6</span><span class="DateRangePickerDay_weekday__5_oP9">чт</span>
              </button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="07.06.2024"><span>7</span><span class="DateRangePickerDay_weekday__5_oP9">пт</span>
              </button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34 DateRangePickerDay_weekend__cpxdC"
                      data-date="08.06.2024"><span>8</span><span class="DateRangePickerDay_weekday__5_oP9">сб</span>
              </button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34 DateRangePickerDay_weekend__cpxdC"
                      data-date="09.06.2024"><span>9</span><span class="DateRangePickerDay_weekday__5_oP9">вс</span>
              </button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="10.06.2024"><span>10</span><span
                  class="DateRangePickerDay_weekday__5_oP9">пн</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="11.06.2024"><span>11</span><span
                  class="DateRangePickerDay_weekday__5_oP9">вт</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="12.06.2024"><span>12</span><span
                  class="DateRangePickerDay_weekday__5_oP9">ср</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="13.06.2024"><span>13</span><span
                  class="DateRangePickerDay_weekday__5_oP9">чт</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="14.06.2024"><span>14</span><span
                  class="DateRangePickerDay_weekday__5_oP9">пт</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34 DateRangePickerDay_weekend__cpxdC"
                      data-date="15.06.2024"><span>15</span><span
                  class="DateRangePickerDay_weekday__5_oP9">сб</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34 DateRangePickerDay_weekend__cpxdC"
                      data-date="16.06.2024"><span>16</span><span
                  class="DateRangePickerDay_weekday__5_oP9">вс</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="17.06.2024"><span>17</span><span
                  class="DateRangePickerDay_weekday__5_oP9">пн</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="18.06.2024"><span>18</span><span
                  class="DateRangePickerDay_weekday__5_oP9">вт</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="19.06.2024"><span>19</span><span
                  class="DateRangePickerDay_weekday__5_oP9">ср</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="20.06.2024"><span>20</span><span
                  class="DateRangePickerDay_weekday__5_oP9">чт</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="21.06.2024"><span>21</span><span
                  class="DateRangePickerDay_weekday__5_oP9">пт</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34 DateRangePickerDay_weekend__cpxdC"
                      data-date="22.06.2024"><span>22</span><span
                  class="DateRangePickerDay_weekday__5_oP9">сб</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34 DateRangePickerDay_weekend__cpxdC"
                      data-date="23.06.2024"><span>23</span><span
                  class="DateRangePickerDay_weekday__5_oP9">вс</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="24.06.2024"><span>24</span><span
                  class="DateRangePickerDay_weekday__5_oP9">пн</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="25.06.2024"><span>25</span><span
                  class="DateRangePickerDay_weekday__5_oP9">вт</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="26.06.2024"><span>26</span><span
                  class="DateRangePickerDay_weekday__5_oP9">ср</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="27.06.2024"><span>27</span><span
                  class="DateRangePickerDay_weekday__5_oP9">чт</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="28.06.2024"><span>28</span><span
                  class="DateRangePickerDay_weekday__5_oP9">пт</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34 DateRangePickerDay_weekend__cpxdC"
                      data-date="29.06.2024"><span>29</span><span
                  class="DateRangePickerDay_weekday__5_oP9">сб</span></button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34 DateRangePickerDay_weekend__cpxdC"
                      data-date="30.06.2024"><span>30</span><span
                  class="DateRangePickerDay_weekday__5_oP9">вс</span></button>
            </div>
          </div>
          <div class="DateRangePickerMonth_container__6QV6K">
            <button type="button"
                    class="Button_button__znY7u Button_lightTheme__mOFjb Button_small__9c32H DateRangePickerMonth_month__4hpUD">
              июль
            </button>
            <div class="DateRangePickerMonth_days__TVJ3i">
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="01.07.2024"><span>1</span><span class="DateRangePickerDay_weekday__5_oP9">пн</span>
              </button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="02.07.2024"><span>2</span><span class="DateRangePickerDay_weekday__5_oP9">вт</span>
              </button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="03.07.2024"><span>3</span><span class="DateRangePickerDay_weekday__5_oP9">ср</span>
              </button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="04.07.2024"><span>4</span><span class="DateRangePickerDay_weekday__5_oP9">чт</span>
              </button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="05.07.2024"><span>5</span><span class="DateRangePickerDay_weekday__5_oP9">пт</span>
              </button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34 DateRangePickerDay_weekend__cpxdC"
                      data-date="06.07.2024"><span>6</span><span class="DateRangePickerDay_weekday__5_oP9">сб</span>
              </button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34 DateRangePickerDay_weekend__cpxdC"
                      data-date="07.07.2024"><span>7</span><span class="DateRangePickerDay_weekday__5_oP9">вс</span>
              </button>
              <button type="button"
                      class="Button_button__znY7u Button_lightTheme__mOFjb Button_medium__UJyor DateRangePickerDay_container__4AR34"
                      data-date="08.07.2024"><span>8</span><span class="DateRangePickerDay_weekday__5_oP9">пн</span>
              </button>
            </div>
          </div>
        </div>
        <div
            class="DateRangePickerControl_container__alln0 DateRangePickerControl_visible__keqMU DateRangePickerControl_next__ztckQ">
          <button type="button"
                  class="Button_button__znY7u Button_lightTheme__mOFjb Button_small__9c32H DateRangePickerControl_button__ZQHIi">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
              <path
                  d="M9 5.012c0-.9 1.077-1.35 1.706-.714l3.096 3.13C15.934 9.582 17 10.66 17 12c0 1.34-1.066 2.417-3.198 4.573l-3.096 3.129c-.63.637-1.706.186-1.706-.715 0-.267.105-.524.293-.714l3.095-3.13c1.106-1.118 1.79-1.815 2.223-2.389.39-.517.39-.697.39-.751v-.006c0-.053 0-.234-.39-.751-.432-.574-1.117-1.272-2.223-2.39L9.293 5.727A1.016 1.016 0 0 1 9 5.012Z"></path>
            </svg>
          </button>
        </div>
      </div>
    </div>
    <div class="FiltersContainer_container__2b5CT">
      <div class="FiltersContainer_filters__slbB5">
        <div class="Select_rawInputContainer__6AfZX">
          <div class="Tag_container__aLkPQ SelectRawInput_container__DRZ5_"><span class="Tag_children__r1Gec">Площадка</span>
            <div class="Tag_iconContainer__NZM9K">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16"
                   class="Tag_icon__yzj_d Tag_triangleIcon__iuwqO">
                <path fill="currentColor"
                      d="M8.78 10.024a1 1 0 0 1-1.56 0l-1.92-2.4A1 1 0 0 1 6.08 6h3.84a1 1 0 0 1 .78 1.625l-1.92 2.399Z"></path>
              </svg>
            </div>
          </div>
        </div>
        <div class="Tag_container__aLkPQ Tag_cursorPointer__QGOOJ"><span class="Tag_children__r1Gec">Пушкинская карта</span>
        </div>
        <div class="Select_rawInputContainer__6AfZX">
          <div class="Tag_container__aLkPQ SelectRawInput_container__DRZ5_"><span class="Tag_children__r1Gec">По рейтингу</span>
            <div class="Tag_iconContainer__NZM9K">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16"
                   class="Tag_icon__yzj_d Tag_triangleIcon__iuwqO">
                <path fill="currentColor"
                      d="M8.78 10.024a1 1 0 0 1-1.56 0l-1.92-2.4A1 1 0 0 1 6.08 6h3.84a1 1 0 0 1 .78 1.625l-1.92 2.399Z"></path>
              </svg>
            </div>
          </div>
        </div>

        <SelectorComponent />
      </div>

    </div>
    <div class="PreviewGrid_container__3Lop_ _collectionAlias__previewGrid__PXzav">

      <ListComponent />

      <button type="button"
              class="Button_button__znY7u Button_secondaryTheme___s8Ei Button_medium__UJyor PreviewGrid_loadMoreButton__9tcw_">
        Показать ещё
      </button>
    </div>
  </main>
</template>

<style scoped>

</style>