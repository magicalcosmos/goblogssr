/** 获取文章目录 */
const categoryListSchema = `
  query categoryList($currentPage: Int!, $pageSize: Int!) {
    categoryList(input: {currentPage: $currentPage, pageSize: $pageSize}) {
      categories {
        id
        name
      }
      page {
        currentPage
        pageSize
        total
      }
    }
  }
`;


/** 创建文章目录 */
const createCategorySchema = `
  mutation createCategory($name: String!) {
    createCategory(input: { name: $name }) {
      name
    }
  }
`;

/** 更新文章目录 */
const updateCategorySchema = `
  mutation updateCategory($name: String!, $id: Int!) {
    updateCategory(input: { name: $name, id: $id }) {
      name
    }
  }
`;


/** 删除文章目录 */
const deleteCategorySchema = `
  mutation deleteCategory($id: Int!) {
    deleteCategory(input: { id: $id })
  }
`;

export {
  categoryListSchema,
  createCategorySchema,
  updateCategorySchema,
  deleteCategorySchema,
};
