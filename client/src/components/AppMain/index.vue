<template lang="html">
  <section class="app-main">
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

<style>
  .app-main {
    margin: 0 auto;
    max-width: 1170px;
    padding: 10px;
  }
</style>
