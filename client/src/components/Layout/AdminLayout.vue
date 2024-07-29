<template lang="html">
  <section>
    <AdminHeader></AdminHeader> 
    <AppMain></AppMain>
  </section>
</template>
<script>
  import AdminHeader from '../Header/AdminHeader';
  import AppMain from '../AppMain';
  export default {
    components: {
      AdminHeader,
      AppMain
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

  :deep(.pi) {
  font-size: 12px !important;
}

</style>
