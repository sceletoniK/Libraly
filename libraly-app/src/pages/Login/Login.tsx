import React, {useState, useEffect} from 'react';
import { LockOutlined, UserOutlined } from '@ant-design/icons';
import { Button, Checkbox, Form, Input, Alert } from 'antd';
import './Login.css';
import { validateHeaderName } from 'http';
import Password from 'antd/es/input/Password';



const Login: React.FC = () => {

    const [err, setErr] = useState<string>('');
    const [error, setError] = useState<Error | null>(null);
    const [isLoaded, setIsLoaded] = useState(true);
    

    const FinishLogin = (values: any) => {
        setIsLoaded(false);
        fetch("http://localhost:8080/login",
            {
                method: 'POST',
                body: JSON.stringify({
                    login: values.username,
                    password: values.password
                })
            })
            .then(res => res.json())
            .then(
                (result) => {
                    setIsLoaded(true);
                    if (result == 'Unauthorized') {
                        setErr('Неправильный логин или пароль')
                    }
                    else{
                        localStorage.setItem('token', result);
                    }
                },
                (error) => {
                    setIsLoaded(true);
                    setErr(error.message);
                }
            )
    }

    const handleClose = () => {
        setErr('');
    };

    return (
        <div style={{ display: 'flex', justifyContent: 'center' }}>
            <Form
                name="normal_login"
                className="login-form"
                initialValues={{ remember: true }}
                onFinish={FinishLogin}

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
                    <Button type="primary" htmlType="submit" className="login-form-button" size='large'>
                        Log in
                    </Button>
                    Or <a href="">register now!</a>
                </Form.Item>

                {err !== '' && (
                <Alert message = {err} type="error" closable afterClose={handleClose} />
                )}
            </Form>
        </div>
    );
};

export default Login;