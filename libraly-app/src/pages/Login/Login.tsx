import React, {useState} from 'react';
import axios, {AxiosError} from 'axios'
import {LockOutlined, UserOutlined} from '@ant-design/icons';
import {Button, Form, Input, Alert} from 'antd';
import './Login.css';
import {useNavigate} from "react-router-dom";


const Login: React.FC = () => {

    const [err, setErr] = useState<string>('');
    const navigate = useNavigate();

    const FinishLogin = (values: any) => {
        axios.post("http://localhost:8080/login",
            {
                login: values.username,
                password: values.password
            },
            {
                headers: {
                    'Content-Type': 'application/json'
                }
            })
            .then((result) => {
                localStorage.setItem('token', result.data);
                navigate('/account')
            })
            .catch((error: AxiosError) => {
                if (error.response && error.response.status === 401) {
                    setErr('Неправильный логин или пароль')
                }
                else {
                    console.log(error);
                    setErr('Пятисотая на проде');
                }
            });
    }

    const handleClose = () => {
        setErr('');
    };

    return (
        <div style={{display: 'flex', justifyContent: 'center'}}>
            <Form
                name="normal_login"
                className="login-form"
                initialValues={{remember: true}}
                onFinish={FinishLogin}

            >
                <Form.Item
                    name="username"
                    rules={[{required: true, message: 'Please input your Username!'}]}
                >
                    <Input prefix={<UserOutlined className="site-form-item-icon"/>} placeholder="Username"/>
                </Form.Item>
                <Form.Item
                    name="password"
                    rules={[{required: true, message: 'Please input your Password!'}]}
                >
                    <Input
                        prefix={<LockOutlined className="site-form-item-icon"/>}
                        type="password"
                        placeholder="Password"
                    />
                </Form.Item>

                <Form.Item>
                    <Button type="primary" htmlType="submit" className="login-form-button" size='large'>
                        Log in
                    </Button>
                    Or <a href="/reg">register now!</a>
                </Form.Item>

                {err !== '' && (
                    <Alert message={err} type="error" closable afterClose={handleClose}/>
                )}
            </Form>
        </div>
    );
};

export default Login;