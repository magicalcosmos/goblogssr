import { Ajax } from '@/utils';
const Posts = {
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
    return Ajax.mutation(`
      create user {
        
      }
    `)
  }
};
export default Posts;