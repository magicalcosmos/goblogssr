import { Ajax } from '@/utils';
import { EnquirySchemas } from '@/schemas';
const Message = {
  /**
   * Save the posts
   * @param {*} params 
   */
  save(params) {
    return Ajax.mutation({
      apiName: 'saveEnquiry',
      variables: params
    }, EnquirySchemas.saveEnquirySchema)
  }
};
export default Message;
