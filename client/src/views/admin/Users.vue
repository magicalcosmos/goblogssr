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
        paginator
        lazy
        :rows="10" 
        dataKey="id"
        :totalRecords="page.total"
        paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
        :rowsPerPageOptions="[5,10,25]"
        currentPageReportTemplate="Showing {first} to {last} of {totalRecords} categories" 
        responsiveLayout="scroll"
        @page="handlePage($event)"
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
          className="ckb"
          :exportable="false"
        ></Column>
        <Column
          header="User Name"
          :sortable="true"
        >
          <template #body="slotProps">
            <a class="post-title" :href="`/#/admin/categories/${slotProps.data.id}`">{{ slotProps.data.username }}</a>
          </template>
        </Column>
        <Column
          header="Email"
        >
          <template #body="slotProps">
            <a class="post-title" :href="`/#/admin/categories/${slotProps.data.id}`">{{ slotProps.data.email }}</a>
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
      v-if="isCreatePost"
      :visible.sync="isCreatePost"
      :modal="true"
      class="dlg"
    >
      <template #header>
        <h3>{{ dialogTitle }}</h3>
      </template>
      <div class="content">
        <ul class="content-list">
          <li>
            <label>UserName:</label><InputText v-model="params.username"/>
          </li>
          <li>
            <label>Email:</label><InputText v-model="params.email"/>
          </li>
          <li>
            <label>Password: </label><InputText type="password" v-model="params.password"/>
          </li>
        </ul>
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
  import { User } from '@/api';
  import SHA256 from '@/utils/sha256';

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
        dialogTitle: 'Create a new User',
        posts: [],
        filters: {
          global: {
            value: ''
          }
        },
        selectedPosts: [],
        isCreatePost: false,
        isDeletePost: false,
        currentPost: null,
        page: {
          currentPage: 1,
          pageSize: 10,
          total: 0,
        },
        params: {
          username: '',
          password: '',
          email: '',
        }
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
       this.dialogTitle = isCreate ? 'Create a new User' : 'Update User';
      },

      openNew() {
        this.params = {};
        this.currentPost = null;
        this.isCreatePost = true;
        this.getDialogTitle(true);
      },

      confirmCreate() {
        const params = Object.assign({}, this.params, { password: SHA256(this.params.password)});
        if (this.currentPost) {
          params.id = this.currentPost.id;
        }
        User[this.currentPost ? 'update' : 'save'](params).then(() => {
          this.currentPost = null
          this.isCreatePost = false;
          this.getList();
        });
      },
      
      editPost(data) {
        this.currentPost = data;
        this.isCreatePost = true;
        this.params = Object.assign(this.params, data);
        this.getDialogTitle(false);
      },

      deletePost(data) {
        this.currentPost = data;
        this.isDeletePost = true;
      },

      confirmDeletePost() {
        User.del({
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
        User.list(this.page).then(res => {
          this.posts = res.userList.users;
          this.page = Object.assign(this.page, res.userList.page);
        });
      },

      handlePage(event) {
        this.page.currentPage = ++event.page;
        this.page.pageSize = event.rows;
        this.getList();
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

  .content-list li {
    margin-top: 10px;
  }
  
  .content-list li label{
    display: inline-block;
    width: 120px;
  }

  .content-list li input {
    width: 30rem;
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
