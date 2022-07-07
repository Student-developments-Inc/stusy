<template>
  <div>
    <div class="tab__header">
      <a href="#"
         class="tab__link">
        <span class="moving-icon" v-if="role === `teacher`">
          <img src="@/assets/moving.svg">
        </span>
        <div class="mainTitle" @click.prevent="active = !active">
        <span class="avatar-arrow" v-bind:style="[!active?{'transform':'rotate(-90deg)'}:{}]">
          <img src="@/assets/arrow.svg"/>
        </span>
          <strong>{{ title }}</strong>
        </div>
        <div class="createNewViewSection" v-if="role === `teacher`">
          <span @click="menuAction" v-if="!menu">
            <img src="@/assets/addedV2.svg"/>
          </span>
          <ul class="sub-menu" v-if="menu">
            <li @click="menuAction">
              <a>
                <img class="viewSection-arrow" src="@/assets/arrow.svg">
                <span>Вид раздела</span>
              </a>
            </li>
            <li><a>Тест</a></li>
            <li @click="isModalOpen = true"><a>Ссылка</a></li>
            <li><a>Файлы</a></li>
            <li><a>Текст</a></li>
          </ul>
        </div>
        <span class="redactor-icon-pen" v-if="role === `teacher`">
          <img src="@/assets/pen.svg"/>
        </span>
        <span class="trash-icon" v-if="role === `teacher`">
          <img src="@/assets/trash.svg"/>
        </span>
      </a>
    </div>
    <transition name="fade">
      <div class="tab__content" v-show="active">
        <slot/>
      </div>
    </transition>
    <ModalWindow v-if="isModalOpen">
      <div class="modal-window">
        <div>
          <img src="@/assets/diskette.svg" @click="isModalOpen = false">
          <h1>Создание ссылки:</h1>
        </div>
        <label class="input-form">
          <label>Название</label>
          <input type="text" alt="coursesName"/>
        </label>
        <label class="input-form">
          <label>Ссылка</label>
          <input type="text" alt="coursesName"/>
        </label>
      </div>
    </ModalWindow>
  </div>
</template>

<script>
import ModalWindow from "@/components/ModalWindow";

export default {
  name: "AccordionMenu",
  components: {ModalWindow},
  props: [
    "title",
    "role"
  ],
  data() {
    return {
      menu: false,
      active: false,
      isModalOpen: false
    };
  },
  methods: {
    menuAction() {
      this.menu = !this.menu;
    }
  }
};
</script>

<style scoped>
ul {
  list-style-type: none;
}

.moving-icon {
  margin-right: 28px;
  cursor: pointer !important;
}

.avatar-arrow {
  filter: invert(22%) sepia(61%) saturate(6241%) hue-rotate(228deg) brightness(102%) contrast(92%);
  margin-right: 13px;
  transition: .3s;
}

.viewSection-arrow {
  filter: invert(100%) sepia(0%) saturate(7492%) hue-rotate(22deg) brightness(106%) contrast(100%);
  margin-right: 16px;
}

.redactor-icon-pen {
  margin-right: 12px;
}

.trash-icon {
  margin-left: 5px;
}

.mainTitle {
  display: flex;
  align-items: center;
  margin-right: 9px;
}

.tab__link {
  text-decoration: none;
  color: var(--dark);
}

.tab__content {
  margin-left: 22px;
}

.tab__header a {
  display: flex;
  flex-direction: row;
  align-items: center;
}

.tab__header a strong {
  font-style: normal;
  font-weight: 400;
  font-size: 20px;
  line-height: 23px;
}

.createNewViewSection {
  height: 24px;
  min-width: 24px;
  margin-right: 12px;
}

.sub-menu {
  position: absolute;
  display: flex;
  flex-direction: column;
  border-radius: 15px;
  background-color: var(--blue);
  font-style: normal;
  font-weight: 400;
  font-size: 18px;
  line-height: 21px;
  color: var(--light);
}

.sub-menu li {
  padding: 7px 6px 7px 45px;
}

.sub-menu > li:first-child {
  padding: 16px 6px 8px 17px;
  border-radius: 15px 15px 0 0;
  border-bottom: 1px solid #1E42BB;
}

.sub-menu > li:last-child {
  padding: 7px 6px 23px 45px;
  border-radius: 0 0 15px 15px;
}

.sub-menu li:hover {
  background: var(--blue-mate);
}

.modal-window {
  display: flex;
  flex-direction: column;
  gap: 17px;
  background-color: var(--light);
  border-radius: 15px;
  padding: 30px 20px;
}

.modal-window h1 {
  margin-bottom: 10px;
  font-style: normal;
  font-weight: 400;
  font-size: 36px;
  line-height: 41px;
}

.modal-window label {
  font-style: normal;
  font-weight: 400;
  font-size: 20px;
  line-height: 23px;
}

.fade-enter-active {
  animation: scale-up-ver-top .2s cubic-bezier(0.390, 0.575, 0.565, 1.000) both;
}

@keyframes scale-up-ver-top {
  0% {
    transform: scaleY(0.4);
    transform-origin: 100% 0;
  }
  100% {
    transform: scaleY(1);
    transform-origin: 100% 0;
  }
}
</style>