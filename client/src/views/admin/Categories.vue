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
        :paginator="true"
        :rows="10" 
        paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
        :rowsPerPageOptions="[5,10,25]"
        currentPageReportTemplate="Showing {first} to {last} of {totalRecords} categories" 
        responsiveLayout="scroll"
      >
        <template #header>
            <div class="table-header flex flex-column md:flex-row md:justify-content-between">
                <h5 class="mb-2 md:m-0 md:align-self-center">Manage Categories</h5>
                <span class="p-input-icon-left">
                    <i class="pi pi-search" />
                    <InputText v-model="filters['global'].value" placeholder="Search..." />
                </Span>
            </div>
        </template>
        <Column
          selectionMode="multiple"
          :exportable="false"
        ></Column>
        <Column
          header="Name"
          :sortable="true"
        >
          <template #body="slotProps">
            <a class="post-title" :href="`/#/admin/categories/${slotProps.data.id}`">{{ slotProps.data.name }}</a>
          </template>
        </Column>
        <Column
          :exportable="false"
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
              @click="deletePost(slotProps.data)"
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
        <InputText class="title" v-model="title" placeholder="type post category name" />
      </div>
      <template #footer>
        <Button label="Cancel" class="p-button-text" @click="cancelCreate" />
        <Button :label="currentPost ? 'Update' : 'Create'" @click="confirmCreate" />
      </template>
    </Dialog>
    <Dialog
      :visible.sync="isDeletePost"
      :modal="true"
      class="dlg"
    >
      <template #header>
        <h3>Are you sure want to delete {{ currentPost.name }}</h3>
      </template>
      <div class="content">This canot be undone</div>
      <template #footer>
        <Button label="Cancel" class="p-button-text" @click="cancelDelete" />
        <Button label="Delete" @click="confirmDeletePost" />
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
  import { Category } from '@/api';

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
        dialogTitle: 'Create a new Post Category',
        posts: [],
        filters: {
          global: {
            value: ''
          }
        },
        selectedPosts: [],
        isCreatePost: false,
        isDeletePost: false,
        currentPost: null 
      }
    },
    methods: {
      cancelCreate() {
        this.isCreatePost = false;
        this.currentPost = null;
      },

      cancelDelete() {
        this.isDeletePost = false;
        this.currentPost = null;
      },

      getDialogTitle(isCreate) {
       this.dialogTitle = isCreate ? 'Create a new Post Category' : 'Update Post Category';
      },

      openNew() {
        this.isCreatePost = true;
      },

      confirmCreate() {
        const params = {
          name: this.title,
        };
        if (this.currentPost) {
          params.id = this.currentPost.id;
        }
        Category[this.currentPost ? 'update' : 'save'](params).then(() => {
          this.currentPost = null
          this.isCreatePost = false;
          this.getList();
        });
      },
      
      editPost(data) {
        this.currentPost = data;
        this.title = data.name;
        this.isCreatePost = true;
      },

      deletePost(data) {
        this.currentPost = data;
        this.isDeletePost = true;
      },

      confirmDeletePost() {
        Category.del({
          id: this.currentPost.id,
        }).then(res => {
          this.currentPost = null;
          this.isDeletePost = false;
          this.getList();
        });
      },

      confirmDeleteSelected() {},

      exportCSV() {},

      getList() {
        Category.list().then(res => {
          this.posts = res.data.data.categoryList;
        });
      }
    },

    mounted() {
      this.getList();
    }
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
