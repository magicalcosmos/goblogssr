import ajax from '@/utils/ajax';
const Login = {
  /**
   * 登录
   * @param {*} params 
   */
  login(params) {
    params.apiName = 'createUser';
    return ajax.mutation(params, `
      ${params.apiName}(input:{username: "${params.username}", password: "${params.password}"}) {
        id,
        username,
        password
      }
    `);
  }
};
export default Login;