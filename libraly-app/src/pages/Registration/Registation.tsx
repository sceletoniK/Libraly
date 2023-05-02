import type {CascaderProps} from 'antd';
import {
    Alert,
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
import { useNavigate } from 'react-router-dom';
import axios from "axios";
import React, {useState} from 'react';
import './Registration.css'


const formItemLayout = {
    labelCol: {
        xs: {span: 24},
        sm: {span: 8},
    },
    wrapperCol: {
        xs: {span: 24},
        sm: {span: 16},
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
    const [err, setErr] = useState<string>('');
    const navigate = useNavigate();

    const onFinish = (values: any) => {
        axios.post(
            'http://localhost:8080/register',
            {
                login: values.username,
                password: values.password
            })
            .then((result) => {
                switch (result.status) {
                    case 418:
                        setErr('Неправильный логин или пароль')
                        break;

                    case 200:
                        navigate('/account')
                        break;

                    default:
                        setErr('Пятисотая на проде');
                        break;
                }
            })
            .catch((error) => {
                console.log(error)
                setErr('Пятисотая на проде');
            })

    };

    const handleClose = () => {
        setErr('');
    };

    return (
        <div style={{display: 'flex', justifyContent: 'center'}}>
            <Form
                {...formItemLayout}
                form={form}
                name="register"
                onFinish={onFinish}
                initialValues={{residence: ['zhejiang', 'hangzhou', 'xihu'], prefix: '86'}}
                style={{width: 450}}
                scrollToFirstError
            >
                <Form.Item
                    name="nickname"
                    label="Nickname"
                    tooltip="What do you want others to call you?"
                    rules={[{required: true, message: 'Please input your nickname!', whitespace: true}]}
                >
                    <Input/>
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
                                let regexp = new RegExp('^.*(?=.{7,})(?=.*[a-zA-Z])(?=.*\d)(?=.*[@!#$%&? "]).*$'),
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
                    <Input.Password/>
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
                        ({getFieldValue}) => ({
                            validator(_, value) {
                                if (!value || getFieldValue('password') === value) {
                                    return Promise.resolve();
                                }
                                return Promise.reject(new Error('The two passwords that you entered do not match!'));
                            },
                        }),
                    ]}
                >
                    <Input.Password/>
                </Form.Item>

                <Form.Item {...tailFormItemLayout}>
                    <Button type="primary" htmlType="submit" size='large'>
                        Register
                    </Button>
                </Form.Item>

                {err !== '' && (
                    <Alert message={err} type="error" closable afterClose={handleClose}/>
                )}
            </Form>
        </div>
    );
};

export default Registration;