<template lang="html">
  <div>
  <editor
    :initialValue="editorText"
    :options="editorOptions"
    height="500px"
    initialEditType="wysiwyg"
    previewStyle="vertical"
  />
  </div>
</template>
<script>
import { Editor } from '@toast-ui/vue-editor';

export default {
  components: {
    editor: Editor
  },
  data() {
    return {
      editorText: 'This is initialValue.',
      editorOptions: {
        hideModeSwitch: true
      }
    };
  }
};
</script>