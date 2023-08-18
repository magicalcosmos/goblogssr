import { Ajax } from '@/utils';
import { CategorySchemas } from '@/schemas';
const Category = {
  /**
   * query posts
   * @param {*} params 
   */
  list(params) {
    return Ajax.query({
      apiName: 'categoryList',
      variables: params
    }, CategorySchemas.categoryListSchema)
  },

  /**
   * Save the posts
   * @param {*} params 
   */
  save(params) {
    return Ajax.mutation({
      apiName: 'createCategory',
      variables: params
    }, CategorySchemas.createCategorySchema)
  },


  /**
   * update the posts 
   * @param {*} params 
   */
  update(params) {
    return Ajax.mutation({
      apiName: 'updateCategory',
      variables: params
    }, CategorySchemas.updateCategorySchema)
  },

  /**
   * delete the posts 
   * @param {*} params 
   */
  del(params) {
    return Ajax.mutation({
      apiName: 'deleteCategory',
      variables: params
    }, CategorySchemas.deleteCategorySchema)
  }
};
export default Category;
