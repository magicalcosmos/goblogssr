<template lang="html">
  <section id="appMainId" class="app-main">
    <transition name="fade" mode="out-in">
      <router-view></router-view>
    </transition>
  </section>
</template>

<script>
  export default {
    data: () => {
      return {
        
      }
    },
    methods: {
      
    }
  };
</script>

<style lang="scss">
  .app-main {
    margin: 0 auto;
    max-width: 1170px;
    padding: 10px;
    height: 100%;
  }
</style>
