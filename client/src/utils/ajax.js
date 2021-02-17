import axios from 'axios';
import ENV from '../../env/local';
const Ajax = {
  /**
   * 公用
   * @param {*} type 
   * @param {*} graphQL 
   */
  common(type, graphQL) {
    return axios({
      url: ENV.API_BASE,
      method: 'post',
      query: `${type}{${graphQL}}`
    });
  },
  /**
   * 查询
   * @param {*} graphQL 
   */
  query(graphQL) {
    this.common('query', graphQL);
  },
  /**
   * 持久化
   * @param {*} graphQL 
   */
  mutation(graphQL) {
    this.common('mutation', graphQL);
  }
  
};
export default Ajax;