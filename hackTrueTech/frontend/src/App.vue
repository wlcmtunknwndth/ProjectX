<script setup>
import ShowMoreComponent from "./components/ShowMoreComponent.vue";
</script>

<template>
  <div id="__next">
    <div class="CoreLayout_container__2ZTtd">
      <div class="CoreLayout_innerLayout__TKiOh">
        <header id="header" class="CoreLayoutHeader_container__6x02M">
          <div class="flexboxgrid_container__YXz7C CoreLayoutHeader_gridContainer__k84do">
            <router-link :to="{ name: 'home' }" class="CoreLayoutHeader_logoLink__DvDCy">
              <img src="/images/header/Logo.svg" title="MTC LIVE" alt="MTC LIVE">

            </router-link>
            <div class="CoreLayoutHeader_search__K4SKu">
              <div class="Search_container__XfiK9 CoreLayoutSearch_search__aTT2H">
                <div
                    class="Search_search__cRxhy Search_search__cRxhy-single Search_search__cRxhy-allow-clear Search_search__cRxhy-show-search"
                    maxlength="200">
                  <div class="Search_search__cRxhy-selector"><span class="Search_search__cRxhy-selection-search"><input
                      type="search" autocomplete="off" class="Search_search__cRxhy-selection-search-input"
                      role="combobox"
                      aria-expanded="false" aria-haspopup="listbox" aria-owns="rc_select_15_list"
                      aria-autocomplete="list"
                      aria-controls="rc_select_15_list" maxlength="200" value="" id="rc_select_15"></span><span
                      class="Search_search__cRxhy-selection-placeholder"><span class="Search_placeholder__zzA45">Событие, персона, площадка</span></span>
                  </div>
                </div>
                <div class="Search_searchCancel___gQIu">Отмена</div>
              </div>
            </div>
            <div class="CoreLayoutHeader_mainControlsContainer__b9Esz"><a class="CoreLayoutSupportLink_container__SEAma"
                                                                          href="https://support.mts.ru/mts_mts_live"
                                                                          target="_blank" rel="noopener noreferrer">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16">
                <path fill="currentColor" fill-rule="evenodd"
                      d="M3.124 3.124c-.875.875-.94 1.747-1.07 3.491C2.02 7.067 2 7.534 2 8c0 .466.02.933.054 1.385.13 1.744.195 2.616 1.07 3.491s1.747.94 3.491 1.07C7.067 13.98 7.534 14 8 14c.466 0 .933-.02 1.385-.054 1.744-.13 2.616-.195 3.491-1.07s.94-1.747 1.07-3.491C13.98 8.933 14 8.466 14 8c0-.466-.02-.933-.054-1.385-.13-1.744-.195-2.616-1.07-3.491s-1.747-.94-3.491-1.07A18.735 18.735 0 0 0 8 2c-.466 0-.933.02-1.385.054-1.744.13-2.616.195-3.491 1.07Zm3.14 1.558A17.35 17.35 0 0 1 8 4.583c.548 0 1.095.038 1.641.089.276.025.676.06 1.028.376.285.255.404.581.48.945.04.197.072.427.072.674 0 .227-.027.443-.062.632-.073.397-.188.763-.502 1.04-.375.328-.82.35-1.125.364l-.158.008-.067.004h-.645c-.02.152-.124.619-.682.619-.468 0-.756-.386-.701-.82.032-.256.113-.564.36-.825.251-.264.562-.358.815-.403.23-.041.513-.056.802-.07l.086-.005c.103-.005.207-.01.31-.021a2.5 2.5 0 0 0 .069-.523c0-.164-.03-.325-.066-.484-.067-.01-.134-.016-.2-.022l-.018-.002A15.46 15.46 0 0 0 8 6.084c-.582 0-1.172.047-1.581.09a.171.171 0 0 0-.036.008.185.185 0 0 1-.042.008.75.75 0 1 1-1.439-.408 1.41 1.41 0 0 1 .21-.476c.079-.115.178-.21.251-.273.073-.063.183-.147.309-.207.188-.09.386-.123.591-.144Zm2.636 6.25a.9.9 0 1 0-1.8 0 .9.9 0 0 0 1.8 0Z"
                      clip-rule="evenodd"></path>
              </svg>
              <span>Поддержка</span></a>
              <button class="CoreLayoutLocationButton_button__MwMD8" title="Выбор региона" type="button">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 8 8"
                     class="CoreLayoutLocationButton_icon__XrVWK CoreLayoutLocationButton_geoIcon__8_vId">
                  <path fill="currentColor"
                        d="M6.731.145a.889.889 0 0 1 1.125 1.124L5.847 7.295a.889.889 0 0 1-1.605.176L2.965 5.343a.89.89 0 0 0-.305-.305L.53 3.758a.889.889 0 0 1 .176-1.605L6.731.145z"></path>
                </svg>
                <span class="CoreLayoutLocationButton_label__XAMlw">Санкт-Петербург</span>
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16"
                     class="CoreLayoutLocationButton_icon__XrVWK CoreLayoutLocationButton_triangleIcon__Agacp">
                  <path fill="currentColor"
                        d="M8.78 10.024a1 1 0 0 1-1.56 0l-1.92-2.4A1 1 0 0 1 6.08 6h3.84a1 1 0 0 1 .78 1.625l-1.92 2.399Z"></path>
                </svg>
              </button>
            </div>
            <div class="CoreLayoutHeader_userControlsContainer__lvTR_">
              <div class="IdentityProfileWidget_container__SWzSr">
                <div class="prp-profile-b2c-widget svelte-23o2lu" data-mts-theme="light" data-header="">
                  <div class="prp-profile-b2c-account svelte-23o2lu">
                    <button class="prp-button -default svelte-1tuwkvw" mousedown="false">Войти
                      <div class="prp-button__round round svelte-1tuwkvw"></div>
                    </button>
                  </div>
                </div>
              </div>
            </div>
            <nav class="CoreLayoutMenu_container__ddij2 CoreLayoutHeader_menu__gLs6J">
              <ul class="CoreLayoutMenu_menu__BfEW2 CoreLayoutMenu_hasMore__mMTuu">
                <li class="CoreLayoutMenuItem_container__luLtl"><a class="CoreLayoutMenuItem_link__Bo_mY" target="_self"
                                                                   href="/sankt-peterburg">Все события</a></li>
                <li class="CoreLayoutMenuItem_container__luLtl"><a class="CoreLayoutMenuItem_link__Bo_mY" target="_self"
                                                                   href="/sankt-peterburg/venues/concert-hall/mts-live-hall-spb"><img
                    src="https://live.mts.ru/api/bff-sections/sections/25/icon" alt="МТС Live Холл"
                    class="CoreLayoutMenuItem_linkIcon__uXxfT">МТС Live Холл</a></li>
                <li class="CoreLayoutMenuItem_container__luLtl"><a
                    class="CoreLayoutMenuItem_link__Bo_mY" target="_self" href="/sankt-peterburg/collections/concerts">Концерты</a>
                </li>
                <li class="CoreLayoutMenuItem_container__luLtl"><a class="CoreLayoutMenuItem_link__Bo_mY" target="_self"
                                                                   href="/sankt-peterburg/collections/theater">Театры</a>
                </li>
                <li class="CoreLayoutMenuItem_container__luLtl"><a class="CoreLayoutMenuItem_link__Bo_mY" target="_self"
                                                                   href="/sankt-peterburg/collections/musicals">Мюзиклы</a>
                </li>
                <li class="CoreLayoutMenuItem_container__luLtl"><a class="CoreLayoutMenuItem_link__Bo_mY" target="_self"
                                                                   href="/sankt-peterburg/collections/standup">Стендап</a>
                </li>
                <li class="CoreLayoutMenuItem_container__luLtl"><a class="CoreLayoutMenuItem_link__Bo_mY" target="_self"
                                                                   href="/sankt-peterburg/collections/children">Детям</a>
                </li>
                <ShowMoreComponent />
              </ul>
            </nav>
          </div>
        </header>

        <RouterView />

        <footer class="CoreLayoutFooter_container__awfyg">
          <div class="flexboxgrid_container__YXz7C">
            <div class="CoreLayoutFooterDownloadAppBlock_container__EvCRI"><span
                class="LazyImage_container__ukc21 CoreLayoutFooterDownloadAppBlock_qrImage__efHCr"><picture
                class="LazyImage_originalContainer__Kj81p LazyImage_originalVisible__Xn1yN"><img loading="eager"
                                                                                                 class="LazyImage_img__Nz285"
                                                                                                 src="/images/footer/QR@2x.webp"
                                                                                                 style="transform: scale(1);"></picture></span>
              <div class="CoreLayoutFooterDownloadAppBlock_downloadContainer__2pPGF"><h2
                  class="CoreLayoutFooterDownloadAppBlock_title__IcYof">Скачать приложение МТС&nbsp;Live</h2>
                <p class="CoreLayoutFooterDownloadAppBlock_description__7hEZs">Удобные сервисы: навигация на площадках,
                  заказ из бара и другое</p>
                <div class="CoreLayoutFooterDownloadAppBlock_serviceListContainer__GYEpl">
                  <a href="https://appbazar.ru/preview/?id=d9234d2c-b833-411c-9942-4f08063b9cd6">
                    <span class="LazyImage_container__ukc21 CoreLayoutFooterDownloadAppService_image__kBy2Y">
                      <picture class="LazyImage_originalContainer__Kj81p LazyImage_originalVisible__Xn1yN">
                        <img loading="eager" class="LazyImage_img__Nz285" src="/images/footer/AppBazar@2x.webp"
                             style="transform: scale(1);">
                      </picture>
                    </span>
                  </a>
                  <a href="https://play.google.com/store/apps/details?id=ru.mts.live.app&amp;hl=ru&amp;gl=US&amp;pli=1">
                    <span class="LazyImage_container__ukc21 CoreLayoutFooterDownloadAppService_image__kBy2Y">
                      <picture class="LazyImage_originalContainer__Kj81p LazyImage_originalVisible__Xn1yN">
                        <img loading="eager" class="LazyImage_img__Nz285" src="/images/footer/GooglePlay@2x.webp"
                             style="transform: scale(1);">
                      </picture>
                    </span>
                  </a>
                  <a href="https://apps.apple.com/ru/app/%D0%BC%D1%82%D1%81-live-%D0%B0%D1%84%D0%B8%D1%88%D0%B0-%D0%B8-%D0%B1%D0%B8%D0%BB%D0%B5%D1%82%D1%8B/id1539924684">
                    <span class="LazyImage_container__ukc21 CoreLayoutFooterDownloadAppService_image__kBy2Y">
                      <picture class="LazyImage_originalContainer__Kj81p LazyImage_originalVisible__Xn1yN">
                        <img loading="eager" class="LazyImage_img__Nz285" src="/images/footer/AppStore@2x.webp"
                             style="transform: scale(1);">
                      </picture>
                    </span>
                  </a>
                </div>
              </div>
            </div>
            <hr class="CoreLayoutFooter_separator___OEZP">
            <div class="CoreLayoutFooter_title__m7RBb">МТС Live</div>
            <div class="CoreLayoutFooterMenuBlock_container__eksk1"><h3 class="CoreLayoutFooterMenuBlock_title__4NK0Z">
              Твоя персональная подборка событий</h3>
              <ul class="CoreLayoutFooterMenuBlock_menu__df59_">
                <li class="CoreLayoutFooterMenuBlock_menuItem__V99sD"><a title="Все события"
                                                                         class="CoreLayoutFooter_link___ElmY"
                                                                         target="_self" href="/sankt-peterburg">Все
                  события</a></li>
                <li class="CoreLayoutFooterMenuBlock_menuItem__V99sD"><a title="МТС Live Холл"
                                                                         class="CoreLayoutFooter_link___ElmY"
                                                                         target="_self"
                                                                         href="/sankt-peterburg/venues/concert-hall/mts-live-hall-spb">МТС
                  Live Холл</a></li>
                <li class="CoreLayoutFooterMenuBlock_menuItem__V99sD"><a title="Концерты"
                                                                         class="CoreLayoutFooter_link___ElmY"
                                                                         target="_self"
                                                                         href="/sankt-peterburg/collections/concerts">Концерты</a>
                </li>
                <li class="CoreLayoutFooterMenuBlock_menuItem__V99sD"><a title="Театры"
                                                                         class="CoreLayoutFooter_link___ElmY"
                                                                         target="_self"
                                                                         href="/sankt-peterburg/collections/theater">Театры</a>
                </li>
                <li class="CoreLayoutFooterMenuBlock_menuItem__V99sD"><a title="Мюзиклы"
                                                                         class="CoreLayoutFooter_link___ElmY"
                                                                         target="_self"
                                                                         href="/sankt-peterburg/collections/musicals">Мюзиклы</a>
                </li>
                <li class="CoreLayoutFooterMenuBlock_menuItem__V99sD"><a title="Стендап"
                                                                         class="CoreLayoutFooter_link___ElmY"
                                                                         target="_self"
                                                                         href="/sankt-peterburg/collections/standup">Стендап</a>
                </li>
                <li class="CoreLayoutFooterMenuBlock_menuItem__V99sD"><a title="Детям"
                                                                         class="CoreLayoutFooter_link___ElmY"
                                                                         target="_self"
                                                                         href="/sankt-peterburg/collections/children">Детям</a>
                </li>
                <li class="CoreLayoutFooterMenuBlock_menuItem__V99sD"><a title="Подарочная карта"
                                                                         class="CoreLayoutFooter_link___ElmY"
                                                                         target="_self" href="/giftcard">Подарочная
                  карта</a></li>
                <li class="CoreLayoutFooterMenuBlock_menuItem__V99sD"><a title="LifeStyle"
                                                                         class="CoreLayoutFooter_link___ElmY"
                                                                         target="_self" href="/lifestyle">LifeStyle</a>
                </li>
                <li class="CoreLayoutFooterMenuBlock_menuItem__V99sD"><a title="Шоу"
                                                                         class="CoreLayoutFooter_link___ElmY"
                                                                         target="_self"
                                                                         href="/sankt-peterburg/collections/show">Шоу</a>
                </li>
                <li class="CoreLayoutFooterMenuBlock_menuItem__V99sD"><a title="Фестивали"
                                                                         class="CoreLayoutFooter_link___ElmY"
                                                                         target="_self"
                                                                         href="/sankt-peterburg/collections/festivals">Фестивали</a>
                </li>
                <li class="CoreLayoutFooterMenuBlock_menuItem__V99sD"><a title="Выставки"
                                                                         class="CoreLayoutFooter_link___ElmY"
                                                                         target="_self"
                                                                         href="/sankt-peterburg/collections/exhibitions">Выставки</a>
                </li>
                <li class="CoreLayoutFooterMenuBlock_menuItem__V99sD"><a title="Экскурсии"
                                                                         class="CoreLayoutFooter_link___ElmY"
                                                                         target="_self"
                                                                         href="/sankt-peterburg/collections/excursions">Экскурсии</a>
                </li>
                <li class="CoreLayoutFooterMenuBlock_menuItem__V99sD"><a title="Спорт"
                                                                         class="CoreLayoutFooter_link___ElmY"
                                                                         target="_self"
                                                                         href="/sankt-peterburg/collections/sport">Спорт</a>
                </li>
                <router-link :to="{ name: 'main' }">
                  <li class="CoreLayoutFooterMenuBlock_menuItem__V99sD"><a title="Спорт"
                                                                           class="CoreLayoutFooter_link___ElmY"
                                                                           target="_self"
                                                                           style="line-height: 20px;"
                                                                           href="/sankt-peterburg/collections/sport">Доступно
                    для людей с ограниченными возможностями</a>
                  </li>
                </router-link>
                <li class="CoreLayoutFooterMenuBlock_menuItem__V99sD"><a title="Все площадки"
                                                                         class="CoreLayoutFooter_link___ElmY"
                                                                         target="_self" href="/sankt-peterburg/venues">Все
                  площадки</a></li>
                <li class="CoreLayoutFooterMenuBlock_menuItem__V99sD"><a title="Пользовательское соглашение"
                                                                         class="CoreLayoutFooter_link___ElmY"
                                                                         target="_self" href="/users/agreement">Пользовательское
                  соглашение</a></li>
                <li class="CoreLayoutFooterMenuBlock_menuItem__V99sD"><a title="МТС Premium"
                                                                         class="CoreLayoutFooter_link___ElmY"
                                                                         target="_self" href="/premium">МТС Premium</a>
                </li>
                <li class="CoreLayoutFooterMenuBlock_menuItem__V99sD"><a title="Возврат билетов"
                                                                         class="CoreLayoutFooter_link___ElmY"
                                                                         target="_blank"
                                                                         href="https://support.mts.ru/mts_mts_live/Bileti/kak-vernut-bilet-i-chto-delat-esli-meropriyatie-otmenili">Возврат
                  билетов</a></li>
                <li class="CoreLayoutFooterMenuBlock_menuItem__V99sD"><a title="Поддержка"
                                                                         class="CoreLayoutFooter_link___ElmY"
                                                                         target="_blank"
                                                                         href="https://support.mts.ru/mts_mts_live">Поддержка</a>
                </li>
                <li class="CoreLayoutFooterMenuBlock_menuItem__V99sD"><a title="Корпоративным клиентам"
                                                                         class="CoreLayoutFooter_link___ElmY"
                                                                         target="_blank"
                                                                         href="https://corporatesales.live.mts.ru/">Корпоративным
                  клиентам</a></li>
                <li class="CoreLayoutFooterMenuBlock_menuItem__V99sD"><a title="Организаторам"
                                                                         class="CoreLayoutFooter_link___ElmY"
                                                                         target="_blank" href="https://b2b-live.mts.ru">Организаторам</a>
                </li>
                <li class="CoreLayoutFooterMenuBlock_menuItem__V99sD"><a title="Партнерская программа"
                                                                         class="CoreLayoutFooter_link___ElmY"
                                                                         target="_blank"
                                                                         href="https://affiliates.live.mts.ru/">Партнерская
                  программа</a></li>
                <li class="CoreLayoutFooterMenuBlock_menuItem__V99sD"><a title="Рекомендательные технологии"
                                                                         class="CoreLayoutFooter_link___ElmY"
                                                                         target="_blank"
                                                                         href="https://live.mts.ru/image/pdf/pravila-primeneniya-rekomendatelnyh-tekhnologii-mts-laiv.pdf">Рекомендательные
                  технологии</a></li>
              </ul>
            </div>
            <hr class="CoreLayoutFooter_separator___OEZP">
            <div class="CoreLayoutFooterReferenceInfoBlock_container__iiunD"><a href="tel:+7 (800) 250-20-86"
                                                                                class="CoreLayoutFooter_link___ElmY CoreLayoutFooterReferenceInfoBlock_phone__yLFW3">+7
              (800) 250-20-86</a>
              <div class="CoreLayoutFooterReferenceInfoBlock_links__dpNyU"><a href="https://mts.ru/personal/kontaktu/"
                                                                              title="Контакты"
                                                                              class="CoreLayoutFooter_link___ElmY CoreLayoutFooterReferenceInfoBlock_link__V8ZpI">Контакты</a><a
                  href="https://mts.ru/about/investoram-i-akcioneram/korporativnoe-upravlenie/raskritie-informacii"
                  title="Раскрытие информации"
                  class="CoreLayoutFooter_link___ElmY CoreLayoutFooterReferenceInfoBlock_link__V8ZpI">Раскрытие
                информации</a><a
                  href="https://moskva.mts.ru/about/investoram-i-akcioneram/korporativnoe-upravlenie/dokumenti-pao-mts"
                  title="Документы ПАО МТС"
                  class="CoreLayoutFooter_link___ElmY CoreLayoutFooterReferenceInfoBlock_link__V8ZpI">Документы ПАО
                МТС</a><a href="https://mts.ru/personal/cookie/" title="Политика обработки cookies"
                          class="CoreLayoutFooter_link___ElmY CoreLayoutFooterReferenceInfoBlock_link__V8ZpI">Политика
                обработки cookies</a><a href="https://mts.ru/about/obshhestvennosti-i-smi/komplaens-i-delovaya-etika"
                                        title="Комплаенс и деловая этика"
                                        class="CoreLayoutFooter_link___ElmY CoreLayoutFooterReferenceInfoBlock_link__V8ZpI">Комплаенс
                и деловая этика</a></div>
            </div>
            <div class="CoreLayoutFooter_blocks__7H9Cu">
              <div><h4 class="CoreLayoutFooterSocialBlock_title___W09A">Социальные сети</h4>
                <div class="CoreLayoutFooterSocialBlock_socialList__3g7c6"><a
                    class="CoreLayoutFooterSocialBlock_link__YQrIP" href="https://vk.com/mts_live" target="_blank"
                    rel="noopener noreferrer" title="vk">
                  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 40 40"
                       class="CoreLayoutFooterSocialBlock_icon__MyevH CoreLayoutFooterSocialBlock_vk__GTqFk">
                    <path fill="currentColor"
                          d="M.18 15.384C.613 9.57.83 6.664 3.747 3.746 6.664.828 9.57.612 15.385.179A62.456 62.456 0 0 1 20 0c1.554 0 3.109.068 4.616.18 5.814.433 8.72.65 11.638 3.567 2.918 2.918 3.134 5.824 3.567 11.638.112 1.507.18 3.061.18 4.616 0 1.555-.068 3.11-.18 4.616-.433 5.814-.65 8.72-3.567 11.638-2.918 2.918-5.824 3.134-11.638 3.567-1.507.112-3.062.18-4.616.18-1.555 0-3.11-.068-4.616-.18-5.814-.433-8.72-.65-11.638-3.567C.829 33.336.613 30.43.18 24.616A62.458 62.458 0 0 1 0 20c0-1.555.068-3.11.18-4.616zm10.012-2.224H7.3c-.826 0-.991.4-.991.835 0 .767.991 4.606 4.56 9.68 2.397 3.455 5.768 5.34 8.825 5.34 1.835 0 2.066-.416 2.066-1.134v-2.62c0-.835.165-1.002.76-1.002.43 0 1.157.217 2.876 1.886a35.49 35.49 0 0 1 1.386 1.492c.839.944 1.225 1.379 2.002 1.379h2.892c.826 0 1.24-.417.991-1.235-.248-.835-1.19-2.003-2.43-3.422-.344-.426-.788-.862-1.176-1.243-.355-.35-.664-.652-.806-.86-.43-.567-.33-.817 0-1.301 0 0 3.52-5.008 3.883-6.727.182-.617 0-1.068-.875-1.068H28.37a1.247 1.247 0 0 0-.766.207c-.226.15-.398.37-.49.628 0 0-1.488 3.605-3.553 5.958-.678.684-.992.901-1.356.901-.165 0-.446-.217-.446-.834v-5.792c0-.734-.215-1.068-.826-1.068h-4.545c-.462 0-.743.334-.743.668 0 .267.153.457.349.698.318.392.746.92.807 2.14v4.305c0 .952-.165 1.118-.528 1.118-.992 0-3.372-3.638-4.793-7.794-.264-.801-.545-1.135-1.289-1.135z"></path>
                  </svg>
                </a><a class="CoreLayoutFooterSocialBlock_link__YQrIP" href="https://t.me/mts_live" target="_blank"
                       rel="noopener noreferrer" title="telegram">
                  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 40 40"
                       class="CoreLayoutFooterSocialBlock_icon__MyevH CoreLayoutFooterSocialBlock_tg__t0nYI">
                    <path fill="currentColor"
                          d="M.18 15.387C.611 9.574.827 6.667 3.745 3.75 6.664.832 9.57.615 15.384.183A62.461 62.461 0 0 1 20 .003c1.555 0 3.11.068 4.616.18 5.814.432 8.72.649 11.638 3.566 2.918 2.918 3.134 5.825 3.567 11.638.112 1.507.18 3.062.18 4.616 0 1.555-.068 3.11-.18 4.616-.433 5.814-.65 8.72-3.567 11.638-2.918 2.918-5.824 3.134-11.638 3.567-1.507.112-3.061.18-4.616.18-1.555 0-3.11-.068-4.616-.18-5.814-.433-8.72-.649-11.638-3.567C.828 33.34.612 30.433.179 24.62A62.458 62.458 0 0 1 0 20.003c0-1.554.068-3.109.18-4.616zm26.438 13.268 3.384-15.958c.3-1.407-.506-1.957-1.427-1.612L8.684 18.75c-1.357.531-1.337 1.292-.231 1.637l5.676 2.127 5.27-3.832a.774.774 0 1 1 1.041 1.146l-4.004 3.64-.365 5.163c.531 0 .76-.23 1.036-.505l2.488-2.394 5.156 3.8c.946.53 1.612.25 1.867-.876z"></path>
                  </svg>
                </a><a class="CoreLayoutFooterSocialBlock_link__YQrIP" href="https://ok.ru/mts_live" target="_blank"
                       rel="noopener noreferrer" title="ok">
                  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 44 44"
                       class="CoreLayoutFooterSocialBlock_icon__MyevH CoreLayoutFooterSocialBlock_ok__9kBSl">
                    <path fill="currentColor"
                          d="M25.143 14.658c0-1.685-1.42-3.056-3.166-3.056-1.745 0-3.167 1.37-3.167 3.056 0 1.684 1.422 3.056 3.167 3.056 1.746 0 3.166-1.372 3.166-3.056Z"></path>
                    <path fill="currentColor" fill-rule="evenodd"
                          d="M.194 16.923C.67 10.528.908 7.331 4.118 4.121 7.327.912 10.524.674 16.919.198A68.708 68.708 0 0 1 21.997 0c1.71 0 3.42.075 5.078.198 6.395.476 9.592.714 12.802 3.923 3.209 3.21 3.447 6.407 3.923 12.802.123 1.658.198 3.368.198 5.078 0 1.71-.075 3.42-.198 5.078-.476 6.395-.714 9.592-3.923 12.801-3.21 3.21-6.407 3.448-12.802 3.924a68.75 68.75 0 0 1-5.078.198c-1.71 0-3.42-.075-5.078-.198-6.395-.476-9.592-.714-12.801-3.924C.908 36.671.67 33.473.194 27.08A68.703 68.703 0 0 1-.004 22c0-1.71.075-3.42.198-5.078Zm21.785 5.115c4.218 0 7.647-3.31 7.647-7.38 0-4.072-3.429-7.383-7.647-7.383s-7.648 3.311-7.648 7.383c0 4.07 3.43 7.38 7.648 7.38Zm3.111 6.023a14.614 14.614 0 0 0 4.441-1.775c1.047-.637 1.362-1.972.703-2.982-.66-1.013-2.042-1.317-3.09-.68-3.132 1.901-7.166 1.9-10.296 0-1.049-.637-2.431-.333-3.089.68-.66 1.01-.345 2.345.701 2.982a14.638 14.638 0 0 0 4.441 1.775l-4.275 4.127a2.11 2.11 0 0 0 0 3.059c.438.422 1.011.633 1.584.633a2.28 2.28 0 0 0 1.586-.633l4.2-4.056 4.203 4.056a2.296 2.296 0 0 0 3.167 0 2.108 2.108 0 0 0 0-3.059l-4.276-4.127Z"
                          clip-rule="evenodd"></path>
                  </svg>
                </a><a class="CoreLayoutFooterSocialBlock_link__YQrIP" href="https://dzen.ru/mts_live" target="_blank"
                       rel="noopener noreferrer" title="dzen">
                  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 44 44"
                       class="CoreLayoutFooterSocialBlock_icon__MyevH CoreLayoutFooterSocialBlock_dzen__lm6_T">
                    <path fill="currentColor" fill-rule="evenodd"
                          d="M.197 16.924C.673 10.53.911 7.332 4.12 4.122 7.33.913 10.527.675 16.922.2A68.703 68.703 0 0 1 22 .001c1.71 0 3.42.075 5.078.198 6.395.476 9.592.714 12.801 3.923 3.21 3.21 3.448 6.407 3.924 12.802.123 1.658.198 3.368.198 5.078 0 1.71-.075 3.42-.198 5.078-.476 6.395-.714 9.592-3.924 12.801-3.209 3.21-6.406 3.448-12.801 3.924a68.694 68.694 0 0 1-5.078.198c-1.71 0-3.42-.075-5.078-.198-6.395-.476-9.592-.714-12.802-3.924C.911 36.672.673 33.474.197 27.08A68.704 68.704 0 0 1 0 22.002c0-1.71.075-3.42.198-5.078Zm34.236 4.5h-.007c-10.65-.116-11.749-3.52-11.848-11.85a.568.568 0 0 1 .62-.572 13.034 13.034 0 0 1 11.8 11.803.57.57 0 0 1-.566.62ZM22.578 34.426c.099-8.33 1.197-11.734 11.847-11.85.16.006.313.064.423.182.11.118.163.277.149.437A13.033 13.033 0 0 1 23.147 35a.57.57 0 0 1-.568-.574ZM9.003 20.806A13.04 13.04 0 0 1 20.764 9.005a.57.57 0 0 1 .622.572c-.1 8.324-1.196 11.726-11.811 11.848h-.007a.567.567 0 0 1-.565-.619Zm.148 1.951c.108-.118.247-.202.423-.182 10.616.122 11.711 3.524 11.811 11.848a.57.57 0 0 1-.62.572A13.04 13.04 0 0 1 9.001 23.193a.569.569 0 0 1 .149-.436Z"
                          clip-rule="evenodd"></path>
                  </svg>
                </a></div>
              </div>
              <div><h4 class="CoreLayoutFooterAppsBlock_title__Ug5v0">Приложения</h4>
                <div class="CoreLayoutFooterAppsBlock_apps__kNAEC"><a href="https://liveapp.onelink.me/V9Nw/5d42138"
                                                                      class="CoreLayoutFooterAppsBlock_link__pbVma"
                                                                      target="_blank" rel="noopener noreferrer"
                                                                      title="МТС Live">
                  <div class="CoreLayoutFooterAppsBlock_iconContainer__gREFR">
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 57 48"
                         class="CoreLayoutFooterAppsBlock_liveAppIcon__gkY_c">
                      <path fill="currentColor" d="M30.629 48H8.809L0 0h21.82l8.809 48Z"></path>
                      <path fill="currentColor" d="M35.033 23.988 30.629 48h21.82l4.404-24.01h-21.82Z"
                            data-concat="skip"></path>
                    </svg>
                  </div>
                  <div class="CoreLayoutFooterAppsBlock_label__iu3zp">МТС Live</div>
                  <div class="CoreLayoutFooterAppsBlock_download__quS1f">Скачать</div>
                </a><a
                    href="http://mymts.ru/?referrer=appmetrica_tracking_id%3D170041164043440600%26ym_tracking_id%3D4392336556203169459"
                    class="CoreLayoutFooterAppsBlock_link__pbVma" target="_blank" rel="noopener noreferrer"
                    title="Мой МТС">
                  <div class="CoreLayoutFooterAppsBlock_iconContainer__gREFR">
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 44 44"
                         class="CoreLayoutFooterAppsBlock_myMtsIcon__AQKcq">
                      <path fill="currentColor"
                            d="M28.238 6.092v1.762h3.777v7.004h2.115V7.854h3.778V6.092h-9.67ZM14.634 6.119l-2.763 5.761L9.11 6.12H6.092v8.712h2.115V8.596l2.7 5.267h1.928l2.701-5.267v6.235h2.115V6.119h-3.017ZM35.788 35.158a1.45 1.45 0 0 1-.415.85c-.237.217-.53.365-.845.428-.461.092-.93.134-1.4.127-.5.01-.998-.064-1.475-.217a1.84 1.84 0 0 1-.957-.725c-.23-.342-.345-.837-.345-1.467v-.892c0-.634.117-1.13.345-1.472a1.84 1.84 0 0 1 .955-.724 4.56 4.56 0 0 1 1.475-.217 6.65 6.65 0 0 1 1.4.127c.315.063.607.211.845.429.229.228.375.527.414.849h2.117a3.098 3.098 0 0 0-.725-1.821 3.31 3.31 0 0 0-1.59-.975 8.908 8.908 0 0 0-2.46-.296c-1.014 0-1.893.147-2.613.436a3.341 3.341 0 0 0-1.691 1.452c-.373.643-.57 1.502-.582 2.555v.096l.113.005h-.113v.096c.013 1.053.209 1.911.582 2.554a3.34 3.34 0 0 0 1.69 1.453c.722.289 1.602.435 2.614.435a8.885 8.885 0 0 0 2.46-.295 3.298 3.298 0 0 0 1.59-.975 3.095 3.095 0 0 0 .725-1.821l-2.114.005Z"></path>
                    </svg>
                  </div>
                  <div class="CoreLayoutFooterAppsBlock_label__iu3zp">Мой МТС</div>
                  <div class="CoreLayoutFooterAppsBlock_download__quS1f">Скачать</div>
                </a></div>
              </div>
            </div>
            <div class="CoreLayoutFooterCopyrights_container__iijkO">
              <span>© 2024 ПАО «МТС» Все права защищены</span><span>12+</span>
              <div><p>На информационном ресурсе применяются рекомендательные технологии.</p>Единственным пользователем
                доменного имени live.mts.ru является ООО&nbsp;«МДТЗК» на основании договора с ПАО&nbsp;«МТС». ООО&nbsp;«МДТЗК»
                является информационным посредником. В случае, если третье лицо (правообладатель или уполномоченное им
                лицо) считает, что его права на объект интеллектуальной собственности нарушаются, он может направить
                претензию по адресу: ООО&nbsp;«МДТЗК», 109147, г.&nbsp;Москва, вн.тер.г. муниципальный округ Таганский,
                ул. Воронцовская, д.&nbsp;1/3, стр.&nbsp;2А и по e-mail: <a href="mailto:support@ticketland.ru">support@ticketland.ru</a>.
                ©&nbsp;ООО&nbsp;«МДТЗК», 2023
              </div>
            </div>
          </div>
        </footer>
      </div>
    </div>
  </div>
</template>
