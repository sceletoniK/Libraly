import React from 'react';
import { LockOutlined, UserOutlined } from '@ant-design/icons';
import { Button, Checkbox, Form, Input } from 'antd';
import './Login.css';

const Account: React.FC = () => {
  const onFinish = (values: any) => {
    console.log('Received values of form: ', values);
  };

  return (
      <div style={{display: 'flex', justifyContent: 'center'}}>
          <Form
              name="normal_login"
              className="login-form"
              initialValues={{ remember: true }}
              onFinish={onFinish}
              
          >
              <Form.Item
                  name="username"
                  rules={[{ required: true, message: 'Please input your Username!' }]}
              >
                  <Input prefix={<UserOutlined className="site-form-item-icon" />} placeholder="Username" />
              </Form.Item>
              <Form.Item
                  name="password"
                  rules={[{ required: true, message: 'Please input your Password!' }]}
              >
                  <Input
                      prefix={<LockOutlined className="site-form-item-icon" />}
                      type="password"
                      placeholder="Password"
                  />
              </Form.Item>

              <Form.Item>
                  <Button type="primary" htmlType="submit" className="login-form-button">
                      Log in
                  </Button>
                  Or <a href="">register now!</a>
              </Form.Item>
          </Form>
      </div>
  );
};

export default Account;