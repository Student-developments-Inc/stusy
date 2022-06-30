<template>
  <div>
    <div class="tab__header">
      <a href="#"
         class="tab__link"
         @click.prevent="active = !active">
        <span class="avatar-arrow" v-show="!active">
          <img src="@/assets/arrow.svg"/>
        </span>
        <span class="avatar-arrow" v-show="active">
          <img src="@/assets/arrow.svg" style="transform:rotate(-90deg);"/>
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
  margin-right: 13px;
}

.tab__link {
  text-decoration: none;
  color: var(--dark);
}
</style>