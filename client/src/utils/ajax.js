import axios from 'axios';
import ENV from '../../env/local';
const Ajax = {
  /**
   * 公用
   * @param {*} params
   * @param {*} graphQL  return struct
   */
  common(params, graphQL) {
    return axios({
      url: ENV.API_BASE,
      method: 'post',
      data: {
        operationName: params.apiName,
        variables: params.variables || {},
        query: graphQL
      },
    });
  },
  /**
   * 查询
   * @param {*} params 
   * @param {*} graphQL 
   */
  query(params, graphQL) {
    return this.common(params, graphQL);
  },
  /**
   * 持久化
   * @param {*} params 
   * @param {*} graphQL 
   */
  mutation(params, graphQL) {
    return this.common(params, graphQL);
  }
  
};
export default Ajax;