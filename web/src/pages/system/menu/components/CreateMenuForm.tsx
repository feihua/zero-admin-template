import React, { useEffect } from 'react';
import { Form, Input, Modal  } from 'antd';
import { MenuListItem } from '../data.d';

export interface CreateFormProps {
  onCancel: () => void;
  onSubmit: (values: MenuListItem) => void;
  createModalVisible: boolean;
  parentId: number;
}
const FormItem = Form.Item;

const formLayout = {
  labelCol: { span: 7 },
  wrapperCol: { span: 13 },
};

const CreateMenuForm: React.FC<CreateFormProps> = (props) => {
  const [form] = Form.useForm();

  const { onSubmit, onCancel, createModalVisible } = props;

  useEffect(() => {
    if (form && !createModalVisible) {
      form.resetFields();
    }
  }, [props.createModalVisible]);

  useEffect(() => {
    if (props.parentId) {
      form.setFieldsValue({
        parentId: props.parentId,
      });
    }
  }, [props.parentId]);

  const handleSubmit = () => {
    if (!form) return;
    form.submit();
  };

  const handleFinish = (values: MenuListItem) => {
    if (onSubmit) {
      onSubmit(values);
    }
  };

  const renderContent = () => {
    return (
      <>
        <FormItem name="menuName" label="菜单名称">
          <Input id="update-menuName" placeholder={'请输入菜单名称'} />
        </FormItem>
        <FormItem name="parentId" label="父id" hidden>
          <Input id="update-parentId" placeholder={'请输入父id'} />
        </FormItem>
        <FormItem name="menuUrl" label="路由路径">
          <Input id="update-menuUrl" placeholder={'请输入路径'} />
        </FormItem>
        <FormItem name="menuType" label="类型">
          <Input id="update-menuType" placeholder={'请输入类型'} />
        </FormItem>
        <FormItem name="icon" label="图标">
          <Input id="update-icon" placeholder={'请输入图标'} />
        </FormItem>
        <FormItem name="sort" label="排序">
          <Input id="update-sort" placeholder={'请输入排序'} />
        </FormItem>
        <FormItem name="remark" label="备注">
          <Input id="update-remark" placeholder={'请输入备注'} />
        </FormItem>
      </>
    );
  };

  const modalFooter = { okText: '保存', onOk: handleSubmit, onCancel };

  return (
    <Modal
      forceRender
      destroyOnClose
      title="新建菜单"
      visible={createModalVisible}
      {...modalFooter}
    >
      <Form {...formLayout} form={form} onFinish={handleFinish}>
        {renderContent()}
      </Form>
    </Modal>
  );
};

export default CreateMenuForm;
