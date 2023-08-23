<template lang="html">
  <section class="posts">
    <div class="opt-list">
    </div>
    <div class="card">
      <div class="toolbar clearfix">
        <div class="left">
          <Button
            label="New"
            icon="pi pi-plus"
            class="p-button-success mr-2"
            @click="openNew"
          />
          <Button
            label="Delete"
            icon="pi pi-trash"
            class="p-button-danger"
            @click="confirmDeleteSelected"
            :disabled="!selectedPosts || !selectedPosts.length"
          />
        </div>
        <div class="right">
          <Button
            label="Export"
            icon="pi pi-upload"
            class="p-button-help"
            @click="exportCSV($event)"
          />
        </div>
      </div>
      <DataTable
        :value="posts"
        lazy
        paginator
        :rows="10" 
        dataKey="id"
        :totalRecords="page.total"
        paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
        :rowsPerPageOptions="[5,10,25]"
        currentPageReportTemplate="Showing {first} to {last} of {totalRecords} posts" 
        responsiveLayout="scroll"
        @page="handlePage($event)"
      >
        <template #header>
            <div class="table-header flex flex-column md:flex-row md:justify-content-between">
                <h5 class="mb-2 md:m-0 md:align-self-center">Manage Products</h5>
                <span class="p-input-icon-left">
                    <i class="pi pi-search" />
                    <InputText v-model="filters['global'].value" placeholder="Search..." />
                </Span>
            </div>
        </template>
        <Column
          className="ckb"
          selectionMode="multiple"
          :exportable="false"
        ></Column>
        <Column
          header="Title"
          :sortable="true"
        >
          <template #body="slotProps">
            <a class="post-title" :href="`/#/admin/posts/${slotProps.data.id}`">{{ slotProps.data.title }}</a>
          </template>
        </Column>
        <Column
          field="state"
          header="State"
          :sortable="true"
        ></Column>
        <Column
          field="author"
          header="Author"
          :sortable="true"
        ></Column>
        <Column
          field="published_date"
          header="Published Date"
          :sortable="true"
        ></Column>
        <Column
          :exportable="false"
          :styles="{'min-width':'8rem'}"
        >
          <template #body="slotProps">
            <Button
              icon="pi pi-pencil"
              class="p-button-rounded p-button-success mr-2"
              @click="editPost(slotProps.data)"
            />
            <Button
              icon="pi pi-trash"
              class="p-button-rounded p-button-warning"
              @click="confirmDeletePost(slotProps.data)"
            />
          </template>
        </Column>
      </DataTable>
    </div>
    <Dialog
      :visible.sync="isCreatePost"
      :modal="true"
      class="dlg"
    >
      <template #header>
        <h3>{{ dialogTitle }}</h3>
      </template>
      <div class="content">

        <InputText class="title" v-model="title" placeholder="type title" />
      </div>
      <template #footer>
        <Button label="Cancel" class="p-button-text" @click="isCreatePost = false" />
        <Button label="Create" @click="confirmCreate" />
      </template>
    </Dialog>

  </section>
</template>
<script>
  import 'primevue/resources/primevue.min.css';
  import 'primeicons/primeicons.css'
  import 'primevue/resources/themes/lara-light-blue/theme.css';
  import DataTable from 'primevue/datatable';
  import Column from 'primevue/column';
  import InputText from 'primevue/inputtext';
  import Toolbar from 'primevue/toolbar';
  import Button from 'primevue/button';
  import FileUpload from 'primevue/fileupload';
  import Dialog from 'primevue/dialog';
  import { Article } from '@/api';

  export default {
    components: {
      Button,
      Column,
      DataTable,
      Dialog,
      FileUpload,
      InputText,
      Toolbar
    },
    data() {
      return {
        title: '',
        posts: [],
        filters: {
          global: {
            value: ''
          }
        },
        currentPost: null,
        selectedPosts: [],
        isCreatePost: false,
        dialogTitle: 'Create a new Post Article',
        page: {
          currentPage: 1,
          pageSize: 10,
          total: 0,
        },
      }
    },
    methods: {
      openNew() {
        this.isCreatePost = true;
      },

      getDialogTitle(isCreate) {
        this.dialogTitle = isCreate ? 'Create a new Post Article' : 'Update Post Article';
      },

      confirmCreate() {
        const params = {
          title: this.title,
        };
        if (this.currentPost) {
          params.id = this.currentPost.id;
        }
        Article[this.currentPost ? 'update' : 'save'](params).then((data) => {
          this.currentPost = null
          this.isCreatePost = false;
          this.$router.push({
            name: 'PostDetail',
            params: {
              id: data.createArticle.id
            }
          })
        });
      },
      editPost() {},
      confirmDeleteSelected() {
        
      },
      confirmDeletePost() {
        
      },
      exportCSV() {},

      getList() {
        Article.list(this.page).then(res => {
          this.posts = res.articleList.articles;
          this.page = Object.assign(this.page, res.articleList.page);
        });
      },

      handlePage(event) {
        this.page.currentPage = ++event.page;
        this.page.pageSize = event.rows;
        this.getList();
      },
    },
    mounted() {
      this.getList();
    },
  }
</script>
<style scoped>
.table-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.pi-trash {
  color: #ccc;
  font-size: 16px;
}

:deep(.p-datatable) .p-datatable-tbody > tr > td {
  padding: 8px;
}

:deep(.p-datatable) .p-datatable-tbody > tr > td.ckb {
  padding-left: 16px;
}

:deep(.p-button).p-button-icon-only.p-button-rounded {
  height: 20px;
  padding: 0;
  width: 20px;
}

:deep(.p-button).p-button-icon-only.p-button-rounded .pi {
  font-size: 10px;
}

.toolbar {
  background: #f8f9fa;
  border: 1px solid #dee2e6;
  border-radius: 5px;
  margin: 10px 0 20px 0;
  padding: 10px;
}

.toolbar .left {
  float: left;
}

.toolbar .right {
  float: right;
}

:deep(.p-button) {
  font-size: 12px;
  padding: 10px 12px;
}

:deep(.p-datatable-table) td {
  font-size: 14px;
}

.post-title {
  color: #333;
  text-decoration: none;
}

.post-title:hover {
  text-decoration: underline;
}

</style>
