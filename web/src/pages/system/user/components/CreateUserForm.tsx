import React, {useEffect} from 'react';
import {Form, Input, Modal} from 'antd';
import {UserListItem} from '../data.d';

export interface CreateFormProps {
  onCancel: () => void;
  onSubmit: (values: UserListItem) => void;
  createModalVisible: boolean;
}
const FormItem = Form.Item;

const formLayout = {
  labelCol: { span: 7 },
  wrapperCol: { span: 13 },
};

const CreateUserForm: React.FC<CreateFormProps> = (props) => {
  const [form] = Form.useForm();

  const {
    onSubmit,
    onCancel,
    createModalVisible,
  } = props;

  useEffect(() => {
    if (form && !createModalVisible) {
      form.resetFields();
    }

  }, [props.createModalVisible]);


  const handleSubmit = () => {
    if (!form) return;
    form.submit();
  };

  const handleFinish = (values: UserListItem) => {
    if (onSubmit) {
      onSubmit(values);
    }
  };

  const renderContent = () => {
    return (
      <>
        <FormItem
          name="realName"
          label="用户名"
        >
          <Input id="update-name" placeholder={'请输入用户名'}/>
        </FormItem>
        <FormItem
          name="mobile"
          label="手机号"
        >
          <Input id="update-mobile" placeholder={'请输入手机号'}/>
        </FormItem>
        <FormItem
          name="sort"
          label="排序"
        >
          <Input id="update-sort" placeholder={'请输入排序'}/>
        </FormItem>
        <FormItem
          name="remark"
          label="备注"
        >
          <Input.TextArea id="add-remark" placeholder={'请输入备注'} rows={4}/>
        </FormItem>

      </>
    );
  };


  const modalFooter = { okText: '保存', onOk: handleSubmit, onCancel };

  return (
    <Modal
      forceRender
      destroyOnClose
      title="新建用户"
      visible={createModalVisible}
      {...modalFooter}
    >
      <Form
        {...formLayout}
        form={form}
        onFinish={handleFinish}
      >
        {renderContent()}
      </Form>
    </Modal>
  );
};

export default CreateUserForm;
