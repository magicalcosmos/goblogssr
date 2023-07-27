<template lang="html">
  <section class="layout">
    <Header></Header> 
    <AppMain></AppMain>
    <Footer></Footer>
  </section>
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
  }

  @media (min-width: 750px) {
    .layout {
      width: 750px;
    }
  }

  @media (min-width: 992px) {
    .layout {
      width: 970px;
    }
  }

  @media (min-width: 1200px) {
    .layout {
      width: 1170px;
    }
  }

</style>
