/** 获取文章目录 */
const userListSchema = `
  query userList($currentPage: Int!, $pageSize: Int!) {
    userList(input: {currentPage: $currentPage, pageSize: $pageSize}) {
      users {
        id
        username
        email
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
const createUserSchema = `
  mutation createUser($username: String!, $password: String!, $email: String!) {
    createUser(input: { username: $username, password: $password, email: $email }) {
      username
    }
  }
`;

/** 更新文章目录 */
const updateUserSchema = `
  mutation udpateUser($username: String!, $password: String!, $email: String!) {
    updateUser(input: { username: $username, password: $password, email: $email }) {
      username
    }
  }
`;


/** 删除文章目录 */
const deleteUserSchema = `
  mutation deleteUser($id: Int!) {
    deleteUser(input: { id: $id })
  }
`;

export {
  userListSchema,
  createUserSchema,
  updateUserSchema,
  deleteUserSchema,
};
