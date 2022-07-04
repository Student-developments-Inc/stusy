<template>
  <div>
    <div class="tab__header">
      <a href="#"
         class="tab__link"
         @click.prevent="active = !active">
        <span class="avatar-arrow" v-show="!active">
          <img src="@/assets/arrow.svg" style="transform:rotate(-90deg);"/>
        </span>
        <span class="avatar-arrow" v-show="active">
          <img src="@/assets/arrow.svg"/>
        </span>
        <strong>{{ title }}</strong>
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
    "title"
  ],
  data() {
    return {
      active: false
    };
  }
};
</script>

<style scoped>
.avatar-arrow {
  filter: invert(22%) sepia(61%) saturate(6241%) hue-rotate(228deg) brightness(102%) contrast(92%);
  margin-right: 13px;
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