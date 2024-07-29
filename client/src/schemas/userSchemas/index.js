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


/** 创建用户*/
const createUserSchema = `
  mutation createUser($username: String!, $password: String!, $email: String!) {
    createUser(input: { username: $username, password: $password, email: $email }) {
      username
    }
  }
`;

/** 更新用户 */
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

/** 登录 */
const loginUserSchema = `
  mutation loginUser($email: String!, $password: String!) {
    loginUser(input: { email: $email, password: $password }) {
      token
      refreshToken      
      username
      email
    }
  }
`;

export {
  userListSchema,
  createUserSchema,
  updateUserSchema,
  deleteUserSchema,
  loginUserSchema,
};
