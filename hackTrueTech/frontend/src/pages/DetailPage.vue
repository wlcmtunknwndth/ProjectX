<script setup>
import {onMounted} from "vue";

import {useRoute} from 'vue-router'
const route = useRoute()

const id = Number(route.params.id);

import {useEventsStore} from "../store/events.js";
const eventsStore = useEventsStore();

onMounted(() => {
  eventsStore.getDetailEvent(id);
});
</script>

<template>
  <main v-if="eventsStore.detailEvent"
        id="page"
        class="flexboxgrid_container__YXz7C CoreLayout_page__LDIK4 CoreLayout_pagePaddingBottom__vI8wg"
  >
    <div class="AnnouncementDetailsNoServiceFeeBlock_container__sUEEk">
      <div class="flexboxgrid_container__YXz7C"><h2 class="AnnouncementDetailsNoServiceFeeBlock_title__r6aaF">билеты
        без наценки</h2>
        <div class="AnnouncementDetailsNoServiceFeeBlock_subTitle__psDbH">Без комиссии, по цене организатора</div>
      </div>
      <div class="AnnouncementDetailsNoServiceFeeBlock_bg__XUIeW"></div>
    </div>
    <div class="Breadcrumbs_container__N9GSX _alias__breadcrumbs__8o2U_">
      <div class="Breadcrumbs_item__mhV8A"><a class="Breadcrumbs_link__0tBS2" title="Все события"
                                              href="/sankt-peterburg">Все события</a>
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" class="Breadcrumbs_separator__s7VZA">
          <path
              d="M9 5.012c0-.9 1.077-1.35 1.706-.714l3.096 3.13C15.934 9.582 17 10.66 17 12c0 1.34-1.066 2.417-3.198 4.573l-3.096 3.129c-.63.637-1.706.186-1.706-.715 0-.267.105-.524.293-.714l3.095-3.13c1.106-1.118 1.79-1.815 2.223-2.389.39-.517.39-.697.39-.751v-.006c0-.053 0-.234-.39-.751-.432-.574-1.117-1.272-2.223-2.39L9.293 5.727A1.016 1.016 0 0 1 9 5.012Z"></path>
        </svg>
      </div>
      <div class="Breadcrumbs_item__mhV8A">
        <router-link :to="{ name: 'main' }">
          <div class="Breadcrumbs_link__0tBS2" title="Концерты">
            Концерты
          </div>
        </router-link>
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" class="Breadcrumbs_separator__s7VZA">
          <path
              d="M9 5.012c0-.9 1.077-1.35 1.706-.714l3.096 3.13C15.934 9.582 17 10.66 17 12c0 1.34-1.066 2.417-3.198 4.573l-3.096 3.129c-.63.637-1.706.186-1.706-.715 0-.267.105-.524.293-.714l3.095-3.13c1.106-1.118 1.79-1.815 2.223-2.389.39-.517.39-.697.39-.751v-.006c0-.053 0-.234-.39-.751-.432-.574-1.117-1.272-2.223-2.39L9.293 5.727A1.016 1.016 0 0 1 9 5.012Z"></path>
        </svg>
      </div>
      <div class="Breadcrumbs_item__mhV8A">{{ eventsStore.detailEvent.name }}</div>
    </div>
    <div class="AnnouncementDetailsPageView_container__DKRhS">
      <section class="AnnouncementDetailsMediaBlock_container__Y6N0X AnnouncementDetailsPageView_media__022qT">
        <div class="AnnouncementDetailsSlider_container__OHMTq">
          <div
              class="swiper swiper-initialized swiper-horizontal AnnouncementDetailsSlider_slider__s5UZt swiper-backface-hidden">
            <div class="swiper-wrapper" style="cursor: grab; transform: translate3d(0px, 0px, 0px);">
              <div class="swiper-slide AnnouncementDetailsSlider_slide__dgc5O swiper-slide-active"
                   style="margin-right: 22px;">
                <div class="AnnouncementDetailsSlideSelector_container__ZU2W5"><span
                    class="LazyImage_container__ukc21 AnnouncementDetailsImageSlide_container__F2nXZ"><picture
                    class="LazyImage_originalContainer__Kj81p LazyImage_originalVisible__Xn1yN"><img loading="eager"
                                                                                                     class="LazyImage_img__Nz285"
                                                                                                     :src="eventsStore.detailEvent.img_path"
                                                                                                     style="transform: scale(1);"></picture></span>
                </div>
              </div>
            </div>
          </div>
          <div class="AnnouncementDetailsThumbnails_container__GuaPk AnnouncementDetailsThumbnails_hidden__Zf8Cn">
            <div
                class="swiper swiper-initialized swiper-horizontal swiper-free-mode swiper-watch-progress AnnouncementDetailsThumbnails_slider__Qog9H swiper-thumbs">
              <div class="swiper-wrapper" style="transition-duration: 0ms; transition-delay: 0ms;">
                <div class="swiper-slide">
                  <div
                      class="AnnouncementDetailsSlideSelector_container__ZU2W5 AnnouncementDetailsSlideSelector_overlay__H35Bg">
                        <span class="LazyImage_container__ukc21 AnnouncementDetailsImageSlide_container__F2nXZ"><picture
                            class="LazyImage_originalContainer__Kj81p LazyImage_originalVisible__Xn1yN"><img
                            loading="eager" class="LazyImage_img__Nz285"
                            :src="eventsStore.detailEvent.img_path"
                            style="transform: scale(1);"></picture></span></div>
                </div>
              </div>
            </div>
            <button type="button"
                    class="Button_button__znY7u Button_secondaryTheme___s8Ei Button_small__9c32H AnnouncementDetailsThumbnailsNavigation_prevButton__jhYDS">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                <path
                    d="M15 5.012c0-.9-1.077-1.35-1.706-.714l-3.096 3.13C8.066 9.582 7 10.66 7 12c0 1.34 1.066 2.417 3.198 4.573l3.096 3.129c.63.637 1.706.186 1.706-.715 0-.267-.105-.524-.293-.714l-3.095-3.13c-1.106-1.118-1.79-1.815-2.223-2.389-.39-.517-.39-.697-.39-.751v-.006c0-.053 0-.234.39-.751.432-.574 1.117-1.272 2.223-2.39l3.095-3.129c.188-.19.293-.447.293-.715Z"></path>
              </svg>
            </button>
            <button type="button"
                    class="Button_button__znY7u Button_secondaryTheme___s8Ei Button_small__9c32H AnnouncementDetailsThumbnailsNavigation_nextButton__KnrDi">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                <path
                    d="M9 5.012c0-.9 1.077-1.35 1.706-.714l3.096 3.13C15.934 9.582 17 10.66 17 12c0 1.34-1.066 2.417-3.198 4.573l-3.096 3.129c-.63.637-1.706.186-1.706-.715 0-.267.105-.524.293-.714l3.095-3.13c1.106-1.118 1.79-1.815 2.223-2.389.39-.517.39-.697.39-.751v-.006c0-.053 0-.234-.39-.751-.432-.574-1.117-1.272-2.223-2.39L9.293 5.727A1.016 1.016 0 0 1 9 5.012Z"></path>
              </svg>
            </button>
          </div>
        </div>
      </section>
      <aside class="AnnouncementDetailsAsideBlock_container__lWOtS AnnouncementDetailsPageView_aside__PAzGN"><h1
          class="Title_h__gKrut">{{ eventsStore.detailEvent.name }}</h1>
        <div class="AnnouncementDetailsTags_container___fE9G">
          <div class="Badge_container__rAaAq Badge_secondary__1BKTK AnnouncementDetailsTags_tag__Sob7B">
            {{ eventsStore.detailEvent.restrictions }}+
          </div>
          <div
              v-if="eventsStore.detailEvent.feature.includes('blind')"
              class="Badge_container__rAaAq Badge_secondary__1BKTK AnnouncementDetailsTags_tag__Sob7B"
          >
            <svg width="15" height="13" viewBox="0 0 15 13" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path fill-rule="evenodd" clip-rule="evenodd" d="M2.01872 0.443568C2.15686 0.440429 2.22593 0.459266 2.31698 0.515777C2.37977 0.556591 4.87884 3.03996 7.87082 6.03508C11.7984 9.96264 13.3242 11.5073 13.3587 11.5921C13.3838 11.658 13.4058 11.7647 13.4058 11.8275C13.4058 11.8966 13.3713 12.0159 13.3211 12.1164C13.2583 12.2357 13.1861 12.3173 13.0762 12.3926C12.9412 12.4837 12.8941 12.4994 12.7371 12.4962C12.6272 12.4962 12.5236 12.4774 12.4702 12.4429C12.4231 12.4146 11.9397 11.95 10.4076 10.4304L10.0308 10.5811C9.82361 10.6659 9.49082 10.7821 9.29303 10.8417C9.09524 10.9045 8.7907 10.983 8.61803 11.0175C8.44535 11.0521 8.18477 11.096 8.03721 11.1149C7.88965 11.1305 7.57256 11.1462 7.33082 11.1462C7.08907 11.1462 6.74372 11.1243 6.56163 11.0992C6.37954 11.074 6.05617 11.0081 5.83954 10.9547C5.62291 10.9014 5.31837 10.8103 5.16454 10.7538C5.0107 10.6973 4.74698 10.5905 4.58372 10.5152C4.42047 10.443 4.15047 10.308 3.98721 10.2201C3.82396 10.129 3.55396 9.9752 3.3907 9.8716C3.22744 9.77113 2.91663 9.5545 2.7 9.39125C2.48337 9.23113 2.21651 9.01764 2.10349 8.92346C1.99047 8.82613 1.74244 8.60322 1.55407 8.42741C1.36256 8.25159 1.05174 7.9345 0.857093 7.72101C0.662442 7.50752 0.420698 7.22497 0.317093 7.09625C0.213488 6.96752 0.0973256 6.80113 0.0627907 6.72578C0.0282558 6.65357 0 6.54055 0 6.47462C0 6.41183 0.0188372 6.30822 0.0439535 6.24857C0.0659303 6.18892 0.219768 5.97543 0.379884 5.77764C0.54 5.57985 0.919884 5.16857 1.22128 4.86403C1.52581 4.5595 1.9214 4.18903 2.10349 4.03834C2.28558 3.88764 2.60582 3.64276 3.20233 3.21892L2.32326 2.33985C1.73302 1.74962 1.43163 1.4231 1.39709 1.34148C1.37198 1.27868 1.35 1.17194 1.35 1.10601C1.35 1.04008 1.38768 0.914499 1.43477 0.817173C1.5007 0.691592 1.56349 0.619382 1.67965 0.544033C1.80837 0.462405 1.86802 0.443568 2.01558 0.443568H2.01872ZM3.03279 5.04613C2.80675 5.23136 2.38919 5.61752 2.10035 5.90322C1.81151 6.18892 1.57291 6.4338 1.57291 6.45264C1.57291 6.47148 1.73302 6.65357 1.92768 6.86078C2.12547 7.06799 2.45512 7.38822 2.65919 7.57345C2.86326 7.75555 3.16779 8.01613 3.33105 8.14799C3.4943 8.27985 3.8114 8.50904 4.03744 8.65973C4.26035 8.81043 4.65593 9.03962 4.91651 9.17148C5.17396 9.3002 5.5507 9.46032 5.74849 9.52625C5.94628 9.59218 6.22884 9.67066 6.3764 9.7052C6.52396 9.73659 6.80651 9.77741 7.00431 9.79625C7.2021 9.81508 7.47838 9.82136 7.61652 9.81194C7.75465 9.79939 7.98698 9.77113 8.13454 9.74601C8.2821 9.72404 8.50814 9.6738 8.63686 9.63927C8.76559 9.60473 8.97593 9.53566 9.10466 9.48857C9.23024 9.44148 9.34012 9.39753 9.3464 9.38811C9.35268 9.37869 9.11093 9.12439 8.80954 8.82299L8.26326 8.27357C7.74524 8.44938 7.66675 8.46194 7.36535 8.46194C7.09221 8.46194 6.97291 8.44625 6.78454 8.38659C6.65582 8.34578 6.49884 8.28927 6.43919 8.25787C6.37954 8.22648 6.23826 8.13229 6.12838 8.05066C6.01849 7.9659 5.86465 7.81206 5.78617 7.70846C5.70768 7.60485 5.60721 7.43531 5.55698 7.33171C5.50989 7.22811 5.4471 7.02404 5.41884 6.87648C5.37803 6.68811 5.36861 6.53113 5.38117 6.34276C5.39058 6.1795 5.42826 5.97857 5.48163 5.82473L5.56326 5.57357C4.50209 4.51555 4.185 4.20787 4.17558 4.20787C4.16616 4.20787 3.99663 4.3209 3.79884 4.4559C3.60105 4.59404 3.2557 4.85776 3.03279 5.04613ZM7.49093 1.80299C7.90849 1.80927 8.21931 1.83124 8.4171 1.86892C8.58035 1.90031 8.88489 1.96938 9.0921 2.02589C9.29931 2.08241 9.60384 2.18287 9.7671 2.24566C9.93035 2.31159 10.2192 2.43717 10.4107 2.53136C10.6022 2.62241 10.9413 2.80764 11.1642 2.94264C11.3902 3.07764 11.7293 3.30055 11.9177 3.43869C12.1092 3.57683 12.3918 3.79345 12.5456 3.92531C12.6994 4.05403 12.9914 4.30834 13.1892 4.49357C13.387 4.67566 13.7072 5.00217 13.9019 5.21566C14.0965 5.42915 14.3383 5.71171 14.4419 5.84043C14.5455 5.96915 14.6616 6.13555 14.6962 6.20776C14.7307 6.28311 14.759 6.39927 14.759 6.46834C14.759 6.54055 14.7244 6.66299 14.6773 6.75718C14.6334 6.84822 14.5015 7.04287 14.3885 7.19043C14.2755 7.33799 13.8924 7.74927 13.5377 8.10404C13.1829 8.4588 12.8407 8.77904 12.7811 8.81043C12.7026 8.85125 12.6052 8.87008 12.4357 8.87008C12.2662 8.87008 12.1594 8.85125 12.059 8.80415C11.9805 8.76962 11.8737 8.68485 11.8235 8.61892C11.7481 8.5216 11.7262 8.46194 11.7167 8.30811C11.7073 8.15113 11.7199 8.08834 11.7858 7.95962C11.8423 7.84659 11.9805 7.6959 12.2787 7.4259C12.5048 7.21869 12.803 6.92671 12.9412 6.77287C13.0762 6.62217 13.1892 6.48717 13.1892 6.46834C13.1892 6.4495 13.0354 6.27055 12.847 6.06648C12.6618 5.86555 12.3321 5.54217 12.1186 5.35066C11.9051 5.15915 11.588 4.89229 11.4154 4.75729C11.2427 4.62229 10.9601 4.41822 10.7874 4.30206C10.6148 4.18903 10.3322 4.01322 10.1595 3.9159C9.98687 3.82171 9.72628 3.68985 9.57872 3.62706C9.43117 3.56113 9.17059 3.46066 8.99791 3.40415C8.82524 3.3445 8.55524 3.26915 8.4014 3.23462C8.24756 3.20322 7.94931 3.15927 7.7421 3.13729C7.53489 3.11845 7.27431 3.11217 7.16128 3.12159C7.04826 3.13415 6.89442 3.15613 6.81593 3.17496C6.71233 3.20008 6.64012 3.20008 6.53965 3.17496C6.46744 3.15613 6.35756 3.09334 6.29477 3.03682C6.23198 2.98345 6.15349 2.88299 6.1221 2.81706C6.08756 2.75427 6.05303 2.64124 6.04361 2.56589C6.03105 2.48741 6.04047 2.3838 6.06558 2.30845C6.08756 2.23938 6.14407 2.14206 6.18803 2.09182C6.23198 2.04159 6.31989 1.97566 6.38268 1.94427C6.44861 1.90973 6.58361 1.86264 6.68093 1.83752C6.81907 1.80299 7.01058 1.79357 7.49093 1.80299Z" fill="#1D2023"/>
            </svg>
          </div>
          <div
              v-if="eventsStore.detailEvent.feature.includes('disability')"
              class="Badge_container__rAaAq Badge_secondary__1BKTK AnnouncementDetailsTags_tag__Sob7B"
          >
            <svg width="20" height="20" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path fill-rule="evenodd" clip-rule="evenodd" d="M8.86864 3.09494C8.1613 3.28517 7.6899 3.98506 7.79813 4.68431C7.92307 5.49161 8.69236 6.04601 9.47663 5.89392C10.5041 5.69467 10.9806 4.51409 10.3793 3.65697C10.0419 3.17594 9.42324 2.94579 8.86864 3.09494ZM8.91585 6.84175C8.63325 6.92114 8.36803 7.18389 8.29118 7.46068C8.27105 7.53304 8.26043 8.62513 8.26076 10.5835C8.26121 13.4469 8.264 13.601 8.31712 13.7161C8.38473 13.8626 8.48056 13.9549 8.63014 14.0173C8.72723 14.0579 8.99427 14.0642 10.6269 14.0642H12.5115L13.9496 15.5005C15.5505 17.0995 15.5167 17.0722 15.9023 17.0722C16.2521 17.0722 16.5488 16.8927 16.7053 16.5863C16.7685 16.4626 16.7799 16.3983 16.7795 16.1677C16.7791 15.9418 16.7669 15.872 16.7082 15.7591C16.6603 15.667 16.0925 15.0768 14.9539 13.9355L13.2703 12.2482H11.6731H10.0759V11.326V10.4038L11.8238 10.3953L13.5717 10.3868L13.7125 10.304C14.3737 9.91531 14.3383 9.00313 13.6494 8.68006L13.4809 8.60104L11.7813 8.59208L10.0818 8.58312L10.0705 8.04728C10.0607 7.57917 10.0511 7.49355 9.99462 7.36991C9.8027 6.94984 9.33916 6.72289 8.91585 6.84175ZM6.95511 9.07883C5.57644 9.78017 4.62258 11.0259 4.28868 12.561C4.23129 12.8248 4.22158 12.9635 4.22055 13.5345C4.21946 14.1202 4.22778 14.24 4.28977 14.5296C4.75612 16.7085 6.5062 18.2885 8.7045 18.5153C10.1909 18.6687 11.6655 18.1345 12.7321 17.0563C13.0034 16.782 13.3748 16.3208 13.3749 16.258C13.375 16.2029 12.7416 15.5933 12.7161 15.6239C12.7027 15.64 12.6242 15.7552 12.5416 15.8799C12.3519 16.1661 11.826 16.6947 11.5274 16.8992C10.5298 17.5824 9.30878 17.7987 8.14974 17.4976C6.60033 17.095 5.40585 15.7713 5.15334 14.1769C5.07719 13.6961 5.10824 13.0039 5.22628 12.5508C5.5093 11.4644 6.145 10.6156 7.11804 10.0248L7.35191 9.88283V9.40083C7.35191 9.01457 7.34374 8.91884 7.31075 8.91884C7.28811 8.91884 7.12806 8.99081 6.95511 9.07883Z" fill="#1D2023"/>
            </svg>
          </div>
          <div
              v-if="eventsStore.detailEvent.feature.includes('deaf')"
              class="Badge_container__rAaAq Badge_secondary__1BKTK AnnouncementDetailsTags_tag__Sob7B"
          >
            <svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M0.464 0H0.88775L1.11375 0.14125L1.255 0.25425L1.45275 0.48025L14.1935 13.221L14.25 13.3623V13.7578L14.137 13.9838L13.9675 14.0968L13.8828 14.125H13.4873L13.2613 13.9838L9.53225 10.2548L9.391 10.2265L9.3345 10.3678L9.08025 10.735L8.8825 11.074L8.6 11.7238L8.3175 12.5713L8.148 12.882H8.0915L8.035 12.995L7.83725 13.2493L7.61125 13.4753L7.21575 13.7578L6.76375 13.9838L6.5095 14.0685L6.114 14.125H5.323L4.871 14.0403L4.419 13.8425L4.0235 13.5883L3.741 13.334L3.43025 12.9668L3.2325 12.6843L3.03475 12.2323L2.8935 11.6955V10.9893L3.0065 10.622L3.176 10.3113L3.402 10.0288L3.628 9.831L3.88225 9.6615L4.193 9.5485L4.89925 9.46375L5.0405 9.35075L5.1535 9.153V8.84225L5.01225 8.61625L4.871 8.475L4.645 8.22075L4.532 8.10775L4.306 7.71225L4.16475 7.31675L4.10825 7.03425L4.08 4.8025L3.26075 3.98325L3.20425 4.068L3.09125 4.54825L3.0065 5.28275L2.92175 5.45225L2.7805 5.5935L2.611 5.67825H2.272L2.046 5.56525L1.90475 5.3675L1.8765 5.28275V4.66125L1.9895 3.98325L2.13075 3.53125L2.30025 3.07925L2.24375 2.96625L0.1815 0.904L0.125 0.7345V0.339L0.26625 0.113L0.464 0ZM5.23825 5.96075V6.667L5.29475 7.03425L5.4925 7.42975L5.69025 7.65575L5.91625 7.8535L6.1705 8.249L6.2835 8.5315L6.31175 8.67275V9.3225L6.19875 9.6615L6.02925 9.97225L5.85975 10.1983L5.80325 10.2265V10.283L5.69025 10.3395L5.40775 10.509L4.95575 10.622L4.50375 10.6785L4.33425 10.7633L4.193 10.8763L4.05175 11.1305V11.5825L4.1365 11.8368L4.33425 12.2323L4.5885 12.543L4.871 12.769L5.23825 12.9385L5.46425 12.995L5.74675 13.0233L6.14225 12.9668L6.48125 12.8255L6.76375 12.656L6.9615 12.4865L7.21575 12.1475L7.38525 11.6955L7.583 11.1023L7.78075 10.6503L8.17625 9.97225L8.51525 9.492L8.62825 9.3225L8.45875 9.18125L8.261 8.95525L5.2665 5.96075H5.23825Z" fill="black"/>
              <path d="M6.53775 0H7.6395L8.17625 0.0565L8.79775 0.226L9.3345 0.452L9.843 0.7345L10.3233 1.0735L10.6905 1.38425L11.0013 1.695L11.312 2.0905L11.5098 2.373L11.5945 2.51425V2.57075H11.651L11.8488 2.96625L12.0183 3.41825L12.1878 3.98325L12.3008 4.66125V5.73475L12.216 6.2715L12.0465 6.86475L11.764 7.571L11.5098 8.023L11.3968 8.1925H11.3403L11.2838 8.3055L11.0578 8.4185H10.6905L10.4363 8.249V8.1925H10.3798L10.3233 8.05125V7.65575L10.634 7.14725L10.747 6.9495L10.9165 6.4975L11.086 5.84775L11.1143 5.65V4.746L11.0013 4.181L10.8318 3.6725L10.634 3.2205L10.3515 2.79675L10.0973 2.486L9.7865 2.17525L9.617 2.06225L9.47575 1.921L8.96725 1.61025L8.40225 1.356L7.89375 1.21475L7.55475 1.15825H6.65075L6.0575 1.27125L5.46425 1.49725L4.984 1.7515L4.78625 1.89275L4.7015 1.921H4.33425L4.10825 1.808L3.967 1.61025L3.93875 1.5255V1.15825L4.10825 0.904L4.61675 0.59325L5.06875 0.36725L5.52075 0.19775L6.02925 0.0565L6.53775 0Z" fill="black"/>
              <path d="M7.10275 2.17525L7.696 2.23175L8.261 2.40125L8.713 2.62725L9.052 2.8815L9.4475 3.277L9.617 3.5595L9.75825 3.7855L9.8995 4.09625L10.0408 4.5765L10.069 4.77425V5.5935L9.956 6.07375L9.81475 6.41275L9.70175 6.554L9.47575 6.667H9.13675L8.96725 6.58225L8.826 6.441L8.74125 6.2715V5.9325L8.85425 5.5935L8.8825 5.45225V4.859L8.7695 4.49175L8.57175 4.1245L8.3175 3.81375L7.95025 3.5595L7.696 3.47475L7.4135 3.39L6.98975 3.36175L6.53775 3.4465L6.25525 3.53125H5.91625L5.7185 3.4465L5.549 3.24875L5.52075 3.164V2.7685L5.69025 2.51425L5.8315 2.4295L6.36825 2.26L6.65075 2.2035L7.10275 2.17525Z" fill="black"/>
            </svg>
          </div>
          <div
              v-if="eventsStore.detailEvent.feature.includes('neuro')"
              class="Badge_container__rAaAq Badge_secondary__1BKTK AnnouncementDetailsTags_tag__Sob7B"
          >
            <svg width="14" height="16" viewBox="0 0 14 16" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path fill-rule="evenodd" clip-rule="evenodd" d="M5.37728 0.0371673C4.27795 0.180677 3.32504 0.561652 2.46693 1.20076C0.547436 2.63042 -0.356981 5.01631 0.129859 7.36613C0.353206 8.44424 0.990319 9.64222 1.71956 10.3553L1.87821 10.5105V13.2552V16H5.62722H9.37622V14.593V13.1861H10.7977H12.2192V12.0449V10.9037H13.1096C13.5993 10.9037 14 10.8961 14 10.8868C14 10.8776 13.5916 9.90376 13.0924 8.72272L12.1849 6.57542L12.2049 6.33989C12.2281 6.06741 12.1781 5.37281 12.1092 5.00987C11.7955 3.35854 10.8653 1.94898 9.46957 1.00966C8.97039 0.67374 8.21338 0.337319 7.6096 0.183147C6.94965 0.0145934 6.01727 -0.046375 5.37728 0.0371673ZM5.62722 0.950943C3.22801 1.19391 1.33482 3.00032 0.984602 5.38075C0.945425 5.64707 0.937395 5.85811 0.950517 6.27633C0.976729 7.11116 1.11625 7.66606 1.489 8.41804C1.75168 8.94796 2.03111 9.33563 2.46399 9.77075L2.81546 10.124V12.593V15.062H5.62722H8.43897V13.6551V12.2481H9.86047H11.282V11.1069V9.96569H11.9414C12.4566 9.96569 12.5973 9.95716 12.5848 9.92661C12.576 9.90513 12.2707 9.18405 11.9064 8.32424L11.2439 6.76095L11.2638 6.41699C11.3572 4.79664 10.6671 3.19333 9.42812 2.15233C8.6368 1.48743 7.67161 1.07397 6.6516 0.962886C6.36649 0.93187 5.87268 0.926086 5.62722 0.950943ZM5.65846 1.88714C4.71483 2.00232 3.849 2.40831 3.18886 3.04516C2.57212 3.64018 2.20784 4.26428 1.99646 5.08801C1.87615 5.55684 1.84397 6.23124 1.92092 6.67184C2.20725 8.31186 3.68798 9.48814 5.34601 9.39269C6.96089 9.2997 8.24459 8.07752 8.43903 6.44785C8.58443 5.2292 7.85922 4.13878 6.68944 3.81714C6.41713 3.74226 5.78427 3.74148 5.51478 3.81567C4.8314 4.0038 4.2423 4.49999 3.95394 5.1304C3.75518 5.56487 3.69557 6.14954 3.81192 6.52361C4.12208 7.52105 5.37085 7.85422 6.12209 7.13998C6.37449 6.90002 6.53185 6.59402 6.57834 6.25288L6.59646 6.12L6.13527 6.12012L5.67408 6.12022L5.59597 6.28311C5.50128 6.48056 5.39918 6.55941 5.21058 6.58067C5.02819 6.60124 4.81668 6.48718 4.74089 6.32738C4.67306 6.18434 4.6765 5.9188 4.74917 5.68512C5.0718 4.64769 6.37714 4.34475 7.1232 5.13422C7.53234 5.56715 7.62491 6.16249 7.39166 6.86087C7.14431 7.60153 6.52298 8.18164 5.75218 8.39152C5.498 8.46075 4.87626 8.4685 4.62973 8.40553C3.54592 8.12873 2.81546 7.20186 2.81546 6.10346C2.81546 4.20872 4.39901 2.70755 6.28326 2.81604C7.8939 2.90878 9.19474 4.14581 9.37407 5.75525C9.39244 5.92012 9.40747 6.07 9.40747 6.08833C9.40747 6.11084 9.55743 6.11884 9.86828 6.11303L10.3291 6.10436L10.3199 5.86987C10.3066 5.53048 10.2312 5.11643 10.1238 4.79242C9.63756 3.32618 8.41167 2.24322 6.8996 1.94423C6.65141 1.89514 5.88262 1.85978 5.65846 1.88714Z" fill="black"/>
            </svg>
          </div>
        </div>
        <div
            class="AnnouncementDetailsBuyButton_container__69NJq AnnouncementDetailsBuyButton_visible__bPTD1 _alias__buyButton__t_p6x">
          <div class="AnnouncementDetailsBuyButton_button__aRZkX">
            <button type="button" class="Button_button__znY7u Button_primaryTheme__xtq5p Button_medium__UJyor">
              <span>Купить билет</span></button>
          </div>
          <div class="AnnouncementDetailsBuyButton_dynamicButton__8DWTn">
            <button type="button" class="Button_button__znY7u Button_primaryTheme__xtq5p Button_medium__UJyor">
              <span>Купить билет</span></button>
          </div>
        </div>
        <div class="AnnouncementDetailsSharePanel_container__9mTlu">
          <button type="button"
                  class="Button_button__znY7u Button_secondaryTheme___s8Ei Button_medium__UJyor AnnouncementDetailsSharePanel_shareButton__hRnmu">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 25 24"
                 class="AnnouncementDetailsSharePanel_shareIcon__SpiK_">
              <path
                  d="M16.01 8a1 1 0 0 0 .858-1.514l-.164-.274C15.019 3.404 14.176 2 12.93 2c-1.246 0-2.088 1.404-3.773 4.212l-.164.274a1 1 0 0 0 1.715 1.029l.164-.274c.42-.699.763-1.269 1.057-1.734V14a1 1 0 1 0 2 0V5.504c.295.466.639 1.037 1.059 1.737l.164.274A1 1 0 0 0 16.01 8Z"></path>
              <path fill="currentColor"
                    d="M7.947 10.117c.447-.044.941-.082 1.484-.117a1 1 0 1 1 0 2v.002c-.47.031-.9.067-1.29.105-.407.04-.54.057-.653.085a1.602 1.602 0 0 0-.4.17l-.002.001c-.012.008-.046.032-.104.079a4.914 4.914 0 0 0-.603.581c-.05.056-.074.09-.083.1l-.002.003c-.152.258-.182.35-.19.376-.013.044-.024.09-.066.4-.066.48-.107 1.002-.107 1.552s.041 1.07.106 1.546c.042.306.052.349.065.389.007.024.037.116.19.375l.002.003c.008.01.033.043.082.1.077.088.178.194.291.303.113.108.223.205.314.279.058.047.092.07.103.079l.002.001c.228.123.332.154.398.17.112.027.242.044.646.083 1.213.12 2.815.218 4.8.218 1.983 0 3.583-.098 4.795-.217.406-.04.537-.056.65-.084.068-.016.172-.047.4-.17l.002-.001c.011-.008.045-.032.103-.08a4.97 4.97 0 0 0 .604-.58c.049-.057.074-.09.082-.101l.002-.003c.153-.258.183-.351.19-.376.014-.042.024-.086.066-.395.065-.477.107-.998.107-1.548 0-.549-.042-1.07-.107-1.547-.042-.309-.052-.353-.065-.395-.008-.025-.038-.118-.19-.376l-.003-.003c-.008-.01-.033-.044-.082-.1a4.92 4.92 0 0 0-.29-.303 4.942 4.942 0 0 0-.314-.279 1.859 1.859 0 0 0-.103-.079l-.003-.002a1.606 1.606 0 0 0-.399-.17c-.113-.027-.244-.043-.65-.083a41.36 41.36 0 0 0-1.294-.106V12a1 1 0 1 1 0-2c.544.035 1.04.073 1.489.117.745.073 1.118.11 1.832.5.43.234 1.274 1.047 1.524 1.467.415.7.454.982.53 1.545.075.556.125 1.167.125 1.816 0 .65-.05 1.261-.125 1.817-.076.563-.115.845-.53 1.544-.25.42-1.094 1.234-1.524 1.468-.714.39-1.087.426-1.832.5a51.29 51.29 0 0 1-4.99.226c-2.051 0-3.719-.102-4.996-.227-.741-.073-1.112-.11-1.826-.5-.428-.233-1.273-1.047-1.523-1.466-.416-.7-.454-.979-.53-1.537a13.498 13.498 0 0 1-.124-1.816c0-.65.05-1.263.125-1.82.077-.567.115-.85.531-1.55.25-.42 1.094-1.233 1.524-1.468.714-.39 1.088-.426 1.836-.5Z"></path>
            </svg>
          </button>
          <div class="FavouritesItemStatusButton_container__KIJN8 _alias__favouritesItemStatusButton__wRJfb">
            <button class="FavouritesItemStatusButton_button__VRYUO">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 60 60"
                   class="FavouritesItemStatusButton_icon__maJuJ">
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
                  <use clip-path="url(#FavouritesDynamic_svg__c)" data-id="heart"
                       href="#FavouritesDynamic_svg__a"></use>
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
          <div class="FavouritesItemStatusMessage_container__0pBEc _alias__favouritesItemStatusMessage__4bNqE">12
            человек добавили<br>в избранное
          </div>
        </div>
        <section class="CashbackDetailsBlock_container__Vzg6_ CashbackDetailsBlock_guest__xkKkU">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 16 16"
               class="CashbackDetailsInteractive_cashbackIcon__BLJYj"
               style="--icon-gradient: url(#Cashback_svg__a);">
            <path
                d="M11.239 2.575H7.192c-.495 0-.97.197-1.32.545-.35.349-.546.821-.546 1.314v2.013H3.382a1.95 1.95 0 0 0-1.375.567 1.932 1.932 0 0 0 0 2.738c.365.363.86.567 1.375.567h1.945v1.935a1.928 1.928 0 0 0 .551 1.402 1.945 1.945 0 0 0 2.15.433 1.945 1.945 0 0 0 1.055-1.078c.094-.241.14-.498.133-.757V10.32h2.023c.494 0 .969-.196 1.319-.545.35-.348.546-.82.546-1.314V4.434c0-.493-.196-.965-.546-1.314a1.871 1.871 0 0 0-1.32-.545Z"
                style="fill: var(--icon-color, currentcolor);"></path>
            <defs>
              <linearGradient id="Cashback_svg__a" x1="2" x2="20.966" y1="21" y2="4.083"
                              gradientUnits="userSpaceOnUse">
                <stop offset="0.063" stop-color="#9BD5FF"></stop>
                <stop offset="0.443" stop-color="#9B93FF"></stop>
                <stop offset="0.583" stop-color="#9A8DFF"></stop>
                <stop offset="0.849" stop-color="#7C86FE"></stop>
              </linearGradient>
            </defs>
          </svg>
          <div class="CashbackDetailsInteractive_text__Rj0KS">Кешбэк 20%</div>
          <a class="CashbackDetailsInteractive_loginLink__R0vEZ"
             href="/sankt-peterburg/announcements/white-gallows-pervyy-bolshoy-koncert?eventId=26611579#">Войти
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"
                 class="CashbackDetailsInteractive_loginIcon__DzBL_">
              <path
                  d="M9 5.012c0-.9 1.077-1.35 1.706-.714l3.096 3.13C15.934 9.582 17 10.66 17 12c0 1.34-1.066 2.417-3.198 4.573l-3.096 3.129c-.63.637-1.706.186-1.706-.715 0-.267.105-.524.293-.714l3.095-3.13c1.106-1.118 1.79-1.815 2.223-2.389.39-.517.39-.697.39-.751v-.006c0-.053 0-.234-.39-.751-.432-.574-1.117-1.272-2.223-2.39L9.293 5.727A1.016 1.016 0 0 1 9 5.012Z"></path>
            </svg>
          </a>
          <div class="CashbackDetailsBlock_description__q8j7k">Войдите со своим номером, чтобы получать до&nbsp;20%&nbsp;кешбэка
            за каждую покупку
          </div>
        </section>
      </aside>
      <section><h2 class="SimpleAnnouncementSchedule_title__5UNI0">Расписание</h2>
        <div>
          <div class="SimpleAnnouncementScheduleSlider_container__LNOqK">
            <div
                class="swiper swiper-initialized swiper-horizontal swiper-free-mode SimpleAnnouncementScheduleSlider_slider__3MdZt swiper-backface-hidden">
              <div class="swiper-wrapper" style="transform: translate3d(0px, 0px, 0px);">
                <div class="swiper-slide swiper-slide-active" style="width: 218.4px; margin-right: 12px;">
                  <div
                      class="Tag_container__aLkPQ Tag_cursorPointer__QGOOJ Tag_active__sxwVd SimpleAnnouncementScheduleSlide_container__Skk5Y"
                      :title="eventsStore.detailEvent.date">
                    <span class="Tag_children__r1Gec"><span
                      class="SimpleAnnouncementScheduleSlide_date__Qla6M">{{ eventsStore.detailEvent.date }}</span>
                  </span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>
    </div>
    <section class="ReviewsBlock_container__QtjkG"></section>
    <div
        class="swiper swiper-initialized swiper-horizontal CommonSlider_container__hjFVD _alias__slider__xLxq9 swiper-backface-hidden">
      <header class="CommonSliderHeader_headerContaner__yFTv3"><h2 class="CommonSliderHeader_title__fveVH"><a
          class="Button_button__znY7u Button_linkTheme__0OykU Button_small__9c32H" href="/sankt-peterburg/concerts">Концерты
        в Санкт-Петербурге</a></h2>
        <div class="CommonSliderHeader_headerRightContainer__ORTg6">
          <div class="CommonSliderNavigation_container__YyhWk">
            <button type="button" class="Button_button__znY7u Button_secondaryTheme___s8Ei Button_small__9c32H"
                    disabled="">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                <path
                    d="M15 5.012c0-.9-1.077-1.35-1.706-.714l-3.096 3.13C8.066 9.582 7 10.66 7 12c0 1.34 1.066 2.417 3.198 4.573l3.096 3.129c.63.637 1.706.186 1.706-.715 0-.267-.105-.524-.293-.714l-3.095-3.13c-1.106-1.118-1.79-1.815-2.223-2.389-.39-.517-.39-.697-.39-.751v-.006c0-.053 0-.234.39-.751.432-.574 1.117-1.272 2.223-2.39l3.095-3.129c.188-.19.293-.447.293-.715Z"></path>
              </svg>
            </button>
            <button type="button" class="Button_button__znY7u Button_secondaryTheme___s8Ei Button_small__9c32H">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                <path
                    d="M9 5.012c0-.9 1.077-1.35 1.706-.714l3.096 3.13C15.934 9.582 17 10.66 17 12c0 1.34-1.066 2.417-3.198 4.573l-3.096 3.129c-.63.637-1.706.186-1.706-.715 0-.267.105-.524.293-.714l3.095-3.13c1.106-1.118 1.79-1.815 2.223-2.389.39-.517.39-.697.39-.751v-.006c0-.053 0-.234-.39-.751-.432-.574-1.117-1.272-2.223-2.39L9.293 5.727A1.016 1.016 0 0 1 9 5.012Z"></path>
              </svg>
            </button>
          </div>
          <a
              class="Button_button__znY7u Button_secondaryTheme___s8Ei Button_small__9c32H CommonSliderHeader_showMoreButton__gm4NI"
              href="/sankt-peterburg/concerts">Показать еще</a></div>
      </header>
      <div class="swiper-wrapper" style="cursor: grab;">
        <div class="swiper-slide CommonSlider_slide__WGyyy swiper-slide-active"
             style="width: 268.5px; margin-right: 22px;">
          <article class="AnnouncementPreview_container__A430m">
            <div class="AnnouncementPreviewImage_container__Pmp0c"><a
                title="Концерт группы «Три дня дождя». Summer Sound x билайн"
                class="AnnouncementPreviewImage_link__5btQS" data-type="izobrazhenie"
                href="/sankt-peterburg/announcements/tri-dnya-dozhdya-summer-sound-x-bilayn?eventId=18736296"><span
                class="LazyImage_container__ukc21"><picture class="LazyImage_backgroundContainer__UuhKK"><img
                loading="lazy" class="LazyImage_img__Nz285"
                alt="Концерт группы «Три дня дождя». Summer Sound x билайн"
                src="https://live.mts.ru/image/536x360/tri-dnya-dozhdya-summer-sound-x-bilayn-2030c180-8aef-8812-fdd1-10211d05850c.jpg"></picture><picture
                class="LazyImage_originalContainer__Kj81p LazyImage_originalVisible__Xn1yN LazyImage_originalHoverAnimation__jNykx"><img
                loading="eager" class="LazyImage_img__Nz285 LazyImage_original___06Jx"
                src="https://live.mts.ru/image/536x360/tri-dnya-dozhdya-summer-sound-x-bilayn-2030c180-8aef-8812-fdd1-10211d05850c.jpg"
                style="transform: scale(1);"></picture></span></a>
              <div class="FavouritesItemStatusButton_container__KIJN8">
                <button class="FavouritesItemStatusButton_button__VRYUO">
                  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 60 60"
                       class="FavouritesItemStatusButton_icon__maJuJ">
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
                      <use clip-path="url(#FavouritesDynamic_svg__c)" data-id="heart"
                           href="#FavouritesDynamic_svg__a"></use>
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
              <div class="AnnouncementPreviewBages_container__mD1TA"><a
                  title="Концерт группы «Три дня дождя». Summer Sound x билайн" data-type="cena"
                  href="/sankt-peterburg/announcements/tri-dnya-dozhdya-summer-sound-x-bilayn?eventId=18736296">
                <div class="Badge_container__rAaAq Badge_secondary__1BKTK">от 2000 ₽</div>
              </a></div>
              <a title="Концерт группы «Три дня дождя». Summer Sound x билайн"
                 class="AnnouncementPreviewDescription_titleContainer__Ngzmi" data-type="nazvanie_meropriyatiya"
                 href="/sankt-peterburg/announcements/tri-dnya-dozhdya-summer-sound-x-bilayn?eventId=18736296"><h3
                  class="AnnouncementPreviewDescription_title__VoBCR">Концерт группы «Три дня дождя». Summer Sound x
                билайн</h3></a>
              <time class="AnnouncementPreviewDescription_time__UtzSc">31 июля, 19:30</time>
              <div title="Двор Гостинки" class="AnnouncementPreviewDescription_venue__ST9Py"><a
                  href="/sankt-peterburg/venues/drugoe/dvor-gostinki">Двор Гостинки</a></div>
            </div>
          </article>
        </div>
        <div class="swiper-slide CommonSlider_slide__WGyyy swiper-slide-next"
             style="width: 268.5px; margin-right: 22px;">
          <article class="AnnouncementPreview_container__A430m">
            <div class="AnnouncementPreviewImage_container__Pmp0c"><a title="Концерт Дениса Мацуева"
                                                                      class="AnnouncementPreviewImage_link__5btQS"
                                                                      data-type="izobrazhenie"
                                                                      href="/sankt-peterburg/announcements/denis-macuev-new?eventId=18642795"><span
                class="LazyImage_container__ukc21"><picture class="LazyImage_backgroundContainer__UuhKK"><img
                loading="lazy" class="LazyImage_img__Nz285" alt="Концерт Дениса Мацуева"
                src="https://live.mts.ru/image/536x360/denis-macuev-new-69d79afc-8dab-41f3-b45f-f7d7b311d26d.jpg"></picture><picture
                class="LazyImage_originalContainer__Kj81p LazyImage_originalVisible__Xn1yN LazyImage_originalHoverAnimation__jNykx"><img
                loading="eager" class="LazyImage_img__Nz285 LazyImage_original___06Jx"
                src="https://live.mts.ru/image/536x360/denis-macuev-new-69d79afc-8dab-41f3-b45f-f7d7b311d26d.jpg"
                style="transform: scale(1.26);"></picture></span></a>
              <div class="FavouritesItemStatusButton_container__KIJN8">
                <button class="FavouritesItemStatusButton_button__VRYUO">
                  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 60 60"
                       class="FavouritesItemStatusButton_icon__maJuJ">
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
                      <use clip-path="url(#FavouritesDynamic_svg__c)" data-id="heart"
                           href="#FavouritesDynamic_svg__a"></use>
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
              <div class="AnnouncementPreviewBages_container__mD1TA"><a title="Концерт Дениса Мацуева"
                                                                        data-type="cena"
                                                                        href="/sankt-peterburg/announcements/denis-macuev-new?eventId=18642795">
                <div class="Badge_container__rAaAq Badge_secondary__1BKTK">от 2000 ₽</div>
              </a></div>
              <a title="Концерт Дениса Мацуева" class="AnnouncementPreviewDescription_titleContainer__Ngzmi"
                 data-type="nazvanie_meropriyatiya"
                 href="/sankt-peterburg/announcements/denis-macuev-new?eventId=18642795"><h3
                  class="AnnouncementPreviewDescription_title__VoBCR">Концерт Дениса Мацуева</h3></a>
              <time class="AnnouncementPreviewDescription_time__UtzSc">15 мая, 19:00</time>
              <div title="БКЗ «Октябрьский»" class="AnnouncementPreviewDescription_venue__ST9Py"><a
                  href="/sankt-peterburg/venues/concert-hall/bkz-oktyabrskiy">БКЗ «Октябрьский»</a></div>
            </div>
          </article>
        </div>
        <div class="swiper-slide CommonSlider_slide__WGyyy" style="width: 268.5px; margin-right: 22px;">
          <article class="AnnouncementPreview_container__A430m">
            <div class="AnnouncementPreviewImage_container__Pmp0c"><a :title="eventsStore.detailEvent.name"
                                                                      class="AnnouncementPreviewImage_link__5btQS"
                                                                      data-type="izobrazhenie"
                                                                      href="/sankt-peterburg/announcements/white-gallows-pervyy-bolshoy-koncert?eventId=18523565"><span
                class="LazyImage_container__ukc21"><picture class="LazyImage_backgroundContainer__UuhKK"><img
                loading="lazy" class="LazyImage_img__Nz285" :alt="eventsStore.detailEvent.name"
                src="https://live.mts.ru/image/536x360/white-gallows-pervyy-bolshoy-koncert-d497c910-1393-4cf5-b074-f1a7d9e11cf4.jpg"></picture><picture
                class="LazyImage_originalContainer__Kj81p LazyImage_originalVisible__Xn1yN LazyImage_originalHoverAnimation__jNykx"><img
                loading="eager" class="LazyImage_img__Nz285 LazyImage_original___06Jx"
                src="https://live.mts.ru/image/536x360/white-gallows-pervyy-bolshoy-koncert-d497c910-1393-4cf5-b074-f1a7d9e11cf4.jpg"
                style="transform: scale(1);"></picture></span></a>
              <div class="FavouritesItemStatusButton_container__KIJN8">
                <button class="FavouritesItemStatusButton_button__VRYUO">
                  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 60 60"
                       class="FavouritesItemStatusButton_icon__maJuJ">
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
                      <use clip-path="url(#FavouritesDynamic_svg__c)" data-id="heart"
                           href="#FavouritesDynamic_svg__a"></use>
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
              <div class="AnnouncementPreviewBages_container__mD1TA"><a :title="eventsStore.detailEvent.name"
                                                                        data-type="cena"
                                                                        href="/sankt-peterburg/announcements/white-gallows-pervyy-bolshoy-koncert?eventId=18523565">
                <div class="Badge_container__rAaAq Badge_secondary__1BKTK">от 1000 ₽</div>
              </a></div>
              <a :title="eventsStore.detailEvent.name"
                 class="AnnouncementPreviewDescription_titleContainer__Ngzmi" data-type="nazvanie_meropriyatiya"
                 href="/sankt-peterburg/announcements/white-gallows-pervyy-bolshoy-koncert?eventId=18523565"><h3
                  class="AnnouncementPreviewDescription_title__VoBCR">{{ eventsStore.detailEvent.name }}</h3></a>
              <time class="AnnouncementPreviewDescription_time__UtzSc">08 сентября, 20:00</time>
              <div title="Клуб «Космонавт»" class="AnnouncementPreviewDescription_venue__ST9Py"><a
                  href="/sankt-peterburg/venues/club/club-kosmonavt">Клуб «Космонавт»</a></div>
            </div>
          </article>
        </div>
        <div class="swiper-slide CommonSlider_slide__WGyyy" style="width: 268.5px; margin-right: 22px;">
          <article class="AnnouncementPreview_container__A430m">
            <div class="AnnouncementPreviewImage_container__Pmp0c"><a title="Концерт группы «Папин Олимпос»"
                                                                      class="AnnouncementPreviewImage_link__5btQS"
                                                                      data-type="izobrazhenie"
                                                                      href="/sankt-peterburg/announcements/gruppa-papin-olimpos-sankt-peterburg?eventId=18689624"><span
                class="LazyImage_container__ukc21"><picture class="LazyImage_backgroundContainer__UuhKK"><img
                loading="lazy" class="LazyImage_img__Nz285" alt="Концерт группы «Папин Олимпос»"
                src="https://live.mts.ru/image/536x360/gruppa-papin-olimpos-sankt-peterburg-c037cc69-db7b-de59-fee5-09cc7f6f927d.jpg"></picture><picture
                class="LazyImage_originalContainer__Kj81p LazyImage_originalVisible__Xn1yN LazyImage_originalHoverAnimation__jNykx"><img
                loading="eager" class="LazyImage_img__Nz285 LazyImage_original___06Jx"
                src="https://live.mts.ru/image/536x360/gruppa-papin-olimpos-sankt-peterburg-c037cc69-db7b-de59-fee5-09cc7f6f927d.jpg"
                style="transform: scale(1);"></picture></span></a>
              <div class="FavouritesItemStatusButton_container__KIJN8">
                <button class="FavouritesItemStatusButton_button__VRYUO">
                  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 60 60"
                       class="FavouritesItemStatusButton_icon__maJuJ">
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
                      <use clip-path="url(#FavouritesDynamic_svg__c)" data-id="heart"
                           href="#FavouritesDynamic_svg__a"></use>
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
              <div class="AnnouncementPreviewBages_container__mD1TA"><a title="Концерт группы «Папин Олимпос»"
                                                                        data-type="cena"
                                                                        href="/sankt-peterburg/announcements/gruppa-papin-olimpos-sankt-peterburg?eventId=18689624">
                <div class="Badge_container__rAaAq Badge_secondary__1BKTK">от 1400 ₽</div>
              </a></div>
              <a title="Концерт группы «Папин Олимпос»" class="AnnouncementPreviewDescription_titleContainer__Ngzmi"
                 data-type="nazvanie_meropriyatiya"
                 href="/sankt-peterburg/announcements/gruppa-papin-olimpos-sankt-peterburg?eventId=18689624"><h3
                  class="AnnouncementPreviewDescription_title__VoBCR">Концерт группы «Папин Олимпос»</h3></a>
              <time class="AnnouncementPreviewDescription_time__UtzSc">24 мая, 20:00</time>
              <div title="Клуб Sound (Санкт-Петербург)" class="AnnouncementPreviewDescription_venue__ST9Py"><a
                  href="/sankt-peterburg/venues/club/klub-sound-sankt-peterburg">Клуб Sound (Санкт-Петербург)</a>
              </div>
            </div>
          </article>
        </div>
        <div class="swiper-slide CommonSlider_slide__WGyyy" style="width: 268.5px; margin-right: 22px;">
          <article class="AnnouncementPreview_container__A430m">
            <div class="AnnouncementPreviewImage_container__Pmp0c"><a title="Концерт Dead Blonde"
                                                                      class="AnnouncementPreviewImage_link__5btQS"
                                                                      data-type="izobrazhenie"
                                                                      href="/sankt-peterburg/announcements/dead-blonde-sankt-peterburg?eventId=18566370"><span
                class="LazyImage_container__ukc21"><picture class="LazyImage_backgroundContainer__UuhKK"><img
                loading="lazy" class="LazyImage_img__Nz285" alt="Концерт Dead Blonde"
                src="https://live.mts.ru/image/536x360/dead-blonde-sankt-peterburg-c1ca941c-a567-bded-d29c-3f16c01de32e.jpg"></picture><picture
                class="LazyImage_originalContainer__Kj81p LazyImage_originalVisible__Xn1yN LazyImage_originalHoverAnimation__jNykx"><img
                loading="eager" class="LazyImage_img__Nz285 LazyImage_original___06Jx"
                src="https://live.mts.ru/image/536x360/dead-blonde-sankt-peterburg-c1ca941c-a567-bded-d29c-3f16c01de32e.jpg"
                style="transform: scale(1);"></picture></span></a>
              <div class="FavouritesItemStatusButton_container__KIJN8">
                <button class="FavouritesItemStatusButton_button__VRYUO">
                  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 60 60"
                       class="FavouritesItemStatusButton_icon__maJuJ">
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
                      <use clip-path="url(#FavouritesDynamic_svg__c)" data-id="heart"
                           href="#FavouritesDynamic_svg__a"></use>
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
              <div class="AnnouncementPreviewBages_container__mD1TA"><a title="Концерт Dead Blonde" data-type="cena"
                                                                        href="/sankt-peterburg/announcements/dead-blonde-sankt-peterburg?eventId=18566370">
                <div class="Badge_container__rAaAq Badge_secondary__1BKTK">от 1999 ₽</div>
              </a></div>
              <a title="Концерт Dead Blonde" class="AnnouncementPreviewDescription_titleContainer__Ngzmi"
                 data-type="nazvanie_meropriyatiya"
                 href="/sankt-peterburg/announcements/dead-blonde-sankt-peterburg?eventId=18566370"><h3
                  class="AnnouncementPreviewDescription_title__VoBCR">Концерт Dead Blonde</h3></a>
              <time class="AnnouncementPreviewDescription_time__UtzSc">19 мая, 20:00</time>
              <div title="Клуб «А2»" class="AnnouncementPreviewDescription_venue__ST9Py"><a
                  href="/sankt-peterburg/venues/club/club-a2">Клуб «А2»</a></div>
            </div>
          </article>
        </div>
        <div class="swiper-slide CommonSlider_slide__WGyyy" style="width: 268.5px; margin-right: 22px;">
          <article class="AnnouncementPreview_container__A430m">
            <div class="AnnouncementPreviewImage_container__Pmp0c"><a title="Концерт группы «Lumen» &amp; Orchestra"
                                                                      class="AnnouncementPreviewImage_link__5btQS"
                                                                      data-type="izobrazhenie"
                                                                      href="/sankt-peterburg/announcements/lumen-orchestra?eventId=18634218"><span
                class="LazyImage_container__ukc21"><picture class="LazyImage_backgroundContainer__UuhKK"><img
                loading="lazy" class="LazyImage_img__Nz285" alt="Концерт группы «Lumen» &amp; Orchestra"
                src="https://live.mts.ru/image/536x360/lumen-orchestra-7d7d2a85-5feb-9cf6-61e3-b16f9a173b0f.jpg"></picture><picture
                class="LazyImage_originalContainer__Kj81p LazyImage_originalVisible__Xn1yN LazyImage_originalHoverAnimation__jNykx"><img
                loading="eager" class="LazyImage_img__Nz285 LazyImage_original___06Jx"
                src="https://live.mts.ru/image/536x360/lumen-orchestra-7d7d2a85-5feb-9cf6-61e3-b16f9a173b0f.jpg"
                style="transform: scale(1);"></picture></span></a>
              <div class="FavouritesItemStatusButton_container__KIJN8">
                <button class="FavouritesItemStatusButton_button__VRYUO">
                  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 60 60"
                       class="FavouritesItemStatusButton_icon__maJuJ">
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
                      <use clip-path="url(#FavouritesDynamic_svg__c)" data-id="heart"
                           href="#FavouritesDynamic_svg__a"></use>
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
              <div class="AnnouncementPreviewBages_container__mD1TA"><a
                  title="Концерт группы «Lumen» &amp; Orchestra" data-type="cena"
                  href="/sankt-peterburg/announcements/lumen-orchestra?eventId=18634218">
                <div class="Badge_container__rAaAq Badge_secondary__1BKTK">от 2800 ₽</div>
              </a></div>
              <a title="Концерт группы «Lumen» &amp; Orchestra"
                 class="AnnouncementPreviewDescription_titleContainer__Ngzmi" data-type="nazvanie_meropriyatiya"
                 href="/sankt-peterburg/announcements/lumen-orchestra?eventId=18634218"><h3
                  class="AnnouncementPreviewDescription_title__VoBCR">Концерт группы «Lumen» &amp; Orchestra</h3></a>
              <time class="AnnouncementPreviewDescription_time__UtzSc">17 мая, 20:00</time>
              <div title="МТС Live Холл Санкт-Петербург" class="AnnouncementPreviewDescription_venue__ST9Py"><a
                  href="/sankt-peterburg/venues/concert-hall/mts-live-hall-spb">МТС Live Холл Санкт-Петербург</a>
              </div>
            </div>
          </article>
        </div>
        <div class="swiper-slide CommonSlider_slide__WGyyy" style="width: 268.5px; margin-right: 22px;">
          <article class="AnnouncementPreview_container__A430m">
            <div class="AnnouncementPreviewImage_container__Pmp0c"><a title="Концерт «GSPD. Summer Sound x билайн»"
                                                                      class="AnnouncementPreviewImage_link__5btQS"
                                                                      data-type="izobrazhenie"
                                                                      href="/sankt-peterburg/announcements/gspd-summer-sound-x-bilayn?eventId=18736306"><span
                class="LazyImage_container__ukc21"><picture class="LazyImage_backgroundContainer__UuhKK"><img
                loading="lazy" class="LazyImage_img__Nz285" alt="Концерт «GSPD. Summer Sound x билайн»"
                src="https://live.mts.ru/image/536x360/gspd-summer-sound-x-bilayn-0d523804-67ba-fc6b-ae80-a89ee1c9ba97.jpg"></picture><picture
                class="LazyImage_originalContainer__Kj81p LazyImage_originalVisible__Xn1yN LazyImage_originalHoverAnimation__jNykx"><img
                loading="eager" class="LazyImage_img__Nz285 LazyImage_original___06Jx"
                src="https://live.mts.ru/image/536x360/gspd-summer-sound-x-bilayn-0d523804-67ba-fc6b-ae80-a89ee1c9ba97.jpg"
                style="transform: scale(1);"></picture></span></a>
              <div class="FavouritesItemStatusButton_container__KIJN8">
                <button class="FavouritesItemStatusButton_button__VRYUO">
                  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 60 60"
                       class="FavouritesItemStatusButton_icon__maJuJ">
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
                      <use clip-path="url(#FavouritesDynamic_svg__c)" data-id="heart"
                           href="#FavouritesDynamic_svg__a"></use>
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
              <div class="AnnouncementPreviewBages_container__mD1TA"><a
                  title="Концерт «GSPD. Summer Sound x билайн»" data-type="cena"
                  href="/sankt-peterburg/announcements/gspd-summer-sound-x-bilayn?eventId=18736306">
                <div class="Badge_container__rAaAq Badge_secondary__1BKTK">от 1500 ₽</div>
              </a></div>
              <a title="Концерт «GSPD. Summer Sound x билайн»"
                 class="AnnouncementPreviewDescription_titleContainer__Ngzmi" data-type="nazvanie_meropriyatiya"
                 href="/sankt-peterburg/announcements/gspd-summer-sound-x-bilayn?eventId=18736306"><h3
                  class="AnnouncementPreviewDescription_title__VoBCR">Концерт «GSPD. Summer Sound x билайн»</h3></a>
              <time class="AnnouncementPreviewDescription_time__UtzSc">05 августа, 19:30</time>
              <div title="Двор Гостинки" class="AnnouncementPreviewDescription_venue__ST9Py"><a
                  href="/sankt-peterburg/venues/drugoe/dvor-gostinki">Двор Гостинки</a></div>
            </div>
          </article>
        </div>
        <div class="swiper-slide CommonSlider_slide__WGyyy" style="width: 268.5px; margin-right: 22px;">
          <article class="AnnouncementPreview_container__A430m">
            <div class="AnnouncementPreviewImage_container__Pmp0c"><a title="Концерт Теоны Контридзе"
                                                                      class="AnnouncementPreviewImage_link__5btQS"
                                                                      data-type="izobrazhenie"
                                                                      href="/sankt-peterburg/announcements/koncert-teony-kontridze?eventId=18004438"><span
                class="LazyImage_container__ukc21"><picture class="LazyImage_backgroundContainer__UuhKK"><img
                loading="lazy" class="LazyImage_img__Nz285" alt="Концерт Теоны Контридзе"
                src="https://live.mts.ru/image/536x360/koncert-teony-kontridze-863dc19c-12aa-6ff7-f604-eb80acabde95.jpg"></picture><picture
                class="LazyImage_originalContainer__Kj81p LazyImage_originalVisible__Xn1yN LazyImage_originalHoverAnimation__jNykx"><img
                loading="eager" class="LazyImage_img__Nz285 LazyImage_original___06Jx"
                src="https://live.mts.ru/image/536x360/koncert-teony-kontridze-863dc19c-12aa-6ff7-f604-eb80acabde95.jpg"
                style="transform: scale(1);"></picture></span></a>
              <div class="FavouritesItemStatusButton_container__KIJN8">
                <button class="FavouritesItemStatusButton_button__VRYUO">
                  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 60 60"
                       class="FavouritesItemStatusButton_icon__maJuJ">
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
                      <use clip-path="url(#FavouritesDynamic_svg__c)" data-id="heart"
                           href="#FavouritesDynamic_svg__a"></use>
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
              <div class="AnnouncementPreviewBages_container__mD1TA"><a title="Концерт Теоны Контридзе"
                                                                        data-type="cena"
                                                                        href="/sankt-peterburg/announcements/koncert-teony-kontridze?eventId=18004438">
                <div class="Badge_container__rAaAq Badge_secondary__1BKTK">от 1700 ₽</div>
              </a></div>
              <a title="Концерт Теоны Контридзе" class="AnnouncementPreviewDescription_titleContainer__Ngzmi"
                 data-type="nazvanie_meropriyatiya"
                 href="/sankt-peterburg/announcements/koncert-teony-kontridze?eventId=18004438"><h3
                  class="AnnouncementPreviewDescription_title__VoBCR">Концерт Теоны Контридзе</h3></a>
              <time class="AnnouncementPreviewDescription_time__UtzSc">20 сентября, 20:00</time>
              <div title="МТС Live Холл Санкт-Петербург" class="AnnouncementPreviewDescription_venue__ST9Py"><a
                  href="/sankt-peterburg/venues/concert-hall/mts-live-hall-spb">МТС Live Холл Санкт-Петербург</a>
              </div>
            </div>
          </article>
        </div>
      </div>
    </div>
  </main>
</template>

<style scoped>

</style>