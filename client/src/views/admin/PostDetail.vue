<template lang="html">
  <ul class="post-detail">
    <li class="posts-item">
      <label>{{$t('markdown.title')}}:</label>
      <InputText class="post-title"/>
    </li>
    <li class="posts-item">
      <label>{{$t('common.state')}}:</label>
      <Dropdown
        class="dropdown"
        v-model="state"
        :options="states"
        optionLabel="name"
        optionValue="code"
        :showClear="true"
      />
    </li>
    <li class="posts-item">
      <label>{{$t('common.author')}}:</label>
      <Dropdown
        class="dropdown"
        v-model="user"
        :options="users"
        optionLabel="name"
        optionValue="code"
        :showClear="true"
      />
    </li>
    <li class="posts-item">
      <label for="calendar">{{$t('common.publish_date')}}:</label>
      <Calendar
        id="calendar"
        v-model="publish_date"
        dateFormat="yy-mm-dd"
        :showIcon="true"
      />
      <Button
        :label="$t('button.today')"
        class="save p-button-outlined p-button-info"
        @click="today"
      />
    </li>
    <li class="posts-item">
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
    </li>

    <li class="posts-item">
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
    </li>
    <li class="posts-item">
      <label>{{$t('backend.category')}}:</label>
      <Dropdown
        class="dropdown"
        v-model="category"
        :options="categories"
        optionLabel="name"
        optionValue="code"
        :showClear="true"
      />
    </li>
    <li class="posts-item operate">
      <Button
        class="save"
        :label="$t('button.save')"
        @click="savePost"
      />
      <a href="#" class="reset">reset changes</a>
      <a href="#" class="delete-post">delete-post</a>
    </li>
  </ul>
</template>
<script>
import '@toast-ui/editor/dist/toastui-editor.css';
import '@toast-ui/chart/dist/toastui-chart.css';
import 'prismjs/themes/prism.css';
import '@toast-ui/editor-plugin-table-merged-cell/dist/toastui-editor-plugin-table-merged-cell.css';
import 'primevue/resources/primevue.min.css';
import 'primeicons/primeicons.css'
import 'primevue/resources/themes/lara-light-blue/theme.css';

import Calendar from 'primevue/calendar';
import { Editor } from '@toast-ui/vue-editor';
import chart from '@toast-ui/editor-plugin-chart';
import codeSyntaxHighlight from '@toast-ui/editor-plugin-code-syntax-highlight';
import Prism from 'prismjs'
import colorSyntax from '@toast-ui/editor-plugin-color-syntax';
import tableMergedCell from '@toast-ui/editor-plugin-table-merged-cell';
import uml from '@toast-ui/editor-plugin-uml';
import InputText from 'primevue/inputtext';
import Dropdown from 'primevue/dropdown';
import Button from 'primevue/button';

export default {
  components: {
    Calendar,
    Editor,
    InputText,
    Dropdown,
    Button
  },
  data() {
    return {
      publish_date: null, 
      state: 1,
      states: [
        {
          name: 'Draft',
          code: 1
        },
        {
          name: 'Published',
          code: 2
        },
        {
          name: 'Archived',
          code: 3
        }
      ],
      user: 2,
      users: [
        {
          name: 'Admin User',
          code: 1
        },
        {
          name: 'Brody Liao',
          code: 2
        }
      ],
      category: 1,
      categories: [
        {
          name: 'Docker',
          code: 1
        },
        {
          name: 'Frontend',
          code: 2
        },
        {
          name: 'Golang',
          code: 3
        },
        {
          name: 'Ubuntu',
          code: 4
        },
        {
          name: 'other',
          code: 5
        },
      ],
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
    today() {
      const today = new Date();
      const month = today.getUTCMonth() + 1;
      const day = today.getUTCDate();
      this.publish_date = `${today.getUTCFullYear()}-${ month > 10 ? month : '0' + month }-${day > 10 ? day : '0' + day}`;
    },
    savePost() {}
  }
};
</script>
<style scoped>
  
  .post-detail {
    
  }

  .post-detail .post-title {
    width: 100%;
  }


  .post-detail .posts-item {
    margin-top: 30px;
  }

  .post-detail .posts-item label{
    display: block;
    margin-bottom: 5px;
  }

  .dropdown {
    width: 100%;
  }
  
  .post-detail .posts-item.operate {
    border-top: 2px solid #e5e5e5;
    overflow: hidden;
    padding: 20px;
  }

  .post-detail .posts-item.operate .reset, .post-detail .posts-item.operate .delete-post {
    color: #999;
    text-decoration: none;
  }

  .post-detail .posts-item.operate .reset:hover, .post-detail .posts-item.operate .delete-post:hover {
    text-decoration: underline;
  }

  .post-detail .posts-item.operate .reset {
    display: inline-block;
    margin: 12px 0 0 10px;
    vertical-align: top;
  }

  .post-detail .posts-item.operate .delete-post {
    float: right;
  }
</style>
