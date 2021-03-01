import { Ajax } from '@/utils';
import * as CategoryGraphQL from '@/api/schema/category.graphql'
const Category = {
  /**
   * query posts
   * @param {*} params 
   */
  query(params) {

  },
  /**
   * Save the posts which user submit
   * @param {*} params 
   */
  save(params) {
    debugger;
    return Ajax.mutation({
      apiName: 'createCategory',
      variables: params
    }, CategoryGraphQL.createCategory)
  }
};
export default Category;