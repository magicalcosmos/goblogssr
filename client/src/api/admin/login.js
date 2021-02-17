import ajax from '@/utils/ajax';
const Login = {
  /**
   * 登录
   * @param {*} params 
   */
  login(params) {
    debugger;
    return ajax.mutation(`
      UserLogin {
        
      }
    `);
  }
};
export default Login;