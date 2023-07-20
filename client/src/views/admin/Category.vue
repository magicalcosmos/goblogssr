<template>
  <section class="category">
    <div class="category-item">
      <label>{{$t('backend.category')}}：</label><input type="text" v-model="category"/>
    </div>
    <div class="category-item">
      <button @click="handleSaveCategory">{{$t('button.save')}}</button>
    </div>
  </section>
</template>
<script>
import { Category } from '@/api'
export default {
  data() {
    return {
      category: ''
    }
  },
  methods: {
    handleSaveCategory() {
      Category.save({
        name: this.category
      }).then(() => {
        debugger;
      })
    }
  }
}
</script>