<template>
  <div>
    <div class="tab__header">
      <a href="#"
         class="tab__link"
         @click.prevent="active = !active">
        <span class="moving-icon" v-if="role === `teacher`">
          <img src="@/assets/moving.svg">
        </span>
        <span class="avatar-arrow" v-show="!active">
          <img src="@/assets/arrow.svg" style="transform:rotate(-90deg);"/>
        </span>
        <span class="avatar-arrow" v-show="active">
          <img src="@/assets/arrow.svg"/>
        </span>
        <strong class="mainTitle">{{ title }}</strong>
        <span class="redactor-icon" v-if="role === `teacher`">
          <img src="@/assets/addedV2.svg"/>
        </span>
        <span class="redactor-icon" v-if="role === `teacher`">
          <img src="@/assets/pen.svg"/>
        </span>
        <span class="trash-icon" v-if="role === `teacher`">
          <img src="@/assets/trash.svg"/>
        </span>
      </a>
    </div>
    <div class="tab__content" v-show="active">
      <slot/>
    </div>
  </div>
</template>

<script>
export default {
  name: "AccordionMenu",
  props: [
    "title",
    "role"
  ],
  data() {
    return {
      active: false
    };
  }
};
</script>

<style scoped>
.moving-icon {
  margin-right: 28px;
  cursor: pointer !important;
}

.avatar-arrow {
  filter: invert(22%) sepia(61%) saturate(6241%) hue-rotate(228deg) brightness(102%) contrast(92%);
  margin-right: 13px;
}

.redactor-icon {
  margin-right: 12px;
}

.trash-icon {
  margin-left: 5px;
}

.mainTitle {
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

</style>