/** 获取文章目录 */
const categoryListSchema = `
  query categoryList() {
    categoryList() {
      id
      name
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
  mutation updateCategory($name: String!, $id: ID!) {
    updateCategory(input: { name: $name, id: $id }) {
      name
    }
  }
`;


/** 删除文章目录 */
const deleteCategorySchema = `
  mutation deleteCategory($id: ID!) {
    deleteCategory(input: { id: $id })
  }
`;

export {
  categoryListSchema,
  createCategorySchema,
  updateCategorySchema,
  deleteCategorySchema,
};
