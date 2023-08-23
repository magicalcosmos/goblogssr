/** 获取文章目录 */
const articleListSchema = `
  query articleList(
    $currentPage: Int!,
    $pageSize: Int!
  ) {
    articleList(input: {
      currentPage: $currentPage,
      pageSize: $pageSize
    }) {
      articles {
        id
        title
        brief
        publishAt
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
const createArticleSchema = `
  mutation createArticle(
    $title: String!
  ) {
    createArticle(input: {
      title: $title
    }) {
      id
    }
  }
`;

/** 更新文章目录 */
const updateArticleSchema = `
  mutation updateArticle(
    $title: String!,
    $id: Int!,
    $categoryId: Int!,
    $userId: Int!,
    $brief: String!,
    $content: String!,
    $published: Int!
    $publishAt: Time!
  ) {
    updateArticle(input: {
      title: $title,
      id: $id,
      categoryId: $categoryId,
      userId: $userId,
      brief: $brief,
      content: $content,
      published: $published,
      publishAt: $publishAt
    }) {
      title
    }
  }
`;


/** 删除文章目录 */
const deleteArticleSchema = `
  mutation deleteArticle($id: Int!) {
    deleteArticle(input: { id: $id })
  }
`;

/** 根据ID文章目录 */
const getArticleByIdSchema = `
  query getArticleById($id: Int!) {
    getArticleById(input: { id: $id }) {
      id
      categoryId
      userId
      title
      brief
      content
      published
      publishAt
    }
  }
`;


export {
  articleListSchema,
  getArticleByIdSchema,
  createArticleSchema,
  updateArticleSchema,
  deleteArticleSchema,
};
