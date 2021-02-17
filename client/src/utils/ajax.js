import axios from 'axios';
import ENV from '../../env/local';
const Ajax = {
  /**
   * 公用
   * @param {*} type 
   * @param {*} apiName 
   * @param {*} graphQL 
   */
  common(type, params, graphQL) {
    return axios({
      url: ENV.API_BASE,
      method: 'post',
      data: {
        operationName: params.apiName,
        variables: {},
        query: `${type} ${params.apiName} { ${graphQL} }`
      },
    });
  },
  /**
   * 查询
   * @param {*} params 
   * @param {*} graphQL 
   */
  query(params, graphQL) {
    return this.common('query', graphQL);
  },
  /**
   * 持久化
   * @param {*} params 
   * @param {*} graphQL 
   */
  mutation(params, graphQL) {
    return this.common('mutation', params, graphQL);
  }
  
};
export default Ajax;