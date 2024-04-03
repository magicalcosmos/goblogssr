/** 保存寻求信息 */
const saveEnquirySchema = `
  mutation saveEnquiry(
    $name: String!,
    $email: String!,
    $phone: String!,
    $type: String!,
    $message: String!
  ) {
    saveEnquiry(input: {
      name: $name,
      email: $email,
      phone: $phone,
      type: $type,
      message: $message
    }) {
      id
    }
  }
`;


export {
  saveEnquirySchema,
};
