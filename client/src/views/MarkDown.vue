<template lang="html">
  <div>
  <editor
    :initialValue="viewerText"
    :options="editorOptions"
    @load="onEditorLoad"
    @focus="onEditorFocus"
    @blur="onEditorBlur"
    @change="onEditorChange"
    @stateChange="onEditorStateChange"
    height="500px"
  />

  </div>
</template>
<script>
import { Editor } from '@toast-ui/vue-editor';
import chart from '@toast-ui/editor-plugin-chart';
import codeSyntaxHighlight from '@toast-ui/editor-plugin-code-syntax-highlight';
import hljs from 'highlight.js';
import colorSyntax from '@toast-ui/editor-plugin-color-syntax';
import tableMergedCell from '@toast-ui/editor-plugin-table-merged-cell';
import uml from '@toast-ui/editor-plugin-uml'
export default {
  components: {
    Editor
  },
  data() {
    return {
      editorText: '',
      editorOptions: {
        plugins: [chart, [codeSyntaxHighlight, { hljs }], colorSyntax, tableMergedCell, uml]
      }
    };
  },
  methods: {
    // scroll() {
    //   this.$refs.toastuiEditor.invoke('scrollTop', 10);
    // },
    // moveTop() {
    //   this.$refs.toastuiEditor.invoke('moveCursorToStart');
    // },
    // getHtml() {
    //   let html = this.$refs.toastuiEditor.invoke('getHtml');
    //   this.viewerText = html
    // },
    onEditorLoad() {
        // implement your code
    },
    onEditorFocus() {
        // implement your code
    },
    onEditorBlur() {
        // implement your code
    },
    onEditorChange() {
        // implement your code
    },
    onEditorStateChange() {
        // implement your code
    },
  }
};
</script>