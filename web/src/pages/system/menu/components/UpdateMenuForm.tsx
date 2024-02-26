import React, { useEffect } from 'react';
import { Form, Input, Modal } from 'antd';
import { MenuListItem } from '../data.d';

export interface UpdateFormProps {
  onCancel: () => void;
  onSubmit: (values: Partial<MenuListItem>) => void;
  updateModalVisible: boolean;
  currentData: Partial<MenuListItem>;
}
const FormItem = Form.Item;

const formLayout = {
  labelCol: { span: 7 },
  wrapperCol: { span: 13 },
};

const UpdateMenuForm: React.FC<UpdateFormProps> = (props) => {
  const [form] = Form.useForm();

  const { onSubmit, onCancel, updateModalVisible, currentData } = props;

  useEffect(() => {
    if (form && !updateModalVisible) {
      form.resetFields();
    }
  }, [props.updateModalVisible]);

  useEffect(() => {
    if (currentData) {
      form.setFieldsValue({
        ...currentData,
      });
    }
  }, [props.currentData]);

  const handleSubmit = () => {
    if (!form) return;
    form.submit();
  };

  const handleFinish = (values: { [key: string]: any }) => {
    if (onSubmit) {
      onSubmit(values);
    }
  };

  const renderContent = () => {
    return (
      <>
        <FormItem name="id" label="主键" hidden>
          <Input id="update-id" placeholder="请输入主键" />
        </FormItem>
        <FormItem name="menuName" label="菜单名称">
          <Input id="update-menuName" placeholder={'请输入菜单名称'} />
        </FormItem>
        <FormItem name="parentId" label="父id" hidden>
          <Input id="update-parentId" placeholder={'请输入父id'} />
        </FormItem>
        <FormItem name="menuUrl" label="路由路径">
          <Input id="update-menuUrl" placeholder={'请输入路径'} />
        </FormItem>
        <FormItem name="apiUrl" label="接口url">
          <Input id="update-apiUrl" placeholder={'请输入路径'} />
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
        <FormItem name="statusId" label="状态">
          <Input id="update-statusId" placeholder={'请输入状态'} />
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
      title="修改菜单"
      visible={updateModalVisible}
      {...modalFooter}
    >
      <Form {...formLayout} form={form} onFinish={handleFinish}>
        {renderContent()}
      </Form>
    </Modal>
  );
};

export default UpdateMenuForm;
