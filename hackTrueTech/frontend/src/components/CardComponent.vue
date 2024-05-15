<script setup>
import {ref} from "vue";

defineProps({
  event: Object,
})

const addressLink = ref('#');
const link = ref('#');
</script>

<template>
  <article
      class="AnnouncementPreview_container__A430m AnnouncementPreview_mobile__V6O8j AnnouncementPreview_contentVisibility__424Pd">
    <div class="AnnouncementPreviewImage_container__Pmp0c">
      <router-link :to="{ name: 'detail', params: { id: event.id } }">
        <div
            :title="event.name"
            class="AnnouncementPreviewImage_link__5btQS"
            data-type="izobrazhenie"
        >
          <span class="LazyImage_container__ukc21">
          <picture class="LazyImage_backgroundContainer__UuhKK">
            <img
                loading="lazy"
                class="LazyImage_img__Nz285"
                :alt="event.name"
                :src="event.img_path"
            />
          </picture>
          <picture
              class="LazyImage_originalContainer__Kj81p LazyImage_originalVisible__Xn1yN LazyImage_originalHoverAnimation__jNykx">
            <img
                loading="eager"
                class="LazyImage_img__Nz285 LazyImage_original___06Jx"
                :src="event.img_path"
                style="transform: scale(1);"
            />
          </picture>
        </span>
        </div>
      </router-link>
      <div class="FavouritesItemStatusButton_container__KIJN8">
        <button class="FavouritesItemStatusButton_button__VRYUO">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 60 60" class="FavouritesItemStatusButton_icon__maJuJ">
            <defs>
              <clipPath id="FavouritesDynamic_svg__c">
                <use href="#FavouritesDynamic_svg__a"></use>
              </clipPath>
              <filter id="FavouritesDynamic_svg__b" width="28.002" height="24.746" x="-2" y="0.252"
                      color-interpolation-filters="sRGB" filterUnits="userSpaceOnUse">
                <feFlood flood-opacity="0" result="BackgroundImageFix"></feFlood>
                <feColorMatrix in="SourceAlpha" result="hardAlpha"
                               values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"></feColorMatrix>
                <feOffset dy="1"></feOffset>
                <feGaussianBlur stdDeviation="2"></feGaussianBlur>
                <feComposite in2="hardAlpha" operator="out"></feComposite>
                <feColorMatrix values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.25 0"></feColorMatrix>
                <feBlend in2="BackgroundImageFix" result="effect1_dropShadow_12565_118238"></feBlend>
                <feBlend in="SourceGraphic" in2="effect1_dropShadow_12565_118238" result="shape"></feBlend>
              </filter>
              <path id="FavouritesDynamic_svg__a"
                    d="M2.421 13c.52 1.264 1.527 2.293 3.019 3.82a27.694 27.694 0 0 0 1.722 1.61c1.987 1.711 2.981 2.567 4.838 2.567 1.857 0 2.85-.856 4.838-2.568a27.622 27.622 0 0 0 1.722-1.61c1.492-1.526 2.498-2.555 3.019-3.819a5.62 5.62 0 0 0 .401-1.81c.027-.408.031-.82 0-1.228-.242-3.198-2.705-5.71-5.707-5.71a5.346 5.346 0 0 0-2.774.778 5.901 5.901 0 0 0-1.508 1.315 5.859 5.859 0 0 0-1.513-1.309 5.342 5.342 0 0 0-2.752-.765C4.718 4.27 2.252 6.793 2.018 10a6.867 6.867 0 0 0-.018.5c0 .848.089 1.694.421 2.5Zm1.637-1.525c.18 1.176.841 1.93 2.812 3.946.51.522 1.05 1.02 1.597 1.493 2.09 1.8 2.534 2.083 3.533 2.083.999 0 1.443-.283 3.532-2.083a25.63 25.63 0 0 0 1.598-1.493c1.95-1.995 2.619-2.754 2.806-3.91.026-.342.064-.686.064-1.03 0-2.497-1.823-4.229-3.727-4.229-1.044 0-2.025.492-2.733 1.358l-1.543 1.888-1.55-1.882c-.706-.858-1.682-1.345-2.721-1.345C5.823 6.27 4 8.002 4 10.5c0 .325.035.65.058.975Z"></path>
            </defs>
            <g data-id="heartContainer" filter="url(#FavouritesDynamic_svg__b)">
              <use clip-path="url(#FavouritesDynamic_svg__c)" data-id="heart" href="#FavouritesDynamic_svg__a"></use>
            </g>
            <circle cx="30" cy="30" r="1.5" fill="#FF0032" data-id="circle"></circle>
            <g data-id="grp7" transform="translate(7 6)">
              <circle cx="2" cy="6" r="2" fill="#9CD8C3" data-id="oval1"></circle>
              <circle cx="5" cy="2" r="2" fill="#8CE8C3" data-id="oval2"></circle>
            </g>
            <g data-id="grp6" transform="translate(0 28)">
              <circle cx="2" cy="7" r="2" fill="#CC8EF5" data-id="oval1"></circle>
              <circle cx="3" cy="2" r="2" fill="#91D2FA" data-id="oval2"></circle>
            </g>
            <g data-id="grp3" transform="translate(52 28)">
              <circle cx="2" cy="7" r="2" fill="#9CD8C3" data-id="oval2"></circle>
              <circle cx="4" cy="2" r="2" fill="#8CE8C3" data-id="oval1"></circle>
            </g>
            <g fill="#CC8EF5" data-id="grp2" transform="translate(44 6)">
              <circle cx="5" cy="6" r="2" data-id="oval2"></circle>
              <circle cx="2" cy="2" r="2" data-id="oval1"></circle>
            </g>
            <g fill="#91D2FA" data-id="grp5" transform="translate(14 50)">
              <circle cx="6" cy="5" r="2" data-id="oval1"></circle>
              <circle cx="2" cy="2" r="2" data-id="oval2"></circle>
            </g>
            <g fill="#F48EA7" data-id="grp4" transform="translate(35 50)">
              <circle cx="6" cy="5" r="2" data-id="oval1"></circle>
              <circle cx="2" cy="2" r="2" data-id="oval2"></circle>
            </g>
            <g fill="#9FC7FA" data-id="grp1" transform="translate(24)">
              <circle cx="2.5" cy="3" r="2" data-id="oval1"></circle>
              <circle cx="7.5" cy="2" r="2" data-id="oval2"></circle>
            </g>
          </svg>
        </button>
      </div>
    </div>
    <div class="AnnouncementPreviewDescription_container__HknfQ">
      <div class="AnnouncementPreviewBages_container__mD1TA">
        <div class="Badge_container__rAaAq Badge_gradient__wAXqY">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 16 16" class="CashbackBadge_icon__KmKJ1"
               style="--icon-gradient: url(#Cashback_svg__a);">
            <path
                d="M11.239 2.575H7.192c-.495 0-.97.197-1.32.545-.35.349-.546.821-.546 1.314v2.013H3.382a1.95 1.95 0 0 0-1.375.567 1.932 1.932 0 0 0 0 2.738c.365.363.86.567 1.375.567h1.945v1.935a1.928 1.928 0 0 0 .551 1.402 1.945 1.945 0 0 0 2.15.433 1.945 1.945 0 0 0 1.055-1.078c.094-.241.14-.498.133-.757V10.32h2.023c.494 0 .969-.196 1.319-.545.35-.348.546-.82.546-1.314V4.434c0-.493-.196-.965-.546-1.314a1.871 1.871 0 0 0-1.32-.545Z"
                style="fill: var(--icon-color, currentcolor);"></path>
            <defs>
              <linearGradient id="Cashback_svg__a" x1="2" x2="20.966" y1="21" y2="4.083" gradientUnits="userSpaceOnUse">
                <stop offset="0.063" stop-color="#9BD5FF"></stop>
                <stop offset="0.443" stop-color="#9B93FF"></stop>
                <stop offset="0.583" stop-color="#9A8DFF"></stop>
                <stop offset="0.849" stop-color="#7C86FE"></stop>
              </linearGradient>
            </defs>
          </svg>
          <span>{{ event.restrictions }}%</span>
        </div>
        <router-link :to="{ name: 'detail', params: { id: event.id } }">
          <div class="Badge_container__rAaAq Badge_secondary__1BKTK">
            от {{ event.price }} ₽
          </div>
        </router-link>
      </div>
      <router-link :to="{ name: 'detail', params: { id: event.id } }">
        <div class="AnnouncementPreviewDescription_titleContainer__Ngzmi">
          <h3 class="AnnouncementPreviewDescription_title__VoBCR">
            {{ event.name }}
          </h3>
        </div>
      </router-link>
      <time class="AnnouncementPreviewDescription_time__UtzSc">
        {{ event.date }}
      </time>
      <div
          :title="event.name"
          class="AnnouncementPreviewDescription_venue__ST9Py"
      >
        <a :href="addressLink">
          {{ event.address }}
        </a>
      </div>
    </div>
  </article>
</template>