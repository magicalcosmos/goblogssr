import { Ajax } from '@/utils';
import { UserSchemas } from '@/schemas';
const User = {
  /**
   * query posts
   * @param {*} params 
   */
  list(params) {
    return Ajax.query({
      apiName: 'userList',
      variables: params
    }, UserSchemas.userListSchema)
  },

  /**
   * Save the posts
   * @param {*} params 
   */
  save(params) {
    return Ajax.mutation({
      apiName: 'createUser',
      variables: params
    }, UserSchemas.createUserSchema)
  },


  /**
   * update the posts 
   * @param {*} params 
   */
  update(params) {
    return Ajax.mutation({
      apiName: 'updateUser',
      variables: params
    }, UserSchemas.updateUserSchema)
  },

  /**
   * delete the posts 
   * @param {*} params 
   */
  del(params) {
    return Ajax.mutation({
      apiName: 'deleteUser',
      variables: params
    }, UserSchemas.deleteUserSchema)
  }
};
export default User;
