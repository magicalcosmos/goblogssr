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
