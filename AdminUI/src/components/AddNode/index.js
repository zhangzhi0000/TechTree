/*
 * Revision History:
 *     Initial: 2017/10/21        Wang RiYu
 */

import React     from 'react';
import PropTypes from 'prop-types';
import {
  Form,
  Input,
  Select,
  Modal,
}                from 'antd';

const Option = Select.Option;
const FormItem = Form.Item;
const formItemLayout = {
  labelCol: { span: 6 },
  wrapperCol: { span: 14 }
};
const options = [
  {
    id: '0',
    label: 'Go'
  }, {
    id: '1',
    label: '语法'
  }, {
    id: '2',
    label: '框架'
  }, {
    id: '3',
    label: '实例'
  }
];

const AddNode = ({
  visible,
  onCancel,
  onAddNode,
  form: {
    getFieldDecorator,
    validateFieldsAndScroll,
    resetFields,
  },
}) => {
  const handleOK = () => {
    validateFieldsAndScroll((errors, values) => {
      if (errors) { return }

      onAddNode(values);
      resetFields();
    })
  }

  const addModalProps = {
    title: <div style={{textAlign: 'center'}}>添加节点</div>,
    width: '500px',
    visible,
    maskClosable: false,
    okText: '添加',
    onOk: handleOK,
    onCancel,
  }

  return (
    <Modal {...addModalProps}>
      <Form layout="horizontal">
        <FormItem label='节点' hasFeedback {...formItemLayout}>
          {
            getFieldDecorator('parent', {
              initialValue: '',
              rules: [
                {
                  required: true,
                },
              ],
            })(
              <Select
                showSearch
                style={{ width: '100%' }}
                placeholder="选择父节点"
                optionFilterProp="children"
                filterOption={(input, option) => option.props.children.indexOf(input) >= 0}>
                {
                  options.map(node => <Option key={node.id} value={node.id}>{node.label}</Option>)
                }
              </Select>
            )
          }
        </FormItem>
        <FormItem label='节点名' hasFeedback {...formItemLayout}>
          {
            getFieldDecorator('title', {
              initialValue: '',
              rules: [
                {
                  required: true,
                  min: 1,
                }
              ]
            })(<Input placeholder='输入节点名' />)
          }
        </FormItem>
        <FormItem label='简介' hasFeedback {...formItemLayout}>
          {
            getFieldDecorator('desc', {
              initialValue: '',
              rules: [
                {
                  required: true,
                  min: 1,
                }
              ]
            })(<Input placeholder='节点介绍' />)
          }
        </FormItem>
      </Form>
    </Modal>
  )
}

AddNode.protoTypes = {
  visible: PropTypes.bool,
  onCancel: PropTypes.func,
  onAddNode: PropTypes.func,
  form: PropTypes.object,
}

export default Form.create()(AddNode)
