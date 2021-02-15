<template lang="html">
  <section class="login">
    <div>
      <input type="text" class="username"/>
    </div>
    <div>
      <input type="password" class="password"/>
    </div>
    <div>
      <input type="submit" :value="$t('login.name')" @click="handleSubmit"/>
    </div>
  </section>
</template>

<script>
  import { mapState, mapMutations } from 'vuex';
  export default {
    methods: {
      handleSubmit() {
        this.$router.push('/markdown');
      }
    }
  };
</script>
