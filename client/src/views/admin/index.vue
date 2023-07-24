<template>
  <section>
    <ul>
      <li><router-link to="/admin/category">{{ $t('backend.category') }}</router-link></li>
      <li><router-link to="/admin/posts">{{ $t('backend.article') }}</router-link></li>
    </ul>
  </section>
</template>
<script>
export default {
  
}
</script>
