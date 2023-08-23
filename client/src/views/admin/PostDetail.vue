<template lang="html">
<section>
  <ul class="post-detail">
    <li class="posts-item">
      <label>{{$t('markdown.title')}}:</label>
      <InputText
        v-model="data.title"
        class="post-title"
      />
    </li>
    <li class="posts-item">
      <label>{{$t('common.state')}}:</label>
      <Dropdown
        class="dropdown"
        v-model="data.published"
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
        v-model="data.userId"
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
        v-model="data.publishAt"
        dateFormat="yy-mm-dd"
        :showIcon="true"
        :showTime="true"
        :showSeconds="true"
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
      ref="refBrief"
      :initialValue="brief"
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
      ref="refContent"
      :initialValue="content"
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
        v-model="data.categoryId"
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
      <a href="#" class="reset" @click="resetPost">reset changes</a>
      <a href="#" class="delete-post" @click="isDeletePost = true">delete-post</a>
    </li>
  </ul>
    <Dialog
      :visible.sync="isDeletePost"
      :modal="true"
      class="dlg"
    >
      <template #header>
        <div class="dialog-title">Are you sure want to delete <span class="special">{{ data.title }}</span></div>
      </template>
      <div class="content">This canot be undone</div>
      <template #footer>
        <div>
          <Button label="Cancel" class="p-button-text" @click="isDeletePost = false" />
          <Button label="Delete" @click="deletePost" />
        </div>
      </template>
    </Dialog>
</section>
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
import Dialog from 'primevue/dialog';
import { User, Article, Category } from '@/api';
import { formatDate, formatToGreenDate } from '@/utils/time';

export default {
  components: {
    Calendar,
    Editor,
    InputText,
    Dropdown,
    Button,
    Dialog,
  },
  data() {
    return {
      data: {},
      publish_date: null, 
      isDeletePost: false,
      states: [
        {
          name: 'Draft',
          code: 0
        },
        {
          name: 'Published',
          code: 1
        },
        {
          name: 'Archived',
          code: 2
        }
      ],
      users: [],
      categories: [],
      brief: '',
      content: '',
      page: {
        currentPage: 1,
        pageSize: 100000,
      },
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
    savePost() {
      this.data.brief = this.$refs.refBrief.invoke('getHTML');
      this.data.content = this.$refs.refContent.invoke('getHTML');
      Article.update(Object.assign(this.data, {
        publishAt: formatToGreenDate(new Date(this.data.publishAt)),
      })).then(() => {
        this.$toast.add({severity:'success', summary: 'Warm Tip', detail:'Save Success', life: 3000});

        this.$router.push({
          name: 'Posts',
        })
      });
    },

    getUserList() {
      User.list(this.page).then(res => {
        this.users = [];
        res.userList.users.forEach(item => {
          this.users.push({
            name: item.username,
            code: item.id,
          });
        });

        this.handleCommonResult();
      });
    },
    getCategoryList() {
      Category.list(this.page).then(res => {
        this.categories = [];
        res.categoryList.categories.forEach(item => {
          this.categories.push({
            name: item.name,
            code: item.id,
          });
        });

        this.handleCommonResult();
      });
    },
    handleCommonResult() {
      if (this.users.length && this.categories.length) {
        this.getArticleById();
      }
    },
    /**
     * 根据ID获取文章
     */
    getArticleById() {
      Article.getArticleById({
        id: this.$route.params.id
      }).then(res => {
        const data = res.getArticleById;
        data.publishAt = formatDate(new Date(data.publishAt));
        this.data = data;
        this.$refs.refBrief.invoke('setHTML', this.data.brief);
        this.$refs.refContent.invoke('setHTML', this.data.content);
      });
    },

    resetPost() {
      this.getArticleById();
    },

    deletePost() {
      Article.del({ id: this.data.id }).then(() => {});
    }
  },
  mounted() {
    this.getCategoryList();
    this.getUserList();
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

  .dialog-title .special {
    font-weight: bold;
  }
</style>
