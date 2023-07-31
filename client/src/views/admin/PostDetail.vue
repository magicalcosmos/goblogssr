<template lang="html">
  <div class="posts">
    <div class="posts-item">
      <label>{{$t('markdown.title')}}:</label>
      <input type="text" class="posts-title"/>
    </div>
    <div class="posts-item">
      <label>{{$t('markdown.content_brief')}}:</label>
      <editor
      :initialValue="editorText"
      :options="editorOptions"
      @load="onEditorLoad"
      @focus="onEditorFocus"
      @blur="onEditorBlur"
      @change="onEditorChange"
      @stateChange="onEditorStateChange"
      height="200px"
      />
    </div>
    <div class="posts-item">
      <label>{{$t('markdown.content_extended')}}:</label>
      <editor
      :initialValue="editorText"
      :options="editorOptions"
      @load="onEditorLoad"
      @focus="onEditorFocus"
      @blur="onEditorBlur"
      @change="onEditorChange"
      @stateChange="onEditorStateChange"
      height="400px"
      />
    </div>
    <div class="posts-item operate">
      <button>{{$t('button.save')}}</button>
    </div>
  </div>
</template>
<script>
import '@toast-ui/editor/dist/toastui-editor.css';
import '@toast-ui/chart/dist/toastui-chart.css';
import 'prismjs/themes/prism.css';
import '@toast-ui/editor-plugin-table-merged-cell/dist/toastui-editor-plugin-table-merged-cell.css';

import { Editor } from '@toast-ui/vue-editor';
import chart from '@toast-ui/editor-plugin-chart';
import codeSyntaxHighlight from '@toast-ui/editor-plugin-code-syntax-highlight';
import Prism from 'prismjs'
import colorSyntax from '@toast-ui/editor-plugin-color-syntax';
import tableMergedCell from '@toast-ui/editor-plugin-table-merged-cell';
import uml from '@toast-ui/editor-plugin-uml';
export default {
  components: {
    Editor
  },
  data() {
    return {
      editorText: '',
      editorOptions: {
        plugins: [chart, [codeSyntaxHighlight, { highlighter: Prism }], colorSyntax, tableMergedCell, uml]
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
