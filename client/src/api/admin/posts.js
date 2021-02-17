import ajax from '@/utils/ajax';
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
    return ajax.mutation(`
      create user {
        
      }
    `)
  }
};
export default Posts;