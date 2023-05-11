import React from "react";
import {Row, Col, Table, Divider, Form, Upload, Input, Button, Checkbox, Tag} from "antd";
import {InboxOutlined} from "@ant-design/icons";
import {ColumnsType} from "antd/es/table";

interface Client {
    id: React.Key,
    name: string
    admin: boolean
    trouble: boolean
}

const ClientManager = () => {
    const data: Client[] = [
        {
            id: '1',
            name: 'admin',
            admin: true,
            trouble: false,
        },
        {
            id: '2',
            name: 'Light Society',
            admin: false,
            trouble: true,
        }
    ]

    const columns: ColumnsType<Client> = [
        {
            title: 'Логин',
            dataIndex:'name'
        },
        {
            title: 'Роль',
            dataIndex: 'admin',
            render: (value, record) => { return record.admin ? <Tag color={'lime'}>Администратор</Tag> : <Tag color={'purple'}>Пользователь</Tag>}
        },
        {
            title: 'Долги',
            dataIndex: 'trouble',
            render: (value, record) => { return record.trouble ? <Tag color={'red'}>Должник</Tag> : <Tag color={'green'}>Долгов нет</Tag>}
        }
    ]

    return (
        <>
            <Form
                labelCol={{ span: 8 }}
                wrapperCol={{ span: 16 }}
            >
                <Row justify={'center'} gutter={20}>
                    <Col span={6} style={{maxWidth: 300}}>
                        <Form.Item name="dragger" valuePropName="fileList" noStyle>
                            <Upload.Dragger name="files" action="/upload.do">
                                <p className="ant-upload-drag-icon">
                                    <InboxOutlined />
                                </p>
                                <p className="ant-upload-text">Перетащите или кликните чтобы загрузить аватар</p>
                                <p className="ant-upload-hint">Аватар должна быть один и в формате JPG\PNG</p>
                            </Upload.Dragger>
                        </Form.Item>
                    </Col>
                    <Col span={11}>
                        <br/>
                        <Form.Item
                            name="Login"
                            label="Логин"
                            tooltip="Имя пользователя в системе"
                            rules={[{ required: true, message: 'Введите логин пользователя!', whitespace: true }]}
                        >
                            <Input />
                        </Form.Item>

                        <Form.Item
                            name="password"
                            label="Пароль"
                            rules={[
                                {
                                    required: true,
                                    message: 'Введите пароль!',
                                },
                            ]}
                            hasFeedback
                        >
                            <Input.Password />
                        </Form.Item>

                        <Form.Item
                            name="confirm"
                            label="Подтвердить пароль"
                            dependencies={['password']}
                            hasFeedback
                            rules={[
                                {
                                    required: true,
                                    message: 'Подтвердите пароль!',
                                },
                                ({ getFieldValue }) => ({
                                    validator(_, value) {
                                        if (!value || getFieldValue('password') === value) {
                                            return Promise.resolve();
                                        }
                                        return Promise.reject(new Error('Пароль не совпадают!'));
                                    },
                                }),
                            ]}
                        >
                            <Input.Password />
                        </Form.Item>
                        <Form.Item label={'Администратор'}>
                            <Checkbox />
                        </Form.Item>
                        <Row justify={'center'}>
                            <Col span={12} style={{display: "flex", justifyContent: 'center'}}>
                                <Button type={'primary'} size={'large'}>
                                    Добавить пользователя
                                </Button>
                            </Col>
                        </Row>
                    </Col>
                </Row>
            </Form>
            <Divider />
            <Row justify={'center'}>
                <Col span={12}>
                    <Table bordered columns={columns} dataSource={data}/>
                </Col>
            </Row>
        </>
    )
}

export default ClientManager