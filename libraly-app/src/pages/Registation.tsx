import type { CascaderProps } from 'antd';
import {
  AutoComplete,
  Button,
  Cascader,
  Checkbox,
  Col,
  Form,
  Input,
  InputNumber,
  Row,
  Select,
} from 'antd';
import React, { useState } from 'react';
import './Registration.css'



const formItemLayout = {
  labelCol: {
    xs: { span: 24 },
    sm: { span: 8 },
  },
  wrapperCol: {
    xs: { span: 24 },
    sm: { span: 16 },
  },
};

const tailFormItemLayout = {
  wrapperCol: {
    xs: {
      span: 24,
      offset: 0,
    },
    sm: {
      span: 16,
      offset: 8,
    },
  },
};

const Registration: React.FC = () => {
  const [form] = Form.useForm();

  const onFinish = (values: any) => {
    console.log('Received values of form: ', values);
  };

  return (
    <div style={{display: 'flex', justifyContent: 'center'}}>
          <Form
              {...formItemLayout}
              form={form}
              name="register"
              onFinish={onFinish}
              initialValues={{ residence: ['zhejiang', 'hangzhou', 'xihu'], prefix: '86' }}
              style={{ width: 450 }}
              scrollToFirstError
          >
              <Form.Item
                  name="nickname"
                  label="Nickname"
                  tooltip="What do you want others to call you?"
                  rules={[{ required: true, message: 'Please input your nickname!', whitespace: true }]}
              >
                  <Input />
              </Form.Item>

              <Form.Item
                  name="password"
                  label="Password"
                  rules={[
                      {
                          required: true,
                          message: 'Please input your password!',
                      },
                      () => ({
                        validator(_, value) {
                            var regexp = new RegExp('^.*(?=.{7,})(?=.*[a-zA-Z])(?=.*\d)(?=.*[@!#$%&? "]).*$'),
                            test = regexp.test(value);
                            if (test || value === '') {
                                return Promise.resolve();
                            }
                            return Promise.reject(new Error('Пароль должен быть не менее 7 символов и содержать: буквы, цифры и спец. символы'));
                        },
                    })
                  ]}
                  hasFeedback
              >
                  <Input.Password />
              </Form.Item>

              <Form.Item
                  name="confirm"
                  label="Confirm Password"
                  dependencies={['password']}
                  hasFeedback
                  rules={[
                      {
                          required: true,
                          message: 'Please confirm your password!',
                      },
                      ({ getFieldValue }) => ({
                          validator(_, value) {
                              if (!value || getFieldValue('password') === value) {
                                  return Promise.resolve();
                              }
                              return Promise.reject(new Error('The two passwords that you entered do not match!'));
                          },
                      }),
                  ]}
              >
                  <Input.Password />
              </Form.Item>

              <Form.Item {...tailFormItemLayout}>
                  <Button type="primary" htmlType="submit" size='large'>
                        Register
                  </Button>
              </Form.Item>
          </Form>
    </div>
  );
};

export default Registration;