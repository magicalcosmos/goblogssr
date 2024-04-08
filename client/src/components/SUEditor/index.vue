<template>
  <div v-if="editor">
    <component
      ref="dynamicRef"
      :is='editor'
      :options="editorOptions"
      :initialValue="getInitialValue"
      :height="getHeight"
      @load="$emit('onEditorLoad')"
      @focus="$emit('onEditorFocus')"
      @blur="$emit('onEditorBlur')"
      @change="$emit('onEditorChange')"
      @stateChange="$emit('onEditorStateChange')"
    ></component>
  </div>
</template>
<script>
  import '@toast-ui/editor/dist/toastui-editor.css';
  import '@toast-ui/chart/dist/toastui-chart.css';
  import '@toast-ui/editor-plugin-table-merged-cell/dist/toastui-editor-plugin-table-merged-cell.css';


  export default {
    data() {
      return {
        isClient: true,
        editor: null,
        editorOptions: {
          plugins: [],
          /* plugins: [chart, [codeSyntaxHighlight, { highlighter: Prism }], colorSyntax, tableMergedCell, uml] */
        }
      };
    },
    props: {
      initialValue: String,
      options: Object,
      height: String,
    },
    methods: {
      invoke(pattern, data) {
        return this.$refs.dynamicRef.invoke(pattern, data);
      }
    },
    computed: {
      getHeight() {
        return this.$props.height;
      },
      getInitialValue(){
        return this.$props.initialValue;
      }
    },
    mounted() {
      if (this.isClient) {
        Promise.all([
          import('@toast-ui/editor-plugin-chart'),
          import('@toast-ui/editor-plugin-code-syntax-highlight'),
          import('prismjs'),
          import('@toast-ui/editor-plugin-color-syntax'),
          import('@toast-ui/editor-plugin-table-merged-cell'),
          import('@toast-ui/editor-plugin-uml'),
          import('@toast-ui/vue-editor'),
        ]).then((promiseData) => {
          this.editorOptions.plugins[0] = promiseData[0].default;
          this.editorOptions.plugins[1] = [ promiseData[1].default, { highlighter: promiseData[2] } ];
          this.editorOptions.plugins[2] = promiseData[3].default;
          this.editorOptions.plugins[3] = promiseData[4].default;
          this.editorOptions.plugins[4] = promiseData[5].default;
          this.$nextTick(() => {
            this.editor = promiseData[6].Editor;
            this.$nextTick(() => {
              this.$emit('loaded');
            });
          });
        });
      }
    }
  };
</script>
