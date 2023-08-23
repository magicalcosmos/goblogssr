import { Ajax } from '@/utils';
import { ArticleSchemas } from '@/schemas';
const Article = {
  /**
   * query posts
   * @param {*} params 
   */
  list(params) {
    return Ajax.query({
      apiName: 'articleList',
      variables: params
    }, ArticleSchemas.articleListSchema)
  },

  /**
   * get article by id
   * @param {*} params 
   */
  getArticleById(params) {
    return Ajax.query({
      apiName: 'getArticleById',
      variables: params
    }, ArticleSchemas.getArticleByIdSchema)
  },

  /**
   * Save the posts
   * @param {*} params 
   */
  save(params) {
    return Ajax.mutation({
      apiName: 'createArticle',
      variables: params
    }, ArticleSchemas.createArticleSchema)
  },


  /**
   * update the posts 
   * @param {*} params 
   */
  update(params) {
    return Ajax.mutation({
      apiName: 'updateArticle',
      variables: params
    }, ArticleSchemas.updateArticleSchema)
  },

  /**
   * delete the posts 
   * @param {*} params 
   */
  del(params) {
    return Ajax.mutation({
      apiName: 'deleteArticle',
      variables: params
    }, ArticleSchemas.deleteArticleSchema)
  }
};
export default Article;
