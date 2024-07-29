<template lang="html">
  <div class="layout">
    <Header></Header> 
    <AppMain></AppMain>
    <Footer></Footer>
  </div>
</template>
<script>
  import Header from '../Header';
  import AppMain from '../AppMain';
  import Footer from '../Footer';
  export default {
    components: {
      Header,
      AppMain,
      Footer
    },
    data: () => {
      return {
        
      }
    },
    methods: {
      
    }
  };
</script>
<style>
  .layout {
    margin: 0 auto;
    height: 100%;
  }

</style>
